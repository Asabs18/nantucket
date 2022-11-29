#include <ctype.h>
#include <inttypes.h>  // stdint formats: PRIinn, e.g., PRIi64
#include <math.h>      // M_PI
#include <stdarg.h>    // va_list()
#include <stdbool.h>
#include <string.h>  // strcmp()

#include "ntp.h"
#include "ntp_assert.h"

int main(int argc, char* argv[]) {
  //
  const int zero = 0;
  int one        = 1;
  const char a   = 'A';
  char z         = 'Z';
  ntp_ntp(false);
  ntp_true(!false);
  ntp_false(!true);
  //
  ntp_type_full_f(stderr, "", "", char, "c", 'a', ==, 'z', ntp_default_f, "");
  ntp_type_full_f(stderr, "", "", char, "c", 'a', !=, tolower(a), ntp_default_f, "not equal");
  ntp_type_full_f(stderr, "", "", char, "c", 'a', <, z, ntp_default_f, "less than %c", z);
  //
  // TODO - ntp_default_string_f()?
  const char string[] = "djg - pl!";
  ntp_type_full_f(stderr, "", "", int, "d", strcmp(string, "pl"), ==, 0, ntp_default_f, "less than %s", string);
  //
  ntp_type_full("", "", int, "d", 1, ==, 1, "format, ...", 1);
  ntp_type_full("", "", int, "d", 1, ==, 0, NULL);
  ntp_type_full("", "", int, "d", 1, ==, 0, "");
  ntp_type_full("", "", int, "d", 1, !=, one, "not equal");
  ntp_type_full("", "", int, "d", 1, <, zero, "less than %d", zero);
  //
  ntp_char('a', ==, 'z', "char", "");
  ntp_char('a', !=, tolower(a), "not equal");
  ntp_char('a', <, z, "less than %c", z);
  //
  ntp_int(1, ==, 0, "%d", __COUNTER__);
  ntp_int(1, !=, one, "not equal");
  ntp_int(1, <, zero, "less than %d", zero);
  //
  ntp_char('a', ==, 'z', "%d", __COUNTER__);
  ntp_char('a', !=, tolower(a), "not equal");
  ntp_char('a', <, z, "less than %c", z);
  //
  ntp_llong((long long) 1, ==, 0, "%d", __COUNTER__);
  ntp_llong((long long) 1, !=, one, "not equal");
  ntp_llong(1ll, <, zero, "less than %d", zero);
  //
  size_t size = 0xdeadbeef;
  ntp_size(size, !=, 0xdeadbeef, "%d", __COUNTER__);
  //
  ntp_double(M_PI, <, 3.14, "%d", __COUNTER__);
  ntp_ptr(NULL, !=, (void*) (&zero - &zero), "%d", __COUNTER__);
  //
  ntp_int64(INT64_MAX, <, 0, "%d", __COUNTER__);

  ntp_string_equal("pl", "djg", NULL);
  ntp_string_not_equal("pl", "djg", ntp_default_f);

  return 0;
}
#ifdef EVENTS
typedef struct {
  va_list ap;
  const char* fmt;
  const char* a;
  const char a_value[A_VALUE_STRING];
  const char* op;
  const char* b;
  const char b_value[B_VALUE_STRING];
  const char* file;
  const char* function;
  int line;
  struct tm* time;
  void* udata;
} ntp_event_t;
#endif  // EVENTS
