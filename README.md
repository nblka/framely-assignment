# Framely assignment Go project

## Disclaimer

This is my very first Go project, so it's probably not very elegant and this doc might be incomplete. Hope you have enough practice in Go to get it work.

## Task description

Please visit [GO_ASSIGNMENT.md](GO_ASSIGNMENT.md)

## Getting the code

```sh
git clone https://github.com/nblka/framely-assignment.git
cd framely-assignment
```

## Repo structure

The `speedtester` directory contains library code.

The `example` directory contains an example of client which uses the library.

## Library API

To use `speedtester` library, import the `github.com/nblka/framely-assignment/speedtester` package:

```go
import "github.com/nblka/framely-assignment/speedtester"
```

To perform speed testing, use the `Test` method:

```go
speedtester.Test(speedtester.SpeedTest)         // for Speedtest.net service
speedtester.Test(speedtester.NetflixFastCom)    // for fast.com service
```

The `Test` method returns a tuple of the Measure struct holding download and upload speeds measured, and an error, if any. There's a helper function to print Measure struct data in a formatted way.

NOTE: The fast.com service measures only the downloading speed, so the uploading speed is set to 0.

You can also refer to example or tests to find out the API usage.

## Running an example

From the root directory:

```sh
cd example
git run example
```

You will probably need to get imported libraries before running an example by using `go get -u` or something like that.

## Running tests

From the root directory:

```sh
cd speedtester
go test -v
```

The test coverage (determined by `go test -cover`) eventually is very low, about 11,7%. I'm not sure why it's so, maybe because the majority of the code is involved in working with imported libraries which I didn't test. Though there's no much code to test at all.

NOTE: The SpeedTest service powered by `speedtest-go` module is not very stable and occasionally fails with "unexpected EOF" error, at least on my machine.

## Running benchmarks

If you are on Windows like me:

```sh
cd speedtester
go test -v -run="none" -bench="."
```

On Linux you can try

```sh
go test -v -run=none -bench=.
```
