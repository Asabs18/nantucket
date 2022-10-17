// Global variables:
// Counters:
#include <sys/types.h>
int nsusceptible, ninfected, nrecovered, nvaccinated, ndead;
// Status indicators:
// Any positive value indicates user has been infected for n days.
// Negative values indicate:
const int is_susceptible = 0;
const int is_recovered   = -1;
const int is_vaccinated  = -2;
const int is_dead        = -3;
// Update status of each person and count them
void update_status(int npeople, int status[], int newstatus[]) {
  int i;
  // Reset counters in preparation for counting:
  nsusceptible = 0;
  ninfected    = 0;
  nrecovered   = 0;
  ndead        = 0;
  for (i = 0; i < npeople; i++) {
    status[i] = newstatus[i];
    if (status[i] == is_susceptible) {
      nsusceptible++;
    } else if (status[i] == is_recovered) {
      nrecovered++;
    } else if (status[i] == is_dead) {
      ndead++;
    } else if (status[i] > 0) {
      ninfected++;
    }
  }
}
void initialize_status(int npeople, int initial_infections, int nvaccinated, int status[]) {
  int i;
  int vmax;
  // Initial infections:
  for (i = 0; i < initial_infections; i++) {
    status[i] = 1;
  }
  // Vaccinations:
  // Don't exceed total number of people!
  vmax = initial_infections + nvaccinated;
  if (vmax > npeople) {
    vmax = npeople;
  }
  for (i = initial_infections; i < vmax; i++) {
    status[i] = is_vaccinated;
  }
  // Everybody else is susceptible:
  for (i = vmax; i < npeople; i++) {
    status[i] = is_susceptible;
  }
}
// MT data structures
//
// communication mechanism
//  - memory - communication is a via local (process) memory (pointer)
//  - file - communication is a via a file (descriptor)
//  - pipe - communication is a  via a pipe (descriptor)
//
// using a pipe w/an eye towards GO channels, and eliminating the need to
// implement a queue, and the associated exclusion/notification mechanisms.
//
// create (loop on failure) thread for each trial, check results or join
// create n workers, given each T/n trials in memory, join
// create n workers, given each T/n trials in file, join
// create n workers, write T trials to pipe, join
typedef enum { memory, pipe, file } channel_t;
typedef struct {
  channel_t chan;
  union {
    struct {
      void* p;
    } memory;
    struct {
      int fd;
    } pipe;
    struct {
      int fd;
      off_t offset;
    } file;
  };
} thread_channel_t;
typedef struct {
  size_t n_trials;
  size_t n_people;
  size_t n_days;
} trial_args_t;
typedef struct {
} trial_results_t;
typedef struct {
  thread_channel_t in;
  trial_args_t trial_args;
} thread_args_t;
typedef struct {
  thread_channel_t out;
  trial_results_t trial_result;
} thread_return_t;
// ANCHOR - void update(const thread_args_t* args)
inline void update(const thread_args_t* args) {
  // for each trial that the worker's responsible for...
  for (size_t trial = 0; trial < args->trial_args.n_trials; trial++) {
    // ...for each day in the trial...
    for (size_t day = 0; day < args->trial_args.n_days; day++) {
      // ...update the status/condition of each person (for each day in the trial).
      for (size_t person = 0; person < args->trial_args.n_people; person++) {
      }
    }
  }
}