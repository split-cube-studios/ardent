package ebiten

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/split-cube-studios/ardent/internal/common"
)

type Animation struct {
	Image
	state string

	w, h                     uint16
	fpsCounter, frameCounter uint16

	anims map[string]common.Animation
	cache map[uint16]*ebiten.Image

	paused bool
}

func (a *Animation) SetState(state string) {
	if a.state == state {
		return
	}

	a.state = state
	a.Reset()
}

func (a *Animation) SetTickCount(count int) {
	fps := a.anims[a.state].Fps

	a.frameCounter = uint16(count) / (60 / fps)
}

func (a *Animation) Play() {
	a.paused = false
}

func (a *Animation) Pause() {
	a.paused = true
}

func (a *Animation) Reset() {
	a.fpsCounter, a.frameCounter = 0, 0
}

func (a *Animation) tick() {
	if a.paused {
		return
	}

	fps := a.anims[a.state].Fps
	if a.fpsCounter == 60/fps {
		a.frameCounter++
		a.fpsCounter = 0
		return
	}

	a.fpsCounter++
}

func (a *Animation) Size() (int, int) {
	return int(a.w), int(a.h)
}

func (a *Animation) getFrame() *ebiten.Image {
	anim, ok := a.anims[a.state]
	if !ok {
		return nil
	}

	var frameKey uint16
	if !anim.Loop && a.frameCounter >= anim.End-anim.Start {
		frameKey = anim.End
	} else {
		length := anim.End - anim.Start
		if length > 0 {
			frameKey = (a.frameCounter % length) + anim.Start
		} else {
			frameKey = anim.Start
		}
	}

	frame, ok := a.cache[frameKey]
	if ok {
		return frame
	}

	w, _ := a.img.Size()
	xtiles := uint16(w) / a.w
	x := (frameKey % xtiles) * a.w
	y := (frameKey / xtiles) * a.h

	img := a.img.SubImage(
		image.Rect(
			int(x),
			int(y),
			int(x+a.w),
			int(y+a.h),
		),
	).(*ebiten.Image)
	a.cache[frameKey] = img

	return img
}
