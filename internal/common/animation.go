package common

// Animation holds extra info for an animated Asset.
type Animation struct {
	Fps, Start, End uint16
	Loop            bool
}
