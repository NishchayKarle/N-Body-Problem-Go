package main

import (
	"fmt"
	"os"
	"proj3/scheduler"
	"strconv"
	"time"
)

const usage = "USAGE: go run editor.go -m <mode: \"s\" or \"ws\" or \"wb\"> -n <number of bodies> " +
	"-i <number of timesteps> -r <record positions> -t <number of threads> " +
	"-p <print config to console>" +
	"\n Minimum value for number of bodies is 2000 and iterations is 10"

func main() {
	mode := "s"
	numBodies := 10_000
	iterations := 100
	recordPositions := "no"
	threadCount := 64
	printConfigToConsole := false
	var err error

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-m" {
			mode = os.Args[i+1]
			i++
		} else if os.Args[i] == "-n" {
			numBodies, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				fmt.Println("Invalid value for number of bodies given")
				panic(err)
			}
			if numBodies < 2000 {
				panic("Minimum value for number of bodies is 2000")
			}
			i++
		} else if os.Args[i] == "-i" {
			iterations, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				fmt.Println("Invalid value for number of iterations given")
				panic(err)
			}
			if iterations < 10 {
				panic("Minimum value for number of iterations is 10")
			}
			i++
		} else if os.Args[i] == "-r" {
			recordPositions = "yes"
		} else if os.Args[i] == "-t" {
			threadCount, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				fmt.Println("Invalid value for number of threads given")
				panic(err)
			}
			i++
		} else if os.Args[i] == "-p" {
			printConfigToConsole = true

		} else {
			fmt.Println("INVALID COMMAND LINE ARGUMENT GIVEN")
			panic(usage)
		}
	}

	if printConfigToConsole {
		fmt.Println("\nRUNNING N-BODY SIMULATION WITH CONFIGURATION:")
		fmt.Println("---------------------------------------------")
		fmt.Println("MODE			: ", mode)
		fmt.Println("NUMBER OF BODIES	: ", numBodies)
		fmt.Println("NUMBER OF TIMESTEPS	: ", iterations)
		fmt.Println("RECORD POSITIONS IN CSV	: ", recordPositions)
		if mode != "s" {
			fmt.Println("NUMBER OF THREADS	: ", threadCount)
		}
		fmt.Println("---------------------------------------------")
	}

	var config scheduler.Config
	config.Mode = mode
	config.NBodies = numBodies
	config.Iterations = iterations
	config.RecordPositions = recordPositions
	config.ThreadCount = threadCount

	start := time.Now()
	{
		scheduler.Schedule(config)
	}
	totalTime := time.Since(start).Seconds()
	avgTime := totalTime / float64(iterations)

	fmt.Printf("TOTAL TIME: %.5fs, AVG TIME: %.5fs\n", totalTime, avgTime)
	if printConfigToConsole {
		fmt.Println("---------------------------------------------")
	}
}
