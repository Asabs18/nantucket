# Notes

## Go Version

### `Cobra`

- [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
  - [Installing](https://github.com/spf13/cobra#installing)
  - [cobra-cli](https://github.com/spf13/cobra-cli/blob/main/README.md)
    - [Configuring the cobra generator](https://github.com/spf13/cobra-cli/blob/main/README.md#configuring-the-cobra-generator)
    - [Add commands to a project](https://github.com/spf13/cobra-cli/blob/main/README.md#add-commands-to-a-project)
    - `simulate` `[-v | --vprob]` `[-t | --tprob]` `input-file`[^simulate]
    - `analyze` `[-v | --vprob]` `[-t | --tprob]` `input-file`[^analyze]
    - `visualize` `[-n | dmin]` `[-x | dmax]` `input-file` `output-file`[^visualize]
      - `[-n | dmin]`
      - `[-x | dmax]`
      - `output-file`
      - Common Options:
        - `-v`
        - `-t`
        - `input-file`

- [Add commands to a project](https://github.com/spf13/cobra-cli/blob/main/README.md#add-commands-to-a-project)

### Modules

- [Tutorial: Getting started with multi-module workspaces - The Go Programming Language](https://go.dev/doc/tutorial/workspaces)

## C Version

### Parallel Opportunities

#### Definitions

- `p`:  represents opportunities for parallelism
  - processes
  - _threads_ (`pthreads`)
  - tasks
  - _coroutines_ (`goroutines`)
- `q`:  _data_ queue
  - shared memory
    - remote (process) queue
  - pipes
    - local (anonymous) /remote (named)
  - global variables
    - `pthread` managed queue(s)
  - `uq`:     unbuffered queue
  - `bq(n)`:  buffered queue
- `t`:  number of trials
- `d`:  number of days
- `p`:  number of people
- `i`:  number of _interactions_ a person has per day
- `status[P]`: status of all people at the _beginning_ of each day (before `i`)
- `newstatus[P]`: status of all people at the _end_ of each day (after `i`)
- determining if person dies (`die_or_not()`) is independent of any other person
- determining if a person becomes infected (`catch_or_not()`) is dependent on `i` other _random_ people
  - `newstatus[x]` is dependent on `i` _previous statuses_
    - `status` or `newstatus`, _if already updated_

#### Trial

- run `n` trials in parallel
  - enqueue `uq`: `t` - `n`
  - dequeue until empty
- one `p` executes `t` / `p` trials in parallel
  - no queue

#### Day

- run `n` days in pipeline
- same as Trials
  
#### Person

- either❓

[^simulate]: _`simulate` `vprob` `tprob` `file`_
[^analyze]:  _`analyze` `vprob` `tprob` `file`_
[^visualize]: _`visualize` `dmin` `dmax` `inputfile` `outputfile`_
