// Code generated by running "go generate". DO NOT EDIT.

// +build ignore

// ----------------------------------------------------------------------------
//      /usr/lib/gcc/i686-linux-gnu/6/include/stddef.h
// ----------------------------------------------------------------------------
/* Copyright (C) 1989-2016 Free Software Foundation, Inc.

This file is part of GCC.

GCC is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 3, or (at your option)
any later version.

GCC is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

Under Section 7 of GPL version 3, you are granted additional
permissions described in the GCC Runtime Library Exception, version
3.1, as published by the Free Software Foundation.

You should have received a copy of the GNU General Public License and
a copy of the GCC Runtime Library Exception along with this program;
see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
<http://www.gnu.org/licenses/>.  */

typedef unsigned char __u_char;
typedef unsigned short int __u_short;
typedef unsigned int __u_int;
typedef unsigned long int __u_long;
typedef signed char __int8_t;
typedef unsigned char __uint8_t;
typedef signed short int __int16_t;
typedef unsigned short int __uint16_t;
typedef signed int __int32_t;
typedef unsigned int __uint32_t;
typedef signed long long int __int64_t;
typedef unsigned long long int __uint64_t;
typedef long long int __quad_t;
typedef unsigned long long int __u_quad_t;
typedef __u_quad_t __dev_t;
typedef unsigned int __uid_t;
typedef unsigned int __gid_t;
typedef unsigned long int __ino_t;
typedef __u_quad_t __ino64_t;
typedef unsigned int __mode_t;
typedef unsigned int __nlink_t;
typedef long int __off_t;
typedef __quad_t __off64_t;
typedef int __pid_t;
typedef struct {
	int __val[2];
} __fsid_t;
typedef long int __clock_t;
typedef unsigned long int __rlim_t;
typedef __u_quad_t __rlim64_t;
typedef unsigned int __id_t;
typedef long int __time_t;
typedef unsigned int __useconds_t;
typedef long int __suseconds_t;
typedef int __daddr_t;
typedef int __key_t;
typedef int __clockid_t;
typedef void *__timer_t;
typedef long int __blksize_t;
typedef long int __blkcnt_t;
typedef __quad_t __blkcnt64_t;
typedef unsigned long int __fsblkcnt_t;
typedef __u_quad_t __fsblkcnt64_t;
typedef unsigned long int __fsfilcnt_t;
typedef __u_quad_t __fsfilcnt64_t;
typedef int __fsword_t;
typedef int __ssize_t;
typedef long int __syscall_slong_t;
typedef unsigned long int __syscall_ulong_t;
typedef __off64_t __loff_t;
typedef __quad_t *__qaddr_t;
typedef char *__caddr_t;
typedef int __intptr_t;
typedef unsigned int __socklen_t;
typedef unsigned int size_t;
typedef __off64_t off_t;
typedef __mode_t mode_t;
extern void *mmap64(void *__addr, size_t __len, int __prot, int __flags, int __fd, __off64_t __offset);
extern int munmap(void *__addr, size_t __len);
extern int mprotect(void *__addr, size_t __len, int __prot);
extern int msync(void *__addr, size_t __len, int __flags);
extern int mlock(void *__addr, size_t __len);
extern int munlock(void *__addr, size_t __len);
extern int mlockall(int __flags);
extern int munlockall(void);
extern int shm_open(char *__name, int __oflag, mode_t __mode);
extern int shm_unlink(char *__name);
#define _BITS_TYPES_H (1)
#define __S16_TYPE short int
#define __U16_TYPE unsigned short int
#define __S32_TYPE int
#define __U32_TYPE unsigned int
#define __SLONGWORD_TYPE long int
#define __ULONGWORD_TYPE unsigned long int
#define __SQUAD_TYPE __quad_t
#define __UQUAD_TYPE __u_quad_t
#define __SWORD_TYPE int
#define __UWORD_TYPE unsigned int
#define __SLONG32_TYPE long int
#define __ULONG32_TYPE unsigned long int
#define __S64_TYPE __quad_t
#define __U64_TYPE __u_quad_t
#define __size_t__
#define __SIZE_T__
#define _SIZE_T
#define _SYS_SIZE_T_H
#define _T_SIZE_
#define _T_SIZE
#define __SIZE_T
#define _SIZE_T_
#define _BSD_SIZE_T_
#define _SIZE_T_DEFINED_
#define _SIZE_T_DEFINED
#define _BSD_SIZE_T_DEFINED_
#define _SIZE_T_DECLARED
#define ___int_size_t_h
#define _GCC_SIZE_T
#define _SIZET_
#define __size_t
#define _SYS_MMAN_H (1)
#define __off_t_defined
#define __mode_t_defined
#define MAP_FAILED ( ( void * ) - 1 )
#define mmap mmap64
#define PROT_READ (1)
#define PROT_WRITE (2)
#define PROT_EXEC (4)
#define PROT_NONE (0)
#define PROT_GROWSDOWN (16777216)
#define PROT_GROWSUP (33554432)
#define MAP_SHARED (1)
#define MAP_PRIVATE (2)
#define MAP_FIXED (16)
#define MS_ASYNC (1)
#define MS_SYNC (4)
#define MS_INVALIDATE (2)
#define MCL_CURRENT (1)
#define MCL_FUTURE (2)
#define MCL_ONFAULT (4)
