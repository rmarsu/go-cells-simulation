package main

import (
	dingo "3d"
	"fmt"
	"math/rand/v2"
	"time"

)

var (
	WIDTH  = 20
     HEIGHT = 20
     
	CELLS_PER_RUN = 20
	CELLS_LIFETIME = 7
	CELLS_DISTANCE = 3
	TIME_BEFORE_HUNGRY = 5
)

type World struct {
	cells [][]*Cell
}

type Cell struct {
	alive bool

	id     int64
	x, y   int
	gender int
	hungry bool
}

func (w *World) GenerateNewWorld() {
	var id int64
	w.cells = make([][]*Cell, HEIGHT)
	for i := range w.cells {
		w.cells[i] = make([]*Cell, WIDTH)
		for j := range w.cells[i] {
			w.cells[i][j] = &Cell{alive: false, id: id, x: i, y: j, gender: 0}
			id++
		}
	}
}

func (w *World) CreateCell(posX int, posY int, gender int) *Cell {
	w.cells[posX][posY].alive = true
	w.cells[posX][posY].gender = gender
	w.cells[posX][posY].hungry = false
	go w.Death(w.cells[posX][posY].id)
	go w.MakeHungry(w.cells[posX][posY].id)
	return w.cells[posX][posY]
}
func PrintEnviroment(w *World) {
	for _, row := range w.cells {
		for _, cell := range row {
			if cell.alive {
				if cell.gender == 1 {
					fmt.Print("Y")
					} else {
					fmt.Print("X")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println() // Newline for better readability.
}

func (w *World) Death(id int64) {
	time.Sleep(time.Second * time.Duration(CELLS_LIFETIME))
	for _, row := range w.cells {
		for _, cell := range row {
			if cell.alive {
				if cell.id == id {
					cell.alive = false
					break
				}
			}
		}

	}
}

func (w *World) MakeHungry(id int64) {
	time.Sleep(time.Second * time.Duration(TIME_BEFORE_HUNGRY))
	for _, row := range w.cells {
		for _, cell := range row {
			if cell.alive {
				if cell.id == id {
					cell.hungry = true
					break
				}
			}
		}

	}
}

func (w *World) CellsMove() {
	for _, row := range w.cells {
		for _, cell := range row {
			if cell.alive {
				nextX, nextY := cell.x+(rand.IntN(3)-1), cell.y+(rand.IntN(3)-1)
				if nextX >= 2 && nextX < HEIGHT-2 && nextY >= 2 && nextY < WIDTH-2 && w.cells[nextX][nextY].alive == false {
					nextCellID := w.cells[nextX][nextY].id
					w.cells[nextX][nextY].alive = true
					w.cells[nextX][nextY].gender = cell.gender
					w.cells[cell.x][cell.y].alive = false
					w.cells[nextX][nextY].id = cell.id
					w.cells[cell.x][cell.y].id = nextCellID
					go w.Death(w.cells[nextX][nextY].id)
					go w.MakeHungry(w.cells[nextX][nextY].id)
				}
			}
		}
	}
}

func (w *World) CheckCellsAround() {
	// Count of alive cells around the current cell.
	fertilization := false
	
	for _, cell := range w.GetAliveCellsPositions() {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if cell.hungry && w.cells[cell.x+i][cell.y+j].alive {
					w.cells[cell.x][cell.y].alive = false
					w.CreateCell(cell.x+i, cell.y+j, cell.gender)
				}
				if cell.x+i >= 2 && cell.x+i < HEIGHT-2 && cell.y+j >= 2 && cell.y+j < WIDTH-2 {
					count := 0
					if w.cells[cell.x+i][cell.y+j].alive {
						count += 1
						if cell.gender != w.cells[cell.x+i][cell.y+j].gender {
							fertilization = true
						}
					}
					
					if (fertilization) && (!w.cells[cell.x+i][cell.y+j].alive) {
						w.CreateCell(cell.x+i, cell.y+j, rand.IntN(2))
						fertilization = false
					}
				}
			}
		}
	}
}

func (w *World) GetAliveCellsPositions() []*Cell {
	var aliveCells []*Cell
	for _, row := range w.cells {
		for _, cell := range row {
			if cell.alive {
				aliveCells = append(aliveCells, cell)
			}
		}

	}
	return aliveCells
}

func main() {
	Config , err := dingo.Initialize()
	if err!= nil {
          fmt.Println("Error initializing config:", err)
     }
	HEIGHT = Config.Height
	WIDTH = Config.Width
	CELLS_LIFETIME = Config.CellsLifetime
	TIME_BEFORE_HUNGRY = Config.CellsTimeHungry
	CELLS_DISTANCE = Config.CellsDistance
	CELLS_PER_RUN = Config.CellsPerRun
	

	world := &World{}
	world.GenerateNewWorld()
	for i := 0; i < CELLS_PER_RUN; i++ {
		world.CreateCell(dingo.RandRange(CELLS_DISTANCE, HEIGHT-CELLS_DISTANCE), dingo.RandRange(CELLS_DISTANCE, WIDTH-CELLS_DISTANCE), rand.IntN(2))
	}

	for {
		world.CellsMove()
		world.CheckCellsAround()
		time.Sleep(time.Second)
		fmt.Print("\033[H\033[2J")
		PrintEnviroment(world) // Print the initial state of the world.
	}
}
