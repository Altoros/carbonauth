// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package virtual

import (
	"math"
)

func init() {
	registerBuiltins(map[int]Opcode{
		dict.SID("__signbit"):  __signbit,
		dict.SID("__signbitf"): __signbitf,
		dict.SID("acos"):       acos,
		dict.SID("asin"):       asin,
		dict.SID("atan"):       atan,
		dict.SID("ceil"):       ceil,
		dict.SID("copysign"):   copysign,
		dict.SID("cos"):        cos,
		dict.SID("cosh"):       cosh,
		dict.SID("exp"):        exp,
		dict.SID("fabs"):       fabs,
		dict.SID("floor"):      floor,
		dict.SID("isinf"):      isinf,
		dict.SID("isinff"):     isinff,
		dict.SID("isinfl"):     isinfl,
		dict.SID("log"):        log,
		dict.SID("log10"):      log10,
		dict.SID("pow"):        pow,
		dict.SID("round"):      round,
		dict.SID("sin"):        sin,
		dict.SID("sinh"):       sinh,
		dict.SID("sqrt"):       sqrt,
		dict.SID("tan"):        tan,
		dict.SID("tanh"):       tanh,
	})
}

// int isinf(double x);
func (c *cpu) isinf() {
	var r int32
	if math.IsInf(readF64(c.sp), 0) {
		r = 1
	}
	writeI32(c.rp, r)
}

// int isinff(float x);
func (c *cpu) isinff() {
	var r int32
	if math.IsInf(float64(readF32(c.sp)), 0) {
		r = 1
	}
	writeI32(c.rp, r)
}

func (c *cpu) acos()     { writeF64(c.rp, math.Acos(readF64(c.sp))) }
func (c *cpu) asin()     { writeF64(c.rp, math.Asin(readF64(c.sp))) }
func (c *cpu) atan()     { writeF64(c.rp, math.Atan(readF64(c.sp))) }
func (c *cpu) ceil()     { writeF64(c.rp, math.Ceil(readF64(c.sp))) }
func (c *cpu) copysign() { writeF64(c.rp, math.Copysign(readF64(c.sp+f64StackSz), readF64(c.sp))) }
func (c *cpu) cos()      { writeF64(c.rp, math.Cos(readF64(c.sp))) }
func (c *cpu) cosh()     { writeF64(c.rp, math.Cosh(readF64(c.sp))) }
func (c *cpu) exp()      { writeF64(c.rp, math.Exp(readF64(c.sp))) }
func (c *cpu) fabs()     { writeF64(c.rp, math.Abs(readF64(c.sp))) }
func (c *cpu) floor()    { writeF64(c.rp, math.Floor(readF64(c.sp))) }
func (c *cpu) log()      { writeF64(c.rp, math.Log(readF64(c.sp))) }
func (c *cpu) log10()    { writeF64(c.rp, math.Log10(readF64(c.sp))) }
func (c *cpu) pow()      { writeF64(c.rp, math.Pow(readF64(c.sp+f64StackSz), readF64(c.sp))) }
func (c *cpu) sin()      { writeF64(c.rp, math.Sin(readF64(c.sp))) }
func (c *cpu) sinh()     { writeF64(c.rp, math.Sinh(readF64(c.sp))) }
func (c *cpu) sqrt()     { writeF64(c.rp, math.Sqrt(readF64(c.sp))) }
func (c *cpu) tan()      { writeF64(c.rp, math.Tan(readF64(c.sp))) }
func (c *cpu) tanh()     { writeF64(c.rp, math.Tanh(readF64(c.sp))) }

// int __signbit(double x);
func (c *cpu) signbit() {
	var r int32
	if math.Signbit(readF64(c.sp)) {
		r = 1
	}
	writeI32(c.rp, r)
}

// int __signbitf(float x);
func (c *cpu) signbitf() {
	var r int32
	if math.Signbit(float64(readF32(c.sp))) {
		r = 1
	}
	writeI32(c.rp, r)
}

// double round(double x);
func (c *cpu) round() {
	v := readF64(c.sp)
	switch {
	case v < 0:
		v = math.Ceil(v - 0.5)
	case v > 0:
		v = math.Floor(v + 0.5)
	}
	writeF64(c.rp, v)
}
