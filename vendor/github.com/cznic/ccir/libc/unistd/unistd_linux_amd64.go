// Code generated by running "go generate". DO NOT EDIT.

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

package unistd

const (
	XF_LOCK                              = 1
	XF_OK                                = 0
	XF_TEST                              = 3
	XF_TLOCK                             = 2
	XF_ULOCK                             = 0
	XR_OK                                = 4
	XSEEK_CUR                            = 1
	XSEEK_END                            = 2
	XSEEK_SET                            = 0
	XSTDERR_FILENO                       = 2
	XSTDIN_FILENO                        = 0
	XSTDOUT_FILENO                       = 1
	XW_OK                                = 2
	XX_OK                                = 1
	X_BITS_TYPES_H                       = 1
	X_POSIX2_C_BIND                      = 199506
	X_POSIX2_C_DEV                       = 199506
	X_POSIX2_C_VERSION                   = 199506
	X_POSIX2_LOCALEDEF                   = 199506
	X_POSIX2_SW_DEV                      = 199506
	X_POSIX2_VERSION                     = 199506
	X_POSIX_VERSION                      = 199506
	X_UNISTD_H                           = 1
	X_XOPEN_CRYPT                        = 1
	X_XOPEN_ENH_I18N                     = 1
	X_XOPEN_LEGACY                       = 1
	X_XOPEN_UNIX                         = 1
	X_XOPEN_VERSION                      = 500
	X_XOPEN_XCU_VERSION                  = 4
	X_XOPEN_XPG2                         = 1
	X_XOPEN_XPG3                         = 1
	X_XOPEN_XPG4                         = 1
	X__POSIX2_THIS_VERSION               = 199506
	X__S32_TYPE                          = 0
	X__SLONG32_TYPE                      = 0
	Xftruncate                           = 0
	Xlockf                               = 0
	Xlseek                               = 0
	Xpread                               = 0
	Xpwrite                              = 0
	Xtruncate                            = 0
	X_CS_GNU_LIBC_VERSION                = 2
	X_CS_GNU_LIBPTHREAD_VERSION          = 3
	X_CS_LFS64_CFLAGS                    = 1004
	X_CS_LFS64_LDFLAGS                   = 1005
	X_CS_LFS64_LIBS                      = 1006
	X_CS_LFS64_LINTFLAGS                 = 1007
	X_CS_LFS_CFLAGS                      = 1000
	X_CS_LFS_LDFLAGS                     = 1001
	X_CS_LFS_LIBS                        = 1002
	X_CS_LFS_LINTFLAGS                   = 1003
	X_CS_PATH                            = 0
	X_CS_POSIX_V6_ILP32_OFF32_CFLAGS     = 1116
	X_CS_POSIX_V6_ILP32_OFF32_LDFLAGS    = 1117
	X_CS_POSIX_V6_ILP32_OFF32_LIBS       = 1118
	X_CS_POSIX_V6_ILP32_OFF32_LINTFLAGS  = 1119
	X_CS_POSIX_V6_ILP32_OFFBIG_CFLAGS    = 1120
	X_CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS   = 1121
	X_CS_POSIX_V6_ILP32_OFFBIG_LIBS      = 1122
	X_CS_POSIX_V6_ILP32_OFFBIG_LINTFLAGS = 1123
	X_CS_POSIX_V6_LP64_OFF64_CFLAGS      = 1124
	X_CS_POSIX_V6_LP64_OFF64_LDFLAGS     = 1125
	X_CS_POSIX_V6_LP64_OFF64_LIBS        = 1126
	X_CS_POSIX_V6_LP64_OFF64_LINTFLAGS   = 1127
	X_CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS    = 1128
	X_CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS   = 1129
	X_CS_POSIX_V6_LPBIG_OFFBIG_LIBS      = 1130
	X_CS_POSIX_V6_LPBIG_OFFBIG_LINTFLAGS = 1131
	X_CS_POSIX_V7_ILP32_OFF32_CFLAGS     = 1132
	X_CS_POSIX_V7_ILP32_OFF32_LDFLAGS    = 1133
	X_CS_POSIX_V7_ILP32_OFF32_LIBS       = 1134
	X_CS_POSIX_V7_ILP32_OFF32_LINTFLAGS  = 1135
	X_CS_POSIX_V7_ILP32_OFFBIG_CFLAGS    = 1136
	X_CS_POSIX_V7_ILP32_OFFBIG_LDFLAGS   = 1137
	X_CS_POSIX_V7_ILP32_OFFBIG_LIBS      = 1138
	X_CS_POSIX_V7_ILP32_OFFBIG_LINTFLAGS = 1139
	X_CS_POSIX_V7_LP64_OFF64_CFLAGS      = 1140
	X_CS_POSIX_V7_LP64_OFF64_LDFLAGS     = 1141
	X_CS_POSIX_V7_LP64_OFF64_LIBS        = 1142
	X_CS_POSIX_V7_LP64_OFF64_LINTFLAGS   = 1143
	X_CS_POSIX_V7_LPBIG_OFFBIG_CFLAGS    = 1144
	X_CS_POSIX_V7_LPBIG_OFFBIG_LDFLAGS   = 1145
	X_CS_POSIX_V7_LPBIG_OFFBIG_LIBS      = 1146
	X_CS_POSIX_V7_LPBIG_OFFBIG_LINTFLAGS = 1147
	X_CS_V5_WIDTH_RESTRICTED_ENVS        = 4
	X_CS_V6_ENV                          = 1148
	X_CS_V6_WIDTH_RESTRICTED_ENVS        = 1
	X_CS_V7_ENV                          = 1149
	X_CS_V7_WIDTH_RESTRICTED_ENVS        = 5
	X_CS_XBS5_ILP32_OFF32_CFLAGS         = 1100
	X_CS_XBS5_ILP32_OFF32_LDFLAGS        = 1101
	X_CS_XBS5_ILP32_OFF32_LIBS           = 1102
	X_CS_XBS5_ILP32_OFF32_LINTFLAGS      = 1103
	X_CS_XBS5_ILP32_OFFBIG_CFLAGS        = 1104
	X_CS_XBS5_ILP32_OFFBIG_LDFLAGS       = 1105
	X_CS_XBS5_ILP32_OFFBIG_LIBS          = 1106
	X_CS_XBS5_ILP32_OFFBIG_LINTFLAGS     = 1107
	X_CS_XBS5_LP64_OFF64_CFLAGS          = 1108
	X_CS_XBS5_LP64_OFF64_LDFLAGS         = 1109
	X_CS_XBS5_LP64_OFF64_LIBS            = 1110
	X_CS_XBS5_LP64_OFF64_LINTFLAGS       = 1111
	X_CS_XBS5_LPBIG_OFFBIG_CFLAGS        = 1112
	X_CS_XBS5_LPBIG_OFFBIG_LDFLAGS       = 1113
	X_CS_XBS5_LPBIG_OFFBIG_LIBS          = 1114
	X_CS_XBS5_LPBIG_OFFBIG_LINTFLAGS     = 1115
	X_PC_2_SYMLINKS                      = 20
	X_PC_ALLOC_SIZE_MIN                  = 18
	X_PC_ASYNC_IO                        = 10
	X_PC_CHOWN_RESTRICTED                = 6
	X_PC_FILESIZEBITS                    = 13
	X_PC_LINK_MAX                        = 0
	X_PC_MAX_CANON                       = 1
	X_PC_MAX_INPUT                       = 2
	X_PC_NAME_MAX                        = 3
	X_PC_NO_TRUNC                        = 7
	X_PC_PATH_MAX                        = 4
	X_PC_PIPE_BUF                        = 5
	X_PC_PRIO_IO                         = 11
	X_PC_REC_INCR_XFER_SIZE              = 14
	X_PC_REC_MAX_XFER_SIZE               = 15
	X_PC_REC_MIN_XFER_SIZE               = 16
	X_PC_REC_XFER_ALIGN                  = 17
	X_PC_SOCK_MAXBUF                     = 12
	X_PC_SYMLINK_MAX                     = 19
	X_PC_SYNC_IO                         = 9
	X_PC_VDISABLE                        = 8
	X_SC_2_CHAR_TERM                     = 95
	X_SC_2_C_BIND                        = 47
	X_SC_2_C_DEV                         = 48
	X_SC_2_C_VERSION                     = 96
	X_SC_2_FORT_DEV                      = 49
	X_SC_2_FORT_RUN                      = 50
	X_SC_2_LOCALEDEF                     = 52
	X_SC_2_PBS                           = 168
	X_SC_2_PBS_ACCOUNTING                = 169
	X_SC_2_PBS_CHECKPOINT                = 175
	X_SC_2_PBS_LOCATE                    = 170
	X_SC_2_PBS_MESSAGE                   = 171
	X_SC_2_PBS_TRACK                     = 172
	X_SC_2_SW_DEV                        = 51
	X_SC_2_UPE                           = 97
	X_SC_2_VERSION                       = 46
	X_SC_ADVISORY_INFO                   = 132
	X_SC_AIO_LISTIO_MAX                  = 23
	X_SC_AIO_MAX                         = 24
	X_SC_AIO_PRIO_DELTA_MAX              = 25
	X_SC_ARG_MAX                         = 0
	X_SC_ASYNCHRONOUS_IO                 = 12
	X_SC_ATEXIT_MAX                      = 87
	X_SC_AVPHYS_PAGES                    = 86
	X_SC_BARRIERS                        = 133
	X_SC_BASE                            = 134
	X_SC_BC_BASE_MAX                     = 36
	X_SC_BC_DIM_MAX                      = 37
	X_SC_BC_SCALE_MAX                    = 38
	X_SC_BC_STRING_MAX                   = 39
	X_SC_CHARCLASS_NAME_MAX              = 45
	X_SC_CHAR_BIT                        = 101
	X_SC_CHAR_MAX                        = 102
	X_SC_CHAR_MIN                        = 103
	X_SC_CHILD_MAX                       = 1
	X_SC_CLK_TCK                         = 2
	X_SC_CLOCK_SELECTION                 = 137
	X_SC_COLL_WEIGHTS_MAX                = 40
	X_SC_CPUTIME                         = 138
	X_SC_C_LANG_SUPPORT                  = 135
	X_SC_C_LANG_SUPPORT_R                = 136
	X_SC_DELAYTIMER_MAX                  = 26
	X_SC_DEVICE_IO                       = 140
	X_SC_DEVICE_SPECIFIC                 = 141
	X_SC_DEVICE_SPECIFIC_R               = 142
	X_SC_EQUIV_CLASS_MAX                 = 41
	X_SC_EXPR_NEST_MAX                   = 42
	X_SC_FD_MGMT                         = 143
	X_SC_FIFO                            = 144
	X_SC_FILE_ATTRIBUTES                 = 146
	X_SC_FILE_LOCKING                    = 147
	X_SC_FILE_SYSTEM                     = 148
	X_SC_FSYNC                           = 15
	X_SC_GETGR_R_SIZE_MAX                = 69
	X_SC_GETPW_R_SIZE_MAX                = 70
	X_SC_HOST_NAME_MAX                   = 180
	X_SC_INT_MAX                         = 104
	X_SC_INT_MIN                         = 105
	X_SC_IOV_MAX                         = 60
	X_SC_IPV6                            = 235
	X_SC_JOB_CONTROL                     = 7
	X_SC_LEVEL1_DCACHE_ASSOC             = 189
	X_SC_LEVEL1_DCACHE_LINESIZE          = 190
	X_SC_LEVEL1_DCACHE_SIZE              = 188
	X_SC_LEVEL1_ICACHE_ASSOC             = 186
	X_SC_LEVEL1_ICACHE_LINESIZE          = 187
	X_SC_LEVEL1_ICACHE_SIZE              = 185
	X_SC_LEVEL2_CACHE_ASSOC              = 192
	X_SC_LEVEL2_CACHE_LINESIZE           = 193
	X_SC_LEVEL2_CACHE_SIZE               = 191
	X_SC_LEVEL3_CACHE_ASSOC              = 195
	X_SC_LEVEL3_CACHE_LINESIZE           = 196
	X_SC_LEVEL3_CACHE_SIZE               = 194
	X_SC_LEVEL4_CACHE_ASSOC              = 198
	X_SC_LEVEL4_CACHE_LINESIZE           = 199
	X_SC_LEVEL4_CACHE_SIZE               = 197
	X_SC_LINE_MAX                        = 43
	X_SC_LOGIN_NAME_MAX                  = 71
	X_SC_LONG_BIT                        = 106
	X_SC_MAPPED_FILES                    = 16
	X_SC_MB_LEN_MAX                      = 108
	X_SC_MEMLOCK                         = 17
	X_SC_MEMLOCK_RANGE                   = 18
	X_SC_MEMORY_PROTECTION               = 19
	X_SC_MESSAGE_PASSING                 = 20
	X_SC_MONOTONIC_CLOCK                 = 149
	X_SC_MQ_OPEN_MAX                     = 27
	X_SC_MQ_PRIO_MAX                     = 28
	X_SC_MULTI_PROCESS                   = 150
	X_SC_NETWORKING                      = 152
	X_SC_NGROUPS_MAX                     = 3
	X_SC_NL_ARGMAX                       = 119
	X_SC_NL_LANGMAX                      = 120
	X_SC_NL_MSGMAX                       = 121
	X_SC_NL_NMAX                         = 122
	X_SC_NL_SETMAX                       = 123
	X_SC_NL_TEXTMAX                      = 124
	X_SC_NPROCESSORS_CONF                = 83
	X_SC_NPROCESSORS_ONLN                = 84
	X_SC_NZERO                           = 109
	X_SC_OPEN_MAX                        = 4
	X_SC_PAGESIZE                        = 30
	X_SC_PASS_MAX                        = 88
	X_SC_PHYS_PAGES                      = 85
	X_SC_PII                             = 53
	X_SC_PII_INTERNET                    = 56
	X_SC_PII_INTERNET_DGRAM              = 62
	X_SC_PII_INTERNET_STREAM             = 61
	X_SC_PII_OSI                         = 57
	X_SC_PII_OSI_CLTS                    = 64
	X_SC_PII_OSI_COTS                    = 63
	X_SC_PII_OSI_M                       = 65
	X_SC_PII_SOCKET                      = 55
	X_SC_PII_XTI                         = 54
	X_SC_PIPE                            = 145
	X_SC_POLL                            = 58
	X_SC_PRIORITIZED_IO                  = 13
	X_SC_PRIORITY_SCHEDULING             = 10
	X_SC_RAW_SOCKETS                     = 236
	X_SC_READER_WRITER_LOCKS             = 153
	X_SC_REALTIME_SIGNALS                = 9
	X_SC_REGEXP                          = 155
	X_SC_REGEX_VERSION                   = 156
	X_SC_RE_DUP_MAX                      = 44
	X_SC_RTSIG_MAX                       = 31
	X_SC_SAVED_IDS                       = 8
	X_SC_SCHAR_MAX                       = 111
	X_SC_SCHAR_MIN                       = 112
	X_SC_SELECT                          = 59
	X_SC_SEMAPHORES                      = 21
	X_SC_SEM_NSEMS_MAX                   = 32
	X_SC_SEM_VALUE_MAX                   = 33
	X_SC_SHARED_MEMORY_OBJECTS           = 22
	X_SC_SHELL                           = 157
	X_SC_SHRT_MAX                        = 113
	X_SC_SHRT_MIN                        = 114
	X_SC_SIGNALS                         = 158
	X_SC_SIGQUEUE_MAX                    = 34
	X_SC_SINGLE_PROCESS                  = 151
	X_SC_SPAWN                           = 159
	X_SC_SPIN_LOCKS                      = 154
	X_SC_SPORADIC_SERVER                 = 160
	X_SC_SSIZE_MAX                       = 110
	X_SC_SS_REPL_MAX                     = 241
	X_SC_STREAMS                         = 174
	X_SC_STREAM_MAX                      = 5
	X_SC_SYMLOOP_MAX                     = 173
	X_SC_SYNCHRONIZED_IO                 = 14
	X_SC_SYSTEM_DATABASE                 = 162
	X_SC_SYSTEM_DATABASE_R               = 163
	X_SC_THREADS                         = 67
	X_SC_THREAD_ATTR_STACKADDR           = 77
	X_SC_THREAD_ATTR_STACKSIZE           = 78
	X_SC_THREAD_CPUTIME                  = 139
	X_SC_THREAD_DESTRUCTOR_ITERATIONS    = 73
	X_SC_THREAD_KEYS_MAX                 = 74
	X_SC_THREAD_PRIORITY_SCHEDULING      = 79
	X_SC_THREAD_PRIO_INHERIT             = 80
	X_SC_THREAD_PRIO_PROTECT             = 81
	X_SC_THREAD_PROCESS_SHARED           = 82
	X_SC_THREAD_ROBUST_PRIO_INHERIT      = 247
	X_SC_THREAD_ROBUST_PRIO_PROTECT      = 248
	X_SC_THREAD_SAFE_FUNCTIONS           = 68
	X_SC_THREAD_SPORADIC_SERVER          = 161
	X_SC_THREAD_STACK_MIN                = 75
	X_SC_THREAD_THREADS_MAX              = 76
	X_SC_TIMEOUTS                        = 164
	X_SC_TIMERS                          = 11
	X_SC_TIMER_MAX                       = 35
	X_SC_TRACE                           = 181
	X_SC_TRACE_EVENT_FILTER              = 182
	X_SC_TRACE_EVENT_NAME_MAX            = 242
	X_SC_TRACE_INHERIT                   = 183
	X_SC_TRACE_LOG                       = 184
	X_SC_TRACE_NAME_MAX                  = 243
	X_SC_TRACE_SYS_MAX                   = 244
	X_SC_TRACE_USER_EVENT_MAX            = 245
	X_SC_TTY_NAME_MAX                    = 72
	X_SC_TYPED_MEMORY_OBJECTS            = 165
	X_SC_TZNAME_MAX                      = 6
	X_SC_T_IOV_MAX                       = 66
	X_SC_UCHAR_MAX                       = 115
	X_SC_UINT_MAX                        = 116
	X_SC_UIO_MAXIOV                      = 60
	X_SC_ULONG_MAX                       = 117
	X_SC_USER_GROUPS                     = 166
	X_SC_USER_GROUPS_R                   = 167
	X_SC_USHRT_MAX                       = 118
	X_SC_V6_ILP32_OFF32                  = 176
	X_SC_V6_ILP32_OFFBIG                 = 177
	X_SC_V6_LP64_OFF64                   = 178
	X_SC_V6_LPBIG_OFFBIG                 = 179
	X_SC_V7_ILP32_OFF32                  = 237
	X_SC_V7_ILP32_OFFBIG                 = 238
	X_SC_V7_LP64_OFF64                   = 239
	X_SC_V7_LPBIG_OFFBIG                 = 240
	X_SC_VERSION                         = 29
	X_SC_WORD_BIT                        = 107
	X_SC_XBS5_ILP32_OFF32                = 125
	X_SC_XBS5_ILP32_OFFBIG               = 126
	X_SC_XBS5_LP64_OFF64                 = 127
	X_SC_XBS5_LPBIG_OFFBIG               = 128
	X_SC_XOPEN_CRYPT                     = 92
	X_SC_XOPEN_ENH_I18N                  = 93
	X_SC_XOPEN_LEGACY                    = 129
	X_SC_XOPEN_REALTIME                  = 130
	X_SC_XOPEN_REALTIME_THREADS          = 131
	X_SC_XOPEN_SHM                       = 94
	X_SC_XOPEN_STREAMS                   = 246
	X_SC_XOPEN_UNIX                      = 91
	X_SC_XOPEN_VERSION                   = 89
	X_SC_XOPEN_XCU_VERSION               = 90
	X_SC_XOPEN_XPG2                      = 98
	X_SC_XOPEN_XPG3                      = 99
	X_SC_XOPEN_XPG4                      = 100
)
