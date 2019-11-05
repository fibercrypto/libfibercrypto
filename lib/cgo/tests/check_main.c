#include <stdio.h>
#include <stdlib.h>

#include <check.h>
#include <string.h>

#include "libfibercrypto.h"
#include "skyassert.h"
#include "skyerrors.h"
#include "skystring.h"
#include "skytest.h"

START_TEST(TestRegisterSkycoinPlugin) { 
  
 }
END_TEST

Suite *check_main(void) {
  Suite *s = suite_create("Load check_main");
  TCase *tc;
  tc = tcase_create("check_main");
  tcase_add_checked_fixture(tc, setup, teardown);
  // tcase_add_test(tc, TestRegisterSkycoinPlugin);
  return s;
}
