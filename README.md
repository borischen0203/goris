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


# Features
- `get` command: Be able to download a gopher image to desktop.
- `get link` command: Be able to get a gopher image link with html.
- `list` command: Be able to display a list with all gopher image names.
- `view` command: Be able to pop up gopher image window.

# How to use

## On macOS via Homebrew
Required
- Install homebrew
- Install chrome browser

Step1:
```bash
$ brew tap borischen0203/goris
```
Step2:
```bash
$ brew install goris
```

### Run demo
#### `list` command
```bash
# show a gopher image file name list
$ goris list
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
#### `get` command
```bash
# download gopher image and save to desktop
$ goris get 5th-element
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

##### `get link` command
```bash
# get gopher image link
$ goris get link 5th-element
Try to Generate gopher link...
<img src="https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png" alt="5th-element">
5th-element! I choose you! Paste above link in the readme!
```

##### `view` command
```bash
# pop up a gopher image window
$ goris view 5th-element
Pop-up 5th-element window!
```

## Run in Docker:
Required
- Install docker

### Run process
Step1: Pull docker image(borischen0203/goris)
```bash
docker pull borischen0203/goris
```
Step2:  Run docker image as below command
```bash
docker run -it --rm borischen0203/goris
```

## Run in Local:

Required
- Install go(version >= 1.6)
- Install `make` cli(https://formulae.brew.sh/formula/make)
```bash
brew install make
```

### Run process
Step1: Clone the repo
```bash
git clone https://github.com/borischen0203/goris.git
```
Step2: Use `make` to execute makefile run build
```bash
make build
```
Step3: Execute build file with or with command
```bash
./goris [command]
```

## Tech stack
- Golang
- Cobra
- Docker
- Github actions
- Shell



### Todo:
- [X] `view` command: Be able to preview gopher image.
- [X] `get link` command: Be able to generate gopher image link with html.

### Reference:
https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43



