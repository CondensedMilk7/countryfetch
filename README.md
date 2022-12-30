# Countryfetch

A cli tool for fetching information about countries. A Go alternative to my original
[countryfetch](https://github.com/CondensedMilk7/countryfetch) which is written in TypeScript (Deno).

This is a work in progress.

## Installation & Usage

Make sure you have [GO](https://go.dev/) installed and run this one-line installer:

````bash
git clone https://github.com/CondensedMilk7/countryfetch-go.git && cd ./countryfetch-go && go build && cp ./countryfetch-go ~/.local/bin/countryfetch
```

If you have the original `countryfetch` and you want to keep it, replace the last argument with something else,
like `countryfetch-go` and use that as a command.

```
USAGE:
  -capital string
    	Find country by given capital
  -name string
    	Find country by given name
  -sync
    	Fetch and save data to cache
EXAMPLE:
  countryfetch -name italy
```

## To do:

- Better error handling
- Flag ASCII art (like in the original)
