// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type CallbackFunc func(wr io.Writer, tyMap map[string]Type, comment, ret, name, rawArgs string)

var compiledFuncs []string = []string{}
var dumpCpuMap = flag.Bool("dump-ops", false, "Generate the opcode map for the switch/case in op.go")

func fileHeader(wr io.Writer, tag string, imports []string) error {
	_, err := fmt.Fprintf(wr, `// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Code generated by running "go generate". DO NOT EDIT.

// +build `+tag+`

package virtual

import (
	"%s"
)

`, strings.Join(imports, "\"\n\t\""))
	return err
}

type Type int

const (
	TyPtr Type = iota
	TyStr
	TyInt32

	TyVoid
	TyError
)

func (t Type) Syscall(val string) string {
	switch t {
	case TyPtr,
		TyStr:
		return val
	case TyInt32:
		return fmt.Sprintf("uintptr(%s)", val)
	default:
		log.Fatal("Cannot syscall type: ", t)
		return ""
	}
}

func (t Type) Pop(sp string) string {
	switch t {
	case TyPtr,
		TyStr:
		return fmt.Sprintf("popPtr(%s)", sp)
	case TyInt32:
		return fmt.Sprintf("popI32(%s)", sp)
	default:
		log.Fatal("Cannot syscall type: ", t)
		return ""
	}
}

func (t Type) Write(target, val string) string {
	switch t {
	case TyPtr,
		TyStr:
		return fmt.Sprintf("writePtr(%s, %s)", target, val)
	case TyInt32:
		return fmt.Sprintf("writeI32(%s, int32(%s))", target, val)
	case TyVoid:
		// void is usually used for a function without a return value
		// so the write is a NOP
		return ""
	default:
		log.Fatal("Cannot handle write type: ", t)
	}
	return ""
}

func (t Type) FormatStr(val interface{}) string {
	arg := "%s"
	switch t {
	case TyPtr:
		arg = "%#x"
	case TyStr:
		arg = "%s"
	case TyInt32:
		arg = "%#x"
	case TyError:
		arg = "%v"
	case TyVoid:
		arg = "%d"
	default:
		log.Fatal("Cannot format type: ", t)
	}
	return arg
}

func (t Type) Format(name string) string {
	if t == TyStr {
		return fmt.Sprintf("GoUTF16String(%s)", name)
	}
	return name
}

func compileWinSyscallStubs(wr io.Writer, tyMap map[string]Type, comment, ret, name, rawArgs string) {
	if _, err := fmt.Fprintf(wr, "// %s \nfunc (c *cpu) %s() {\n\twinStub(\"%s\")\n}\n", comment, name, name); err != nil {
		log.Fatal("cannot write function header: ", err)
	}
}

func compileWinSyscall(wr io.Writer, tyMap map[string]Type, comment, ret, name, rawArgs string) {
	if _, err := fmt.Fprintf(wr, "// %s \nfunc (c *cpu) %s() {\n", comment, name); err != nil {
		log.Fatal("cannot write function header: ", err)
	}

	// handle the arguments in the correct order (reverse, so we get the right elements of the stack)
	args := strings.Split(rawArgs, ",")
	if len(rawArgs) == 0 {
		args = []string{}
	}
	syscallArgs := make([]string, len(args))
	formatStrs := make([]string, len(args))
	printArgs := make([]string, len(args))
	for i := len(args) - 1; i >= 0; i-- {
		preArg := ""
		if i == len(args)-1 {
			preArg = "\tsp := c.sp\n"
		}

		arg := args[i]
		a := strings.Split(strings.TrimSpace(arg), " ")
		if len(a) != 2 {
			log.Fatal("expected length 2: type and argument name for ", arg)
		}
		ty, exists := tyMap[a[0]]
		if !exists {
			log.Fatal("cannot map type: ", a[0], " in ", arg)
		}

		argName := a[1]
		syscallArgs[i] = ty.Syscall(argName)
		formatStrs[i] = ty.FormatStr(argName)
		printArgs[i] = ty.Format(argName)

		_, err := fmt.Fprintf(wr, "%s\tsp, %s := %s\n", preArg, a[1], ty.Pop("sp"))
		if err != nil {
			log.Fatal("cannot generate stack pop: ", err)
		}
	}

	fn := ""
	// fill with null for the respective syscall
	fillArgs := 0
	switch {
	case 0 <= len(args) && len(args) <= 3:
		fn = "Syscall"
		fillArgs = 3 - len(args)
	case 3 < len(args) && len(args) <= 6:
		fn = "Syscall6"
		fillArgs = 6 - len(args)
	case 6 < len(args) && len(args) <= 9:
		fn = "Syscall9"
		fillArgs = 9 - len(args)
	default:
		log.Fatal("Unsupported argument size: ", len(args))
	}

	for i := 0; i < fillArgs; i++ {
		syscallArgs = append(syscallArgs, "0")
	}
	argStr := strings.Join(syscallArgs, ", \n\t\t")
	if _, err := fmt.Fprintf(wr, "\n\tret, _, err := syscall.%s(proc%s.Addr(), %d, %s);\n", fn, name, len(args), argStr); err != nil {
		log.Fatal("could not write syscall: ", err)
	}

	retTy, exists := tyMap[ret]
	if !exists {
		log.Fatal("cannot map return type: ", retTy)
	}

	printArgs = append(printArgs, retTy.Format("ret"))
	if _, err := fmt.Fprintf(wr, "\tif strace {\n\t\tfmt.Fprintf(os.Stderr, \"%s(%s) %s %s\\n\", %s, err)\n\t}\n\tif err != 0 {\n\t\tc.setErrno(err)\n\t}\n",
		name, strings.Join(formatStrs, ", "), retTy.FormatStr(retTy), TyError.FormatStr("error"), strings.Join(printArgs, ", \n\t\t\t")); err != nil {
		log.Fatal("cannot generate strace: ", err)
	}

	if _, err := fmt.Fprintf(wr, "\t%s\n}\n\n", retTy.Write("c.rp", "ret")); err != nil {
		log.Fatal("cannot end function block: ", err)
	}
}

func compileWinFile(wr io.Writer, callback CallbackFunc) {
	bytes, err := ioutil.ReadFile("windows.go")
	if err != nil {
		log.Fatal("Cannot read windows.go: ", err)
	}
	content := string(bytes)

	reTy := regexp.MustCompile("//ty:(.*?): (.*)")
	reSys := regexp.MustCompile("//sys: (.*?) (.*?)\\((.*)\\);")

	// get type mappings for function signatures
	tyMatches := reTy.FindAllStringSubmatch(content, -1)
	tyMap := map[string]Type{}
	for _, match := range tyMatches {
		// the target type
		ty := match[1]
		// the aliases e.g. a list like `HANDLE, LPWXYZ`
		aliases := match[2]

		for _, alias := range strings.Split(aliases, ",") {
			alias = strings.TrimSpace(alias)
			switch ty {
			case "ptr":
				tyMap[alias] = TyPtr
			case "str":
				tyMap[alias] = TyStr
			case "int32":
				tyMap[alias] = TyInt32
			case "void":
				tyMap[alias] = TyVoid
			default:
				log.Fatal("unknown target type: ", ty)
			}
		}
	}

	// compile syscalls
	sysMatches := reSys.FindAllStringSubmatch(content, -1)
	for _, match := range sysMatches {
		// the return type of the function
		ret := strings.TrimSpace(match[1])
		// the function name
		name := strings.TrimSpace(match[2])
		// the arguments
		rawArgs := strings.TrimSpace(match[3])
		compiledFuncs = append(compiledFuncs, name)
		callback(wr, tyMap, match[0], ret, name, rawArgs)
	}
}

func main() {
	var out bytes.Buffer
	var buf bytes.Buffer

	flag.Parse()

	if err := fileHeader(&buf, "windows", []string{"fmt", "syscall", "os"}); err != nil {
		log.Fatal("Cannot write file header: ", err)
	}
	compileWinFile(&out, compileWinSyscall)

	// procCreateFileW             = modkernel32.NewProc("CreateFileW")
	if _, err := fmt.Fprintf(&buf, "var (\n\tmodkernel32           = syscall.NewLazyDLL(\"kernel32.dll\")\n"); err != nil {
		log.Fatal("cannot begin external proc declaration: ", err)
	}

	for _, fn := range compiledFuncs {
		if _, err := fmt.Fprintf(&buf, "\tproc%-30s = modkernel32.NewProc(\"%s\")\n", fn, fn); err != nil {
			log.Fatal("cannot write sid mapping: ", err)
		}
		if *dumpCpuMap {
			fmt.Printf("\t\tcase %s:\n\t\t\tc.builtin(c.%s)\n", fn, fn)
		}
	}

	if _, err := fmt.Fprintf(&buf, ")\n\n"); err != nil {
		log.Fatal("cannot terminate external proc declaration: ", err)
	}

	// register OPCodes
	if _, err := fmt.Fprintf(&buf, "func init() {\n\tregisterBuiltins(map[int]Opcode{\n"); err != nil {
		log.Fatal("cannot write registerBuiltins: ", err)
	}
	for _, fn := range compiledFuncs {
		if _, err := fmt.Fprintf(&buf, "\t\tdict.SID(\"%s\"): %s,\n", fn, fn); err != nil {
			log.Fatal("cannot write sid mapping: ", err)
		}
	}
	if _, err := fmt.Fprintf(&buf, "\t})\n}\n\n"); err != nil {
		log.Fatal("cannot terminate registerBuiltins: ", err)
	}

	// merge headers & compiled code
	if _, err := buf.Write(out.Bytes()); err != nil {
		log.Fatal("cannot merge outputs: ", err)
	}

	if err := ioutil.WriteFile("windows_impl.go", buf.Bytes(), 0655); err != nil {
		log.Fatal("cannot write windows_impl.go: ", err)
	}

	// build stubs for linux
	out.Reset()
	if err := fileHeader(&out, "!windows", []string{"fmt"}); err != nil {
		log.Fatal("Cannot write file header: ", err)
	}
	if _, err := fmt.Fprintf(&out, "func winStub(name string) {\n\tpanic(fmt.Errorf(\"%%s not supported on linux\", name))\n}\n\n"); err != nil {
		log.Fatal("Cannot write winStub!", err)
	}
	compileWinFile(&out, compileWinSyscallStubs)
	if err := ioutil.WriteFile("windows_impl_linux.go", out.Bytes(), 0655); err != nil {
		log.Fatal("cannot write windows_impl_linux.go: ", err)
	}
}
