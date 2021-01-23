package engine

// Sound represents music or a sound effect.
type Sound interface {
	// Play plays the audio from the current
	// position to the end.
	Play()

	// Loop plays the audio from the current
	// position, and repeats from the start
	// after reaching the end.
	Loop()

	// Pause stops the sound for playing,
	// keeping the current position.
	Pause()

	// Reset seeks to the start of the sound.
	// Reset will also pause the track.
	Reset()
}

// SoundControl is a global control
// for all sounds.
type SoundControl interface {
	// SetVolume sets the playback volume
	// for a given sound group, between 0.0 and 1.0
	// inclusively. Using an empty string for the sound
	// group will apply the volume to all groups.
	SetVolume(string, float64)

	// Volume returns the playback volume
	// for a given sound group.
	Volume(string) float64
}
