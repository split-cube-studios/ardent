package common

// Sound is a collection of audio data
type Sound struct {
	// Group is the control group
	// the Sound belongs to.
	Group string

	// Options are the different audio
	// tracks that may represent the same sound.
	// A random option will be selected each time
	// the sound is played.
	Options [][]byte
}
