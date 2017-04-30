// Copyright 2017 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"flag"
	"fmt"
	"go/format"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"log"

	"github.com/cznic/cc"
	"github.com/cznic/ccir"
	"github.com/cznic/internal/buffer"
	"github.com/cznic/ir"
	"github.com/cznic/strutil"
	"github.com/cznic/virtual"
	"github.com/cznic/xc"
)

var (
	cpp      = flag.Bool("cpp", false, "")
	dict     = xc.Dict
	errLimit = flag.Int("errlimit", 10, "")
	filter   = flag.String("re", "", "")
	ndebug   = flag.Bool("ndebug", false, "")
	noexec   = flag.Bool("noexec", false, "")
	oLog     = flag.Bool("log", false, "")
	trace    = flag.Bool("trc", false, "")
	yydebug  = flag.Int("yydebug", 0, "")
)

func findRepo(s string) string {
	s = filepath.FromSlash(s)
	for _, v := range strings.Split(strutil.Gopath(), string(os.PathListSeparator)) {
		p := filepath.Join(v, "src", s)
		fi, err := os.Lstat(p)
		if err != nil {
			continue
		}

		if fi.IsDir() {
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			if p, err = filepath.Rel(wd, p); err != nil {
				log.Fatal(err)
			}

			return p
		}
	}
	return ""
}

func errStr(err error) string {
	switch x := err.(type) {
	case scanner.ErrorList:
		if len(x) != 1 {
			x.RemoveMultiples()
		}
		var b bytes.Buffer
		for i, v := range x {
			if i != 0 {
				b.WriteByte('\n')
			}
			b.WriteString(v.Error())
			if i == 9 {
				fmt.Fprintf(&b, "\n\t... and %v more errors", len(x)-10)
				break
			}
		}
		return b.String()
	default:
		return err.Error()
	}
}

func build(predef string, tus [][]string, opts ...cc.Opt) ([]*cc.TranslationUnit, *virtual.Binary) {
	var b buffer.Bytes
	var lpos token.Position
	if *cpp {
		opts = append(opts, cc.Cpp(func(toks []xc.Token) {
			if len(toks) != 0 {
				p := toks[0].Position()
				if p.Filename != lpos.Filename {
					fmt.Fprintf(&b, "# %d %q\n", p.Line, p.Filename)
				}
				lpos = p
			}
			for _, v := range toks {
				b.WriteString(cc.TokSrc(v))
			}
			b.WriteByte('\n')
		}))
	}

	ndbg := ""
	if *ndebug {
		ndbg = "#define NDEBUG 1"
	}
	var build [][]ir.Object
	tus = append(tus, []string{ccir.CRT0Path})
	var asta []*cc.TranslationUnit
	for _, src := range tus {
		model, err := ccir.NewModel()
		if err != nil {
			log.Fatal(err)
		}

		ast, err := cc.Parse(
			fmt.Sprintf(`
%s
#define __arch__ %s
#define __os__ %s
#include <builtin.h>
%s
`, ndbg, runtime.GOARCH, runtime.GOOS, predef),
			src,
			model,
			append([]cc.Opt{
				cc.AllowCompatibleTypedefRedefinitions(),
				cc.EnableImplicitFuncDef(),
				cc.EnableNonConstStaticInitExpressions(),
				cc.EnableWideBitFieldTypes(),
				cc.ErrLimit(*errLimit),
				cc.SysIncludePaths([]string{ccir.LibcIncludePath}),
			}, opts...)...,
		)
		if s := b.Bytes(); len(s) != 0 {
			log.Printf("\n%s", s)
			b.Close()
		}
		if err != nil {
			log.Fatal(errStr(err))
		}

		asta = append(asta, ast)
		objs, err := ccir.New(ast)
		if err != nil {
			log.Fatal(err)
		}

		if *oLog {
			for i, v := range objs {
				switch x := v.(type) {
				case *ir.DataDefinition:
					fmt.Fprintf(&b, "# [%v]: %T %v %v\n", i, x, x.ObjectBase, x.Value)
				case *ir.FunctionDefinition:
					fmt.Fprintf(&b, "# [%v]: %T %v %v\n", i, x, x.ObjectBase, x.Arguments)
					for i, v := range x.Body {
						fmt.Fprintf(&b, "%#05x\t%v\n", i, v)
					}
				default:
					log.Fatalf("[%v]: %T %v: %v", i, x, x, err)
				}
			}
		}
		for i, v := range objs {
			if err := v.Verify(); err != nil {
				switch x := v.(type) {
				case *ir.FunctionDefinition:
					fmt.Fprintf(&b, "# [%v, err]: %T %v %v\n", i, x, x.ObjectBase, x.Arguments)
					for i, v := range x.Body {
						fmt.Fprintf(&b, "%#05x\t%v\n", i, v)
					}
					log.Fatalf("# [%v]: Verify (A): %v\n%s", i, err, b.Bytes())
				default:
					log.Fatalf("[%v]: %T %v: %v", i, x, x, err)
				}
			}
		}
		build = append(build, objs)
	}

	linked, err := ir.LinkMain(build...)
	if err != nil {
		log.Fatalf("ir.LinkMain: %s\n%s", err, b.Bytes())
	}

	for _, v := range linked {
		if err := v.Verify(); err != nil {
			log.Fatal(err)
		}
	}

	bin, err := virtual.LoadMain(linked)
	if err != nil {
		log.Fatal(err)
	}

	return asta, bin
}

func macros(buf io.Writer, ast *cc.TranslationUnit) {
	var a []string
	for k, v := range ast.Macros {
		if v.Value != nil && v.Type.Kind() != cc.Bool {
			switch fn := v.DefTok.Position().Filename; {
			case
				fn == "builtin.h",
				fn == "<predefine>",
				strings.HasPrefix(fn, "predefined_"):
				// ignore
			default:
				a = append(a, string(dict.S(k)))
			}
		}
	}
	sort.Strings(a)
	for _, v := range a {
		m := ast.Macros[dict.SID(v)]
		if m.Value == nil {
			log.Fatal("TODO")
		}

		switch t := m.Type; t.Kind() {
		case
			cc.Int, cc.UInt, cc.Long, cc.ULong, cc.LongLong, cc.ULongLong,
			cc.Float, cc.LongDouble, cc.Bool:
			fmt.Fprintf(buf, "X%s = %v\n", v, m.Value)
		case cc.Ptr:
			switch t := t.Element(); t.Kind() {
			case cc.Char:
				fmt.Fprintf(buf, "X%s = %q\n", v, dict.S(int(m.Value.(cc.StringLitID))))
			default:
				log.Fatalf("%v", t.Kind())
			}
		default:
			log.Fatalf("%v", t.Kind())
		}
	}

	a = a[:0]
	for _, v := range ast.Declarations.Identifiers {
		switch x := v.Node.(type) {
		case *cc.DirectDeclarator:
			d := x.TopDeclarator()
			id, _ := d.Identifier()
			if x.EnumVal == nil {
				break
			}

			a = append(a, string(dict.S(id)))
		default:
			log.Fatalf("%T", x)
		}
	}
	sort.Strings(a)
	for _, v := range a {
		dd := ast.Declarations.Identifiers[dict.SID(v)].Node.(*cc.DirectDeclarator)
		fmt.Fprintf(buf, "X%s = %v\n", v, dd.EnumVal)
	}
}

func main() {
	const repo = "sqlite.org/sqlite-amalgamation-3180000/"

	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	flag.Parse()
	pth := findRepo(repo)
	if pth == "" {
		log.Fatalf("repository not found: %v", repo)
		return
	}

	asta, bin := build(
		`
		//#define SQLITE_DEBUG 1
		//#define SQLITE_ENABLE_API_ARMOR 1
		#define SQLITE_ENABLE_MEMSYS5 1
		#define SQLITE_USE_URI 1
		`,
		[][]string{
			{"main.c"},
			{filepath.Join(pth, "sqlite3.c")},
		},
		cc.EnableAnonymousStructFields(),
		cc.IncludePaths([]string{pth}),
	)

	var b0 bytes.Buffer
	enc := gob.NewEncoder(&b0)
	if err := enc.Encode(&bin); err != nil {
		log.Fatal(err)
	}

	var b1 bytes.Buffer
	comp := gzip.NewWriter(&b1)
	if _, err := comp.Write(b0.Bytes()); err != nil {
		log.Fatal(err)
	}

	if err := comp.Close(); err != nil {
		log.Fatal(err)
	}

	var b2 buffer.Bytes
	lic, err := ioutil.ReadFile("SQLITE-LICENSE")
	if err != nil {
		log.Fatal(err)
	}

	b2.WriteString(`// Code generated by running "go generate". DO NOT EDIT.
	
/*
	
`)
	b2.Write(lic)
	b2.WriteString(`
*/
	
package bin
	
const (
`)
	macros(&b2, asta[0])
	b2.WriteString(`
	Data = "`)
	b := b1.Bytes()
	for _, v := range b {
		switch {
		case v == '\\':
			b2.WriteString(`\\`)
		case v == '"':
			b2.WriteString(`\"`)
		case v < ' ', v >= 0x7f:
			fmt.Fprintf(&b2, `\x%02x`, v)
		default:
			b2.WriteByte(v)
		}
	}
	b2.WriteString(`"
)
`)
	b3, err := format.Source(b2.Bytes())
	if err != nil {
		b3 = b
	}
	os.MkdirAll("internal/bin", 0775)
	if err := ioutil.WriteFile(fmt.Sprintf("internal/bin/bin_%s_%s.go", runtime.GOOS, runtime.GOARCH), b3, 0664); err != nil {
		log.Fatal(err)
	}

	log.Printf("code %#08x, text %#08x, data %#08x, bss %#08x, pc2func %v, pc2line %v, symbols %v, gz %v\n",
		len(bin.Code), len(bin.Text), len(bin.Data), bin.BSS, len(bin.Functions), len(bin.Lines), len(bin.Sym), b1.Len(),
	)
}
