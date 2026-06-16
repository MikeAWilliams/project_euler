# Project Euler — Go

Solutions to [Project Euler](https://projecteuler.net) problems written in Go.

## Structure

Each problem lives in its own subdirectory (`p01/`, `p02/`, …) with a `util/` package for shared helpers (primes, big integers, etc.).

## Running

Run a single problem:

```bash
go run ./p01/
```

Run all problems:

```bash
./run_all.sh
```

> **Note:** Problems 13 and 22 read from data files. Place them at `p13/p13_data.txt` and `p22/p022_names.txt` before running.
