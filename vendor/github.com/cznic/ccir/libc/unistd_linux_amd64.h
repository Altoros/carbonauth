// Code generated by running "go generate". DO NOT EDIT.

// +build ignore

// ----------------------------------------------------------------------------
//      /usr/include/unistd.h
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

// ----------------------------------------------------------------------------
//      /usr/lib/gcc/x86_64-linux-gnu/6/include/stddef.h
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
typedef signed long int __int64_t;
typedef unsigned long int __uint64_t;
typedef long int __quad_t;
typedef unsigned long int __u_quad_t;
typedef unsigned long int __dev_t;
typedef unsigned int __uid_t;
typedef unsigned int __gid_t;
typedef unsigned long int __ino_t;
typedef unsigned long int __ino64_t;
typedef unsigned int __mode_t;
typedef unsigned long int __nlink_t;
typedef long int __off_t;
typedef long int __off64_t;
typedef int __pid_t;
typedef struct {
	int __val[2];
} __fsid_t;
typedef long int __clock_t;
typedef unsigned long int __rlim_t;
typedef unsigned long int __rlim64_t;
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
typedef long int __blkcnt64_t;
typedef unsigned long int __fsblkcnt_t;
typedef unsigned long int __fsblkcnt64_t;
typedef unsigned long int __fsfilcnt_t;
typedef unsigned long int __fsfilcnt64_t;
typedef long int __fsword_t;
typedef long int __ssize_t;
typedef long int __syscall_slong_t;
typedef unsigned long int __syscall_ulong_t;
typedef __off64_t __loff_t;
typedef __quad_t *__qaddr_t;
typedef char *__caddr_t;
typedef long int __intptr_t;
typedef unsigned int __socklen_t;
typedef __ssize_t ssize_t;
typedef __gid_t gid_t;
typedef __uid_t uid_t;
typedef __off64_t off_t;
typedef __off64_t off64_t;
typedef __useconds_t useconds_t;
typedef __pid_t pid_t;
typedef __intptr_t intptr_t;
typedef __socklen_t socklen_t;
extern int access(char *__name, int __type);
extern __off64_t lseek64(int __fd, __off64_t __offset, int __whence);
extern int close(int __fd);
extern ssize_t read(int __fd, void *__buf, size_t __nbytes);
extern ssize_t write(int __fd, void *__buf, size_t __n);
extern ssize_t pread64(int __fd, void *__buf, size_t __nbytes, __off64_t __offset);
extern ssize_t pwrite64(int __fd, void *__buf, size_t __n, __off64_t __offset);
extern int pipe(int __pipedes[2]);
extern unsigned int alarm(unsigned int __seconds);
extern unsigned int sleep(unsigned int __seconds);
extern __useconds_t ualarm(__useconds_t __value, __useconds_t __interval);
extern int usleep(__useconds_t __useconds);
extern int pause(void);
extern int chown(char *__file, __uid_t __owner, __gid_t __group);
extern int fchown(int __fd, __uid_t __owner, __gid_t __group);
extern int lchown(char *__file, __uid_t __owner, __gid_t __group);
extern int chdir(char *__path);
extern int fchdir(int __fd);
extern char *getcwd(char *__buf, size_t __size);
extern char *getwd(char *__buf);
extern int dup(int __fd);
extern int dup2(int __fd, int __fd2);
extern char **__environ;
extern int execve(char *__path, char *__argv[], char *__envp[]);
extern int execv(char *__path, char *__argv[]);
extern int execle(char *__path, char *__arg, ...);
extern int execl(char *__path, char *__arg, ...);
extern int execvp(char *__file, char *__argv[]);
extern int execlp(char *__file, char *__arg, ...);
extern int nice(int __inc);
extern void _exit(int __status);
extern long int pathconf(char *__path, int __name);
extern long int fpathconf(int __fd, int __name);
extern long int sysconf(int __name);
extern size_t confstr(int __name, char *__buf, size_t __len);
extern __pid_t getpid(void);
extern __pid_t getppid(void);
extern __pid_t getpgrp(void);
extern __pid_t __getpgid(__pid_t __pid);
extern __pid_t getpgid(__pid_t __pid);
extern int setpgid(__pid_t __pid, __pid_t __pgid);
extern int setpgrp(void);
extern __pid_t setsid(void);
extern __pid_t getsid(__pid_t __pid);
extern __uid_t getuid(void);
extern __uid_t geteuid(void);
extern __gid_t getgid(void);
extern __gid_t getegid(void);
extern int getgroups(int __size, __gid_t __list[]);
extern int setuid(__uid_t __uid);
extern int setreuid(__uid_t __ruid, __uid_t __euid);
extern int setgid(__gid_t __gid);
extern int setregid(__gid_t __rgid, __gid_t __egid);
extern __pid_t fork(void);
extern __pid_t vfork(void);
extern char *ttyname(int __fd);
extern int ttyname_r(int __fd, char *__buf, size_t __buflen);
extern int isatty(int __fd);
extern int link(char *__from, char *__to);
extern int symlink(char *__from, char *__to);
extern ssize_t readlink(char *__path, char *__buf, size_t __len);
extern int unlink(char *__name);
extern int rmdir(char *__path);
extern __pid_t tcgetpgrp(int __fd);
extern int tcsetpgrp(int __fd, __pid_t __pgrp_id);
extern char *getlogin(void);
extern int getlogin_r(char *__name, size_t __name_len);
extern int gethostname(char *__name, size_t __len);
extern int chroot(char *__path);
extern char *getpass(char *__prompt);
extern int fsync(int __fd);
extern long int gethostid(void);
extern void sync(void);
extern int getpagesize(void);
extern int getdtablesize(void);
extern int truncate64(char *__file, __off64_t __length);
extern int ftruncate64(int __fd, __off64_t __length);
extern int brk(void *__addr);
extern void *sbrk(intptr_t __delta);
extern int lockf64(int __fd, int __cmd, __off64_t __len);
extern int fdatasync(int __fildes);
extern char *crypt(char *__key, char *__salt);
extern void encrypt(char *__glibc_block, int __edflag);
extern void swab(void *__from, void *__to, ssize_t __n);
extern char *ctermid(char *__s);
extern char *cuserid(char *__s);
extern int pthread_atfork(void (*__prepare) (void), void (*__parent) (void), void (*__child) (void));
typedef long unsigned int size_t;
enum { _PC_LINK_MAX, _PC_MAX_CANON, _PC_MAX_INPUT, _PC_NAME_MAX, _PC_PATH_MAX, _PC_PIPE_BUF, _PC_CHOWN_RESTRICTED, _PC_NO_TRUNC, _PC_VDISABLE, _PC_SYNC_IO, _PC_ASYNC_IO, _PC_PRIO_IO, _PC_SOCK_MAXBUF, _PC_FILESIZEBITS, _PC_REC_INCR_XFER_SIZE, _PC_REC_MAX_XFER_SIZE, _PC_REC_MIN_XFER_SIZE, _PC_REC_XFER_ALIGN, _PC_ALLOC_SIZE_MIN, _PC_SYMLINK_MAX, _PC_2_SYMLINKS };
enum { _SC_ARG_MAX, _SC_CHILD_MAX, _SC_CLK_TCK, _SC_NGROUPS_MAX, _SC_OPEN_MAX, _SC_STREAM_MAX, _SC_TZNAME_MAX, _SC_JOB_CONTROL, _SC_SAVED_IDS, _SC_REALTIME_SIGNALS, _SC_PRIORITY_SCHEDULING, _SC_TIMERS, _SC_ASYNCHRONOUS_IO, _SC_PRIORITIZED_IO, _SC_SYNCHRONIZED_IO, _SC_FSYNC, _SC_MAPPED_FILES, _SC_MEMLOCK, _SC_MEMLOCK_RANGE, _SC_MEMORY_PROTECTION, _SC_MESSAGE_PASSING, _SC_SEMAPHORES, _SC_SHARED_MEMORY_OBJECTS, _SC_AIO_LISTIO_MAX, _SC_AIO_MAX, _SC_AIO_PRIO_DELTA_MAX, _SC_DELAYTIMER_MAX, _SC_MQ_OPEN_MAX, _SC_MQ_PRIO_MAX, _SC_VERSION, _SC_PAGESIZE, _SC_RTSIG_MAX, _SC_SEM_NSEMS_MAX, _SC_SEM_VALUE_MAX, _SC_SIGQUEUE_MAX, _SC_TIMER_MAX, _SC_BC_BASE_MAX, _SC_BC_DIM_MAX, _SC_BC_SCALE_MAX, _SC_BC_STRING_MAX, _SC_COLL_WEIGHTS_MAX, _SC_EQUIV_CLASS_MAX, _SC_EXPR_NEST_MAX, _SC_LINE_MAX, _SC_RE_DUP_MAX, _SC_CHARCLASS_NAME_MAX, _SC_2_VERSION, _SC_2_C_BIND, _SC_2_C_DEV, _SC_2_FORT_DEV, _SC_2_FORT_RUN, _SC_2_SW_DEV, _SC_2_LOCALEDEF, _SC_PII, _SC_PII_XTI, _SC_PII_SOCKET, _SC_PII_INTERNET, _SC_PII_OSI, _SC_POLL, _SC_SELECT, _SC_UIO_MAXIOV, _SC_IOV_MAX = 60, _SC_PII_INTERNET_STREAM, _SC_PII_INTERNET_DGRAM, _SC_PII_OSI_COTS, _SC_PII_OSI_CLTS, _SC_PII_OSI_M, _SC_T_IOV_MAX, _SC_THREADS, _SC_THREAD_SAFE_FUNCTIONS, _SC_GETGR_R_SIZE_MAX, _SC_GETPW_R_SIZE_MAX, _SC_LOGIN_NAME_MAX, _SC_TTY_NAME_MAX, _SC_THREAD_DESTRUCTOR_ITERATIONS, _SC_THREAD_KEYS_MAX, _SC_THREAD_STACK_MIN, _SC_THREAD_THREADS_MAX, _SC_THREAD_ATTR_STACKADDR, _SC_THREAD_ATTR_STACKSIZE, _SC_THREAD_PRIORITY_SCHEDULING, _SC_THREAD_PRIO_INHERIT, _SC_THREAD_PRIO_PROTECT, _SC_THREAD_PROCESS_SHARED, _SC_NPROCESSORS_CONF, _SC_NPROCESSORS_ONLN, _SC_PHYS_PAGES, _SC_AVPHYS_PAGES, _SC_ATEXIT_MAX, _SC_PASS_MAX, _SC_XOPEN_VERSION, _SC_XOPEN_XCU_VERSION, _SC_XOPEN_UNIX, _SC_XOPEN_CRYPT, _SC_XOPEN_ENH_I18N, _SC_XOPEN_SHM, _SC_2_CHAR_TERM, _SC_2_C_VERSION, _SC_2_UPE, _SC_XOPEN_XPG2, _SC_XOPEN_XPG3, _SC_XOPEN_XPG4, _SC_CHAR_BIT, _SC_CHAR_MAX, _SC_CHAR_MIN, _SC_INT_MAX, _SC_INT_MIN, _SC_LONG_BIT, _SC_WORD_BIT, _SC_MB_LEN_MAX, _SC_NZERO, _SC_SSIZE_MAX, _SC_SCHAR_MAX, _SC_SCHAR_MIN, _SC_SHRT_MAX, _SC_SHRT_MIN, _SC_UCHAR_MAX, _SC_UINT_MAX, _SC_ULONG_MAX, _SC_USHRT_MAX, _SC_NL_ARGMAX, _SC_NL_LANGMAX, _SC_NL_MSGMAX, _SC_NL_NMAX, _SC_NL_SETMAX, _SC_NL_TEXTMAX, _SC_XBS5_ILP32_OFF32, _SC_XBS5_ILP32_OFFBIG, _SC_XBS5_LP64_OFF64, _SC_XBS5_LPBIG_OFFBIG, _SC_XOPEN_LEGACY, _SC_XOPEN_REALTIME, _SC_XOPEN_REALTIME_THREADS, _SC_ADVISORY_INFO, _SC_BARRIERS, _SC_BASE, _SC_C_LANG_SUPPORT, _SC_C_LANG_SUPPORT_R, _SC_CLOCK_SELECTION, _SC_CPUTIME, _SC_THREAD_CPUTIME, _SC_DEVICE_IO, _SC_DEVICE_SPECIFIC, _SC_DEVICE_SPECIFIC_R, _SC_FD_MGMT, _SC_FIFO, _SC_PIPE, _SC_FILE_ATTRIBUTES, _SC_FILE_LOCKING, _SC_FILE_SYSTEM, _SC_MONOTONIC_CLOCK, _SC_MULTI_PROCESS, _SC_SINGLE_PROCESS, _SC_NETWORKING, _SC_READER_WRITER_LOCKS, _SC_SPIN_LOCKS, _SC_REGEXP, _SC_REGEX_VERSION, _SC_SHELL, _SC_SIGNALS, _SC_SPAWN, _SC_SPORADIC_SERVER, _SC_THREAD_SPORADIC_SERVER, _SC_SYSTEM_DATABASE, _SC_SYSTEM_DATABASE_R, _SC_TIMEOUTS, _SC_TYPED_MEMORY_OBJECTS, _SC_USER_GROUPS, _SC_USER_GROUPS_R, _SC_2_PBS, _SC_2_PBS_ACCOUNTING, _SC_2_PBS_LOCATE, _SC_2_PBS_MESSAGE, _SC_2_PBS_TRACK, _SC_SYMLOOP_MAX, _SC_STREAMS, _SC_2_PBS_CHECKPOINT, _SC_V6_ILP32_OFF32, _SC_V6_ILP32_OFFBIG, _SC_V6_LP64_OFF64, _SC_V6_LPBIG_OFFBIG, _SC_HOST_NAME_MAX, _SC_TRACE, _SC_TRACE_EVENT_FILTER, _SC_TRACE_INHERIT, _SC_TRACE_LOG, _SC_LEVEL1_ICACHE_SIZE, _SC_LEVEL1_ICACHE_ASSOC, _SC_LEVEL1_ICACHE_LINESIZE, _SC_LEVEL1_DCACHE_SIZE, _SC_LEVEL1_DCACHE_ASSOC, _SC_LEVEL1_DCACHE_LINESIZE, _SC_LEVEL2_CACHE_SIZE, _SC_LEVEL2_CACHE_ASSOC, _SC_LEVEL2_CACHE_LINESIZE, _SC_LEVEL3_CACHE_SIZE, _SC_LEVEL3_CACHE_ASSOC, _SC_LEVEL3_CACHE_LINESIZE, _SC_LEVEL4_CACHE_SIZE, _SC_LEVEL4_CACHE_ASSOC, _SC_LEVEL4_CACHE_LINESIZE, _SC_IPV6 = 235, _SC_RAW_SOCKETS, _SC_V7_ILP32_OFF32, _SC_V7_ILP32_OFFBIG, _SC_V7_LP64_OFF64, _SC_V7_LPBIG_OFFBIG, _SC_SS_REPL_MAX, _SC_TRACE_EVENT_NAME_MAX, _SC_TRACE_NAME_MAX, _SC_TRACE_SYS_MAX, _SC_TRACE_USER_EVENT_MAX, _SC_XOPEN_STREAMS, _SC_THREAD_ROBUST_PRIO_INHERIT, _SC_THREAD_ROBUST_PRIO_PROTECT };
enum { _CS_PATH, _CS_V6_WIDTH_RESTRICTED_ENVS, _CS_GNU_LIBC_VERSION, _CS_GNU_LIBPTHREAD_VERSION, _CS_V5_WIDTH_RESTRICTED_ENVS, _CS_V7_WIDTH_RESTRICTED_ENVS, _CS_LFS_CFLAGS = 1000, _CS_LFS_LDFLAGS, _CS_LFS_LIBS, _CS_LFS_LINTFLAGS, _CS_LFS64_CFLAGS, _CS_LFS64_LDFLAGS, _CS_LFS64_LIBS, _CS_LFS64_LINTFLAGS, _CS_XBS5_ILP32_OFF32_CFLAGS = 1100, _CS_XBS5_ILP32_OFF32_LDFLAGS, _CS_XBS5_ILP32_OFF32_LIBS, _CS_XBS5_ILP32_OFF32_LINTFLAGS, _CS_XBS5_ILP32_OFFBIG_CFLAGS, _CS_XBS5_ILP32_OFFBIG_LDFLAGS, _CS_XBS5_ILP32_OFFBIG_LIBS, _CS_XBS5_ILP32_OFFBIG_LINTFLAGS, _CS_XBS5_LP64_OFF64_CFLAGS, _CS_XBS5_LP64_OFF64_LDFLAGS, _CS_XBS5_LP64_OFF64_LIBS, _CS_XBS5_LP64_OFF64_LINTFLAGS, _CS_XBS5_LPBIG_OFFBIG_CFLAGS, _CS_XBS5_LPBIG_OFFBIG_LDFLAGS, _CS_XBS5_LPBIG_OFFBIG_LIBS, _CS_XBS5_LPBIG_OFFBIG_LINTFLAGS, _CS_POSIX_V6_ILP32_OFF32_CFLAGS, _CS_POSIX_V6_ILP32_OFF32_LDFLAGS, _CS_POSIX_V6_ILP32_OFF32_LIBS, _CS_POSIX_V6_ILP32_OFF32_LINTFLAGS, _CS_POSIX_V6_ILP32_OFFBIG_CFLAGS, _CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS, _CS_POSIX_V6_ILP32_OFFBIG_LIBS, _CS_POSIX_V6_ILP32_OFFBIG_LINTFLAGS, _CS_POSIX_V6_LP64_OFF64_CFLAGS, _CS_POSIX_V6_LP64_OFF64_LDFLAGS, _CS_POSIX_V6_LP64_OFF64_LIBS, _CS_POSIX_V6_LP64_OFF64_LINTFLAGS, _CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS, _CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS, _CS_POSIX_V6_LPBIG_OFFBIG_LIBS, _CS_POSIX_V6_LPBIG_OFFBIG_LINTFLAGS, _CS_POSIX_V7_ILP32_OFF32_CFLAGS, _CS_POSIX_V7_ILP32_OFF32_LDFLAGS, _CS_POSIX_V7_ILP32_OFF32_LIBS, _CS_POSIX_V7_ILP32_OFF32_LINTFLAGS, _CS_POSIX_V7_ILP32_OFFBIG_CFLAGS, _CS_POSIX_V7_ILP32_OFFBIG_LDFLAGS, _CS_POSIX_V7_ILP32_OFFBIG_LIBS, _CS_POSIX_V7_ILP32_OFFBIG_LINTFLAGS, _CS_POSIX_V7_LP64_OFF64_CFLAGS, _CS_POSIX_V7_LP64_OFF64_LDFLAGS, _CS_POSIX_V7_LP64_OFF64_LIBS, _CS_POSIX_V7_LP64_OFF64_LINTFLAGS, _CS_POSIX_V7_LPBIG_OFFBIG_CFLAGS, _CS_POSIX_V7_LPBIG_OFFBIG_LDFLAGS, _CS_POSIX_V7_LPBIG_OFFBIG_LIBS, _CS_POSIX_V7_LPBIG_OFFBIG_LINTFLAGS, _CS_V6_ENV, _CS_V7_ENV };
#define _BITS_TYPES_H (1)
#define __S16_TYPE short int
#define __U16_TYPE unsigned short int
#define __S32_TYPE int
#define __U32_TYPE unsigned int
#define __SLONGWORD_TYPE long int
#define __ULONGWORD_TYPE unsigned long int
#define __SQUAD_TYPE long int
#define __UQUAD_TYPE unsigned long int
#define __SWORD_TYPE long int
#define __UWORD_TYPE unsigned long int
#define __SLONG32_TYPE int
#define __ULONG32_TYPE unsigned int
#define __S64_TYPE long int
#define __U64_TYPE unsigned long int
#define _UNISTD_H (1)
#define _POSIX_VERSION (199506l)
#define __POSIX2_THIS_VERSION (199506l)
#define _POSIX2_VERSION __POSIX2_THIS_VERSION
#define _POSIX2_C_VERSION __POSIX2_THIS_VERSION
#define _POSIX2_C_BIND __POSIX2_THIS_VERSION
#define _POSIX2_C_DEV __POSIX2_THIS_VERSION
#define _POSIX2_SW_DEV __POSIX2_THIS_VERSION
#define _POSIX2_LOCALEDEF __POSIX2_THIS_VERSION
#define _XOPEN_VERSION (500)
#define _XOPEN_XCU_VERSION (4)
#define _XOPEN_XPG2 (1)
#define _XOPEN_XPG3 (1)
#define _XOPEN_XPG4 (1)
#define _XOPEN_UNIX (1)
#define _XOPEN_CRYPT (1)
#define _XOPEN_ENH_I18N (1)
#define _XOPEN_LEGACY (1)
#define STDIN_FILENO (0)
#define STDOUT_FILENO (1)
#define STDERR_FILENO (2)
#define __ssize_t_defined
#define __gid_t_defined
#define __uid_t_defined
#define __off_t_defined
#define __off64_t_defined
#define __useconds_t_defined
#define __pid_t_defined
#define __intptr_t_defined
#define __socklen_t_defined
#define R_OK (4)
#define W_OK (2)
#define X_OK (1)
#define F_OK (0)
#define SEEK_SET (0)
#define SEEK_CUR (1)
#define SEEK_END (2)
#define lseek lseek64
#define pread pread64
#define pwrite pwrite64
#define truncate truncate64
#define ftruncate ftruncate64
#define F_ULOCK (0)
#define F_LOCK (1)
#define F_TLOCK (2)
#define F_TEST (3)
#define lockf lockf64
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
#define _PC_LINK_MAX _PC_LINK_MAX
#define _PC_MAX_CANON _PC_MAX_CANON
#define _PC_MAX_INPUT _PC_MAX_INPUT
#define _PC_NAME_MAX _PC_NAME_MAX
#define _PC_PATH_MAX _PC_PATH_MAX
#define _PC_PIPE_BUF _PC_PIPE_BUF
#define _PC_CHOWN_RESTRICTED _PC_CHOWN_RESTRICTED
#define _PC_NO_TRUNC _PC_NO_TRUNC
#define _PC_VDISABLE _PC_VDISABLE
#define _PC_SYNC_IO _PC_SYNC_IO
#define _PC_ASYNC_IO _PC_ASYNC_IO
#define _PC_PRIO_IO _PC_PRIO_IO
#define _PC_SOCK_MAXBUF _PC_SOCK_MAXBUF
#define _PC_FILESIZEBITS _PC_FILESIZEBITS
#define _PC_REC_INCR_XFER_SIZE _PC_REC_INCR_XFER_SIZE
#define _PC_REC_MAX_XFER_SIZE _PC_REC_MAX_XFER_SIZE
#define _PC_REC_MIN_XFER_SIZE _PC_REC_MIN_XFER_SIZE
#define _PC_REC_XFER_ALIGN _PC_REC_XFER_ALIGN
#define _PC_ALLOC_SIZE_MIN _PC_ALLOC_SIZE_MIN
#define _PC_SYMLINK_MAX _PC_SYMLINK_MAX
#define _PC_2_SYMLINKS _PC_2_SYMLINKS
#define _SC_ARG_MAX _SC_ARG_MAX
#define _SC_CHILD_MAX _SC_CHILD_MAX
#define _SC_CLK_TCK _SC_CLK_TCK
#define _SC_NGROUPS_MAX _SC_NGROUPS_MAX
#define _SC_OPEN_MAX _SC_OPEN_MAX
#define _SC_STREAM_MAX _SC_STREAM_MAX
#define _SC_TZNAME_MAX _SC_TZNAME_MAX
#define _SC_JOB_CONTROL _SC_JOB_CONTROL
#define _SC_SAVED_IDS _SC_SAVED_IDS
#define _SC_REALTIME_SIGNALS _SC_REALTIME_SIGNALS
#define _SC_PRIORITY_SCHEDULING _SC_PRIORITY_SCHEDULING
#define _SC_TIMERS _SC_TIMERS
#define _SC_ASYNCHRONOUS_IO _SC_ASYNCHRONOUS_IO
#define _SC_PRIORITIZED_IO _SC_PRIORITIZED_IO
#define _SC_SYNCHRONIZED_IO _SC_SYNCHRONIZED_IO
#define _SC_FSYNC _SC_FSYNC
#define _SC_MAPPED_FILES _SC_MAPPED_FILES
#define _SC_MEMLOCK _SC_MEMLOCK
#define _SC_MEMLOCK_RANGE _SC_MEMLOCK_RANGE
#define _SC_MEMORY_PROTECTION _SC_MEMORY_PROTECTION
#define _SC_MESSAGE_PASSING _SC_MESSAGE_PASSING
#define _SC_SEMAPHORES _SC_SEMAPHORES
#define _SC_SHARED_MEMORY_OBJECTS _SC_SHARED_MEMORY_OBJECTS
#define _SC_AIO_LISTIO_MAX _SC_AIO_LISTIO_MAX
#define _SC_AIO_MAX _SC_AIO_MAX
#define _SC_AIO_PRIO_DELTA_MAX _SC_AIO_PRIO_DELTA_MAX
#define _SC_DELAYTIMER_MAX _SC_DELAYTIMER_MAX
#define _SC_MQ_OPEN_MAX _SC_MQ_OPEN_MAX
#define _SC_MQ_PRIO_MAX _SC_MQ_PRIO_MAX
#define _SC_VERSION _SC_VERSION
#define _SC_PAGESIZE _SC_PAGESIZE
#define _SC_PAGE_SIZE _SC_PAGESIZE
#define _SC_RTSIG_MAX _SC_RTSIG_MAX
#define _SC_SEM_NSEMS_MAX _SC_SEM_NSEMS_MAX
#define _SC_SEM_VALUE_MAX _SC_SEM_VALUE_MAX
#define _SC_SIGQUEUE_MAX _SC_SIGQUEUE_MAX
#define _SC_TIMER_MAX _SC_TIMER_MAX
#define _SC_BC_BASE_MAX _SC_BC_BASE_MAX
#define _SC_BC_DIM_MAX _SC_BC_DIM_MAX
#define _SC_BC_SCALE_MAX _SC_BC_SCALE_MAX
#define _SC_BC_STRING_MAX _SC_BC_STRING_MAX
#define _SC_COLL_WEIGHTS_MAX _SC_COLL_WEIGHTS_MAX
#define _SC_EQUIV_CLASS_MAX _SC_EQUIV_CLASS_MAX
#define _SC_EXPR_NEST_MAX _SC_EXPR_NEST_MAX
#define _SC_LINE_MAX _SC_LINE_MAX
#define _SC_RE_DUP_MAX _SC_RE_DUP_MAX
#define _SC_CHARCLASS_NAME_MAX _SC_CHARCLASS_NAME_MAX
#define _SC_2_VERSION _SC_2_VERSION
#define _SC_2_C_BIND _SC_2_C_BIND
#define _SC_2_C_DEV _SC_2_C_DEV
#define _SC_2_FORT_DEV _SC_2_FORT_DEV
#define _SC_2_FORT_RUN _SC_2_FORT_RUN
#define _SC_2_SW_DEV _SC_2_SW_DEV
#define _SC_2_LOCALEDEF _SC_2_LOCALEDEF
#define _SC_PII _SC_PII
#define _SC_PII_XTI _SC_PII_XTI
#define _SC_PII_SOCKET _SC_PII_SOCKET
#define _SC_PII_INTERNET _SC_PII_INTERNET
#define _SC_PII_OSI _SC_PII_OSI
#define _SC_POLL _SC_POLL
#define _SC_SELECT _SC_SELECT
#define _SC_UIO_MAXIOV _SC_UIO_MAXIOV
#define _SC_IOV_MAX _SC_IOV_MAX
#define _SC_PII_INTERNET_STREAM _SC_PII_INTERNET_STREAM
#define _SC_PII_INTERNET_DGRAM _SC_PII_INTERNET_DGRAM
#define _SC_PII_OSI_COTS _SC_PII_OSI_COTS
#define _SC_PII_OSI_CLTS _SC_PII_OSI_CLTS
#define _SC_PII_OSI_M _SC_PII_OSI_M
#define _SC_T_IOV_MAX _SC_T_IOV_MAX
#define _SC_THREADS _SC_THREADS
#define _SC_THREAD_SAFE_FUNCTIONS _SC_THREAD_SAFE_FUNCTIONS
#define _SC_GETGR_R_SIZE_MAX _SC_GETGR_R_SIZE_MAX
#define _SC_GETPW_R_SIZE_MAX _SC_GETPW_R_SIZE_MAX
#define _SC_LOGIN_NAME_MAX _SC_LOGIN_NAME_MAX
#define _SC_TTY_NAME_MAX _SC_TTY_NAME_MAX
#define _SC_THREAD_DESTRUCTOR_ITERATIONS _SC_THREAD_DESTRUCTOR_ITERATIONS
#define _SC_THREAD_KEYS_MAX _SC_THREAD_KEYS_MAX
#define _SC_THREAD_STACK_MIN _SC_THREAD_STACK_MIN
#define _SC_THREAD_THREADS_MAX _SC_THREAD_THREADS_MAX
#define _SC_THREAD_ATTR_STACKADDR _SC_THREAD_ATTR_STACKADDR
#define _SC_THREAD_ATTR_STACKSIZE _SC_THREAD_ATTR_STACKSIZE
#define _SC_THREAD_PRIORITY_SCHEDULING _SC_THREAD_PRIORITY_SCHEDULING
#define _SC_THREAD_PRIO_INHERIT _SC_THREAD_PRIO_INHERIT
#define _SC_THREAD_PRIO_PROTECT _SC_THREAD_PRIO_PROTECT
#define _SC_THREAD_PROCESS_SHARED _SC_THREAD_PROCESS_SHARED
#define _SC_NPROCESSORS_CONF _SC_NPROCESSORS_CONF
#define _SC_NPROCESSORS_ONLN _SC_NPROCESSORS_ONLN
#define _SC_PHYS_PAGES _SC_PHYS_PAGES
#define _SC_AVPHYS_PAGES _SC_AVPHYS_PAGES
#define _SC_ATEXIT_MAX _SC_ATEXIT_MAX
#define _SC_PASS_MAX _SC_PASS_MAX
#define _SC_XOPEN_VERSION _SC_XOPEN_VERSION
#define _SC_XOPEN_XCU_VERSION _SC_XOPEN_XCU_VERSION
#define _SC_XOPEN_UNIX _SC_XOPEN_UNIX
#define _SC_XOPEN_CRYPT _SC_XOPEN_CRYPT
#define _SC_XOPEN_ENH_I18N _SC_XOPEN_ENH_I18N
#define _SC_XOPEN_SHM _SC_XOPEN_SHM
#define _SC_2_CHAR_TERM _SC_2_CHAR_TERM
#define _SC_2_C_VERSION _SC_2_C_VERSION
#define _SC_2_UPE _SC_2_UPE
#define _SC_XOPEN_XPG2 _SC_XOPEN_XPG2
#define _SC_XOPEN_XPG3 _SC_XOPEN_XPG3
#define _SC_XOPEN_XPG4 _SC_XOPEN_XPG4
#define _SC_CHAR_BIT _SC_CHAR_BIT
#define _SC_CHAR_MAX _SC_CHAR_MAX
#define _SC_CHAR_MIN _SC_CHAR_MIN
#define _SC_INT_MAX _SC_INT_MAX
#define _SC_INT_MIN _SC_INT_MIN
#define _SC_LONG_BIT _SC_LONG_BIT
#define _SC_WORD_BIT _SC_WORD_BIT
#define _SC_MB_LEN_MAX _SC_MB_LEN_MAX
#define _SC_NZERO _SC_NZERO
#define _SC_SSIZE_MAX _SC_SSIZE_MAX
#define _SC_SCHAR_MAX _SC_SCHAR_MAX
#define _SC_SCHAR_MIN _SC_SCHAR_MIN
#define _SC_SHRT_MAX _SC_SHRT_MAX
#define _SC_SHRT_MIN _SC_SHRT_MIN
#define _SC_UCHAR_MAX _SC_UCHAR_MAX
#define _SC_UINT_MAX _SC_UINT_MAX
#define _SC_ULONG_MAX _SC_ULONG_MAX
#define _SC_USHRT_MAX _SC_USHRT_MAX
#define _SC_NL_ARGMAX _SC_NL_ARGMAX
#define _SC_NL_LANGMAX _SC_NL_LANGMAX
#define _SC_NL_MSGMAX _SC_NL_MSGMAX
#define _SC_NL_NMAX _SC_NL_NMAX
#define _SC_NL_SETMAX _SC_NL_SETMAX
#define _SC_NL_TEXTMAX _SC_NL_TEXTMAX
#define _SC_XBS5_ILP32_OFF32 _SC_XBS5_ILP32_OFF32
#define _SC_XBS5_ILP32_OFFBIG _SC_XBS5_ILP32_OFFBIG
#define _SC_XBS5_LP64_OFF64 _SC_XBS5_LP64_OFF64
#define _SC_XBS5_LPBIG_OFFBIG _SC_XBS5_LPBIG_OFFBIG
#define _SC_XOPEN_LEGACY _SC_XOPEN_LEGACY
#define _SC_XOPEN_REALTIME _SC_XOPEN_REALTIME
#define _SC_XOPEN_REALTIME_THREADS _SC_XOPEN_REALTIME_THREADS
#define _SC_ADVISORY_INFO _SC_ADVISORY_INFO
#define _SC_BARRIERS _SC_BARRIERS
#define _SC_BASE _SC_BASE
#define _SC_C_LANG_SUPPORT _SC_C_LANG_SUPPORT
#define _SC_C_LANG_SUPPORT_R _SC_C_LANG_SUPPORT_R
#define _SC_CLOCK_SELECTION _SC_CLOCK_SELECTION
#define _SC_CPUTIME _SC_CPUTIME
#define _SC_THREAD_CPUTIME _SC_THREAD_CPUTIME
#define _SC_DEVICE_IO _SC_DEVICE_IO
#define _SC_DEVICE_SPECIFIC _SC_DEVICE_SPECIFIC
#define _SC_DEVICE_SPECIFIC_R _SC_DEVICE_SPECIFIC_R
#define _SC_FD_MGMT _SC_FD_MGMT
#define _SC_FIFO _SC_FIFO
#define _SC_PIPE _SC_PIPE
#define _SC_FILE_ATTRIBUTES _SC_FILE_ATTRIBUTES
#define _SC_FILE_LOCKING _SC_FILE_LOCKING
#define _SC_FILE_SYSTEM _SC_FILE_SYSTEM
#define _SC_MONOTONIC_CLOCK _SC_MONOTONIC_CLOCK
#define _SC_MULTI_PROCESS _SC_MULTI_PROCESS
#define _SC_SINGLE_PROCESS _SC_SINGLE_PROCESS
#define _SC_NETWORKING _SC_NETWORKING
#define _SC_READER_WRITER_LOCKS _SC_READER_WRITER_LOCKS
#define _SC_SPIN_LOCKS _SC_SPIN_LOCKS
#define _SC_REGEXP _SC_REGEXP
#define _SC_REGEX_VERSION _SC_REGEX_VERSION
#define _SC_SHELL _SC_SHELL
#define _SC_SIGNALS _SC_SIGNALS
#define _SC_SPAWN _SC_SPAWN
#define _SC_SPORADIC_SERVER _SC_SPORADIC_SERVER
#define _SC_THREAD_SPORADIC_SERVER _SC_THREAD_SPORADIC_SERVER
#define _SC_SYSTEM_DATABASE _SC_SYSTEM_DATABASE
#define _SC_SYSTEM_DATABASE_R _SC_SYSTEM_DATABASE_R
#define _SC_TIMEOUTS _SC_TIMEOUTS
#define _SC_TYPED_MEMORY_OBJECTS _SC_TYPED_MEMORY_OBJECTS
#define _SC_USER_GROUPS _SC_USER_GROUPS
#define _SC_USER_GROUPS_R _SC_USER_GROUPS_R
#define _SC_2_PBS _SC_2_PBS
#define _SC_2_PBS_ACCOUNTING _SC_2_PBS_ACCOUNTING
#define _SC_2_PBS_LOCATE _SC_2_PBS_LOCATE
#define _SC_2_PBS_MESSAGE _SC_2_PBS_MESSAGE
#define _SC_2_PBS_TRACK _SC_2_PBS_TRACK
#define _SC_SYMLOOP_MAX _SC_SYMLOOP_MAX
#define _SC_STREAMS _SC_STREAMS
#define _SC_2_PBS_CHECKPOINT _SC_2_PBS_CHECKPOINT
#define _SC_V6_ILP32_OFF32 _SC_V6_ILP32_OFF32
#define _SC_V6_ILP32_OFFBIG _SC_V6_ILP32_OFFBIG
#define _SC_V6_LP64_OFF64 _SC_V6_LP64_OFF64
#define _SC_V6_LPBIG_OFFBIG _SC_V6_LPBIG_OFFBIG
#define _SC_HOST_NAME_MAX _SC_HOST_NAME_MAX
#define _SC_TRACE _SC_TRACE
#define _SC_TRACE_EVENT_FILTER _SC_TRACE_EVENT_FILTER
#define _SC_TRACE_INHERIT _SC_TRACE_INHERIT
#define _SC_TRACE_LOG _SC_TRACE_LOG
#define _SC_LEVEL1_ICACHE_SIZE _SC_LEVEL1_ICACHE_SIZE
#define _SC_LEVEL1_ICACHE_ASSOC _SC_LEVEL1_ICACHE_ASSOC
#define _SC_LEVEL1_ICACHE_LINESIZE _SC_LEVEL1_ICACHE_LINESIZE
#define _SC_LEVEL1_DCACHE_SIZE _SC_LEVEL1_DCACHE_SIZE
#define _SC_LEVEL1_DCACHE_ASSOC _SC_LEVEL1_DCACHE_ASSOC
#define _SC_LEVEL1_DCACHE_LINESIZE _SC_LEVEL1_DCACHE_LINESIZE
#define _SC_LEVEL2_CACHE_SIZE _SC_LEVEL2_CACHE_SIZE
#define _SC_LEVEL2_CACHE_ASSOC _SC_LEVEL2_CACHE_ASSOC
#define _SC_LEVEL2_CACHE_LINESIZE _SC_LEVEL2_CACHE_LINESIZE
#define _SC_LEVEL3_CACHE_SIZE _SC_LEVEL3_CACHE_SIZE
#define _SC_LEVEL3_CACHE_ASSOC _SC_LEVEL3_CACHE_ASSOC
#define _SC_LEVEL3_CACHE_LINESIZE _SC_LEVEL3_CACHE_LINESIZE
#define _SC_LEVEL4_CACHE_SIZE _SC_LEVEL4_CACHE_SIZE
#define _SC_LEVEL4_CACHE_ASSOC _SC_LEVEL4_CACHE_ASSOC
#define _SC_LEVEL4_CACHE_LINESIZE _SC_LEVEL4_CACHE_LINESIZE
#define _SC_IPV6 _SC_IPV6
#define _SC_RAW_SOCKETS _SC_RAW_SOCKETS
#define _SC_V7_ILP32_OFF32 _SC_V7_ILP32_OFF32
#define _SC_V7_ILP32_OFFBIG _SC_V7_ILP32_OFFBIG
#define _SC_V7_LP64_OFF64 _SC_V7_LP64_OFF64
#define _SC_V7_LPBIG_OFFBIG _SC_V7_LPBIG_OFFBIG
#define _SC_SS_REPL_MAX _SC_SS_REPL_MAX
#define _SC_TRACE_EVENT_NAME_MAX _SC_TRACE_EVENT_NAME_MAX
#define _SC_TRACE_NAME_MAX _SC_TRACE_NAME_MAX
#define _SC_TRACE_SYS_MAX _SC_TRACE_SYS_MAX
#define _SC_TRACE_USER_EVENT_MAX _SC_TRACE_USER_EVENT_MAX
#define _SC_XOPEN_STREAMS _SC_XOPEN_STREAMS
#define _SC_THREAD_ROBUST_PRIO_INHERIT _SC_THREAD_ROBUST_PRIO_INHERIT
#define _SC_THREAD_ROBUST_PRIO_PROTECT _SC_THREAD_ROBUST_PRIO_PROTECT
#define _CS_PATH _CS_PATH
#define _CS_V6_WIDTH_RESTRICTED_ENVS _CS_V6_WIDTH_RESTRICTED_ENVS
#define _CS_POSIX_V6_WIDTH_RESTRICTED_ENVS _CS_V6_WIDTH_RESTRICTED_ENVS
#define _CS_GNU_LIBC_VERSION _CS_GNU_LIBC_VERSION
#define _CS_GNU_LIBPTHREAD_VERSION _CS_GNU_LIBPTHREAD_VERSION
#define _CS_V5_WIDTH_RESTRICTED_ENVS _CS_V5_WIDTH_RESTRICTED_ENVS
#define _CS_POSIX_V5_WIDTH_RESTRICTED_ENVS _CS_V5_WIDTH_RESTRICTED_ENVS
#define _CS_V7_WIDTH_RESTRICTED_ENVS _CS_V7_WIDTH_RESTRICTED_ENVS
#define _CS_POSIX_V7_WIDTH_RESTRICTED_ENVS _CS_V7_WIDTH_RESTRICTED_ENVS
#define _CS_LFS_CFLAGS _CS_LFS_CFLAGS
#define _CS_LFS_LDFLAGS _CS_LFS_LDFLAGS
#define _CS_LFS_LIBS _CS_LFS_LIBS
#define _CS_LFS_LINTFLAGS _CS_LFS_LINTFLAGS
#define _CS_LFS64_CFLAGS _CS_LFS64_CFLAGS
#define _CS_LFS64_LDFLAGS _CS_LFS64_LDFLAGS
#define _CS_LFS64_LIBS _CS_LFS64_LIBS
#define _CS_LFS64_LINTFLAGS _CS_LFS64_LINTFLAGS
#define _CS_XBS5_ILP32_OFF32_CFLAGS _CS_XBS5_ILP32_OFF32_CFLAGS
#define _CS_XBS5_ILP32_OFF32_LDFLAGS _CS_XBS5_ILP32_OFF32_LDFLAGS
#define _CS_XBS5_ILP32_OFF32_LIBS _CS_XBS5_ILP32_OFF32_LIBS
#define _CS_XBS5_ILP32_OFF32_LINTFLAGS _CS_XBS5_ILP32_OFF32_LINTFLAGS
#define _CS_XBS5_ILP32_OFFBIG_CFLAGS _CS_XBS5_ILP32_OFFBIG_CFLAGS
#define _CS_XBS5_ILP32_OFFBIG_LDFLAGS _CS_XBS5_ILP32_OFFBIG_LDFLAGS
#define _CS_XBS5_ILP32_OFFBIG_LIBS _CS_XBS5_ILP32_OFFBIG_LIBS
#define _CS_XBS5_ILP32_OFFBIG_LINTFLAGS _CS_XBS5_ILP32_OFFBIG_LINTFLAGS
#define _CS_XBS5_LP64_OFF64_CFLAGS _CS_XBS5_LP64_OFF64_CFLAGS
#define _CS_XBS5_LP64_OFF64_LDFLAGS _CS_XBS5_LP64_OFF64_LDFLAGS
#define _CS_XBS5_LP64_OFF64_LIBS _CS_XBS5_LP64_OFF64_LIBS
#define _CS_XBS5_LP64_OFF64_LINTFLAGS _CS_XBS5_LP64_OFF64_LINTFLAGS
#define _CS_XBS5_LPBIG_OFFBIG_CFLAGS _CS_XBS5_LPBIG_OFFBIG_CFLAGS
#define _CS_XBS5_LPBIG_OFFBIG_LDFLAGS _CS_XBS5_LPBIG_OFFBIG_LDFLAGS
#define _CS_XBS5_LPBIG_OFFBIG_LIBS _CS_XBS5_LPBIG_OFFBIG_LIBS
#define _CS_XBS5_LPBIG_OFFBIG_LINTFLAGS _CS_XBS5_LPBIG_OFFBIG_LINTFLAGS
#define _CS_POSIX_V6_ILP32_OFF32_CFLAGS _CS_POSIX_V6_ILP32_OFF32_CFLAGS
#define _CS_POSIX_V6_ILP32_OFF32_LDFLAGS _CS_POSIX_V6_ILP32_OFF32_LDFLAGS
#define _CS_POSIX_V6_ILP32_OFF32_LIBS _CS_POSIX_V6_ILP32_OFF32_LIBS
#define _CS_POSIX_V6_ILP32_OFF32_LINTFLAGS _CS_POSIX_V6_ILP32_OFF32_LINTFLAGS
#define _CS_POSIX_V6_ILP32_OFFBIG_CFLAGS _CS_POSIX_V6_ILP32_OFFBIG_CFLAGS
#define _CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS _CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS
#define _CS_POSIX_V6_ILP32_OFFBIG_LIBS _CS_POSIX_V6_ILP32_OFFBIG_LIBS
#define _CS_POSIX_V6_ILP32_OFFBIG_LINTFLAGS _CS_POSIX_V6_ILP32_OFFBIG_LINTFLAGS
#define _CS_POSIX_V6_LP64_OFF64_CFLAGS _CS_POSIX_V6_LP64_OFF64_CFLAGS
#define _CS_POSIX_V6_LP64_OFF64_LDFLAGS _CS_POSIX_V6_LP64_OFF64_LDFLAGS
#define _CS_POSIX_V6_LP64_OFF64_LIBS _CS_POSIX_V6_LP64_OFF64_LIBS
#define _CS_POSIX_V6_LP64_OFF64_LINTFLAGS _CS_POSIX_V6_LP64_OFF64_LINTFLAGS
#define _CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS _CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS
#define _CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS _CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS
#define _CS_POSIX_V6_LPBIG_OFFBIG_LIBS _CS_POSIX_V6_LPBIG_OFFBIG_LIBS
#define _CS_POSIX_V6_LPBIG_OFFBIG_LINTFLAGS _CS_POSIX_V6_LPBIG_OFFBIG_LINTFLAGS
#define _CS_POSIX_V7_ILP32_OFF32_CFLAGS _CS_POSIX_V7_ILP32_OFF32_CFLAGS
#define _CS_POSIX_V7_ILP32_OFF32_LDFLAGS _CS_POSIX_V7_ILP32_OFF32_LDFLAGS
#define _CS_POSIX_V7_ILP32_OFF32_LIBS _CS_POSIX_V7_ILP32_OFF32_LIBS
#define _CS_POSIX_V7_ILP32_OFF32_LINTFLAGS _CS_POSIX_V7_ILP32_OFF32_LINTFLAGS
#define _CS_POSIX_V7_ILP32_OFFBIG_CFLAGS _CS_POSIX_V7_ILP32_OFFBIG_CFLAGS
#define _CS_POSIX_V7_ILP32_OFFBIG_LDFLAGS _CS_POSIX_V7_ILP32_OFFBIG_LDFLAGS
#define _CS_POSIX_V7_ILP32_OFFBIG_LIBS _CS_POSIX_V7_ILP32_OFFBIG_LIBS
#define _CS_POSIX_V7_ILP32_OFFBIG_LINTFLAGS _CS_POSIX_V7_ILP32_OFFBIG_LINTFLAGS
#define _CS_POSIX_V7_LP64_OFF64_CFLAGS _CS_POSIX_V7_LP64_OFF64_CFLAGS
#define _CS_POSIX_V7_LP64_OFF64_LDFLAGS _CS_POSIX_V7_LP64_OFF64_LDFLAGS
#define _CS_POSIX_V7_LP64_OFF64_LIBS _CS_POSIX_V7_LP64_OFF64_LIBS
#define _CS_POSIX_V7_LP64_OFF64_LINTFLAGS _CS_POSIX_V7_LP64_OFF64_LINTFLAGS
#define _CS_POSIX_V7_LPBIG_OFFBIG_CFLAGS _CS_POSIX_V7_LPBIG_OFFBIG_CFLAGS
#define _CS_POSIX_V7_LPBIG_OFFBIG_LDFLAGS _CS_POSIX_V7_LPBIG_OFFBIG_LDFLAGS
#define _CS_POSIX_V7_LPBIG_OFFBIG_LIBS _CS_POSIX_V7_LPBIG_OFFBIG_LIBS
#define _CS_POSIX_V7_LPBIG_OFFBIG_LINTFLAGS _CS_POSIX_V7_LPBIG_OFFBIG_LINTFLAGS
#define _CS_V6_ENV _CS_V6_ENV
#define _CS_V7_ENV _CS_V7_ENV
