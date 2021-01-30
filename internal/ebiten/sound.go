package ebiten

import (
	"bytes"
	"math"
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

	sc     *SoundControl
	volume float64
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

	s.player.SetVolume(s.volume)
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

	s.player.SetVolume(s.volume)
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

	// should not produce error
	_ = s.player.Rewind()
}

// Close implements the Close method of engine.Sound.
func (s *Sound) Close() {

	if s.player == nil {
		return
	}

	s.player.Close()

	if s.sc != nil {
		s.sc.removeSound(s)
	}
}

func (s *Sound) setVolume(v float64) {

	s.volume = v

	if s.player == nil {
		return
	}

	s.player.SetVolume(v)
}

// SoundControl is an ebiten implementation of engine.SoundControl.
type SoundControl struct {
	groups  map[string]map[*Sound]struct{}
	volumes map[string]float64
}

// NewSoundControl returns an instantiated *SoundControl.
func NewSoundControl() *SoundControl {
	return &SoundControl{
		groups:  make(map[string]map[*Sound]struct{}),
		volumes: make(map[string]float64),
	}
}

// SetVolume implements the SetVolume method of engine.SoundControl.
func (sc *SoundControl) SetVolume(group string, v float64) {

	v = math.Max(
		0,
		math.Min(1.0, v),
	)

	sc.volumes[group] = v

	for sound := range sc.groups[group] {
		sound.setVolume(v)
	}
}

// Volume implements the Volume method of engine.SoundControl.
func (sc *SoundControl) Volume(group string) (v float64) {

	var ok bool
	if v, ok = sc.volumes[group]; !ok {
		v = 1.0
	}

	return
}

// nolint:unused
func (sc *SoundControl) addSound(s *Sound) {

	group := sc.groups[s.group]
	if group == nil {
		group = make(map[*Sound]struct{})
		sc.groups[s.group] = group
	}

	group[s] = struct{}{}
}

func (sc *SoundControl) removeSound(s *Sound) {
	if group, ok := sc.groups[s.group]; ok {
		delete(group, s)
	}
}
