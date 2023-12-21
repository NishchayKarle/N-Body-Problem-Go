package scheduler

import (
	"fmt"
	"os"
	"proj3/nbody"
)

func RunSequential(numBodies, iterations int, dt float32, record string) {
	bodies := make([]*nbody.Body, numBodies)

	var file *os.File
	var err error

	// WRITE POSITIONS
	if record == "yes" {
		file, err = os.Create("../scheduler/sequential/nbody.csv")
		if err != nil {
			fmt.Println("ERROR WHEN OPENING FILE \"nbody.csv\"")
			panic(err)
		}
		defer file.Close()
	}

	for i := 0; i < numBodies; i++ {
		nbody.InitPositionsAndVelocities(i, bodies, numBodies)
	}

	for iter := 0; iter < iterations+1; iter++ {
		if record == "yes" {
			if iterations%(iterations/10) == 0 {
				nbody.ParticlePositionsToCSV(file, iter/(iterations/10), bodies, numBodies) // WRITE POSITIONS AFTER ITERATION
			}
		}

		for i := 0; i < numBodies; i++ {
			nbody.ComputeBodyForce(i, bodies, dt, numBodies, 1e-4) // COMPUTE INTERBODY FORCES
		}

		for i := 0; i < numBodies; i++ {
			nbody.IntegratePositions(i, bodies, numBodies, dt) // INTEGRATE POSITIONS
		}
	}
}
