# Project

## Project Description

### Lesson Summary

- [ ] 2023-02-05:
  - Discussed state of `gogo` 
    - djg: propagation of configuration mechanism to all illustrations
    - djg: documentation
  - Discussed next project – some sort of _cellular automata_ or _genetic algorithm_ **playground** (_laboratory_😎❓)
    - [simgo package - github.com/fschuetz04/simgo - Go Packages](https://pkg.go.dev/github.com/fschuetz04/simgo)
    - [About PySim – PySim](https://pysim.org/about/)
    - Game engines – text vs. graphical, rules engines – flexibility, domain specific, knowledge base, grids, Conway’s GOL, etc.
  - Looked at Visualize Spike code and researched how to render HTML at run-time
    - **as: Would a direct image/charting/graphics package be better or too complex for Nantucket’s needs?**
  - Ran `Cute Mold` and `Protozoa` - partial inspiration for new project – on Windows w/o any issues❗
    - djg: GitPod attempt confrms issue is most likely OpenGL on Linux
  - Next
    - as: complete Visualize spike
    - as/djg: discuss next project

- [ ] 2023-01-23:
  - [ ] Discussed our respective sports teams (let me know when you want to learn about 🏒), `TODO.MD` and progress on our respective spikes
  - [ ] Tried to get CSV integration working but ran into bunch of github issues…*luckily, not as bad as it looked*
  - Next
    - [ ] as: update sources/GitHub back to *spec*
    - [ ] as: Edit away all dead/commented code
    - [ ] as: Update CSV code to create one file instead `multiTrials`
    - [ ] as: Eliminate `multiTrials`
    - [ ] as: Complete CSV integration
    - [ ] as: Charting spike

### Backlog

- [ ] as: Try to get go results to match pdf (Rewrite init functions to fix)
- [ ] as: Look into concurrent potential (Trials, Days, People)
- [ ] as: Look over c code (and rest of the project) for guidance   
- [ ] as: Charting spike
- [ ] as: Integrate charting support
- [ ] as: Integrate gogo features (e.g., concurrency, configuration…)

### Todo

- [ ] as: Change structure from serial to concurrent model based on trial
- [ ] as: Spike visualize
- [ ] as: Write visualize
- [ ] as: SPIKE: go routines and channels
- [ ] as: Collect time for each trial and overall time + other statistics/subsequent cmdLineVars
- [ ] as: Unit tests

### In Progress

- [ ] as: update nantucket to use gota and to read and write to csv file
- [ ] as: Decide plotting package and do a small spike
- [ ] as: Unit tests (Use to do spike on data and maybe plotting package)
- [ ] as: Use the output file from cmdLineVars to store csv data
- [ ] djg: Viper spike
- [ ] as: Feedback on gogo
- [ ] as: Use the output file from cmdLineVars to store csv data
- [ ] djg: Surface all constants in the c-version  
- [ ] djg: Spec out MT version  

### Done ✓

- [✓] as: Integrate CSV support
- [✓] as: Check in all updates to Nantucket
- [✓] djg: Reviewed gogo executable (after chmod to enable 'x')
- [✓] djg: Discussed diff. pkgs., five OS algos. & some real world examples
- [✓] as: Clean up printing/output from one trial to multi trials
- [✓] as: Restructure cmdLineVars/environment to make trial/application specific variables
- [✓] as: clean up repo and simulate file
- [✓] as: Update all flags (including default values)
- [✓] as: Read up on cobra(issues in simulate)
- [✓] as: Update go
- [✓] as: Change go setup to work with 2 commands(simulate/visualize)
- [✓] as: Change loop to work with one or more than one trial
- [✓] as: Add Cobra command line
- [✓] as: Refactor Go code
- [✓] as: Add my work into the cobra setup  
- [✓] djg: C versions of *analysis* and *visualize*.  
- [✓] djg: Add initial GO files to repo.  
- [✓] djg: Update package name.  
- [✓] djg: Create global flags.  
- [✓] djg: Create sub-commands.  
