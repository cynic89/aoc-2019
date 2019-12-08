package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INPUT_FILE_DAY_8 = "inputs/day8"
	WIDTH            = 25
	HEIGHT           = 6
	BLACK            = 48
	WHITE            = 49
	TRANSPARENT      = 50
)

type Layer struct {
	pixels [HEIGHT][WIDTH]int
}

func readLayers(file string) ([]*Layer, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	var layers []*Layer
	var currentLayer = new(Layer)
	var i, j int
	layers = append(layers, currentLayer)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		for _, val := range text {
			if j == WIDTH {
				i++
				j = 0
			}

			if i == HEIGHT {
				currentLayer = new(Layer)
				layers = append(layers, currentLayer)
				i = 0
				j = 0
			}

			currentLayer.pixels[i][j] = int(val)
			j++

		}
		break
	}
	return layers, nil
}

func findLayerWithLeastZeros(layers []*Layer) *Layer {
	var minZeros = 100
	var maxZerosLayer *Layer

	for _, layer := range layers {
		var zerosCount int
		for i := 0; i < HEIGHT; i++ {
			for j := 0; j < WIDTH; j++ {
				if layer.pixels[i][j] == 48 {
					zerosCount++
				}
			}
		}

		if zerosCount <= minZeros {
			minZeros = zerosCount
			maxZerosLayer = layer
		}
	}
	return maxZerosLayer
}

func oneMulTwo(layer *Layer) int {
	var oneCount, twoCount int

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if layer.pixels[i][j] == 49 {
				oneCount++
			}

			if layer.pixels[i][j] == 50 {
				twoCount++
			}
		}
	}

	return oneCount * twoCount
}

func decode(layers []*Layer) *Layer {
	var finalImage = new(Layer)

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {

			for _, layer := range layers {
				if layer.pixels[i][j] == BLACK {
					finalImage.pixels[i][j] = BLACK
					break
				}

				if layer.pixels[i][j] == WHITE {
					finalImage.pixels[i][j] = WHITE
					break
				}

				if layer.pixels[i][j] == TRANSPARENT {
					continue
				}
			}

		}
	}
	return finalImage
}

func message(layer *Layer) string {
	var message = ""
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if layer.pixels[i][j] == BLACK {
				message = message + " "
			}

			if layer.pixels[i][j] == WHITE {
				message = message + "1"
			}
		}
		message = message + "\n"
	}

	return message
}

func day8() {
	layers, err := readLayers(INPUT_FILE_DAY_8)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	maxLayer := findLayerWithLeastZeros(layers)
	oneMulTwoCount := oneMulTwo(maxLayer)
	fmt.Println(oneMulTwoCount)

	finalImage := decode(layers)
	fmt.Println(message(finalImage))
}
