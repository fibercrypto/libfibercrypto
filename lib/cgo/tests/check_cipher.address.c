#include <stdio.h>
#include <stdlib.h>

#include <check.h>
#include <string.h>

#include "libfibercrypto.h"
#include "skyassert.h"
#include "skyerrors.h"
#include "skystring.h"
#include "skytest.h"


START_TEST(TestAddressRoundtrip) { 
    ck_assert(1 == 1); 
}
END_TEST

Suite *check_cipher_address(void) {
  Suite *s = suite_create("Load check_cipher_address.address");
  TCase *tc;

  tc = tcase_create("check_cipher.address");
//   tcase_add_checked_fixture(tc, setup, teardown);
  tcase_add_test(tc, TestAddressRoundtrip);

  return s;
}
