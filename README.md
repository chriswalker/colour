# colour
[![CircleCI](https://circleci.com/gh/chriswalker/colour/tree/master.svg?style=svg)](https://circleci.com/gh/chriswalker/colour/tree/master)

A simple library for emitting colour escape sequences in command-line apps. tested on MacOS, should work on Linux. Does not support Windows currently.

## Installation

```bash
go get github.com/chriswalker/colour
```

## Usage

You can wrap arbitrary text in any of the standard 16 terminal colours (Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Bright Black, Bright Red and so on) as follows:

```go
fmt.Println("The following text will be blue: ", colour.Blue("hello"))
fmt.Println("The following text will be bright blue: ", colour.BrightBlue("hello"))
// and so on
```

This is a very lightweight colour library, currently only supporting foreground colours as per my needs. Should you need something beefier, try [fatih/color](https://github.com/fatih/color), from which this draws some inspiration.


