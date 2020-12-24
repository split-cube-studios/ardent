package engine

import "image"

// Hitbox is an object with collision support.
type Hitbox interface {
	Position() Vec2
	HitBounds() image.Rectangle // TODO upgrade to float based approach?
	HitClasses() []string
	Hit(Entity)
}
