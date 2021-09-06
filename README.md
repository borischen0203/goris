# goris
This `goris` command-line tool is mainly used to download a gopher png files from a amazing Go developer Aurélie's repo.
see below link:

https://github.com/scraly/gophers


## Features
- `get` command: Download a gopher image.
- `list` command: List all gopher image names.

## Installation

### On macOS (or Linux) via Homebrew
```bash
> brew tap borischen0203/goris
> brew install goris
```

## How to use

### Run
```bash
> goris list # show a gopher image file name list 

> goris get # download gopher image and save to desktop
```

### Demo
```bash
> goris list 
[Try to get Gopher list...]
5th-element
arrow-gopher
back-to-the-future-v2
baywatch
big-bang-theory
.
.
.

> goris get 5th-element
```

## Tech stacks
- Golang
- Cobra



### Todo:
- `view` command: Be able to preview gopher image.

- `get link` command: Be able to get the link that allow to past in README.md file.




