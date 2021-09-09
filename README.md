<img src="https://raw.githubusercontent.com/scraly/gophers/main/mac-gopher.png" alt="mac-gopher" width=500 height=474 >

<p align="Left">
  <p align="Left">
    <a href="https://github.com/borischen0203/goris/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/borischen0203/goris.svg?logo=github&style=flat-square"></a>
    <a href="https://github.com/borischen0203/goris/actions/workflows/go.yml"><img alt="GitHub release" src="https://github.com/borischen0203/goris/actions/workflows/go.yml/badge.svg?logo=github&style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/borischen0203/goris"><img src="https://goreportcard.com/badge/github.com/borischen0203/goris" alt="Code Status" /></a>
  </p>
</p>

# goris
This `goris` command-line tool is mainly used to download gopher png images from an amazing Go developer Aurélie's repo.
see below link:

https://github.com/scraly/gophers


## Features
- `get` command: Be able to download a gopher image to desktop.
- `get link` command: Be able to get a gopher image link with html.
- `list` command: Be able to display a list with all gopher image names.
- `view` command: Be able to pop up gopher image window.

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
> goris get link # get gopher image link
> goris view # pop up a gopher image window
```

### Demo example
`list` command
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
```
`get` command
```
> goris get 5th-element
Try to get '5th-element' Gopher...
Perfect! Just saved in /Users/boris/desktop/5th-element.png!
 ---------------------
< 5th-element download! >
 ---------------------
        \
         \
         ,_---~~~~~----._
  _,,_,*^____      _____``*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \`/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |
```

`get link` command
```
> goris get link 5th-element
Try to Generate gopher link...
<img src="https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png" alt="5th-element">
5th-element! I choose you! Paste above link in the readme!
```

`view` command
```
> goris view 5th-element
# pop up a gopher image window
```

## Tech stack
- Golang
- Cobra



### Todo:
- [X] `view` command: Be able to preview gopher image.
- [X] `get link` command: Be able to generate gopher image link with html.

### Reference:
https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43



