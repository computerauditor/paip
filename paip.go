package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	var ipString string

	// Check if data is being piped in
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Read from stdin if data is being piped in
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ipString = scanner.Text()
		}
	} else {
		// Check if an IP address is provided as a command line argument
		if len(os.Args) != 2 {
			fmt.Println("Usage: oppa.go <IP>")
			os.Exit(1)
		}
		ipString = os.Args[1]
	}

	// Parse the IP address
	ip := net.ParseIP(ipString)
	if ip == nil {
		fmt.Println("Invalid IP address")
		os.Exit(1)
	}

	// Print the subnet from the given IP address up to 255
	for i := 0; i <= 255; i++ {
		subnetIP := net.IPv4(ip[12], ip[13], ip[14], byte(i))
		fmt.Println(subnetIP)
	}
}
