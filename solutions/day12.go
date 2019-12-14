package main

import (
	"fmt"
	"github.com/cynic89/aoc-2019/util"
	"time"
)

type space struct {
	moons []*moon
}

type moon struct {
	pos threeD
	vel threeD

	rPos threeD
	rVel threeD
}
type threeD struct {
	x, y, z int
}

func (s *space) applyGravity() {
	var others []moon
	for _, m := range s.moons {
		others = append(others, *m)
	}
	for i, self := range s.moons {
		for j, other := range others {
			if i == j {
				continue
			}
			if self.pos.x != other.pos.x {
				self.vel.x += (other.pos.x - self.pos.x) / util.Abs(other.pos.x-self.pos.x)
			}
			if self.pos.y != other.pos.y {
				self.vel.y += (other.pos.y - self.pos.y) / util.Abs(other.pos.y-self.pos.y)
			}
			if self.pos.z != other.pos.z {
				self.vel.z += (other.pos.z - self.pos.z) / util.Abs(other.pos.z-self.pos.z)
			}

			if self.rPos.x != other.rPos.x {
				self.rVel.x += (self.rPos.x - other.rPos.x) / util.Abs(other.rPos.x-self.rPos.x)
			}
			if self.rPos.y != other.rPos.y {
				self.rVel.y += (self.rPos.y - other.rPos.y) / util.Abs(other.rPos.y-self.rPos.y)
			}
			if self.rPos.z != other.rPos.z {
				self.rVel.z += (self.rPos.z - other.rPos.z) / util.Abs(other.rPos.z-self.rPos.z)
			}

		}
	}
}

func (s *space) applyVelocity() {
	for _, self := range s.moons {
		self.pos.x += self.vel.x
		self.pos.y += self.vel.y
		self.pos.z += self.vel.z

		self.rPos.x -= self.rVel.x
		self.rPos.y -= self.rVel.y
		self.rPos.z -= self.rVel.z
	}
}

func (m moon) energy() int {
	return (util.Abs(m.pos.x) + util.Abs(m.pos.y) + util.Abs(m.pos.z)) * (util.Abs(m.vel.x) + util.Abs(m.vel.y) + util.Abs(m.vel.z))
}

func (m moon) string() string {
	return fmt.Sprintf("%d%d%d%d%d%d", m.pos.x, m.pos.y, m.pos.z, m.vel.x, m.vel.y, m.vel.z)
}

func (s *space) energy() (energy int) {
	for _, moon := range s.moons {
		energy += moon.energy()
	}
	return energy
}

func (s *space) string() string {
	var r string
	for _, moon := range s.moons {
		r = r + moon.string()
	}
	return r
}

func (s *space) calculateEnergy(steps int) int {
	for i := 0; i < steps; i++ {
		s.applyGravity()
		s.applyVelocity()
	}

	return s.energy()
}

func (s *space) driftUntilReachPreviousState() (steps int64) {

	energyMap := make(map[int][]string)
	for {
		s.applyGravity()
		s.applyVelocity()
		e := s.energy()
		st := s.string()
		if contains(st, energyMap[e]) {
			return steps
		}
		energyMap[e] = append(energyMap[e], st)

		steps++
	}

}

func (s *space) driftUntilReachPreviousStateOptimized() (steps int64) {
	fmt.Printf("Current step = %d at %v\n", steps, time.Now())
	for {
		zeroVal := threeD{0, 0, 0}
		mid := true
		s.applyGravity()
		s.applyVelocity()
		for _, m := range s.moons {
			if m.vel != zeroVal {
				mid = false
			}
		}

		if mid {
			fmt.Println(s.moons[0].rPos)
			fmt.Println(s.moons[0].pos)
			steps++
			return steps
		}
		steps++

		if steps/100000000 > 1 && steps%100000000 == 0 {
			fmt.Printf("Current step = %d at %v\n", steps, time.Now())
		}
	}
	return steps
}

func contains(v string, l []string) bool {
	for _, e := range l {
		if e == v {
			return true
		}
	}
	return false
}

/*
	&moon{pos: threeD{-1, 0, 2}},
	&moon{pos: threeD{2, -10, -7}},
	&moon{pos: threeD{4, -8, 8}},
	&moon{pos: threeD{3, 5, -1}},

	&moon{pos: threeD{-8, -10, 0}, rPos: threeD{-8, -10, 0}},
	&moon{pos: threeD{5, 5, 10}, rPos: threeD{5, 5, 10}},
	&moon{pos: threeD{2, -7, 3}, rPos: threeD{2, -7, 3}},
	&moon{pos: threeD{9, -8, -3}, rPos: threeD{9, -8, -3}},


		&moon{pos: threeD{16, -11, 2}, rPos: threeD{16, -11, 2}},
		&moon{pos: threeD{0, -4, 7}, rPos: threeD{0, -4, 7}},
		&moon{pos: threeD{6, 4, -10}, rPos: threeD{6, 4, -10}},
		&moon{pos: threeD{-3, -2, -4}, rPos: threeD{-3, -2, -4}},

	&moon{pos: threeD{-3, 15, -11}, rPos: threeD{-3, 15, -11}},
		&moon{pos: threeD{3, 13, -19}, rPos: threeD{3, 13, -19}},
		&moon{pos: threeD{-13, 18, -2}, rPos: threeD{-13, 18, -2}},
		&moon{pos: threeD{6, 0, -1}, rPos: threeD{6, 0, -1}},

<x=-3, y=15, z=-11>
<x=3, y=13, z=-19>
<x=-13, y=18, z=-2>
<x=6, y=0, z=-1>

<x=16, y=-11, z=2>
<x=0, y=-4, z=7>
<x=6, y=4, z=-10>
<x=-3, y=-2, z=-4>
*/

func day12() {

	var moons = []*moon{
		&moon{pos: threeD{-3, 15, -11}, rPos: threeD{-3, 15, -11}},
		&moon{pos: threeD{3, 13, -19}, rPos: threeD{3, 13, -19}},
		&moon{pos: threeD{-13, 18, -2}, rPos: threeD{-13, 18, -2}},
		&moon{pos: threeD{6, 0, -1}, rPos: threeD{6, 0, -1}},
	}

	s := space{moons: moons}
	e := s.calculateEnergy(1000)
	fmt.Println(e)


	//steps := s.driftUntilReachPreviousState()
	//fmt.Println(steps)

	var moons2 = []*moon{
		&moon{pos: threeD{-3, 5, -11}, rPos: threeD{-3, 5, -11}},
		&moon{pos: threeD{3, 13, -19}, rPos: threeD{3, 13, -19}},
		&moon{pos: threeD{-13, 18, -2}, rPos: threeD{-13, 18, -2}},
		&moon{pos: threeD{6, 0, -1}, rPos: threeD{6, 0, -1}},
	}

	space2 := space{moons: moons2}

	steps := space2.driftUntilReachPreviousStateOptimized()
	fmt.Println(steps)
}
