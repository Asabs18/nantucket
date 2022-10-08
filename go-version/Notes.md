# Notes

## `Cobra`

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

## `GO`

- [Tutorial: Getting started with multi-module workspaces - The Go Programming Language](https://go.dev/doc/tutorial/workspaces)

[^simulate]: _`simulate` `vprob` `tprob` `file`_
[^analyze]:  _`analyze` `vprob` `tprob` `file`_
[^visualize]: _`visualize` `dmin` `dmax` `inputfile` `outputfile`_
