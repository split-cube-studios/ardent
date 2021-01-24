package ebiten

import (
	"bytes"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
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

	option := s.options[rand.Intn(len(s.options))]
	sound, _ := vorbis.Decode(s.context, bytes.NewReader(option))

	// NOTE we may want to cache the players
	s.player, _ = audio.NewPlayer(
		s.context,
		sound,
	)

	s.player.Play()
}

// Loop implements the Loop method of engine.Sound.
func (s *Sound) Loop() {

	if len(s.options) == 0 {
		return
	}

	if s.player != nil {
		s.player.Close()
	}

	option := s.options[rand.Intn(len(s.options))]
	sound, _ := vorbis.Decode(s.context, bytes.NewReader(option))

	loop := audio.NewInfiniteLoop(
		sound,
		sound.Length(),
	)

	// shouldn't return an error
	s.player, _ = audio.NewPlayer(s.context, loop)

	s.player.Play()
}

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
