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

START_TEST(TestNewGenericOutput) {
  printf("Load TestNewGenericOutput \n");
  core__AltcoinMetadata fakeMeta;
}
END_TEST

Suite *check_util_coin(void) {
  Suite *s = suite_create("Load check_util_coin");
  TCase *tc;
  tc = tcase_create("check_util_coin");
  tcase_add_checked_fixture(tc, setup, teardown);
  tcase_add_test(tc, TestNewGenericOutput);
  suite_add_tcase(s, tc);
  return s;
}
