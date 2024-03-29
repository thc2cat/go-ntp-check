package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

var Version = "go-ntp-check-v0.4" // Also defined via Makefile and git tag

// History :
// v0.4 : using generics Abs

// Verify that local time does not differs much from ntp server
// exit 1 if ntp skew >  delta*scale values, verbose display time diffs

func main() {
	// ntp lib use: import "github.com/beevik/ntp"
	ntpServer := flag.String("server", "time.cloudflare.com", "NTP server")
	verbose := flag.Bool("v", false, "verbose mode")
	scale := flag.String("scale", "s", "skew scale [s|ms]")
	deltat := flag.Int("delta", 2, "max skew in scale units")

	flag.Parse()

	_, _ = ntp.Time(*ntpServer)
	_ = time.Now()

	ntpTime, err := ntp.Time(*ntpServer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	now := time.Now()
	delta := ntpTime.Sub(now)

	var scaleunit time.Duration
	switch *scale {
	case "ms":
		scaleunit = time.Millisecond
	default:
		*scale = "s"
		scaleunit = time.Second
	}

	if *verbose {
		fmt.Printf("%s getting Ntp time from %s\n", Version, *ntpServer)
		fmt.Printf("Ntp time\t: %v\n", ntpTime.Format(time.UnixDate))
		fmt.Printf("Local time\t: %v\n", now.Local().Format(time.UnixDate))
		fmt.Printf("Delta max set to %d%s, current is %v\n", *deltat, *scale, delta.Round(time.Millisecond))
	}

	if Abs(delta) > (time.Duration)(*deltat)*scaleunit {
		fmt.Printf("Local clock is ntp-desynchronised from %s : delta is %v \nS",
			*ntpServer,
			delta.Round(time.Millisecond))
		os.Exit(1)
	} else {
		if *verbose {
			fmt.Printf("Status : OK\n")
		}
	}
}

// Testing generics Abs
type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Abs returns the absolute value of x.
func Abs[T number](x T) T {
	if x < 0 {
		return -x
	}
	return x

}

// Faster abs http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
// func abs(n int64) int64 {
// 	y := n >> 63
// 	return (n ^ y) - y
// }
