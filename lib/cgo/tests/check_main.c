#include <stdio.h>
#include <stdlib.h>

#include <check.h>
#include <string.h>

#include "fcassert.h"
#include "fcerrors.h"
#include "fcstring.h"
#include "fctest.h"
#include "libfibercrypto.h"

START_TEST(TestRegisterSkycoinPlugin) {
  printf("Load TestRegisterSkycoinPlugin \n");

  GoString SkycoinTicker = {"SKY", 3};
  GoString CoinHoursTicker = {"SCH", 3};
  GoString CalculatedHoursTicker = {"SCH#ACC", 7};
  GoString_ Temp;
  FC_util_AltcoinCaption(SkycoinTicker, &Temp);
  ck_assert_str_eq(Temp.p, "Skycoin");
  FC_util_AltcoinCaption(CoinHoursTicker, &Temp);
  ck_assert_str_eq(Temp.p, "Coin Hours");
  FC_util_AltcoinCaption(CalculatedHoursTicker, &Temp);
  ck_assert_str_eq(Temp.p, "Calculated Hours");
}
END_TEST

Suite *check_main(void) {
  Suite *s = suite_create("Load check_main");
  TCase *tc;
  tc = tcase_create("check_main");
  tcase_add_checked_fixture(tc, setup, teardown);
  tcase_add_test(tc, TestRegisterSkycoinPlugin);
  suite_add_tcase(s, tc);
  return s;
}
