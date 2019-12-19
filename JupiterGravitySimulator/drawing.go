package main

import (
    "image"
    "canvas"
)

// let's place our drawing functions here.

//AnimateSystem takes a slice of Universe objects along with a canvas width
//parameter and a frequency parameter and generates a slice of images
//corresponding to drawing each Universe every frequency generations
//on a canvasWidth x canvasWidth canvas
func AnimateSystem(timePoints []Universe, canvasWidth int, frequency int) []image.Image {
    images := make([]image.Image, 0)

    for i := range timePoints {
        if i % frequency == 0 {
            images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
        }
    }

    return images
}

//DrawToCanvas generates the image corresponding to a canvas after drawing a
//Universe object's bodies on a square canvas that is canvasWidth pixels x
//canvasWidth pixels
func DrawToCanvas(u Universe, canvasWidth int) image.Image {
    // set a new square canvas
    c := canvas.CreateNewPalettedCanvas(canvasWidth, canvasWidth, nil)

    // create a black background
    c.SetFillColor(canvas.MakeColor(0, 0, 0))
    c.ClearRect(0, 0, canvasWidth, canvasWidth)
    c.Fill()

    // range over all the bodies and draw them.
    for i, b := range u.bodies {
        c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
        cx := (b.position.x/u.width)*float64(canvasWidth)
        cy := (b.position.y/u.width)*float64(canvasWidth)
        r := (b.radius/u.width)*float64(canvasWidth)
        if i == 0 { // Jupiter
            c.Circle(cx, cy, r)
        } else {
            //moons need to have radius scaled by factor of 10 to be visible.
            r *= 10.0
            c.Circle(cx, cy, r)
        }
        c.Fill()
    }
    // we want to return an image!
    return canvas.GetImage(c)
}
