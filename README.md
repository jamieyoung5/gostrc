# gostrc

[![Ci](https://github.com/jamieyoung5/gostrc/actions/workflows/test.yml/badge.svg)](https://github.com/jamieyoung5/gostrc/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/jamieyoung5/gostrc/graph/badge.svg?token=YOUR_TOKEN_ID)](https://codecov.io/gh/jamieyoung5/gostrc)

This repository is a compilation of various algorithm and data structure implementations in Go, developed during university or for personal projects.

## Implementations

This collection currently includes:

- **Dancing Links (Algorithm X):** An implementation of Knuth's Algorithm X using the Dancing Links (DLX) technique for solving exact cover problems. (Found in `dlx/dlx.go`, `dlx/column.go`)
- **Thread-Safe Stack:** A generic, concurrent-safe Stack implementation. (Found in `stack.go`)
- **Thread-Safe Circular Queue:** A generic, concurrent-safe, fixed-size circular queue. (Found in `circular_queue.go`)
- **Slice Utilities:** A set of generic helper functions for common slice operations, such as:
    - `RandomSubset`
    - `Reverse`
    - `Equal`
    - `CountDuplicates`
    - `Counts` (Found in `sliceutil/sliceutil.go`)
- **String Utilities:** helper functions for common string related operations, such as:
    - `MaxLen` (Found in `strutil/strutil.go`)

## Installation

```bash
go get https://github.com/jamieyoung5/gostrc
```
