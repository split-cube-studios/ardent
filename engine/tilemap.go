package engine

// Tilemap is a placeholder for various tilemap implementations.
type Tilemap interface{}

// TileOverlapEvent updates renderer state in the case of a tile overlap.
type TileOverlapEvent func(bool, Image, interface{}) interface{}
