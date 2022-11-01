#if !defined(_ARGS_H_)
  #define _ARGS_H_

  #include <stddef.h>

  #include "log.h"
// channel types where threads find their trial data
typedef enum { memory_channel, pipe_channel, fifo_channel, mq_channel, file_channel } channel_t;
// simulation arguments
const int default_number_of_trials        = 100;  // HACK: 1000;
const size_t default_number_of_threads    = 0;    // HACK: 4;
const size_t default_size_of_thread_stack = (8 * 1024);
const channel_t default_in_channel        = memory_channel;
const channel_t default_out_channel       = memory_channel;
const int default_number_bins             = 50;
const int default_log_level               = LOG_TRACE;  // HACK: adjust appropriately, add to CLI
// trial arguments
const int default_max_exposures_per_day        = 20;
const double default_max_random_sick_days      = 3.0;
const int default_number_of_people_per_trial   = 10000;  // 1e+4;
const double default_death_prob                = 0.015;
const double default_vaccination_prob          = 0.00;
const double default_transmission_prob         = 0.015;
const int default_number_of_trial_days         = 100;
const float default_percent_initial_infections = 0.01;
const int default_number_of_sick_days          = 14;

#endif  // _ARGS_H_