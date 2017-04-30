// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package virtual

import "syscall"

func init() {
	registerBuiltins(map[int]Opcode{
		dict.SID("fcntl"):  fcntl,
		dict.SID("open"):   open64,
		dict.SID("open64"): open64,
	})
}

// int fcntl(int fildes, int cmd, ...);
func (c *cpu) fcntl() { panic("unreachable") }

// int open64(const char *pathname, int flags, ...);
func (c *cpu) open64() {
	ap := c.rp - ptrStackSz
	pathname := readPtr(ap)
	ap -= i32StackSz
	flags := readI32(ap)
	ap -= i32StackSz
	mode := readU32(ap)

	path := GoString(pathname)

	h, err := syscall.Open(path, int(flags), mode)
	if err != nil {
		c.thread.setErrno(err)
	}
	// For compatibility reasons a HANDLE (atleast for file types) is always 32-bits
	// so this truncating from uintptr -> int32 is safe.
	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa384203%28v=vs.85%29.aspx
	writeI32(c.rp, int32(h))
}
