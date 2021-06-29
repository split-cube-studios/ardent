package engine

type Button struct {
	BaseUIComponent

	img Image
}

func NewButton(img Image) *Button {
	return &Button{
		img: img,
	}
}

func (b *Button) Draw() Image {
	return b.img
}
