#if !defined(_ARGS_H_)
  #define _ARGS_H_

const int max_exposures                = 20;
const double max_random_sick_days      = 3.0;
const int max_people                   = 1e+4;
const double default_death_prob        = 0.015;
const int number_of_trial_days         = 100;
const float percent_initial_infections = 0.01;
const int number_of_sick_days          = 14;
const int number_of_trials             = 1000;

  // HACK: define as used for array declaration
  //  const int nbins = 50;
  #define nbins 50

#endif  // _ARGS_H_