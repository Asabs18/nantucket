#include "random.h"
#include "virus.h"
#include <assert.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
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
  // Insert the rest of the program here!
  //assert(argc == 4);
  vprob = 0.0;// atof(argv[1]);
  tprob = 0.015;//atof(argv[2]);
  output = fopen(/*argv[3]*/ "simulate-0.0-0.015.dat", "w");

  nvaccinated = vprob * npeople;
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
  }

  printf("Population: %d\n", npeople);
  printf("Vaccination Probability: %lf\n", vprob);
  printf("Transmission Probability: %lf\n", tprob);
  printf("Initial Infections: %d\n", initial_infections);
  printf("Simulation Period: %d days\n", ndays);
  printf("Number of Recovered: %d\n", nrecovered);
  printf("Number of Dead: %d\n", ndead);
  printf("Case Fatality Rate: %lf\n", (double)ndead / (nrecovered + ndead));
}
