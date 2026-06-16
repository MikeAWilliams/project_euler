# Project Euler — Go

Solutions to [Project Euler](https://projecteuler.net) problems written in Go.

## Structure

Each problem lives in its own subdirectory (`p01/`, `p02/`, …) with a `util/` package for shared helpers (primes, big integers, etc.).

## Running

Run a single problem:

```bash
cd p01 && go run .
```

Run all problems:

```bash
./run_all.sh
```

> **Note:** Problems 13 and 22 read from data files (`p13_data.txt`, `p022_names.txt`). Place those files in their respective subdirectories before running.
