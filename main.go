package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// Verifie que le temps ntp et local ne s'ecartent pas plus de 5 sec.
// exit 1 if ntp skew > 5secs, verbose display diffs
func main() {
	// ntp lib use: import "github.com/beevik/ntp"
	ntpServer := flag.String("server", "ntp.uvsq.fr", "NTP server")
	verbose := flag.Bool("v", false, "verbose mode")
	flag.Parse()

	ntpTime, err := ntp.Time(*ntpServer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	now := time.Now()
	delta := ntpTime.Sub(now)

	if *verbose {
		fmt.Printf("Getting Ntp time from %s\n", *ntpServer)
		fmt.Printf("Ntp time\t: %v\n", ntpTime.Format(time.UnixDate))
		fmt.Printf("Local time\t: %v\n", now.Local().Format(time.UnixDate))
		fmt.Printf("Delta is ")
		if delta < time.Second {
			fmt.Printf("%v \n", delta.Round(time.Millisecond))
		} else {
			fmt.Printf("%v \n", delta.Round(time.Second))
		}
	}

	if delta > 5*time.Second {
		fmt.Printf("Deviation Error : Clock skew from ntp.uvsq.fr is %v \n",
			delta.Round(time.Second))
		os.Exit(1)
	}

}
