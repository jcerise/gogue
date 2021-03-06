package noises

import (
	"github.com/gogue-framework/gogue/gamemap"
	"math"
)

// NoiseGenerator generates noises on tiles in a circle around an entity
type NoiseGenerator struct {
	cosTable map[int]float64
	sinTable map[int]float64
}

// InitializeNoiseGenerator creates and stores cos and sin tables for 360 degrees, so they don't have to be regenerated
// each time this object is used.
func (f *NoiseGenerator) InitializeNoiseGenerator() {

	f.cosTable = make(map[int]float64)
	f.sinTable = make(map[int]float64)

	for i := 0; i < 360; i++ {
		ax := math.Sin(float64(i) / (float64(180) / math.Pi))
		ay := math.Cos(float64(i) / (float64(180) / math.Pi))

		f.sinTable[i] = ax
		f.cosTable[i] = ay
	}
}

// RayCastSound casts out rays each degree in a 360 circle from the entity. If a ray passes over a floor (does not block sound)
// tile, keep going, up to the maximum distance the sound can travel from the entity. If the ray intersects a wall
// (blocks sound), stop, as the sound will not penetrate the wall. Every tile that the sound carries through will
// get a noise value corresponding to the entity, and the value of the sound. Sound degrades the further from the
// source it is.
func (f *NoiseGenerator) RayCastSound(entity, entityX, entityY int, intensity float64, gameMap *gamemap.GameMap) {

	for i := 0; i < 360; i++ {

		ax := f.sinTable[i]
		ay := f.cosTable[i]

		x := float64(entityX)
		y := float64(entityY)

		// Mark the entities current position as the source of the noise. This tile will have the full noise intensity
		// value for this frame
		tile := gameMap.Tiles[entityX][entityY]
		tile.Noises[entity] = intensity

		// Reduce the intensity by a value of 1, and then start raycasting. For each tile away from the source (the
		// entities location), reduce the intensity by 1. Once the intensity is 0, stop.
		reducedIntensity := intensity - 1

		for j := reducedIntensity; j > 0; j-- {
			x -= ax
			y -= ay

			roundedX := int(round(x))
			roundedY := int(round(y))

			if x < 0 || x > float64(gameMap.Width-1) || y < 0 || y > float64(gameMap.Height-1) {
				// If the ray is cast outside of the map, stop
				break
			}

			tile := gameMap.Tiles[roundedX][roundedY]
			tile.Noises[entity] = j

			if gameMap.Tiles[roundedX][roundedY].BlocksNoises == true {
				// The ray hit a tile that does not transmit sound, go no further
				break
			}
		}
	}
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}
