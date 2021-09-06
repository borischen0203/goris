<img src="https://raw.githubusercontent.com/scraly/gophers/main/mac-gopher.png" alt="mac-gopher">

<p align="Left">
  <p align="Left">
    <a href="https://github.com/borischen0203/goris/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/borischen0203/goris.svg?logo=github&style=flat-square"></a>
    <a href="https://github.com/borischen0203/goris/actions/workflows/go.yml"><img alt="GitHub release" src="https://github.com/borischen0203/goris/actions/workflows/go.yml/badge.svg?logo=github&style=flat-square"></a>
    <!-- <a href="https://goreportcard.com/report/github.com/borischen0203/goris"><img src="https://goreportcard.com/badge/github.com/borischen0203/goris" alt="Code Status" /></a> -->
  </p>
</p>

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
> goris list # show a gopher image file name list

> goris get # download gopher image and save to desktop, or get a gopher link
```

### Demo
```bash
> goris list
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
Try to get '5th-element' Gopher...
Perfect! Just saved in /Users/boris/desktop/5th-element.png!

> goris get link 5th-element
Try to Generate gopher link...
<img src="https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png" alt="5th-element">
5th-element! I choose you! Paste above link in the readme!
```

## Tech stack
- Golang
- Cobra



### Todo:
- `view` command: Be able to preview gopher image.

### Reference:
https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43



