package main

import (
	"flag"
	"fmt"

	"github.com/jasonmccallister/hosts"
)

func main() {
	ip := flag.String("ip", "", "the IP to find")
	hostname := flag.String("hostname", "", "hostname to find")
	flag.Parse()

	// show records by ip
	if *ip != "" {
		records, err := hosts.FindIP(*ip)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("searching for ip", *ip)

		for _, record := range records {
			fmt.Println(record)
		}
	}

	if *hostname != "" {
		hostRecords, err := hosts.FindHost(*hostname)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("searching for host", *hostname)

		for _, hr := range hostRecords {
			fmt.Println(hr)
		}
	}
}
