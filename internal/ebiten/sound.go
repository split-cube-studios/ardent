package ebiten

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// Sound is an ebiten implementation of engine.Sound.
type Sound struct {
	group   string
	options [][]byte
	context *audio.Context
	player  *audio.Player
}

// Play implements the Play method of engine.Sound.
func (s *Sound) Play() {

	if len(s.options) == 0 {
		return
	}

	if s.player != nil {
		s.player.Close()
	}

	// NOTE we may want to cache the players
	s.player = audio.NewPlayerFromBytes(
		s.context,
		s.options[rand.Intn(len(s.options))],
	)

	s.player.Play()
}

// Loop implements the Loop method of engine.Sound.
func (s *Sound) Loop() {}

// Pause implements the Pause method of engine.Sound.
func (s *Sound) Pause() {

	if s.player == nil {
		return
	}

	s.player.Pause()
}

// Reset implements the Reset method of engine.Sound.
func (s *Sound) Reset() {

	if s.player == nil {
		return
	}

	s.player.Rewind()
}

// SoundControl is an ebiten implementation of engine.SoundControl.
type SoundControl struct{}

// SetVolume implements the SetVolume method of engine.SoundControl.
func (sc *SoundControl) SetVolume(group string, v float64) {}

// Volume implements the Volume method of engine.SoundControl.
func (sc *SoundControl) Volume(group string) float64 {
	return 1.0
}
