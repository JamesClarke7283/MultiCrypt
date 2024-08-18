package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func runBenchmark(binary string) time.Duration {
	start := time.Now()
	cmd := exec.Command(binary)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error starting the application: %v\n", err)
		os.Exit(1)
	}

	// Wait for the application to fully start (adjust this time as needed)
	time.Sleep(2 * time.Second)

	err = cmd.Process.Kill()
	if err != nil {
		fmt.Printf("Error killing the application: %v\n", err)
		os.Exit(1)
	}

	return time.Since(start)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run pgo_benchmark.go <non-pgo-binary> <pgo-binary>")
		os.Exit(1)
	}

	nonPGOBinary := os.Args[1]
	pgoBinary := os.Args[2]

	fmt.Println("Running benchmark for non-PGO binary...")
	nonPGOTime := runBenchmark(nonPGOBinary)

	fmt.Println("Running benchmark for PGO binary...")
	pgoTime := runBenchmark(pgoBinary)

	timeDiff := nonPGOTime - pgoTime
	percentChange := (float64(timeDiff) / float64(nonPGOTime)) * 100

	fmt.Printf("\nResults:\n")
	fmt.Printf("Non-PGO startup time: %v\n", nonPGOTime)
	fmt.Printf("PGO startup time: %v\n", pgoTime)
	fmt.Printf("Time difference: %v\n", timeDiff)
	fmt.Printf("Percentage change: %.2f%%\n", percentChange)

	if timeDiff > 0 {
		fmt.Printf("PGO version started %.2f%% faster\n", percentChange)
	} else {
		fmt.Printf("PGO version started %.2f%% slower\n", -percentChange)
	}
}
