package engine

// Sound represents music or a sound effect.
//
// A Sound may have multiple options for its audio track,
// randomly selection one each time the Sound is played.
//
// Sounds also belong to a specified control group.
// Groups can be modified in batches by a SoundControl.
type Sound interface {
	// Play plays the audio from the current
	// position to the end. An error may be returned
	// if a sound cannot be properly decoded.
	Play() error

	// Loop plays the audio from the current
	// position, and repeats from the start
	// after reaching the end. An error may be returned
	// if a sound cannot be properly decoded.
	Loop() error

	// Pause stops the sound for playing,
	// keeping the current position.
	Pause()

	// Reset seeks to the start of the sound.
	// Reset will also pause the track.
	Reset()

	// Close releases the underlying audio assets.
	// The user must call Close when they are done using the Sound.
	// Close does not need to be called when reusing
	// a given sound.
	Close()
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
