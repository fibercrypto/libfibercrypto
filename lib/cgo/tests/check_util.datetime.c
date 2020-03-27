#include <stdio.h>
#include <stdlib.h>

#include <check.h>
#include <string.h>
#include <time.h>

#include "fcassert.h"
#include "fcerrors.h"
#include "fcstring.h"
#include "fctest.h"
#include "libfibercrypto.h"

START_TEST(TestParseDate)
{
  printf("Load TestParseDate \n");
  GoUint64_ some_date = 1577665665;
}
END_TEST

Suite *check_util_datetime(void)
{
  Suite *s = suite_create("Load check_util_datetime");
  TCase *tc;
  tc = tcase_create("check_util_datetime");
  tcase_add_checked_fixture(tc, setup, teardown);
  tcase_add_test(tc, TestParseDate);
  suite_add_tcase(s, tc);
  return s;
}
