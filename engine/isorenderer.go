package engine

// IsoRenderer is an isometric renderer.
type IsoRenderer interface {
	Renderer
	SetTilemap(*Tilemap)
}
