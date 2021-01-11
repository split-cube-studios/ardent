//+build !headless

package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/split-cube-studios/ardent/engine/input"
	"github.com/split-cube-studios/ardent/engine/input/raw"
)

var toEbitenKey = map[int]ebiten.Key{
	raw.Key0:            ebiten.Key0,
	raw.Key1:            ebiten.Key1,
	raw.Key2:            ebiten.Key2,
	raw.Key3:            ebiten.Key3,
	raw.Key4:            ebiten.Key4,
	raw.Key5:            ebiten.Key5,
	raw.Key6:            ebiten.Key6,
	raw.Key7:            ebiten.Key7,
	raw.Key8:            ebiten.Key8,
	raw.Key9:            ebiten.Key9,
	raw.KeyA:            ebiten.KeyA,
	raw.KeyB:            ebiten.KeyB,
	raw.KeyC:            ebiten.KeyC,
	raw.KeyD:            ebiten.KeyD,
	raw.KeyE:            ebiten.KeyE,
	raw.KeyF:            ebiten.KeyF,
	raw.KeyG:            ebiten.KeyG,
	raw.KeyH:            ebiten.KeyH,
	raw.KeyI:            ebiten.KeyI,
	raw.KeyJ:            ebiten.KeyJ,
	raw.KeyK:            ebiten.KeyK,
	raw.KeyL:            ebiten.KeyL,
	raw.KeyM:            ebiten.KeyM,
	raw.KeyN:            ebiten.KeyN,
	raw.KeyO:            ebiten.KeyO,
	raw.KeyP:            ebiten.KeyP,
	raw.KeyQ:            ebiten.KeyQ,
	raw.KeyR:            ebiten.KeyR,
	raw.KeyS:            ebiten.KeyS,
	raw.KeyT:            ebiten.KeyT,
	raw.KeyU:            ebiten.KeyU,
	raw.KeyV:            ebiten.KeyV,
	raw.KeyW:            ebiten.KeyW,
	raw.KeyX:            ebiten.KeyX,
	raw.KeyY:            ebiten.KeyY,
	raw.KeyZ:            ebiten.KeyZ,
	raw.KeyApostrophe:   ebiten.KeyApostrophe,
	raw.KeyBackslash:    ebiten.KeyBackslash,
	raw.KeyBackspace:    ebiten.KeyBackspace,
	raw.KeyCapsLock:     ebiten.KeyCapsLock,
	raw.KeyComma:        ebiten.KeyComma,
	raw.KeyDelete:       ebiten.KeyDelete,
	raw.KeyDown:         ebiten.KeyDown,
	raw.KeyEnd:          ebiten.KeyEnd,
	raw.KeyEnter:        ebiten.KeyEnter,
	raw.KeyEqual:        ebiten.KeyEqual,
	raw.KeyEscape:       ebiten.KeyEscape,
	raw.KeyF1:           ebiten.KeyF1,
	raw.KeyF2:           ebiten.KeyF2,
	raw.KeyF3:           ebiten.KeyF3,
	raw.KeyF4:           ebiten.KeyF4,
	raw.KeyF5:           ebiten.KeyF5,
	raw.KeyF6:           ebiten.KeyF6,
	raw.KeyF7:           ebiten.KeyF7,
	raw.KeyF8:           ebiten.KeyF8,
	raw.KeyF9:           ebiten.KeyF9,
	raw.KeyF10:          ebiten.KeyF10,
	raw.KeyF11:          ebiten.KeyF11,
	raw.KeyF12:          ebiten.KeyF12,
	raw.KeyGraveAccent:  ebiten.KeyGraveAccent,
	raw.KeyHome:         ebiten.KeyHome,
	raw.KeyInsert:       ebiten.KeyInsert,
	raw.KeyKP0:          ebiten.KeyKP0,
	raw.KeyKP1:          ebiten.KeyKP1,
	raw.KeyKP2:          ebiten.KeyKP2,
	raw.KeyKP3:          ebiten.KeyKP3,
	raw.KeyKP4:          ebiten.KeyKP4,
	raw.KeyKP5:          ebiten.KeyKP5,
	raw.KeyKP6:          ebiten.KeyKP6,
	raw.KeyKP7:          ebiten.KeyKP7,
	raw.KeyKP8:          ebiten.KeyKP8,
	raw.KeyKP9:          ebiten.KeyKP9,
	raw.KeyKPAdd:        ebiten.KeyKPAdd,
	raw.KeyKPDecimal:    ebiten.KeyKPDecimal,
	raw.KeyKPDivide:     ebiten.KeyKPDivide,
	raw.KeyKPEnter:      ebiten.KeyKPEnter,
	raw.KeyKPEqual:      ebiten.KeyKPEqual,
	raw.KeyKPMultiply:   ebiten.KeyKPMultiply,
	raw.KeyKPSubtract:   ebiten.KeyKPSubtract,
	raw.KeyLeft:         ebiten.KeyLeft,
	raw.KeyLeftBracket:  ebiten.KeyLeftBracket,
	raw.KeyMenu:         ebiten.KeyMenu,
	raw.KeyMinus:        ebiten.KeyMinus,
	raw.KeyNumLock:      ebiten.KeyNumLock,
	raw.KeyPageDown:     ebiten.KeyPageDown,
	raw.KeyPageUp:       ebiten.KeyPageUp,
	raw.KeyPause:        ebiten.KeyPause,
	raw.KeyPeriod:       ebiten.KeyPeriod,
	raw.KeyPrintScreen:  ebiten.KeyPrintScreen,
	raw.KeyRight:        ebiten.KeyRight,
	raw.KeyRightBracket: ebiten.KeyRightBracket,
	raw.KeyScrollLock:   ebiten.KeyScrollLock,
	raw.KeySemicolon:    ebiten.KeySemicolon,
	raw.KeySlash:        ebiten.KeySlash,
	raw.KeySpace:        ebiten.KeySpace,
	raw.KeyTab:          ebiten.KeyTab,
	raw.KeyUp:           ebiten.KeyUp,
	raw.KeyAlt:          ebiten.KeyAlt,
	raw.KeyControl:      ebiten.KeyControl,
	raw.KeyShift:        ebiten.KeyShift,
}

var _ input.KeySource = (*KeySource)(nil)

type KeySource struct{}

func (s *KeySource) IsAnyPressed() bool {
	for _, v := range toEbitenKey {
		if ebiten.IsKeyPressed(v) {
			return true
		}
	}

	return false
}

func (s *KeySource) IsAnyJustPressed() bool {
	for _, v := range toEbitenKey {
		if inpututil.IsKeyJustPressed(v) {
			return true
		}
	}

	return false
}

func (s *KeySource) StateOf(in input.Input) input.State {
	var value float64

	if s.IsPressed(in) {
		value = 1.0
	}

	return input.State{
		Type:        input.Keyboard,
		Input:       in,
		Value:       value,
		JustPressed: s.IsJustPressed(in),
	}
}

func (s *KeySource) IsPressed(in input.Input) bool {
	return ebiten.IsKeyPressed(toEbitenKey[in])
}

func (s *KeySource) IsJustPressed(in input.Input) bool {
	return inpututil.IsKeyJustPressed(toEbitenKey[in])
}

func (s *KeySource) IsJustReleased(in input.Input) bool {
	return inpututil.IsKeyJustReleased(toEbitenKey[in])
}
