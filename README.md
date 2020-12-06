# gofinstats [![GoDoc](https://godoc.org/github.com/yaustn/gofinstats?status.svg)](https://godoc.org/github.com/yaustn/gofinstats) [![Go Report Card](https://goreportcard.com/badge/github.com/yaustn/gofinstats)](https://goreportcard.com/report/github.com/yaustn/gofinstats)

A golang package with different statistic calculators, generally used in financial contexts.

## Usage

Import the package
```
go get -u github.com/yaustn/gofinstats
```

## Simple Moving Average

> A **simple moving average** (SMA) is the unweighted mean of the previous *n* data.
>
> \- [_from Wikipedia_](https://en.wikipedia.org/wiki/Moving_average#Simple_moving_average)

Notes: 
- O(1) calculations
- float64 is not very precise
- Not threadsafe
