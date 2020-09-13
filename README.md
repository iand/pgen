# pgen
A Go package for deterministic generation of random-like numbers.

This package is designed for simulations and procedural generation where you need to generate random-like features from a known seed. 

A standard number generator produces random numbers in a sequence so if a code refactor changes the
order in which you call the rng you'll get a different result. pgen allows you to specify which
number in the random sequence is to be used for a specific procedurally generated feature. This
preserves behaviour even if code refactoring changes the order in which the simulation is generated.

[![Build Status](https://travis-ci.org/iand/pgen.svg?branch=master)](https://travis-ci.org/iand/pgen)

## Installation

Simply run

    go get github.com/iand/pgen

Documentation is at [http://godoc.org/github.com/iand/pgen](http://godoc.org/github.com/iand/pgen)


## Usage

Example of a procedural level generator:

    package main

    import (
        "github.com/iand/pgen"
        "math/rand"
    )

    func generateLevel(level int64) {
        // These constants control which random number is returned by the generator
        const (
            numberOfPits = iota
            numberOfEnemies
            spawnPointX
            spawnPointY
            legendaryLootChance
        )
        
        gen := pgen.New(level)

        // Generate up to 8 pits
        for i := 0; i < gen.Intn(numberOfPits, 8) {
            generatePit()
        }

        // Generate up to 4 enemies
        for i := 0; i < gen.Intn(numberOfEnemies, 4) {
            generateEnemy()
        }

        // Starting location on the 64x64 grid
        x, y := gen.Intn(spawnPointX,64), gen.Intn(spawnPointY,64)
        placePlayer(x, y)

        // Is there a legendary item?
        if rand.Float64() <= gen.Intn(legendaryLootChance).Float64() {
            placeLoot()
        }
    }       

The key advantage of a procedural generator over an ordinary random number generator is that the order of calls doesn't matter
and adding extra calls in between others will not affect the random numbers. This makes it robust against ongoing changes in the
program. In the example above, adding new type of enemy later on won't break any existing levels that have been generated.


## License

This is free and unencumbered software released into the public domain. For more
information, see <http://unlicense.org/> or the accompanying [`UNLICENSE`](UNLICENSE) file.
