package main

import (
    "math"
)

//let's place our gravity simulation functions here.

//SimulateGravity takes an initial universe as well as a number of generations and a time interval parameter.
//It returns a slice of Universe objects corresponding to running a gravity simulation over the appropriate number of generations, estimated for the given time interval (in seconds).
func SimulateGravity(initialUniverse Universe, numGens int, time float64) []Universe {
    timePoints := make([]Universe, numGens+1)

    //set the initial universe
    timePoints[0] = initialUniverse

    // call UpdateUniverse numGens times
    for i := 1; i <= numGens; i++ {
        timePoints[i] = UpdateUniverse(timePoints[i-1], time)
    }

    return timePoints
}

//UpdateUniverse takes a Universe object and a time interval (in seconds) and returns the updated universe simulated after the given time interval.
func UpdateUniverse(currentUniverse Universe, time float64) Universe {
    newUniverse := CopyUniverse(currentUniverse)

    // for every body in new universe, update its fields that will change (position, velocity, acceleration)
    // newUniverse.bodies has type []Body and contains the bodies we want to access
    for i := range newUniverse.bodies {
        // b is my current body = newUniverse.bodies[i]
        newUniverse.bodies[i].acceleration = UpdateAcceleration(newUniverse.bodies, newUniverse.bodies[i])

        newUniverse.bodies[i].velocity = UpdateVelocity(newUniverse.bodies[i], time)

        newUniverse.bodies[i].position = UpdatePosition(newUniverse.bodies[i], time)
    }

    return newUniverse
}

//UpdateVelocity takes a Body object and a time interval in seconds.  It updates the velocity of that object over the time interval using Newtonian dynamics equations.
func UpdateVelocity(b Body, time float64) OrderedPair {
    var v OrderedPair

    v.x = b.velocity.x + b.acceleration.x * time
    v.y = b.velocity.y + b.acceleration.y * time

    return v
}

//UpdatePosition takes a Body object and a time interval in seconds.  It updates the position of that object over the time interval using Newtonian dynamics equations.
func UpdatePosition(b Body, time float64) OrderedPair {
    var p OrderedPair

    p.x = b.position.x + b.velocity.x * time + 0.5*b.acceleration.x*(time*time)

    p.y = b.position.y + b.velocity.y * time + 0.5*b.acceleration.y*(time*time)

    return p
}

//UpdateAcceleration takes a slice of Body objects and a Body object and computes the acceleration of the body object based on the force of gravity of everything in the universe.
func UpdateAcceleration(bodies []Body, b Body) OrderedPair {
    var accel OrderedPair

    force := ComputeNetForce(bodies, b)

    // apply Newton's second law: F=ma, and so a = F/m
    // do this over components of force.
    accel.x = force.x/b.mass
    accel.y = force.y/b.mass

    return accel
}

//ComputeNetForce takes a slice of Body objects and a separate body object. It returns the net force of gravity (as an ordered pair) that the bodies exert on the given Body object.
func ComputeNetForce(bodies []Body, b Body) OrderedPair {
    var netForce OrderedPair

    // range over all bodies in the slice and compute the force that they exert on the given body
    for i := range bodies {
        if bodies[i] != b {
            force := ComputeForce(bodies[i], b)
            //add contribution of force to netForce
            netForce.x += force.x
            netForce.y += force.y
        }
    }

    return netForce
}

//ComputeForce takes two bodies and computes force exerted due to gravity on the second object due to the first.
func ComputeForce(b2, b Body) OrderedPair {
    var force OrderedPair
    // force.x, force.y get default values of 0, 0
    d := Distance(b2.position, b.position)

    // apply formula for magnitude of gravitational force
    F := G*b2.mass*b.mass/(d*d)

    deltaX := b2.position.x - b.position.x
    deltaY := b2.position.y - b.position.y

    // now we set the x and y components of force
    force.x = F*(deltaX/d) // F * cos(angle at which b2 lies compared to b)
    force.y = F*(deltaY/d) // F * sin(angle at which b2 lies compared to b)

    return force
}

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
    // this is the distance formula from days of precalculus long ago ...
    deltaX := p1.x - p2.x
    deltaY := p1.y - p2.y
    return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

//CopyUniverse takes a Universe object and produces a new Universe object with all fields copied over.
func CopyUniverse(currentUniverse Universe) Universe {
    var newUniverse Universe

    // copy over the fields of currentUniverse into newUniverse
    newUniverse.width = currentUniverse.width
    // this wouldn't duplicate the bodies ... it would give me the same bodies: newUniverse.bodies = currentUniverse.bodies
    // instead: I need to range over all the bodies and create new ones

    // declare the bodies of the new universe
    numBodies := len(currentUniverse.bodies)
    newUniverse.bodies = make([]Body, numBodies)

    for i := range currentUniverse.bodies {
        var b Body
        // set the fields of b equal to those of the current body (currentUniverse.bodies[i])
        currBody := currentUniverse.bodies[i]
        b.name = currBody.name
        b.red = currBody.red
        b.green = currBody.green
        b.blue = currBody.blue
        b.mass = currBody.mass
        b.radius = currBody.radius
        b.position.x = currBody.position.x
        b.position.y = currBody.position.y
        b.acceleration.x = currBody.acceleration.x
        b.acceleration.y = currBody.acceleration.y
        b.velocity.x = currBody.velocity.x
        b.velocity.y = currBody.velocity.y

        newUniverse.bodies[i] = b
    }

    return newUniverse
}
