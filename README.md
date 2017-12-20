# Gogue - Roguelike toolkit for Go

Gogue aims to create a simple to use toolkit for creating Roguelike games in the Go language. It uses BearLibTerminal for rendering, so that will be required to use the toolkit.
This is by no means a complete toolkit, but its got (and will have) a bunch of things I've found handy for building roguelikes. Hopefully someone else finds them handy as well.

Development is on-going, so use at your own risk.

## Features

This feature list is incomplete, as various pieces are still in development.

- Terminal creation and management (using BearlibTerminal)
    - Glyph rendering
    - Input handling
- Dynamic input registration system
- Lightweight Entity/Component/System implementation
- Map generation
- Scrolling camera
- Field of View (only raycasting at the moment, but more to come)
- UI
    - Logging
    - Menus
- Pathfinding
- Random number generation
- Random name generation

... and whatever else I deem useful

## Getting Started

Standard Go package install - `go get github.com/jcerise/gogue`

### Prerequisites

BearLibTerminal is required to use this package. I'll put specific platform install instructions here, as it varies from platform to platform, and is not well documented (I feel).
