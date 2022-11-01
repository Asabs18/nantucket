#include <assert.h>
#include <pthread.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>

#include "args.h"
#include "log.h"
#include "opt.h"
#include "random.h"
#include "virus.h"

// MT data structures
//
// communication mechanism
//  - memory - communication is a via local (process) memory (pointer)
//  - file - communication is a via a file (descriptor)
//  - pipe - communication is a  via a pipe (descriptor)
//
// using a pipe w/an eye towards GO channels, and eliminating the need to
// implement a queue, and the associated exclusion/notification mechanisms.
// create (loop on failure) thread for each trial, check results or join
// create n workers, given each T/n trials in memory, join
// create n workers, given each T/n trials in file, join
// create n workers, write T trials to pipe, join

// channel meta-data:
typedef struct {
  channel_t chan;
  union {
    // array of trial data
    struct {
      void* p;
      size_t count;
    } memory;
    // pipe queue
    struct {
      int fd;
    } pipe;
    // fifo queue
    struct {
      int fd;
    } fifo;
    // message queue (posix) queue
    struct {
      //   int fd;
      //   off_t offset;
    } mq;
    // file queue
    struct {
      //   int fd;
      //   off_t offset;
    } file;
  };
} thread_channel_t;
// trial data
typedef struct {
  size_t max_exposures_per_day;
  double max_random_sick_days;
  size_t number_of_people_per_trial;
  double death_prob;
  size_t number_of_trial_days;
  float percent_initial_infections;
  size_t number_of_sick_days;
  size_t initial_vaccinated;
} trial_arg_t;
//
typedef struct {
  size_t number_of_dead;
  size_t number_recovered;
  time_t time;
} trial_result_t;
//
typedef struct {
  const thread_channel_t* in;
  thread_channel_t* out;
} thread_arg_t;
//
typedef struct {
  trial_arg_t trial_arg;
  size_t number_of_threads;
  size_t stack_size;
  size_t number_of_trials;
} app_arg_t;
//
bool get_next_trial(const thread_channel_t* in, const trial_arg_t* trial_arg, size_t* processed_trials) {
  switch (in->chan) {
    case memory_channel: {
      const trial_arg_t* trial_args = in->memory.p;
      size_t number_of_trials       = in->memory.count;
      // for each trial that the worker's responsible for...
      if (*processed_trials < number_of_trials) {
        // for (size_t trial = 0; trial < number_of_trials; trial++) {
        trial_arg = &trial_args[*processed_trials];
        (*processed_trials)++;
        return true;
      }
      return false;
    } break;
    case pipe_channel: {
      int fd      = in->pipe.fd;
      ssize_t ret = read(fd, &trial_arg, sizeof &trial_arg);
      assert(0 == ret || sizeof &trial_arg == ret);
      return 0 != ret;
      assert(false);
    } break;
    default:
      assert(false);
  }
  assert(false);
}
//
bool put_next_trial(const thread_channel_t* out, const trial_result_t* trial_result, size_t processed_trials) {
  //   bool put_next_trial(const thread_channel_t* out, trial_result_t* trial_results[], size_t processed_trials, ...) {
  switch (out->chan) {
    case memory_channel: {
      trial_result_t** trial_results = out->memory.p;
      size_t number_of_trials        = out->memory.count;
      // for each trial that the worker's responsible for...
      if (processed_trials < number_of_trials) {
        *trial_results[processed_trials] = *trial_result;
        return true;
      }
      return false;
    } break;
    case pipe_channel: {
      int fd      = out->pipe.fd;
      ssize_t ret = write(fd, &(*trial_result), sizeof *trial_result);
      assert(0 == ret || sizeof *trial_result == ret);
      return 0 != ret;
      assert(false);
    } break;
    default:
      assert(false);
  }
  assert(false);
}
//
void* update(void* args) {
  thread_arg_t* thread_args   = (thread_arg_t*) args;
  const thread_channel_t* in  = thread_args->in;
  const thread_channel_t* out = thread_args->out;
  size_t processed_trials     = 0;
  trial_arg_t trial_arg       = { 0 };
  //   for each trial that the worker's responsible for...
  while (get_next_trial(in, &trial_arg, &processed_trials)) {
    const trial_result_t trial_result = { -1, -1 };
    // ...for each day in the trial...
    for (size_t day = 0; day < trial_arg.number_of_trial_days; day++) {
      // ...update the status/condition of each person (for each day in the trial).
      for (size_t person = 0; person < trial_arg.number_of_people_per_trial; person++) {
        // newstatus = f(status, &trial_result);
      }
    }
    put_next_trial(out, &trial_result, processed_trials);
  }
  pthread_exit(0);
}

//   size_t number_of_trials    = 0;
//   switch (thread_args->in->chan) {
//     case memory_channel: {
//       //   trial_args_t* trial_args = thread_args->in->memory.p;
//       //   number_of_trials         = thread_args->in->memory.count;
//       //   // for each trial that the worker's responsible for...
//       //   for (size_t trial = 0; trial < number_of_trials; trial++) {
//       //     trial_args_t* trial_arg = &trial_args[trial];
//       //     // ...for each day in the trial...
//       //     for (size_t day = 0; day < trial_arg->number_of_trial_days; day++) {
//       //       // ...update the status/condition of each person (for each day in the trial).
//       //       for (size_t person = 0; person < trial_arg->number_of_people_per_trial; person++) {
//       //         // newstatus = f(status);
//       //       }
//       //     }
//       //   }
//     } break;
//     case pipe_channel: {
//       int fd = thread_args->in->pipe.fd;
//     } break;
//     default:
//       assert(false);
//   }
  // Global variables:
  // Counters:
  int nsusceptible, ninfected, nrecovered, nvaccinated, ndead;

  int catch_or_not(double tprob, int npeople, int status[]) {
    // Insert function here!
    int nexposures = default_max_exposures_per_day * rand01();
    for (size_t i = 0; i < nexposures; i++) {
      int otherperson = rand01() * npeople;
      if (status[otherperson] > 0) {
        if (rand01() < tprob) return 1;
      }
    }
    return is_susceptible;
  }
  int die_or_not(double dprob, int sick_days, int person, int npeople, int status[]) {
    // Insert function here!
    if (rand01() < dprob) {
      return is_dead;
    }
    if (status[person] > sick_days + default_max_random_sick_days * normal()) {
      return is_recovered;
    }
    return status[person] + 1;
  }
  // HACK: int main(int argc, char* argv[]) {
  int main() {
    // HACK: init w/defaults
    app_arg_t app_arg = { .number_of_threads = default_number_of_threads,
                          .stack_size        = default_size_of_thread_stack,
                          .number_of_trials  = default_number_of_trials,
                          .trial_arg         = { .max_exposures_per_day      = default_max_exposures_per_day,
                                                 .max_random_sick_days       = default_max_random_sick_days,
                                                 .number_of_people_per_trial = default_number_of_people_per_trial,
                                                 .death_prob                 = default_death_prob,
                                                 .number_of_trial_days       = default_number_of_trial_days,
                                                 .percent_initial_infections = default_percent_initial_infections,
                                                 .number_of_sick_days        = default_number_of_sick_days,
                                                 .initial_vaccinated         = default_vaccination_prob * default_number_of_people_per_trial } };
    //
    const int npeople = default_number_of_people_per_trial;
    int status[npeople];
    int newstatus[npeople];
    int day, person;
    double vprob;                                 // Vaccination probability.
    double tprob;                                 // Transmission probability
    double dprob           = default_death_prob;  // Death probability per day
    int ndays              = default_number_of_trial_days;
    int initial_infections = default_percent_initial_infections * npeople;
    int sick_days          = default_number_of_sick_days;
    FILE* output;
    double sum  = 0;
    double sum2 = 0;
    int trial;
    int ntrials = default_number_of_trials;
    // Insert the rest of the program here!
    //  assert(4 == argc);
    vprob       = default_vaccination_prob;   // atof(argv[1]);
    tprob       = default_transmission_prob;  // atof(argv[2]);
    output      = fopen(/*argv[3]*/ "djg.dat", "w");
    nvaccinated = vprob * npeople;
    // app_arg UT
    assert(app_arg.number_of_threads == default_number_of_threads);
    assert(app_arg.number_of_trials == ntrials);
    assert(app_arg.stack_size == default_size_of_thread_stack);
    const trial_arg_t* trial_arg = &app_arg.trial_arg;
    assert(trial_arg->max_exposures_per_day == default_max_exposures_per_day);
    assert(trial_arg->max_random_sick_days == default_max_random_sick_days);
    assert(trial_arg->number_of_people_per_trial == npeople);
    assert(trial_arg->death_prob == dprob);
    assert(trial_arg->number_of_trial_days == default_number_of_trial_days);
    assert(trial_arg->percent_initial_infections == default_percent_initial_infections);
    assert(trial_arg->number_of_sick_days == sick_days);
    assert(trial_arg->initial_vaccinated == vprob * npeople);
    //
    //
    for (size_t thread = 0; thread < app_arg.number_of_threads; thread++) {
    }

    // update w/CLI values
    // app_arg.[...] = ...
    trial_arg_t* trial_args[app_arg.number_of_trials];
    for (size_t i = 0; i < app_arg.number_of_trials; i++) {
      trial_args[i] = &app_arg.trial_arg;
    }
    trial_result_t trial_results[app_arg.number_of_trials];
    memset(trial_results, 0, sizeof trial_results);
    //
    thread_channel_t in     = { .chan = memory_channel, .memory = { .p = trial_args, .count = app_arg.number_of_trials } };
    thread_channel_t out    = { .chan = memory_channel, .memory = { .p = trial_results, .count = app_arg.number_of_trials } };
    thread_arg_t thread_arg = { .in = &in, .out = &out };

    for (trial = 0; trial < ntrials; trial++) {
      if (trial % 20 == 0) {
        printf("Processing trial %d\n", trial);
      }
      initialize_status(npeople, initial_infections, nvaccinated, status);

      for (day = 0; day < ndays; day++) {
        for (person = 0; person < npeople; person++) {
          if (is_susceptible == status[person]) {
            newstatus[person] = catch_or_not(tprob, npeople, status);
          } else if (status[person] > 0) {
            newstatus[person] = die_or_not(dprob, sick_days, person, npeople, status);
          } else {
            newstatus[person] = status[person];
          }
        }
        update_status(npeople, status, newstatus);
        // fprintf(output, "%d %d %d %d %d %d\n", day, nsusceptible, nrecovered, nvaccinated, ninfected, ndead);
      }
      sum += ndead;
      sum2 += (ndead * ndead);
      fprintf(output, "%d %d %d %d %d %d\n", trial, nsusceptible, nrecovered, nvaccinated, ninfected, ndead);
    }
    printf("\n");

    printf("Population: %d\n", npeople);
    printf("Vaccination Probability: %lf\n", vprob);
    printf("Transmission Probability: %lf\n", tprob);
    printf("Initial Infections: %d\n", initial_infections);
    printf("Simulation Period: %d days\n", ndays);
    printf("Number of Trials: %d\n", ntrials);
    printf("Average Number of Deaths per Trial: %lf\n", sum / ntrials);
    printf("Standard Deviation of Deaths per Trial: %lf\n", standard_deviation(sum, sum2, npeople, status));
    // printf("Number of Recovered: %d\n", nrecovered);
    // printf("Number of Dead: %d\n", ndead);
    // printf("Case Fatality Rate: %lf\n", (double) ndead / (nrecovered + ndead));
  }
