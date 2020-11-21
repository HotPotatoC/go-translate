# Go Translate CLI (Google Translate)

Google translate via CLI made with Go

## Installation

```
$ go install github.com/HotPotatoC/go-translate
```

## Overview

```bash
Usage: go-translate [options]
Options:
  -s string
        Source language [en] (default "en")
  -st string
        Text to translate
  -t string
        Target language [id] (default "id")
```

## Usage

```
$ go-translate -st "Hello, World\!"
Halo, Dunia!

$ go-translate -t ja -st "Good morning friend"
おはよう友達

$ go-translate -t no -st "I am going to buy a cup of coffee"
Jeg skal kjøpe en kopp kaffe

$ go-translate -s id -t en -st "Apa kabar bro?"
How are you, brother?
```
