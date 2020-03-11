#include <stdio.h>
#include <stdlib.h>

#include <check.h>
#include <string.h>

#include "fcassert.h"
#include "fcerrors.h"
#include "fcstring.h"
#include "fctest.h"
#include "libfibercrypto.h"

typedef struct
{
  unsigned char *name;
  GoString args;
  util__GenericAddress want;
} test_case;

START_TEST(TestNewGenericAddress)
{
  printf("Load TestNewGenericAddress \n");

  test_case tests[3];

  tests[0].name = "valid_Addrs";
  tests[0].args.p = "2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt";
  tests[0].args.n = 35;
  tests[0].want.Address.p = "2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt";
  tests[0].want.Address.n = 35;

  for (int i = 0; i < 1; i++)
  {
    printf("Load %d\n", i);
    test_case tt = tests[i];
    util__GenericAddress got;
    GoUint32_ err = FC_util_NewGenericAddress(tt.args, &got);
    ck_assert_int_eq(err, FC_OK);
    ck_assert_msg(isutil__GenericAddress_Eq(got, tt.want), "Fail isutil__GenericAddress_Eq");

    GoUint8 buffer_got_slice[1024];
    GoSlice_ got_slice = {buffer_got_slice, 0, 1024};
    GoUint8 buffer_args_slice[1024];
    GoSlice_ args_slice = {buffer_args_slice, 0, 1024};
    err = FC_util_GenericAddress_Bytes(&got, &got_slice);
    ck_assert_int_eq(err, FC_OK);
    FC_StringtoByte(tt.args, &args_slice);
    ck_assert(isGoSlice_Eq(&got_slice, &args_slice));

    GoUint8 buffer_got_string[1024];
    GoString_ got_string = {&buffer_got_string, 0};
    err = FC_util_GenericAddress_String(&got, &got_string);
    ck_assert_int_eq(err, FC_OK);
    ck_assert_str_eq(got_string.p, tt.args.p);

    GoUint8_ got_isbip32;
    err = FC_util_GenericAddress_IsBip32(&got, &got_isbip32);
    ck_assert_int_eq(FC_OK, err);
    ck_assert_int_eq(got_isbip32, 0);

    GoUint8 got_null;
    err = FC_util_GenericAddress_Null(&got, &got_null);
    ck_assert_int_eq(err, FC_OK);
    if (strcmp(tt.name, "invalid_Addrs"))
    {
      ck_assert_int_eq(got_null, 0);
    }
    else
    {
      ck_assert_int_eq(got_null, 1);
    }
  }

  ck_assert_int_eq(0, FC_OK);
}
END_TEST

Suite *check_util_cipher(void)
{
  Suite *s = suite_create("Load check_util_cipher");
  TCase *tc;
  tc = tcase_create("check_util_cipher");
  tcase_add_checked_fixture(tc, setup, teardown);
  tcase_add_test(tc, TestNewGenericAddress);
  suite_add_tcase(s, tc);
  return s;
}
