package nbody

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

type Body struct {
	x, y, z    float32 // POSITIONS
	vx, vy, vz float32 // VELOCITIES
}

// return a new body
func NewBody() *Body {
	return &Body{}
}

// write to csv
func ParticlePositionsToCSV(file *os.File, iteration int,
	bodies []*Body, numBodies int) {
	for i := 0; i < numBodies; i++ {
		_, err := fmt.Fprintf(file, "%d, %e, %e, %e\n",
			iteration, bodies[i].x, bodies[i].y, bodies[i].z)

		if err != nil {
			fmt.Println("ERROR WHEN WRITING TO FILE \"nbody.csv\"")
			panic(err)
		}
	}
}

// initialize n bodies with random positions and velocities
func InitPositionsAndVelocities(id int, bodies []*Body, numBodies int) {
	random := func(a, b float32) float32 {
		return a + rand.Float32()*b
	}

	bodies[id] = NewBody()

	if id%3 == 0 {
		bodies[id].x = -1000.0 + random(-2.2, 3.3)
		bodies[id].y = 0.0 + random(-2.2, 3.3)
		bodies[id].z = 0.0 + random(-2.2, 3.3)
	} else if id%3 == 1 {
		bodies[id].x = 0.0 + random(-2.2, 3.3)
		bodies[id].y = 0.0 + random(-2.2, 3.3)
		bodies[id].z = -1000.0 + random(-2.2, 3.3)
	} else {
		bodies[id].x = 0.0 + random(-2.2, 3.3)
		bodies[id].y = 1000.0 + random(-2.2, 3.3)
		bodies[id].z = 0.0 + random(-2.2, 3.3)
	}

	bodies[id].vx = 0.0
	bodies[id].vy = 0.0
	bodies[id].vz = 0.0
}

// compute interbody forces
func ComputeBodyForce(id int, bodies []*Body, dt float32,
	numBodies int, softeningFactor float32) {
	var Fx, Fy, Fz float32
	for j := 0; j < numBodies; j++ {
		dx := bodies[j].x - bodies[id].x
		dy := bodies[j].y - bodies[id].y
		dz := bodies[j].z - bodies[id].z

		distSqr := dx*dx + dy*dy + dz*dz + softeningFactor
		invrDist := float32(1.0 / math.Pow(float64(distSqr), 0.5))
		invrDist3 := invrDist * invrDist * invrDist

		Fx += dx * invrDist3
		Fy += dy * invrDist3
		Fz += dz * invrDist3
	}

	bodies[id].vx += dt * Fx
	bodies[id].vy += dt * Fy
	bodies[id].vz += dt * Fz

}

// integrate postions
func IntegratePositions(id int, bodies []*Body, numBodies int, dt float32) {
	bodies[id].x += bodies[id].vx * dt
	bodies[id].y += bodies[id].vy * dt
	bodies[id].z += bodies[id].vz * dt
}
