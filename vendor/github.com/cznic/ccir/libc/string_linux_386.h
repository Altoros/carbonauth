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

// ----------------------------------------------------------------------------
//      /usr/include/string.h
// ----------------------------------------------------------------------------
/* Copyright (C) 1991-2016 Free Software Foundation, Inc.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <http://www.gnu.org/licenses/>.  */

typedef unsigned int size_t;
extern void *memcpy(void *__dest, void *__src, size_t __n);
extern void *memmove(void *__dest, void *__src, size_t __n);
extern void *memccpy(void *__dest, void *__src, int __c, size_t __n);
extern void *memset(void *__s, int __c, size_t __n);
extern int memcmp(void *__s1, void *__s2, size_t __n);
extern void *memchr(void *__s, int __c, size_t __n);
extern char *strcpy(char *__dest, char *__src);
extern char *strncpy(char *__dest, char *__src, size_t __n);
extern char *strcat(char *__dest, char *__src);
extern char *strncat(char *__dest, char *__src, size_t __n);
extern int strcmp(char *__s1, char *__s2);
extern int strncmp(char *__s1, char *__s2, size_t __n);
extern int strcoll(char *__s1, char *__s2);
extern size_t strxfrm(char *__dest, char *__src, size_t __n);
extern char *strdup(char *__s);
extern char *strchr(char *__s, int __c);
extern char *strrchr(char *__s, int __c);
extern size_t strcspn(char *__s, char *__reject);
extern size_t strspn(char *__s, char *__accept);
extern char *strpbrk(char *__s, char *__accept);
extern char *strstr(char *__haystack, char *__needle);
extern char *strtok(char *__s, char *__delim);
extern char *__strtok_r(char *__s, char *__delim, char **__save_ptr);
extern char *strtok_r(char *__s, char *__delim, char **__save_ptr);
extern size_t strlen(char *__s);
extern char *strerror(int __errnum);
extern void __bzero(void *__s, size_t __n);
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
#define NULL ( ( void * ) 0 )
#define _STRING_H (1)
