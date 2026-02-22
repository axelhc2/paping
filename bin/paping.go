package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
)

func main() {
	if len(os.Args) < 4 || os.Args[2] != "-p" {
		fmt.Println("Usage: paping <IP> -p <PORT>")
		return
	}

	target := os.Args[1]
	port := os.Args[3]
	address := net.JoinHostPort(target, port)

	fmt.Printf("paping v1.0.0 - Credit AXEL CHETAIL https://infrawire.fr\n\n")
	fmt.Printf("Connecting to %s on TCP %s:\n\n", target, port)

	var attempted, connected, failed int
	var min, max, sum float64

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if attempted > 0 {
			avg := 0.0
			if connected > 0 { avg = sum / float64(connected) }
			fmt.Printf("\nConnection statistics:\n")
			fmt.Printf("\tAttempted = %d, Connected = %d, Failed = %d (%.2f%%)\n", 
				attempted, connected, failed, (float64(failed)/float64(attempted))*100)
			fmt.Printf("Approximate connection times:\n")
			fmt.Printf("\tMinimum = %.2fms, Maximum = %.2fms, Average = %.2fms\n", min, max, avg)
		}
		os.Exit(0)
	}()

	for {
		attempted++
		start := time.Now()
		
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		duration := float64(time.Since(start).Microseconds()) / 1000

		if err != nil {
			failed++
			fmt.Printf("%sConnection timed out%s\n", colorRed, colorReset)
		} else {
			connected++
			// Coloration spécifique demandée : IP, MS, Protocol et Port en vert uniquement
			fmt.Printf("Connected to %s%s%s: time=%s%.2fms%s protocol=%sTCP%s port=%s%s%s\n", 
				colorGreen, target, colorReset, 
				colorGreen, duration, colorReset, 
				colorGreen, colorReset, 
				colorGreen, port, colorReset)
			conn.Close()

			if min == 0 || duration < min { min = duration }
			if duration > max { max = duration }
			sum += duration
		}

		elapsed := time.Since(start)
		if elapsed < time.Second {
			time.Sleep(time.Second - elapsed)
		}
	}
}
