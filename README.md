# Framey assignment Go project

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

NOTE: The fast.com service library (`go-fast`) don't return one speed, instead it returns several speed values in series. To get one value from these, I calculate a median.

You can also refer to example or tests to find out the API usage.

## Running an example

From the root directory:

```sh
cd example
git run example
```

You will probably need to get imported libraries before running an example by using `go get -u` or something like that.

<details>
  <summary>Sample example output (click to expand)</summary>
  
```sh
> go run example
speedtest service run
User *.*.*.*, (ER-Telecom) [*.*, *.*]
Servers {[[5740]     1.82km 
Tomsk (Russia) by Rostelecom
 [20145]   143.13km
Kemerovo (Russia) by MegaTelecom
 [14402]   209.47km
Novosibirsk (Russia) by EDINOS
 [31968]   213.21km
Novosibirsk (Russia) by ADMAN LLC
 [6430]   250.49km
Novosibirsk (Russia) by Tele2 Russia
 [46230]   250.49km
Novosibirsk (Russia) by INLINE LTD
 [14386]   284.19km
Sharypovo (Russia) by Sibline
 [10311]   295.93km
Kiselevsk (Russia) by Electron-Service Ltd.
 [37192]   354.86km
Nauchny (Russia) by FGBUN CrAO RAN
 [1833]   357.61km
Barnaul (Russia) by JSC Zap-Sib TransTeleCom
]}
Targets [[5740]     1.82km 
Tomsk (Russia) by Rostelecom
]
Starting download speed test... done
Starting upload speed test... done
Downloading speed 39.27868439851729 Mbps, uploading speed 37.43262930940951 Mbps <nil>
fast.com service run
Got urls: [https://ipv4-c003-rix001-retn-isp.1.oca.nflxvideo.net/speedtest?c=ru&n=5670ipv4-c002-mow001-retn-isp.1.oca.nflxvideo.net/speedtest?c=ru&n=56707&v=29&e=1641998123tn-isp.1.oca.nflxvideo.net/speedtest?c=ru&n=56707&v=23&e=1641998123&t=CaglHJJQjmF806D7deo.net/speedtest?c=ru&n=56707&v=30&e=1641998123&t=lvXJhocFwHHV5hffElXKPx6uQGwjKzWKMZi=ru&n=56707&v=14&e=1641998123&t=e66_tNCQ0hKLzpfwHzUMX4y-cvyl6kFRYp0r1Q]
Starting speed test...
Got 44.99 Mbps
Got 42.19 Mbps
Got 40.84 Mbps
Got 41.00 Mbps
Got 40.93 Mbps
Got 40.54 Mbps
Got 40.69 Mbps
Got 40.59 Mbps
Got 40.53 Mbps
Got 40.35 Mbps
Got 40.40 Mbps
Got 40.35 Mbps
Got 40.26 Mbps
done
Downloading speed 40.686933333333336 Mbps, uploading speed 0 Mbps <nil>
```

</details>

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
