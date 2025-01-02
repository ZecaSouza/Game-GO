package game

import (
	"my-game/assets/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bounds := image.Bounds()

	halfw := float64(bounds.Dx()) / 2 //metade largura nave
	halfh := float64(bounds.Dy()) / 2 //metade altura nave

	position.X -= halfw
	position.Y -= halfh

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 10.0

	l.position.Y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(l.position.X, l.position.Y)
	screen.DrawImage(l.image, op)
}

func (l *Laser) Collider() *Rect {
	bounds := l.image.Bounds()


	return NewRect(l.position.X, 
		           l.position.Y, 
					float64(bounds.Dx()),
					float64(bounds.Dy()),
				)
}