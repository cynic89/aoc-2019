package main

import (
	"bufio"
	"fmt"
	"github.com/cynic89/aoc-2019/util"
	"os"
	"sort"
)

const (
	INPUT_FILE_DAY_18       = "inputs/day18"
	INPUT_FILE_DAY_18_PART2 = "inputs/day18_part2"
)

type tunnel struct {
	layout        map[util.Coordinates]int32
	keys          map[int32]bool
	entraces      []util.Coordinates
	keysCollected map[int32]bool
	keyLocs       map[int32]util.Coordinates
	size          util.Coordinates
}

type passage struct {
	steps    int
	keys     map[int32]bool
	pos      util.Coordinates
	from     util.Coordinates
	quadrant int
}

type passageId struct {
	pos  util.Coordinates
	from util.Coordinates
	keys string
}

func NewPassage(steps int, keys map[int32]bool, pos, from util.Coordinates, quadrant int) passage {
	return passage{steps: steps, keys: keys, pos: pos, from: from, quadrant: quadrant}
}

func GetPassageId(p passage) passageId {
	psgId := passageId{pos: p.pos, from: p.from}
	var keys []int32
	for k, _ := range p.keys {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var keysStr string
	for _, k := range keys {
		keysStr = keysStr + fmt.Sprintf("%c", k)
	}
	psgId.keys = keysStr
	return psgId
}

func readTunnelLayout(path string) (tunnel, error) {
	f, err := os.Open(path)
	if err != nil {
		return tunnel{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var tunnel = tunnel{layout: make(map[util.Coordinates]int32), keys: make(map[int32]bool),
		keysCollected: make(map[int32]bool), keyLocs: make(map[int32]util.Coordinates)}
	var x, y int
	for scanner.Scan() {
		line := scanner.Text()
		x = 0
		for _, c := range line {
			if c == '@' {
				tunnel.entraces = append(tunnel.entraces, util.Coordinates{x, y})
			}
			if c >= 'a' && c <= 'z' {
				tunnel.keys[c] = true
				tunnel.keyLocs[c] = util.Coordinates{x, y}
			}
			tunnel.layout[util.Coordinates{x, y}] = c
			x++
		}
		y++
	}
	tunnel.size.X = x
	tunnel.size.Y = y
	return tunnel, nil
}

func findAllKeys(t tunnel) {
	var queue []passage
	for i, entrace := range t.entraces {
		queue = append(queue, NewPassage(0, make(map[int32]bool), entrace, entrace, i))
	}
	visited := make(map[passageId]bool)
	completed := make(map[int]bool)
	var count int
	for {
		count++
		p := queue[0]
		psgId := GetPassageId(p)
		c := t.layout[p.pos]

		if len(t.keysCollected) == len(t.keys) && allCompleted(completed, t) {
			fmt.Printf("Found all keys after %d steps and %d iterations \n", findNumOfSteps(queue, t), count)
			break
		}
		if allKeysAcquired(p, t) {
			queue = append(queue, p)
			queue = queue[1:]
			completed[qudrant(p.pos, t)] = true
			continue
		}

		if isDoor(c) {
			if qudrant(p.pos, t) == qudrant(t.keyLocs[c+32], t) {
				if !p.keys[c+32] {
					queue = queue[1:]
					continue
				}
			} else {
				if !t.keysCollected[c+32] {
					queue = append(queue, p)
					queue = queue[1:]
					continue
				}
			}
		}

		if visited[psgId] {
			queue = queue[1:]
			continue
		}

		visited[psgId] = true

		pos := p.pos
		queue = tryPassage(util.Coordinates{pos.X, pos.Y - 1}, t, queue, p)
		queue = tryPassage(util.Coordinates{pos.X, pos.Y + 1}, t, queue, p)
		queue = tryPassage(util.Coordinates{pos.X + 1, pos.Y}, t, queue, p)
		queue = tryPassage(util.Coordinates{pos.X - 1, pos.Y}, t, queue, p)

		queue = queue[1:]
	}
}

func allCompleted(c map[int]bool, t tunnel) bool {
	completed := true
	for i := 0; i < len(t.entraces); i++ {
		if !c[i] {
			completed = false
		}
	}
	return completed
}

func findNumOfSteps(queue []passage, t tunnel) int {
	steps := make(map[int]int)
	for _, p := range queue {
		if allKeysAcquired(p, t) {
			q := qudrant(p.pos, t)
			s, ok := steps[q]
			if ok {
				if p.steps < s {
					steps[q] = p.steps
				}
			} else {
				steps[q] = p.steps
			}
		}
	}

	totalSteps := 0
	for _, v := range steps {
		totalSteps += v
	}

	return totalSteps
}

func allKeysAcquired(p passage, t tunnel) bool {
	q := qudrant(p.pos, t)
	for k, v := range t.keyLocs {
		if q == qudrant(v, t) {
			if !p.keys[k] {
				return false
			}
		}
	}
	return true
}

func tryPassage(pos util.Coordinates, t tunnel, queue []passage, p passage) []passage {
	c := t.layout[pos]
	if isWall(c) {
		return queue
	}

	if isDoor(c) {
		return append(queue, NewPassage(p.steps+1, p.keys, pos, p.pos, p.quadrant))
	}

	if isKey(c) {
		t.keysCollected[c] = true
		newKeys := copyKeys(p.keys)
		newKeys[c] = true
		return append(queue, NewPassage(p.steps+1, newKeys, pos, p.pos, p.quadrant))
	}

	return append(queue, NewPassage(p.steps+1, p.keys, pos, p.pos, p.quadrant))
}

func copyKeys(keys map[int32]bool) map[int32]bool {
	newKeys := make(map[int32]bool)
	for k, v := range keys {
		newKeys[k] = v
	}
	return newKeys
}

func isKey(c int32) bool {
	return c >= 97 && c <= 122
}

func isDoor(c int32) bool {
	return c >= 65 && c <= 90
}

func isWall(c int32) bool {
	return c == '#'
}

func qudrant(pos util.Coordinates, t tunnel) int {
	if len(t.entraces) == 1 {
		return 0
	}
	if pos.X <= t.size.X/2 && pos.Y <= t.size.Y/2 {
		return 0
	}

	if pos.X > t.size.X/2 && pos.Y <= t.size.Y/2 {
		return 1
	}

	if pos.X <= t.size.X/2 && pos.Y > t.size.Y/2 {
		return 2
	}

	if pos.X > t.size.X/2 && pos.Y > t.size.Y/2 {
		return 3
	}
	return -1
}

func day18() {
	fmt.Println("Running day18")
	tunnel, err := readTunnelLayout(INPUT_FILE_DAY_18)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	findAllKeys(tunnel)

	tunnel, err = readTunnelLayout(INPUT_FILE_DAY_18_PART2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	findAllKeys(tunnel)
}
