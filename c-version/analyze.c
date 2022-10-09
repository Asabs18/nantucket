#include <math.h>
double standard_deviation(double sum, double sum_squared, int npeople,
                          int status[]) {
  double mean = sum / npeople;
  double diff_mean_sq = 0.0;
  for (int i = 0; i < npeople; ++i) {
    double diff_mean = status[i] - mean;
    diff_mean_sq += (diff_mean * diff_mean);
  }
  return sqrt(diff_mean_sq / npeople);
}
#include "simulate.c"