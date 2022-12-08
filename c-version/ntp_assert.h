#ifndef _NTP_ASSERT_H_
#define _NTP_ASSERT_H_

//
#if !defined(ntp_type_full)
  #define ntp_type_full(prefix, suffix, T, fmt, a, op, b, format, ...)                                                                          \
    ntp_type_full_f(stderr, prefix, suffix, T, fmt, a, op, b, ntp_assert_f, format, ##__VA_ARGS__)
#endif  // ntp_type_full
//
#if !defined(ntp_ntp)
  #define ntp_ntp(expr, ...) ntp_ntp_f(stderr, expr, ntp_assert_f, ##__VA_ARGS__)
#endif  // ntp_ntp

#include "ntp.h"
//
inline void ntp_assert_f(FILE* stream, const char* a, const char* op, const char* b, bool result, const char a_value[A_VALUE_STRING],
                         const char b_value[B_VALUE_STRING], time_t time, const char* file, const char* function, int line, const char* format,
                         ...) {
  va_list args;
  va_start(args, format);
  fprintf(stream, result ? TRUE_RESULT_COLOR : FALSE_RESULT_COLOR);
  fprintf(stream, "%s:%d :: %s :: (%s %s %s) :: (%s %s %s) :: ", file, line, function, a, op, b, a_value, op, b_value);
  vfprintf(stream, format, args);
  fprintf(stream, "\n");
  fprintf(stream, NO_COLOR);
  fflush(stream);
  va_end(args);
}
extern inline void ntp_assert_f(FILE* stream, const char* a, const char* op, const char* b, bool result, const char a_value[A_VALUE_STRING],
                                const char b_value[B_VALUE_STRING], time_t time, const char* file, const char* function, int line,
                                const char* format, ...);
#endif  // _NTP_ASSERT_H_
