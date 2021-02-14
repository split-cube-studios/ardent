package engine

// Collider resolves collisions.
type Collider interface {
	SetTilemap(*Tilemap)
	Resolve(Vec2, Vec2) Vec2
}
