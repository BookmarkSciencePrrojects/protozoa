package constants

// Constants
const (
	// Drawing constants
	GridWidth    = 180
	GridHeight   = 120
	GridUnitSize = 6.6667
	ScreenWidth  = 1200
	ScreenHeight = 800

	// Environment constants
	NumFood         = 500
	MaxFoodLifespan = 600

	// Organism constants
	NumInitialOrganisms         = 10
	MaxOrganismsAllowed         = 5000
	InitialHealth               = 100
	MaxHealth                   = 100
	HealthChangePerTurn         = -1
	HealthChangeFromMoving      = -1
	HealthChangeFromEating      = 100
	HealthChangeFromReproducing = -50
	HealthThresholdForEating    = 0

	// Sequence constants
	MaxSequenceNodes  = 30
	MaxNodesToMutate  = 2
	PercentActions    = 0.5
	PercentConditions = 0.5

	// Time trial constants
	OrganismAgeToEndSimulation = 1000
	MaxCyclesToRunHeadless     = 100000

	// Reporting constants
	PopulationDifferenceToReport = 5
)
