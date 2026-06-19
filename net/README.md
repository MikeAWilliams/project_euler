# Project Euler Solutions in C#/.NET

Solutions to Project Euler problems implemented in C# using .NET.

## Structure

Each problem has its own folder (e.g., `p01/`, `p02/`, etc.) with:
- `Program.cs` - Solution code
- `p0X.csproj` - Project file

Shared helper code lives in `Shared/` (namespace `Shared`), a class library
that problems can reference instead of duplicating logic. 

## Running a Solution

```bash
make run 01
```

## Creating a New Problem

```bash
make new 02         # new console project
make new 02 shared  # new console project that references the shared library
```

## Linking an Existing Problem to the Shared Library

```bash
make link 02
```

## Running Tests

Tests for the shared library live in `Shared.Tests/` (xUnit).

```bash
make test
```

## Requirements

- .NET 10.0 or compatible version
