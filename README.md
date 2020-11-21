# Go Translate CLI (Google Translate)

Google translate via CLI made with Go

## Overview

```bash
Usage: translate [options]
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
$ go run main.go -st "Hello, World\!"
Halo, Dunia!

$ go run main.go -t ja -st "Good morning friend"
おはよう友達

$ go run main.go -t no -st "I am going to buy a cup of coffee"
Jeg skal kjøpe en kopp kaffe

$ go run main.go -s id -t en -st "Apa kabar bro?"
How are you, brother?
```
