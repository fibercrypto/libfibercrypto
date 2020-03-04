#include <stdio.h>
#include <stdlib.h>

#include <check.h>
#include <string.h>

#include "fcassert.h"
#include "fcerrors.h"
#include "fcstring.h"
#include "fctest.h"
#include "libfibercrypto.h"

START_TEST(TestRegisterSkycoinPlugin)
{
  printf("Load TestRegisterSkycoinPlugin \n");

  ck_assert_int_eq(0, 0);
}
END_TEST

Suite *check_main(void)
{
  Suite *s = suite_create("Load check_main");
  TCase *tc;
  tc = tcase_create("check_main");
  tcase_add_checked_fixture(tc, setup, teardown);
  tcase_add_test(tc, TestRegisterSkycoinPlugin);
  suite_add_tcase(s, tc);
  return s;
}
