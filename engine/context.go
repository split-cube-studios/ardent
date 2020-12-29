package engine

import (
	"image"
	"math"
)

// Context is a rendering context.
type Context struct {
	Renderer
	Collider

	partitionMap *PartitionMap

	entities   map[string][]Entity
	entitySwap []Entity
}

// NewContext creates a Context with the given Renderer and Collider.
func NewContext(renderer Renderer, collider Collider) *Context {
	ctx := &Context{
		Renderer: renderer,
		Collider: collider,
		entities: make(map[string][]Entity),
	}

	// TODO configurable values
	ctx.partitionMap = NewPartitionMap(250, 1000)

	return ctx
}

// AddEntity adds Entities to the Context.
func (c *Context) AddEntity(entities ...Entity) {
	c.entitySwap = append(c.entitySwap, entities...)
}

// GetEntities gets the Context's Entities.
func (c *Context) GetEntities(class string) []Entity {
	return c.entities[class]
}

// Tick updates the Context's internal state.
func (c *Context) Tick() {
	for i, e := range c.entitySwap {
		if c.Collider != nil {
			e.SetCollider(c.Collider)
		}

		c.AddImage(e.Images()...)
		c.partitionMap.Add(e)

		c.entitySwap[i] = nil
	}

	c.entitySwap = c.entitySwap[:0]

	vp := c.Viewport()
	pos := Vec2{
		X: float64(vp.Min.X),
		Y: float64(vp.Min.Y),
	}

	// cell dist to load from partition map
	pcells := int(math.Max(
		float64(vp.Dx()),
		float64(vp.Dy()),
	))/250 - 1

	c.partitionMap.Tick(pos, pcells, c.updateEntities)
}

func (c *Context) updateEntities(entries []PartitionEntry) {
	for _, entry := range entries {
		entry.(Entity).Tick()

		hitbox, ok := entry.(Hitbox)
		if !ok {
			continue
		}

		for _, targetClass := range hitbox.HitClasses() {
			entries := c.partitionMap.Class(targetClass)
			for _, targetEntry := range entries {
				targetHitbox, ok := targetEntry.(Hitbox)
				if !ok {
					continue
				}

				sourcePos := hitbox.Position()
				sourceBounds := hitbox.HitBounds()
				targetPos := targetHitbox.Position()
				targetBounds := targetHitbox.HitBounds()

				if image.Rect(
					int(sourcePos.X)+sourceBounds.Min.X,
					int(sourcePos.Y)+sourceBounds.Min.Y,
					int(sourcePos.X)+sourceBounds.Max.X,
					int(sourcePos.Y)+sourceBounds.Max.Y,
				).Overlaps(
					image.Rect(
						int(targetPos.X)+targetBounds.Min.X,
						int(targetPos.Y)+targetBounds.Min.Y,
						int(targetPos.X)+targetBounds.Max.X,
						int(targetPos.Y)+targetBounds.Max.Y,
					),
				) {
					hitbox.Hit(targetEntry.(Entity))
				}
			}
		}
	}
}
