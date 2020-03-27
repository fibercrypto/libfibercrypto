#ifndef TEST_MAIN_H
#define TEST_MAIN_H
#include "libfibercrypto.h"
#include "fcerrors.h"
#include "fcstring.h"
#include "fctest.h"
#include "fctypes.h"
#include <check.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

Suite *check_util_cipher(void);
Suite *check_util_datetime(void);

#endif