package engine

// Animation is a series of frames that play in sequence.
type Animation interface {
	SetState(string)

	SetTickCount(int)
	Play()
	Pause()
	Reset()

	Image
}
