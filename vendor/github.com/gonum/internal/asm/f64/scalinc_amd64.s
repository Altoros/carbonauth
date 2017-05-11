// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Some of the loop unrolling code is copied from:
// http://golang.org/src/math/big/arith_amd64.s
// which is distributed under these terms:
//
// Copyright (c) 2012 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//+build !noasm,!appengine

#include "textflag.h"

// func DscalInc(alpha float64, x []float64, n, incX uintptr)
TEXT ·ScalInc(SB), NOSPLIT, $0
	MOVHPD alpha+0(FP), X7
	MOVLPD alpha+0(FP), X7
	MOVQ   x+8(FP), R8
	MOVQ   n+32(FP), DX
	MOVQ   incX+40(FP), R10

	MOVQ $0, SI
	MOVQ R10, AX // nextX = incX
	SHLQ $1, R10 // incX *= 2

	SUBQ $2, DX // n -= 2
	JL   tail   // if n < 0

loop:
	// x[i] *= alpha unrolled 2x.
	MOVHPD 0(R8)(SI*8), X0
	MOVLPD 0(R8)(AX*8), X0
	MULPD  X7, X0
	MOVHPD X0, 0(R8)(SI*8)
	MOVLPD X0, 0(R8)(AX*8)

	ADDQ R10, SI // ix += incX
	ADDQ R10, AX // nextX += incX

	SUBQ $2, DX // n -= 2
	JGE  loop   // if n >= 0 goto loop

tail:
	ADDQ $2, DX // n += 2
	JLE  end    // if n <= 0

	// x[i] *= alpha for the last iteration if n is odd.
	MOVSD 0(R8)(SI*8), X0
	MULSD X7, X0
	MOVSD X0, 0(R8)(SI*8)

end:
	RET