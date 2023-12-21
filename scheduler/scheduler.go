package scheduler

type Config struct {
	Mode string // Represents which scheduler scheme to use
	// If Mode == "s" run the sequential version
	// If Mode == "ws" run the work-stealing parallel version
	// If Mode == "wb" run the work-balancing parallel version
	// These are the only values for Version
	NBodies         int    // Number of Particles
	Iterations      int    // Number of iterations to simulate
	RecordPositions string // Record positions of the Bodies in a csv file
	// If RecordPositions = "yes" record positions
	// Or else don't record positions
	ThreadCount int // Number of go routines for the parallel versions
}

// Run the correct version based on the Mode field of the configuration value
func Schedule(config Config) {
	if config.Mode == "s" {
		RunSequential(
			config.NBodies,
			config.Iterations,
			0.01,
			config.RecordPositions,
		)
	} else if config.Mode == "ws" || config.Mode == "wb" {
		RunParallel(
			config.NBodies,
			config.Iterations,
			0.01,
			config.RecordPositions,
			config.ThreadCount,
			config.Mode,
		)
	} else {
		panic("Invalid scheduling scheme: " + config.Mode)
	}
}
