package dingo

type Config struct {
	Width  int
	Height int

	CellsPerRun      int
	CellsLifetime     int
	CellsDistance     int
	CellsTimeHungry int
}