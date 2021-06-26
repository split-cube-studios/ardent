# Ardent
[![Go Reference](https://pkg.go.dev/badge/github.com/split-cube-studios/ardent.svg)](https://pkg.go.dev/github.com/split-cube-studios/ardent) ![Go](https://github.com/split-cube-studios/ardent/workflows/Go/badge.svg)

Ardent is a cross-platform 2D game engine, initially developed for the game Ardent Woods.

Please note that Ardent is currently pre-stable! Breaking changes can happen at any time. Something something WIP.

## Features

Alongside basic features such as user input and rendering, Ardent offers a few key features:
- Custom asset files
- 2D and isometric rendering
- Simple collisions
- Spatial partitioning for large maps
- State machines
- Context based steering algorithm
- And much more!

## Assets

Ardent provides its own asset file format. This allows asset files to indicate information about animations,
image atlases, sound effect variants, and more.

A tool called `aautil` (ardent asset utility) is provided to create asset files.

To install `aautil`, run `go install ./cmd/aautil`.

`aautil` can take one or more paths as arguments for asset file folder locations. YAML files are used for configuration. Checkout some of the examples for samples of configuration files.

## Discord

Come chat with us in the `#ardent` channel on [Discord](https://discord.gg/dUqS7RfSqv)!
