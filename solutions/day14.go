package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	INPUT_FILE_DAY_14 = "inputs/day14"
)

type chemical struct {
	units       float64
	name        string
	ingredients []chemical
}

func parseFuelData() (fuelMap map[string]chemical, err error) {
	f, err := os.Open(INPUT_FILE_DAY_14)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	fuelMap = make(map[string]chemical)

	for scanner.Scan() {
		var ingredients []chemical
		line := scanner.Text()
		reaction := strings.Split(line, "=>")
		for _, ing := range strings.Split(reaction[0], ",") {
			ing = strings.TrimSpace(ing)
			chem := strings.Split(ing, " ")
			unit, _ := strconv.Atoi(chem[0])
			name := chem[1]
			ingredients = append(ingredients, chemical{units: float64(unit), name: name})
		}

		ch := strings.TrimSpace(reaction[1])
		chem := strings.Split(ch, " ")
		unit, _ := strconv.Atoi(chem[0])
		name := chem[1]
		fuelMap[name] = chemical{units: float64(unit), name: name, ingredients: ingredients}
	}
	return fuelMap, nil

}

func solve(fuelMap map[string]chemical, extra map[string]float64) (ores float64) {
	f := fuelMap["FUEL"]
	var queue = []chemical{f}

	for len(queue) > 0 {
		ch := queue[0]
		reaction := fuelMap[ch.name]
		if extra[ch.name] >= ch.units {
			extra[ch.name] -= ch.units
			queue = queue[1:]
			continue
		}
		requiredReactions := reactionsRequired(ch, reaction, extra)
		if len(reaction.ingredients) == 1 && reaction.ingredients[0].name == "ORE" {
			oresToAdd := requiredReactions * reaction.ingredients[0].units
			ores += oresToAdd
			queue = queue[1:]
			continue
		}

		for _, ing := range reaction.ingredients {
			requiredIng := chemical{units: ing.units * requiredReactions, name: ing.name}
			queue = append(queue, requiredIng)
		}

		queue = queue[1:]

	}

	return ores

}

func reactionsRequired(c, reaction chemical, extra map[string]float64) float64 {
	u := math.Ceil(c.units / reaction.units)
	v := math.Ceil((c.units - extra[c.name]) / reaction.units)
	r := math.Min(u, v)

	if u <= v {
		extra[c.name] += (r * reaction.units) - c.units
	} else {
		extra[c.name] = (r * reaction.units) - (c.units - extra[c.name])
	}

	return r
}

func maxFuel(oresAvailable float64, fuelMap map[string]chemical, extra map[string]float64) float64 {
	var fuel, oresUsed float64
	for oresUsed < oresAvailable {
		oresUsed += solve(fuelMap, extra)
		fuel++
	}
	fmt.Println(int64(oresUsed))
	return fuel
}

func day14() {
	fuelMap, err := parseFuelData()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	extra := make(map[string]float64)
	ores := solve(fuelMap, extra)
	fmt.Println(int(ores))
	m := maxFuel(1000000000000.0, fuelMap, extra)
	fmt.Println(int64(m-2))//donno why
}
