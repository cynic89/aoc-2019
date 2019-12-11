package main

import (
	"bufio"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/cynic89/aoc-2019/util"
	"math"
	"os"
	"sort"
)

const (
	INPUT_FILE_DAY_10 = "inputs/day10"
)

type asteroid struct {
	loc        util.Coordinates
	invisibles map[util.Coordinates]bool
	visibles   map[util.Coordinates]bool
}

type layout struct {
	asteroids     map[util.Coordinates]*asteroid
	width, height int
}

func plotAsteroids() (layout, error) {
	f, err := os.Open(INPUT_FILE_DAY_10)
	if err != nil {
		return layout{}, err
	}

	defer func() { f.Close() }()

	scanner := bufio.NewScanner(f)
	var layout layout
	var x, y int
	layout.asteroids = make(map[util.Coordinates]*asteroid)
	for scanner.Scan() {
		text := scanner.Text()
		x = 0
		for _, c := range text {
			if c == '#' {
				loc := util.Coordinates{x, y}
				layout.asteroids[loc] = &asteroid{loc: loc,
					invisibles: make(map[util.Coordinates]bool), visibles: make(map[util.Coordinates]bool)}
			}
			x++
		}
		y++

	}
	layout.width = x
	layout.height = y
	return layout, nil
}

func (l *layout) calculateView() {

	for _, asteroid := range l.asteroids {
		for i := 1; i <= l.width+1; i++ {
			l.circle(asteroid, i)
		}
	}
}

func (l *layout) calculateViewForAsteroid(a *asteroid) {

	for i := 1; i <= l.width+1; i++ {
		l.circle(a, i)
	}
}

func (l *layout) circle(asteroid *asteroid, radius int) {
	startingLoc := util.Coordinates{asteroid.loc.X + radius, asteroid.loc.Y}
	path := path{radius: radius, center: asteroid.loc, startingLoc: startingLoc, currentLoc: startingLoc}
	for !path.complete {
		if l.outOfBounds(path.currentLoc) {
			path.move()
			continue
		}

		a, ok := l.asteroids[path.currentLoc]
		if !ok {
			path.move()
			continue
		}

		if !asteroid.invisibles[path.currentLoc] {
			asteroid.visibles[path.currentLoc] = true
			l.findInvisibles(asteroid, a)
		}
		path.move()

	}

}

func (l *layout) findInvisibles(self, other *asteroid) {
	i := other.loc.X
	j := other.loc.Y
	diff := calcDiff(self.loc, other.loc)
	for {
		i = i + diff.X
		j = j + diff.Y
		if !(i < l.width && j < l.height && i >= 0 && j >= 0) {
			break
		}
		self.invisibles[util.Coordinates{i, j}] = true
	}
}

func (l *layout) outOfBounds(pos util.Coordinates) bool {
	if pos.X >= l.width || pos.Y >= l.height || pos.X < 0 || pos.Y < 0 {
		return true
	}
	return false
}

type path struct {
	lastMove    string
	radius      int
	currentLoc  util.Coordinates
	startingLoc util.Coordinates
	center      util.Coordinates
	complete    bool
}

func (p *path) move() {

	defer func() {
		if p.currentLoc == p.startingLoc {
			if p.lastMove != "" {
				p.complete = true
			}
			return
		}
	}()

	if p.complete {
		return
	}

	if p.lastMove == "" {
		p.down()
		return
	}

	if p.currentLoc == p.startingLoc {
		if p.lastMove != "" {
			p.complete = true
		}
		return
	}

	if p.lastMove == "down" {
		if p.edge() {
			p.left()
		} else {
			p.down()
		}
		return
	}

	if p.lastMove == "left" {
		if p.edge() {
			p.up()
		} else {
			p.left()
		}
		return
	}

	if p.lastMove == "up" {
		if p.edge() {
			p.right()
		} else {
			p.up()
		}
		return
	}

	if p.lastMove == "right" {
		if p.edge() {
			p.down()
		} else {
			p.right()
		}
		return
	}

}

func (p *path) edge() bool {
	diff := util.Coordinates{p.currentLoc.X - p.center.X, p.currentLoc.Y - p.center.Y}
	return util.Abs(diff.X) == util.Abs(diff.Y)
}

func (p *path) down() {
	p.lastMove = "down"
	p.currentLoc.Y = p.currentLoc.Y + 1

}

func (p *path) left() {
	p.lastMove = "left"
	p.currentLoc.X = p.currentLoc.X - 1
}

func (p *path) up() {
	p.lastMove = "up"
	p.currentLoc.Y = p.currentLoc.Y - 1
}

func (p *path) right() {
	p.lastMove = "right"
	p.currentLoc.X = p.currentLoc.X + 1
}

func calcDiff(src, target util.Coordinates) (diff util.Coordinates) {
	diff = util.Coordinates{target.X - src.X, target.Y - src.Y}
	absDiff := util.Coordinates{util.Abs(diff.X), util.Abs(diff.Y)}

	lcm := util.LCM(absDiff.X, absDiff.Y)
	diff.X = diff.X / lcm
	diff.Y = diff.Y / lcm

	return diff

}

func (l *layout) bestAsteroidLocation() (util.Coordinates, int) {
	var maxVisible int
	var maxVisibleCoordinate util.Coordinates
	for _, a := range l.asteroids {
		if len(a.visibles) > maxVisible {
			maxVisible = len(a.visibles)
			maxVisibleCoordinate = a.loc
		}
	}

	return maxVisibleCoordinate, maxVisible
}

func sortByAngle(a *asteroid) []util.Coordinates {

	var quadrants [4][]util.Coordinates

	for loc, _ := range a.visibles {
		if (a.loc.Angle(loc) <= math.Pi/2) && (a.loc.Angle(loc) > 0) {
			quadrants[0] = append(quadrants[0], loc)
		}

		if (a.loc.Angle(loc) <= 0) && (a.loc.Angle(loc) > -math.Pi/2) {
			quadrants[1] = append(quadrants[1], loc)
		}

		if (a.loc.Angle(loc) <= -math.Pi/2) && (a.loc.Angle(loc) < 0) {
			quadrants[2] = append(quadrants[2], loc)
		}

		if (a.loc.Angle(loc) > math.Pi/2) && (a.loc.Angle(loc) <= math.Pi) {
			quadrants[3] = append(quadrants[3], loc)
		}
	}

	sort.SliceStable(quadrants[0], func(i, j int) bool {
		return math.Pi/2-a.loc.Angle(quadrants[0][i]) < math.Pi/2-a.loc.Angle(quadrants[0][j])
	})

	sort.SliceStable(quadrants[1], func(i, j int) bool {
		return math.Pi/2-a.loc.Angle(quadrants[1][i]) < math.Pi/2-a.loc.Angle(quadrants[1][j])
	})

	sort.SliceStable(quadrants[2], func(i, j int) bool {
		return math.Pi/2-a.loc.Angle(quadrants[2][i]) < math.Pi/2-a.loc.Angle(quadrants[2][j])
	})

	sort.SliceStable(quadrants[3], func(i, j int) bool {
		return math.Pi/2-a.loc.Angle(quadrants[3][i]) < math.Pi/2-a.loc.Angle(quadrants[3][j])
	})

	allQ := append(append(append(quadrants[0], quadrants[1]...), quadrants[2]...), quadrants[3]...)
	return allQ

}

func (l *layout) vapourize(a *asteroid, n int) []util.Coordinates {
	var vaporizedAsteroids []util.Coordinates
	for len(l.asteroids) > 1 {

		vaporizedAsteroids = append(vaporizedAsteroids, sortByAngle(a)...)
		a.visibles = make(map[util.Coordinates]bool)
		a.invisibles = make(map[util.Coordinates]bool)
		for _, loc := range vaporizedAsteroids {
			//l.render()
			delete(l.asteroids, loc)
		}
		l.calculateViewForAsteroid(a)
	}
	return vaporizedAsteroids
}

func (l *layout) render() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	var x, y int
	for y < l.height {
		x = 0
		for x < l.width {
			if l.asteroids[util.Coordinates{x, y}] != nil {
				tm.Printf("#")
			} else {
				tm.Printf(".")
			}
			x++
		}
		tm.Println()
		y++
	}
	tm.Flush()
}

func day10() {
	l, err := plotAsteroids()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	l.calculateView()
	loc, count := l.bestAsteroidLocation()
	fmt.Printf("Loc = %d, count = %d\n", loc, count)

	vaporizedAsteroids := l.vapourize(l.asteroids[loc], 200)
	fmt.Printf("The 200th vaporized asteroid is at %v\n", vaporizedAsteroids[199])
}
