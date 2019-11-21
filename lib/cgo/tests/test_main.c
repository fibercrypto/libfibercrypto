#include "test_main.h"
#include <check.h>
// run suite
int main(void) {
  int number_failed = 0;
  SRunner *sr = srunner_create(check_main());
  srunner_set_fork_status(sr, CK_NOFORK);
  srunner_run_all(sr, CK_VERBOSE);
  number_failed = srunner_ntests_failed(sr);
  srunner_free(sr);
  sr = NULL;
  return (number_failed == 0) ? EXIT_SUCCESS : EXIT_FAILURE;
}
