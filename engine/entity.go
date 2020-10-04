package engine

type Entity interface {
	Tick()

	SetCollider(Collider)

	Position() Vec2

	AddImage(...Image)
	Images() []Image

	Dispose()
	IsDisposed() bool
}

type CoreEntity struct {
	Vec2
	prevPos, Origin Vec2

	images []Image

	collider Collider
	disposed bool
}

func (e *CoreEntity) Tick() {
	if e.collider != nil {
		e.Vec2 = e.collider.Resolve(e.prevPos, e.Vec2)
		e.prevPos = e.Vec2
	}

	for _, img := range e.images {
		w, h := img.Size()
		img.Translate(
			e.X-e.Origin.X*float64(w),
			e.Y-e.Origin.Y*float64(h),
		)
	}
}

func (e *CoreEntity) SetCollider(collider Collider) {
	e.collider = collider
}

func (e *CoreEntity) Position() Vec2 {
	return e.Vec2
}

func (e *CoreEntity) AddImage(image ...Image) {
	e.images = append(e.images, image...)
}

func (e *CoreEntity) Images() []Image {
	return e.images
}

func (e *CoreEntity) Dispose() {
	e.disposed = true
}

func (e *CoreEntity) IsDisposed() bool {
	return e.disposed
}
