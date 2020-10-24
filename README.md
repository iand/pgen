# pgen

[![Build Status](https://travis-ci.org/iand/pgen.svg?branch=master)](https://travis-ci.org/iand/pgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/iand/pgen)](https://goreportcard.com/report/github.com/iand/pgen)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/iand/pgen)

## Overview

A Go package for deterministic generation of random-like numbers.

This package is designed for simulations and procedural generation where you need to generate random-like features from a known seed. 

A standard number generator produces random numbers in a sequence so if a code refactor changes the
order in which you call the rng you'll get a different result. pgen allows you to specify which
number in the random sequence is to be used for a specific procedurally generated feature. This
preserves behaviour even if code refactoring changes the order in which the simulation is generated.

## Installation

Simply run the following in your module:

    go get github.com/iand/pgen

Documentation is at [http://godoc.org/github.com/iand/pgen](http://godoc.org/github.com/iand/pgen)


## Usage

Example of a procedural level generator:

    package main

    import (
        "github.com/iand/pgen"
        "math/rand"
    )

    // generateLevel generates a new game level based on the level seed. It will always
    // generate the same level no matter how many other levels were generated before
    // or if the order of generation is changed.
    func generateLevel(levelSeed int64) {
        // These constants control which random number is returned by the generator
        const (
            numberOfPits = iota
            numberOfEnemies
            spawnPointX
            spawnPointY
            legendaryLootChance
        )
        
        gen := pgen.New(levelSeed)

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


## License

This is free and unencumbered software released into the public domain. For more
information, see <http://unlicense.org/> or the accompanying [`UNLICENSE`](UNLICENSE) file.
