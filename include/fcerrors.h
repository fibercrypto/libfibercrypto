#include <signal.h>

#if __APPLE__
#include "TargetConditionals.h"
#endif

#if __linux__
#define FC_ABORT SIGABRT
#elif __APPLE__
#if TARGET_OS_MAC
#define FC_ABORT 2
#endif
#endif

#ifndef SKY_ERRORS_H
#define SKY_ERRORS_H

// Generic error conditions
#define FC_OK 0
#define FC_ERROR 0x7FFFFFFF
#define FC_BAD_HANDLE 0x7F000001
#define FC_INVALID_TIMESTRING 0x7F000002
#endif
