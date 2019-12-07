package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type object struct {
	val       string
	parent    *object
	orbitters []*object
}

const INPUT_FILE_DAY_6 = "inputs/day6"

func parseInput(file string) (object, error) {
	f, err := os.Open(file)
	if err != nil {
		return object{}, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	o := object{val: "root"}
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, ")")
		o.addObject(vals[0], vals[1])
	}

	return o, nil
}

func (o *object) addObject(parent, child string) {
	var parentNode *object
	_, parentNode, err := o.findObject(parent)
	if err != nil {
		parentNode = &object{val: parent, parent: o}
		o.orbitters = append(o.orbitters, parentNode)
	}

	oldParent, childNode, err := o.findObject(child)
	if err != nil {
		childNode = &object{val: child, parent: parentNode}
		parentNode.orbitters = append(parentNode.orbitters, childNode)
		return
	}
	childNode.parent = parentNode
	parentNode.orbitters = append(parentNode.orbitters, childNode)

	var newOrbitters []*object
	for _, ob := range oldParent.orbitters {
		if ob.val != child {
			newOrbitters = append(newOrbitters, ob)
		}
	}
	oldParent.orbitters = newOrbitters

}

func (o *object) findObject(val string) (*object, *object, error) {
	var queue = []*object{o}
	for len(queue) > 0 {
		obj := queue[0]
		for _, child := range obj.orbitters {
			if child.val == val {
				return obj, child, nil
			}
			queue = append(queue, child)
		}
		queue = queue[1:]
	}
	return nil, nil, fmt.Errorf("%s not found", val)
}

type node struct {
	obj   *object
	depth int
}

func (o *object) print() {
	var queue []node
	queue = append(queue, node{o, 0})
	for len(queue) > 0 {
		n := queue[0]
		obj := n.obj
		fmt.Printf(" %s is at depth %d", obj.val, n.depth-1)
		fmt.Println()
		for _, child := range obj.orbitters {
			queue = append(queue, node{child, n.depth + 1})
		}
		queue = queue[1:]
	}

}

func (o *object) orbits() int {
	var queue []node
	var orbitCount int
	queue = append(queue, node{o, 0})
	for len(queue) > 0 {
		n := queue[0]
		obj := n.obj
		if n.depth > 0 {
			orbitCount = orbitCount + (n.depth - 1)
		}
		for _, child := range obj.orbitters {
			queue = append(queue, node{child, n.depth + 1})
		}
		queue = queue[1:]
	}
	return orbitCount
}

func (o *object) path(source, target *object) int {
	var (
		currentObj = source
		sourcePath []*object
		targetPath []*object
	)
	for currentObj.parent != nil {
		parent := currentObj.parent
		sourcePath = append([]*object{parent}, sourcePath...)
		currentObj = parent
	}

	currentObj = target

	for currentObj.parent != nil {
		parent := currentObj.parent
		targetPath = append([]*object{parent}, targetPath...)
		currentObj = parent
	}

	var divergeLoc int
	for i, _ := range sourcePath {
		if targetPath[i].val != sourcePath[i].val {
			divergeLoc = i
			break
		}
	}

	distance := (len(sourcePath) - divergeLoc) + (len(targetPath) - divergeLoc)
	return distance
}

func day6() {
	obj, err := parseInput(INPUT_FILE_DAY_6)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//obj.print()

	orbitCount := obj.orbits()
	fmt.Printf("Orbits Count = %d \n", orbitCount)

	_, src, _ := obj.findObject("YOU")
	_, target, _ := obj.findObject("SAN")
	distance := obj.path(src, target)

	fmt.Printf("Distance from YOU to SAN = %d \n", distance)

}
