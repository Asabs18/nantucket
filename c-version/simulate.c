#include "random.h"
#include "virus.h"
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
int catch_or_not(double tprob, int npeople, int status[]) {
  // Insert function here!
  int nexposures = 20 * rand01();
  for (size_t i = 0; i < nexposures; i++) {
    int otherperson = rand01() * npeople;
    if (status[otherperson] > 0) {
      if (rand01() < tprob)
        return 1;
    }
  }
  return is_susceptible;
}
int die_or_not(double dprob, int sick_days, int person, int npeople,
               int status[]) {
  // Insert function here!
  if (rand01() < dprob) {
    return is_dead;
  }
  if (status[person] > sick_days + 3.0 * normal()) {
    return is_recovered;
  }
  return status[person] + 1;
}
int main(int argc, char *argv[]) {
  const int npeople = 1e+4;
  int status[npeople];
  int newstatus[npeople];
  int day, person;
  double vprob;         // Vaccination probability.
  double tprob;         // Transmission probability
  double dprob = 0.015; // Death probability per day
  int ndays = 100;
  int initial_infections = 0.01 * npeople;
  int sick_days = 14;
  FILE *output;
#if defined _ANALYZE_CPP_
  double sum = 0;
  double sum2 = 0;
  int trial;
  int ntrials = 1000;
#endif // _ANALYZE_CPP_
  // Insert the rest of the program here!
  assert(4 == argc);
  vprob = atof(argv[1]);
  tprob = atof(argv[2]);
  output = fopen(argv[3], "w");

  nvaccinated = vprob * npeople;
#if defined _ANALYZE_CPP_
  for (trial = 0; trial < ntrials; trial++) {
    printf("Processing trial %d\n", trial);
#endif // _ANALYZE_CPP_
    initialize_status(npeople, initial_infections, nvaccinated, status);

    for (day = 0; day < ndays; day++) {
      for (person = 0; person < npeople; person++) {
        if (is_susceptible == status[person]) {
          newstatus[person] = catch_or_not(tprob, npeople, status);
        } else if (status[person] > 0) {
          newstatus[person] =
              die_or_not(dprob, sick_days, person, npeople, status);
        } else {
          newstatus[person] = status[person];
        }
      }
      update_status(npeople, status, newstatus);
      fprintf(output, "%d %d %d %d %d %d\n", day, nsusceptible, nrecovered,
              nvaccinated, ninfected, ndead);
#if defined _ANALYZE_CPP_
#else
#endif // _ANALYZE_CPP_
    }
#if defined _ANALYZE_CPP_
    sum += ndead;
    sum2 += (ndead * ndead);
    fprintf(output, "%d %d %d %d %d %d\n", trial, nsusceptible, nrecovered,
            nvaccinated, ninfected, ndead);
  }
  printf("\n");
#endif // _ANALYZE_CPP_

  printf("Population: %d\n", npeople);
  printf("Vaccination Probability: %lf\n", vprob);
  printf("Transmission Probability: %lf\n", tprob);
  printf("Initial Infections: %d\n", initial_infections);
  printf("Simulation Period: %d days\n", ndays);
#if defined _ANALYZE_CPP_
  printf("Number of Trials: %d\n", ntrials);
  printf("Average Number of Deaths per Trial: %lf\n", sum / ntrials);
  printf("Standard Deviation of Deaths per Trial: %lf\n",
         standard_deviation(sum, sum2, npeople, status));
#else
  printf("Number of Recovered: %d\n", nrecovered);
  printf("Number of Dead: %d\n", ndead);
  printf("Case Fatality Rate: %lf\n", (double)ndead / (nrecovered + ndead));
#endif // _ANALYZE_CPP_
}
