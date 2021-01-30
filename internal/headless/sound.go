package headless

// Sound is a headless implementation of engine.Sound.
type Sound struct{}

// Play implements the Play method of engine.Sound.
func (s Sound) Play() {}

// Loop implements the Loop method of engine.Sound.
func (s Sound) Loop() {}

// Pause implements the Pause method of engine.Sound.
func (s Sound) Pause() {}

// Reset implements the Reset method of engine.Sound.
func (s Sound) Reset() {}

// Close implements the Close method of engine.Sound.
func (s Sound) Close() {}

// SoundControl is a headless implementation of engine.SoundControl.
type SoundControl struct{}

// SetVolume implements the SetVolume method of engine.SoundControl.
func (sc SoundControl) SetVolume(group string, v float64) {}

// Volume implements the Volume method of engine.SoundControl.
func (sc SoundControl) Volume(group string) float64 {
	return 1.0
}
