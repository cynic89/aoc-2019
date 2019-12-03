package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	INPUT_FILE_3 = "inputs/day3"
)

type position struct {
	x, y int
}

type move struct {
	direction string
	distance  int
}

type wire struct {
	id              string
	moves           []move
	currentPosition position
	totalDistance   int
}

var pathWire = make(map[position][]*wire)
var pathDistance = make(map[position][]int)

func readWire() ([]wire, error) {
	f, err := os.Open(INPUT_FILE_3)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var wires []wire
	scanner := bufio.NewScanner(f)
	var count int
	for scanner.Scan() {
		var moves []move
		text := scanner.Text()
		vals := strings.Split(text, ",")
		for _, val := range vals {
			dir := val[:1]
			distStr := val[1:]
			dist, _ := strconv.Atoi(distStr)
			moves = append(moves, move{direction: dir, distance: dist})
		}

		wires = append(wires, wire{id: fmt.Sprintf("wire%d", count), moves: moves})
		count++
	}
	return wires, nil
}

func trace(w wire) {
	for _, move := range w.moves {
		currPos := position{w.currentPosition.x, w.currentPosition.y}

		switch move.direction {
		case "R":
			lim := currPos.x + move.distance
			for i := currPos.x; i < lim; i++ {
				w.totalDistance = w.totalDistance + 1
				w.currentPosition.x = w.currentPosition.x + 1
				updatePath(&w)
			}

		case "L":
			lim := w.currentPosition.x - move.distance
			for i := currPos.x; i > lim; i-- {
				w.totalDistance = w.totalDistance + 1
				w.currentPosition.x = w.currentPosition.x - 1
				updatePath(&w)
			}
		case "U":
			lim := w.currentPosition.y + move.distance
			for i := currPos.y; i < lim; i++ {
				w.totalDistance = w.totalDistance + 1
				w.currentPosition.y = w.currentPosition.y + 1
				updatePath(&w)
			}
		case "D":
			lim := w.currentPosition.y - move.distance
			for i := currPos.y; i > lim; i-- {
				w.totalDistance = w.totalDistance + 1
				w.currentPosition.y = w.currentPosition.y - 1
				updatePath(&w)
			}
		}
	}
}

func updatePath(w *wire) {
	pos := w.currentPosition
	wires, ok := pathWire[pos]
	if !ok {
		pathWire[pos] = []*wire{w}
		pathDistance[pos] = []int{w.totalDistance}
	} else {
		pathWire[pos] = append(wires, w)
		pathDistance[pos] = append(pathDistance[pos], w.totalDistance)
	}
}

func intersectionWithMinManhattanDistance() int {
	var minDistance int
	for path, wires := range pathWire {
		if len(wires) == 2 && wires[0].id != wires[1].id {

			d := abs(path.x) + abs(path.y)

			if minDistance == 0 {
				minDistance = d
			}
			if d < minDistance {
				minDistance = d
			}
		}
	}

	return minDistance
}

func intersectionWithMinTotalTravelledSteps() int {
	var minDistance int
	for p, wires := range pathWire {
		if len(wires) == 2 && wires[0].id != wires[1].id {

			d := pathDistance[p][0] + pathDistance[p][1]

			if minDistance == 0 {
				minDistance = d
			}
			if d < minDistance {
				minDistance = d
			}
		}
	}

	return minDistance
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func main() {

	wires, err := readWire()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	trace(wires[0])
	trace(wires[1])

	dPart1 := intersectionWithMinManhattanDistance()
	fmt.Println(dPart1)

	dPart2 := intersectionWithMinTotalTravelledSteps()
	fmt.Println(dPart2)

}
