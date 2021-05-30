package engine

// LayoutFunc is a function type
// that is responsible for handling
// screen resizing.
type LayoutFunc func(w int, h int) (nw int, nh int)

// LayoutFill returns a LayoutFunc that fills the screen
// without stretching. Scaling is based on an original
// virtual width and virtual height.
func LayoutFill(virtualWidth, virtualHeight int) LayoutFunc {
	return func(w, h int) (nw, nh int) {
		if w > h {
			nw = (virtualHeight * w) / h
			nh = virtualHeight
		} else {
			nw = virtualWidth
			nh = (virtualWidth * h) / w
		}
		return
	}
}

// LayoutFit returns a LayoutFunc that fits the virtual viewport
// to the screen. Black bars may be shown around the edges
// if the screen's aspect ratio does not match the virtual aspect ratio.
func LayoutFit(virtualWidth, virtualHeight int) LayoutFunc {
	return func(_, _ int) (int, int) {
		return virtualWidth, virtualHeight
	}
}

// LayoutDefault is the default LayoutFunc. This LayoutFunc
// is used when none is provided. The virtual resolution
// is set to the screen resolution.
func LayoutDefault(w, h int) (int, int) {
	return w, h
}
