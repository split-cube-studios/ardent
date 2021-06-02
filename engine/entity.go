package engine

// Entity is a basic game entity.
type Entity interface {
	Tick()

	SetCollider(Collider)

	Position() Vec2

	AddImage(...Image)
	Images() []Image

	Class() string

	Dispose()
	IsDisposed() bool
}

// CoreEntity is a default Entity implementation.
type CoreEntity struct {
	Vec2
	prevPos Vec2

	// Direction is the current cardinal direction
	// the entity is facing.
	Direction byte

	images []Image

	collider Collider
	disposed bool
}

// Tick updates the CoreEntity's position.
func (e *CoreEntity) Tick() {
	if e.collider != nil {
		e.Vec2 = e.collider.Resolve(e.prevPos, e.Vec2)
	}

	e.prevPos = e.Vec2

	for _, img := range e.images {
		img.Translate(e.X, e.Y)
	}
}

// SetCollider sets the CoreEntity's Collider.
func (e *CoreEntity) SetCollider(collider Collider) {
	e.collider = collider
}

// Position gets the CoreEntity's current position.
func (e *CoreEntity) Position() Vec2 {
	return e.Vec2
}

// AddImage adds an Image to the CoreEntity.
func (e *CoreEntity) AddImage(image ...Image) {
	for _, img := range image {
		img.Translate(e.X, e.Y)
	}

	e.images = append(e.images, image...)
}

// Images gets the CoreEntity's Images.
func (e *CoreEntity) Images() []Image {
	return e.images
}

// Dispose marks the CoreEntity as disposed, and disposes its Images.
func (e *CoreEntity) Dispose() {
	e.disposed = true
	for _, img := range e.images {
		img.Dispose()
	}
}

// IsDisposed checks if the CoreEntity has been disposed.
func (e *CoreEntity) IsDisposed() bool {
	return e.disposed
}

// MoveTowards moves the CoreEntity in the direction of angle
// by distance dist. The Direction field is updated with the
// current closest cardinal direction.
func (e *CoreEntity) MoveTowards(angle, dist float64) {
	e.Direction = AngleToCardinal(angle)
	e.Vec2 = e.Translate(angle, dist)
}
