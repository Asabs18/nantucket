#ifndef _NTP_H_
#define _NTP_H_

#define NTP_DEBUG

#include <stdio.h>  // FILE*, snprintf()
#include <time.h>   // time_t
// Expression based formatted print macro & function
// [Expression-Oriented Programming in C: The FMT Macro - DEV Community](https://tinyl.io/7SFv)
#define FMT(buffer, fmt, ...) fmt_str(sizeof(buffer), (buffer), (fmt), __VA_ARGS__)
inline static char* fmt_str(size_t len, char buffer[/*restrict*/ static len], const char fmt[/*restrict*/], ...) {
  va_list ap;
  va_start(ap, fmt);
  vsnprintf(buffer, len, fmt, ap);
  va_end(ap);

  return buffer;
}
// Terminal Support
#define ESC "\x1b"
#define CSI ESC "["
// Styles
#define BOLD      CSI "1m"
#define ITALIC    CSI "3m"
#define NO_ITALIC CSI "23m"
// Colors
#define NO_COLOR     ESC "[0m"
#define RED          ESC "[31m"
#define GREEN        ESC "[32m"
#define BROWN        ESC "[33m"
#define BLUE         ESC "[34m"
#define MAGENTA      ESC "[35m"
#define CYAN         ESC "[36m"
#define GRAY         ESC "[37m"
#define RGB(r, g, b) CSI "38;2;" #r ";" #g ";" #b "m"
#define YELLOW       RGB(255, 255, 0)
#define PEACH        RGB(255, 203, 164)  // 255, 218, 185
// printf() Format Modifiers
#define NTP_SIZE_MODIFIER  "z"
#define NTP_CHAR_MODIFIER  "hh"
#define NTP_SHORT_MODIFIER "h"
// Configuration
#if !defined(TRUE_RESULT_COLOR)
  #define TRUE_RESULT_COLOR GREEN
#endif
#if !defined(FALSE_RESULT_COLOR)
  #define FALSE_RESULT_COLOR RED
#endif
#if !defined(VALUE_STRING)
  #define VALUE_STRING (32)
#endif
#if !defined(EXPR_VALUE_STRING)
  #define EXPR_VALUE_STRING VALUE_STRING
#endif
#if !defined(TIMESTAMP_STRING)
  #define TIMESTAMP_STRING VALUE_STRING
#endif
#if !defined(A_VALUE_STRING)
  #define A_VALUE_STRING VALUE_STRING
#endif
#if !defined(B_VALUE_STRING)
  #define B_VALUE_STRING VALUE_STRING
#endif
// Callback function prototype
typedef void (*ntp_f)(FILE* stream, const char* a, const char* op, const char* b, bool result, const char a_value[A_VALUE_STRING],
                      const char b_value[B_VALUE_STRING], time_t time, const char* file, const char* function, int line, const char* format,
                      ...);
// Create timestamp string
inline char* ntp_default_timestamp(time_t time, char timestamp[TIMESTAMP_STRING]) {
  struct tm buf;
  localtime_r(&time, &buf);
  strftime(timestamp, TIMESTAMP_STRING, "%F:%T", &buf);
  return timestamp;
}
extern inline char* ntp_default_timestamp(time_t time, char timestamp[TIMESTAMP_STRING]);
// Default callback function
inline void ntp_default_f(FILE* stream, const char* a, const char* op, const char* b, bool result, const char a_value[A_VALUE_STRING],
                          const char b_value[B_VALUE_STRING], time_t time, const char* file, const char* function, int line, const char* format,
                          ...) {
  va_list args;
  va_start(args, format);
  fprintf(stream, result ? TRUE_RESULT_COLOR : FALSE_RESULT_COLOR);
  fprintf(stream, "%15s::%s:%s:%s:%d: (%s %s %s) :: (%s %s %s) :: (", __func__, ntp_default_timestamp(time, (char[TIMESTAMP_STRING]){ '\0' }),
          file, function, line, a, op, b, a_value, op, b_value);
  vfprintf(stream, format, args);
  fprintf(stream, ")\n");
  fprintf(stream, NO_COLOR);
  fflush(stream);
  va_end(args);
}
extern inline void ntp_default_f(FILE* stream, const char* a, const char* op, const char* b, bool result, const char a_value[A_VALUE_STRING],
                                 const char b_value[B_VALUE_STRING], time_t time, const char* file, const char* function, int line,
                                 const char* format, ...);
// Full NTP type macro w/callback function
#define ntp_type_full_f(stream, prefix, suffix, T, fmt, a, op, b, f, format, ...)                                                               \
  do {                                                                                                                                          \
    T ntp_tmp_a_ = ((T) (a));                                                                                                                   \
    T ntp_tmp_b_ = ((T) (b));                                                                                                                   \
    if (!(ntp_tmp_a_ op ntp_tmp_b_)) {                                                                                                          \
      (f)(stream, #a, #op, #b, false, FMT((char[A_VALUE_STRING]){ 0 }, prefix "%" fmt suffix, ntp_tmp_a_),                                      \
          FMT((char[B_VALUE_STRING]){ 0 }, prefix "%" fmt suffix, ntp_tmp_b_), time(NULL), __FILE__, __func__, __LINE__, format,                \
          ##__VA_ARGS__);                                                                                                                       \
    } else {                                                                                                                                    \
      (f)(stream, #a, #op, #b, true, FMT((char[A_VALUE_STRING]){ 0 }, prefix "%" fmt suffix, ntp_tmp_a_),                                       \
          FMT((char[B_VALUE_STRING]){ 0 }, prefix "%" fmt suffix, ntp_tmp_b_), time(NULL), __FILE__, __func__, __LINE__, format,                \
          ##__VA_ARGS__);                                                                                                                       \
    }                                                                                                                                           \
  } while (0)
// Full NTP type macro w/default callback function
#if !defined(ntp_type_full)
  #define ntp_type_full(prefix, suffix, T, fmt, a, op, b, format, ...)                                                                          \
    ntp_type_full_f(stderr, prefix, suffix, T, fmt, a, op, b, ntp_default_f, format, ##__VA_ARGS__)
#endif  // ntp_type_full
// Single expression (Boolean) NTP type macro w/callback function
#define ntp_ntp_f(stream, expr, f, format, ...)                                                                                                 \
  do {                                                                                                                                          \
    bool ntp_tmp_expr_ = (expr);                                                                                                                \
    if (!(ntp_tmp_expr_)) {                                                                                                                     \
      (f)(stream, #expr, "", "", false, "false", "", time(NULL), __FILE__, __func__, __LINE__, format, #__VA_ARGS__);                           \
    } else {                                                                                                                                    \
      (f)(stream, #expr, "", "", true, "true", "", time(NULL), __FILE__, __func__, __LINE__, #__VA_ARGS__);                                     \
    }                                                                                                                                           \
  } while (0)
// Single expression (Boolean) NTP type macro w/default callback function
#if !defined(ntp_ntp)
  #define ntp_ntp(expr, ...) ntp_ntp_f(stderr, expr, ntp_default_f, ##__VA_ARGS__)
#endif  // ntp_ntp
// Boolean tracepoints.
#define ntp_true(expr, ...)  ntp_ntp(true == expr, __VA_ARGS__)
#define ntp_false(expr, ...) ntp_ntp(false == expr, __VA_ARGS__)
// Base tracepoint
#define ntp_type(T, fmt, a, op, b, ...) ntp_type_full("", "", T, fmt, a, op, b, __VA_ARGS__)
// Builtin integer type specific tracepoints.
#define ntp_char(a, op, b, ...)   ntp_type_full("'\\x", "'", char, "02" NTP_CHAR_MODIFIER "x", a, op, b, __VA_ARGS__)
#define ntp_uchar(a, op, b, ...)  ntp_type_full("'\\x", "'", unsigned char, "02" NTP_CHAR_MODIFIER "x", a, op, b, __VA_ARGS__)
#define ntp_short(a, op, b, ...)  ntp_type(short, NTP_SHORT_MODIFIER "d", a, op, b, __VA_ARGS__)
#define ntp_ushort(a, op, b, ...) ntp_type(unsigned short, NTP_SHORT_MODIFIER "u", a, op, b, __VA_ARGS__)
#define ntp_int(a, op, b, ...)    ntp_type(int, "d", a, op, b, __VA_ARGS__)
#define ntp_uint(a, op, b, ...)   ntp_type(unsigned int, "u", a, op, b, __VA_ARGS__)
#define ntp_long(a, op, b, ...)   ntp_type(long int, "ld", a, op, b, __VA_ARGS__)
#define ntp_ulong(a, op, b, ...)  ntp_type(unsigned long int, "lu", a, op, b, __VA_ARGS__)
#define ntp_llong(a, op, b, ...)  ntp_type(long long int, "lld", a, op, b, __VA_ARGS__)
#define ntp_ullong(a, op, b, ...) ntp_type(unsigned long long int, "llu", a, op, b, __VA_ARGS__)
// size_t type specific tracepoints.
#define ntp_size(a, op, b, ...) ntp_type(size_t, NTP_SIZE_MODIFIER "u", a, op, b, __VA_ARGS__)
// Builtin float/double type specific tracepoints.
#define ntp_float(a, op, b, ...)  ntp_type(float, "f", a, op, b, __VA_ARGS__)
#define ntp_double(a, op, b, ...) ntp_type(double, "g", a, op, b, __VA_ARGS__)
// Builtin double type specific equality tracepoints.
#define ntp_double_equal(a, b, precision)
// Stdint type specific tracepoints.
#define ntp_int8(a, op, b, ...)   ntp_type(int8_t, PRIi8, a, op, b, __VA_ARGS__)
#define ntp_uint8(a, op, b, ...)  ntp_type(uint8_t, PRIu8, a, op, b, __VA_ARGS__)
#define ntp_int16(a, op, b, ...)  ntp_type(int16_t, PRIi16, a, op, b, __VA_ARGS__)
#define ntp_uint16(a, op, b, ...) ntp_type(uint16_t, PRIu16, a, op, b, __VA_ARGS__)
#define ntp_int32(a, op, b, ...)  ntp_type(int32_t, PRIi32, a, op, b, __VA_ARGS__)
#define ntp_uint32(a, op, b, ...) ntp_type(uint32_t, PRIu32, a, op, b, __VA_ARGS__)
#define ntp_int64(a, op, b, ...)  ntp_type(int64_t, PRIi64, a, op, b, __VA_ARGS__)
#define ntp_uint64(a, op, b, ...) ntp_type(uint64_t, PRIu64, a, op, b, __VA_ARGS__)
// String specific tracepoints.
#define ntp_string_equal(a, b, ...)           ntp_int(strcmp((a), (b)), ==, 0, ##__VA_ARGS__)
#define ntp_string_not_equal(a, b, ...)       ntp_false((strcmp((a), (b)) == 0), __VA_ARGS__)
#define ntp_memory_equal(size, a, b, ...)     ntp_true((memcmp((size), (a), (b)) == 0), __VA_ARGS__)
#define ntp_memory_not_equal(size, a, b, ...) ntp_false((memcmp((size), (a), (b)) == 0), __VA_ARGS__)
// Pointer specific tracepoints.
#define ntp_ptr(a, op, b, ...)       ntp_type(const void*, "p", a, op, b, __VA_ARGS__)
#define ntp_ptr_equal(a, b, ...)     ntp_ptr(a, ==, b, __VA_ARGS__)
#define ntp_ptr_not_equal(a, b, ...) ntp_ptr(a, !=, b, __VA_ARGS__)
#define ntp_null(ptr, ...)           ntp_ptr(ptr, ==, NULL, __VA_ARGS__)
#define ntp_not_null(ptr, ...)       ntp_ptr(ptr, !=, NULL, __VA_ARGS__)
#define ntp_ptr_null(ptr, ...)       ntp_ptr(ptr, ==, NULL, __VA_ARGS__)
#define ntp_ptr_not_null(ptr, ...)   ntp_ptr(ptr, !=, NULL, __VA_ARGS__)

#endif  // _NTP_H_
