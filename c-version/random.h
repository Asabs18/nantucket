#include <stdlib.h>
#include <time.h>

// Random number between zero and one
double rand01() {
  static int needsrand = 1;
  if (needsrand) {
    srand(time(NULL));
    needsrand = 0;
  }
  return (rand() / (1.0 + RAND_MAX));
}
// Normal distribution:
double normal() {
  int nroll = 12;
  double sum = 0;
  int i;
  for (i = 0; i < nroll; i++) {
    sum += rand01();
  }
  return (sum - 6.0);
}
