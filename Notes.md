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

### Loggers

- [yksz/c-logger: A simple logging library for C/C++](https://github.com/yksz/c-logger)
- [rxi/log.c: A simple logging library implemented in C99](https://github.com/rxi/log.c)

### Command Line Parsers

- [getopt(3) - Linux manual page](https://www.man7.org/linux/man-pages/man3/getopt.3.html)
- [Argp (The GNU C Library)](https://www.gnu.org/software/libc/manual/html_node/Argp.html)
- [opt command line parser - Table of Contents](https://public.lanl.gov/jt/Software/opt/opt_toc.html)

### Parallel Opportunities (Multi-Threading / `pthreads`)

- data _axis_
  - trials
  - days
  - people
- _critical data_: `status[` _i + 1_ `] <- status [` _i_ `]`, that is:

```C
for each t in trials:
  for each d in days:       // in the trial
    for each p in persons:  // on each day in each trial
      status[ p + 1 ] <- status [ p ]
```

#### [Models](https://maxim.int.ru/bookshelf/PthreadsProgram/htm/r_19.html)

- Boss/Worker
- Peer
- Pipeline

```C
// communication mechanism
//  - memory - communication is a via local (process) memory (pointer)
//  - file - communication is a via a file (descriptor)
//  - pipe - communication is a  via a pipe (descriptor)
// using a pipe w/an eye towards GO channels, and eliminating the need to
// implement a queue, and the associated exclusion/notification mechanisms.
enum {
  memory,
  file,
  pipe
} channel_t;

typedef struct {
  channel_t chan;

} thread_args_t;

typedef struct {
  channel_t chan;
  union {
    // pipe
    int fd;
    // memory
    void* p;
    // file
    struct {
      int fd;
      off_t offset;
    }
  }
}

typedef struct {

} thread_args_t;

typedef struct {

} thread_returns_t;
```

- `status[P]`: status of all people at the _beginning_ of each day (before `i`)
- `newstatus[P]`: status of all people at the _end_ of each day (after `i`)
- determining if person dies (`die_or_not()`) is independent of any other person
- determining if a person becomes infected (`catch_or_not()`) is dependent on `i` other _random_ people
  - `newstatus[x]` is dependent on `i` _previous statuses_
    - `status` or `newstatus`, _if already updated_

[^simulate]: _`simulate` `vprob` `tprob` `file`_
[^analyze]:  _`analyze` `vprob` `tprob` `file`_
[^visualize]: _`visualize` `dmin` `dmax` `inputfile` `outputfile`_
