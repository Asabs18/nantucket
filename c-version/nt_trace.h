// [log.c/log.h at master · rxi/log.c](https://github.com/rxi/log.c/blob/master/src/log.h)

#if !defined _NT_TRACE_H_
  #define _NT_TRACE_H_

  #include <stdarg.h>
  #include <stdbool.h>
  #include <stdio.h>
  #include <time.h>

  #define NT_TRACE_VERSION "0.1.0"

typedef struct {
  va_list ap;
  const char* fmt;
  const char* fmt_prefix;
  const char* fmt_suffix;
  const char* a;  // BACK: capture value
  const char* op; // BACK: enum?
  const char* b;  // BACK: capture value
  const char* function;
  const char* file;
  struct tm* time;
  void* udata;
  int line;
  int level;
} ntt_event_t;

typedef void (*ntt_trace_f)(ntt_event_t* ev);
typedef void (*ntt_lock_f)(bool lock, void* udata);

  #define LEVEL_NAME_PREFIX_UC(name)  NTT##name
  #define LEVEL_NAME_PREFIX_LCC       ntt
  #define EXPAND_AS_ENUMERATION(a, b) a,
  #define EXPAND_AS_TRACER_FUNCTIONS(a, b)                                                                                                      \
    static inline void f_##b(int level, const char* file, const char* function, int line, const char* fmt, ...) {                               \
      f_ntt_ntt(a, file, function, line, fmt);                                                                                                  \
    }

  #define LEVEL_TABLE(LEVEL)                                                                                                                    \
    LEVEL(NTT_TRACE, ntt_trace)                                                                                                                 \
    LEVEL(NTT_DEBUG, ntt_debug)                                                                                                                 \
    LEVEL(NTT_INFO, ntt_info)                                                                                                                   \
    LEVEL(NTT_WARN, ntt_warn)                                                                                                                   \
    LEVEL(NTT_ERROR, ntt_error)

//
enum { LEVEL_TABLE(EXPAND_AS_ENUMERATION) NUM_STATES };
void f_ntt_ntt(int level, const char* file, const char* function, int line, const char* fmt, ...);
// LEVEL_TABLE(EXPAND_AS_TRACER_FUNCTIONS); HACK: why? diff?

//   #define ntt_trace(...) ntt_ntt(NTT_TRACE, __FILE__, __PRETTY_FUNCTION__, __LINE__, __VA_ARGS__)
  #define ntt_level(level, ...) f_ntt_ntt(level, __FILE__, __PRETTY_FUNCTION__, __LINE__, __VA_ARGS__)
  #define ntt_trace(...)        ntt_level(NTT_TRACE, __FILE__, __PRETTY_FUNCTION__, __LINE__, __VA_ARGS__)
  #define ntt_debug(...)        ntt_ntt(NTT_DEBUG, __FILE__, __PRETTY_FUNCTION__, __LINE__, __VA_ARGS__)
  #define ntt_info(...)         ntt_ntt(NTT_INFO, __FILE__, __PRETTY_FUNCTION__, __LINE__, __VA_ARGS__)
  #define ntt_warn(...)         ntt_ntt(NTT_WARN, __FILE__, __PRETTY_FUNCTION__, __LINE__, __VA_ARGS__)
  #define ntt_error(...)        ntt_ntt(NTT_ERROR, __FILE__, __PRETTY_FUNCTION__, __LINE__, __VA_ARGS__)
  #define ntt_fatal(...)        ntt_ntt(NTT_FATAL, __FILE__, __PRETTY_FUNCTION__, __LINE__, __VA_ARGS__)

const char* ntt_level_string(int level);
void ntt_set_lock(ntt_lock_f fn, void* udata);
void ntt_set_level(int level);
void ntt_set_quiet(bool enable);
int ntt_add_callback(ntt_trace_f fn, void* udata, int level);
int ntt_add_fp(FILE* fp, int level);

#endif  // _NT_TRACE_H_