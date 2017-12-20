package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const simSteps = 1000

func main() {
	ps := parseInput("./input")
	pt := make(particles, len(ps))
	copy(pt, ps)
	ps.simulate()
	fmt.Printf("The particle closest to <0,0,0> has the id %d.\n", ps.getMinID())
	remainingP := pt.simulateWithCollisonRemoval()
	fmt.Printf("After collisions there are %d particels left.\n", remainingP)
}

type particle struct {
	id  int
	pos [3]int
	vel [3]int
	acc [3]int
}

func (p *particle) tick() {
	p.vel = [3]int{p.acc[0] + p.vel[0], p.acc[1] + p.vel[1], p.acc[2] + p.vel[2]}
	p.pos = [3]int{p.pos[0] + p.vel[0], p.pos[1] + p.vel[1], p.pos[2] + p.vel[2]}
}

type particles []particle

func (ps particles) simulate() {
	for i := 0; i < simSteps; i++ {
		for i := range ps {
			ps[i].tick()
		}
	}
}

func (ps particles) simulateWithCollisonRemoval() int {
	removed := make(map[particle]bool)
	for i := 0; i < simSteps; i++ {
		hit := make(map[[3]int]int)
		for i := range ps {
			if !removed[ps[i]] {
				ps[i].tick()
				if v, ok := hit[ps[i].pos]; ok {
					removed[ps[i]] = true
					removed[ps[v]] = true
				}
				hit[ps[i].pos] = ps[i].id
			}

		}
	}
	return len(ps) - len(removed)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (ps particles) getMinID() (id int) {
	min := 1<<63 - 1
	for _, p := range ps {
		dist := abs(p.pos[0]) + abs(p.pos[1]) + abs(p.pos[2])
		if dist < min && dist > 0 {
			min = dist
			id = p.id
		}
	}
	return
}

func parseInput(file string) particles {
	get := func(s string) int {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return v
	}

	parts := make([]particle, 0)
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()
	s := bufio.NewScanner(fileHandle)
	i := 0
	for s.Scan() {
		line := s.Text()
		t := strings.Split(line, ",")

		x, y, z := get(t[0][3:]), get(t[1]), get(t[2][:len(t[2])-1])
		vx, vy, vz := get(t[3][4:]), get(t[4]), get(t[5][:len(t[5])-1])
		ax, ay, az := get(t[6][4:]), get(t[7]), get(t[8][:len(t[8])-1])

		p := particle{id: i,
			pos: [3]int{x, y, z},
			vel: [3]int{vx, vy, vz},
			acc: [3]int{ax, ay, az}}
		parts = append(parts, p)
		i++
	}
	return parts
}
