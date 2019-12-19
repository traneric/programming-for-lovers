package main

import (
	"math"
	"fmt"
	"os"
	"strconv"
	"gifhelper"
)

//G is the gravitational constant in the gravitational force equation.
//It is declared as a "global" constant that can be accessed by all functions.
const G = 1.4 * 6.67408e-11

//data setup.
type Body struct {
	name                             string
	mass, radius                     float64
	position, velocity, acceleration OrderedPair
	red, green, blue                 uint8
}

type OrderedPair struct {
	x, y float64
}

type Universe struct {
	bodies []Body
	width  float64
}

func main() {
	// declaring objects
	var jupiter, io, europa, ganymede, callisto Body

	jupiter.name = "Jupiter"
	io.name = "Io"
	europa.name = "Europa"
	ganymede.name = "Ganymede"
	callisto.name = "Callisto"

	jupiter.red, jupiter.green, jupiter.blue = 223, 227, 202
	io.red, io.green, io.blue = 249, 249, 165
	europa.red, europa.green, europa.blue = 132, 83, 52
	ganymede.red, ganymede.green, ganymede.blue = 76, 0, 153
	callisto.red, callisto.green, callisto.blue = 0, 153, 76

	jupiter.mass = 1.898 * math.Pow(10, 27)
	io.mass = 8.9319 * math.Pow(10, 22)
	europa.mass = 4.7998 * math.Pow(10, 22)
	ganymede.mass = 1.4819 * math.Pow(10, 23)
	callisto.mass = 1.0759 * math.Pow(10, 23)

	jupiter.radius = 71000000
	io.radius = 1821000
	europa.radius = 1569000
	ganymede.radius = 2631000
	callisto.radius = 2410000

	jupiter.position.x, jupiter.position.y = 2000000000, 2000000000
	io.position.x, io.position.y = 2000000000-421600000, 2000000000
	europa.position.x, europa.position.y = 2000000000, 2000000000+670900000
	ganymede.position.x, ganymede.position.y = 2000000000+1070400000, 2000000000
	callisto.position.x, callisto.position.y = 2000000000, 2000000000-1882700000

	jupiter.velocity.x, jupiter.velocity.y = 0, 0
	io.velocity.x, io.velocity.y = 0, -17320
	europa.velocity.x, europa.velocity.y = -13740, 0
	ganymede.velocity.x, ganymede.velocity.y = 0, 10870
	callisto.velocity.x, callisto.velocity.y = 8200, 0

	// declaring universe and setting its fields.
	var jupiterSystem Universe
	jupiterSystem.width = 4000000000
	jupiterSystem.bodies = []Body{jupiter, io, europa, ganymede, callisto}

	fmt.Println("Let's read in some command line arguments.")

	//numGens is os.Args[1], time is os.Args[2]
	numGens, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic("Error: Problem reading numGens argument.")
	}

	time, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil {
		panic("Error: Problem reading time argument.")
	}

	//os.Args[3] is canvas width in pixels
	canvasWidth, err3 := strconv.Atoi(os.Args[3])
	if err3 != nil {
		panic("Error: Problem reading canvas width argument.")
	}

	//os.Args[4] is frequency (number of generations) with which we draw
	//universes to images.
	frequency, err4 := strconv.Atoi(os.Args[4])
	if err4 != nil {
		panic("Error: Problem reading frequency argument.")
	}

	fmt.Println("Ready to implement gravity simulator!")

	timePoints := SimulateGravity(jupiterSystem, numGens, time)

	fmt.Println("Simulation run successfully!")

	images := AnimateSystem(timePoints, canvasWidth, frequency)

	fmt.Println("Images drawn!")

	filename := "JupiterMoons"

	gifhelper.ImagesToGIF(images, filename)

	fmt.Println("Animated GIF produced. Exiting normally.")

	// example: .\JupiterGravitySimulator 2000 60 500 50
	// 2000 generations, 60 seconds, 500 canvas width, draw every 50th gen

}
