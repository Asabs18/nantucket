// Global variables:
// Counters:
int nsusceptible, ninfected, nrecovered, nvaccinated, ndead;
// Status indicators:
// Any positive value indicates user has been infected for n days.
// Negative values indicate:
const int is_susceptible = 0;
const int is_recovered = -1;
const int is_vaccinated = -2;
const int is_dead = -3;
// Update status of each person and count them
void update_status(int npeople, int status[], int newstatus[]) {
  int i;
  // Reset counters in preparation for counting:
  nsusceptible = 0;
  ninfected = 0;
  nrecovered = 0;
  ndead = 0;
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
void initialize_status(int npeople, int initial_infections, int nvaccinated,
                       int status[]) {
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
