#include <assert.h>
#include <ctype.h>
#include <inttypes.h>  // stdint formats: PRIinn, e.g., PRIi64
#include <math.h>      // M_PI
#include <stdarg.h>    // va_list()
#include <stdbool.h>
#include <stdlib.h>  // free()
#include <string.h>  // strcmp()

// #include "ntp.h"
#include "ntp_assert.h"

#define ntp_test_assert(p, s)                                                                                                                   \
  {                                                                                                                                             \
    const char* _ntp_ps = strrchr(p, '(');                                                                                                      \
    const char* _ntp_pe = strrchr(p, ')');                                                                                                      \
    assert(0 == strncmp(s, _ntp_ps, (size_t) (_ntp_pe - _ntp_ps)));                                                                             \
  }

int main(int argc, char* argv[]) {
  size_t ptr_size = 0xdeadbeef;
  char* ptr       = NULL;
  FILE* file      = open_memstream(&ptr, &ptr_size);
  ntp_type_full_f(file, "", "", char, "c", 'a', <, 'z', ntp_default_f, "%c is less than %c", 'a', 'z');
  ntp_ntp_f(file, false, ntp_default_f, "ntp_ntp(false)");

  const char* _ntp_pe = strrchr(ptr, ')');
  size_t foo = strlen("ntp_ntp(false)");
  const char* _ntp_ps = _ntp_pe - foo;
  assert(0 == strncmp("ntp_ntp(false)", _ntp_ps, (size_t) (_ntp_pe - _ntp_ps)));

  // ntp_test_assert(ptr, "(a is less than z)");
  // ntp_ntp(false, "ntp_ntp(false)");
  // ntp_ntp_f(file, false, ntp_default_f, "ntp_ntp(false)");
  // ntp_test_assert(ptr, "ntp_ntp(false)");
  // ntp_type_full_f(file, "", "", char, "c", 'z', <, 'a', ntp_default_f, "%c is less than %c", 'z', 'a');
  // ntp_test_assert(ptr, "(z is less than a)");
  fclose(file);
  free(ptr);
  return EXIT_SUCCESS;
  //
  const int zero = 0;
  int one        = 1;
  const char a   = 'A';
  char z         = 'Z';
  fflush(file);
  ntp_ntp(false, "ntp_ntp(false)");
  // ntp_test_assert(ptr, "(ntp_ntp(false))");
  ntp_true(!false, "");
  ntp_false(!true, "");
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
  ntp_string_not_equal("pl", "djg", "");

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
