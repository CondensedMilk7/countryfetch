# Countryfetch

A cli tool for fetching information about countries. A Go alternative to my original
[countryfetch](https://github.com/CondensedMilk7/countryfetch) which is written in TypeScript (Deno).

This is a work in progress.
![](./media/countryfetch-go.jpg)

# Installation & Usage

Make sure you have [GO](https://go.dev/) installed and run this one-line installer:

```bash
git clone https://github.com/CondensedMilk7/countryfetch-go.git && cd ./countryfetch-go && go install
```

If you have the original `countryfetch` and you want to keep it, do this instead:

```bash
git clone https://github.com/CondensedMilk7/countryfetch-go.git && cd ./countryfetch-go && go build -o countryfetch-go && cp ./countryfetch-go ~/.local/bin/

```

```
USAGE:
  -capital string
    	Find country by given capital.
  -flag
    	Include ASCII flag in the operation. Can be used in combination with -sync & -name.
  -name string
    	Find country by given name.
  -sync
    	Fetch and save data to cache.
EXAMPLE:
    countryfetch -name italy -flag
	    Fetch information about Italy, including its flag.
    countryfetch -sync -flag
	    Store information of all countries in cache, including generated flag ASCII art.
    countryfetch -capital "kuala lumpur"
	    Fetch information about the country of given capital
```

Note that string arguments with multiple words have to be enclosed in brackets.

# Tasks

- Fetch flags remotely (directly, not from stored cache)
- Fetch just the flag for a given country
