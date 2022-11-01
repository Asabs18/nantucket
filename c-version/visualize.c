#include <assert.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "args.h"

int main(int argc, char *argv[]) {
  int nbins = default_number_bins;
  int bins[nbins];
  memset(bins, 0, sizeof bins);
  double dmin, dmax;
  double binsize;
  int trial;
  int nsusceptible, nrecovered, nvaccinated, ninfected, ndead;
  int binno;
  int overunderflow = 0;
  int i;
  FILE *input;
  FILE *output;
  // Insert Program Here
  assert(5 == argc);
  dmin  = atof(argv[1]);
  dmax  = atof(argv[2]);
  input = fopen(argv[3], "r");
  assert(input);
  output = fopen(argv[4], "w");
  assert(output);
  binsize = (dmax - dmin) / nbins;
  while (6 == fscanf(input, "%d %d %d %d %d %d\n", &trial, &nsusceptible, &nrecovered, &nvaccinated, &ninfected, &ndead)) {
    if (dmin <= ndead && ndead <= dmax) {
      binno = (ndead - dmin) / binsize;
      bins[binno]++;

    } else {
      overunderflow++;
    }
  }
  fclose(input);
  for (int i = 0; i < nbins; i++) {
    int ndead = dmin + binsize * (0.5 + i);
    fprintf(output, "%d %d\n", ndead, bins[i]);
  }
  fprintf(output, "# Underflows/Overflows: %d\n", overunderflow);
  fclose(output);

  return EXIT_SUCCESS;
}
