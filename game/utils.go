package game

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Vector struct {
	X float64
	Y float64
}

type Rect struct {
	X float64
	Y float64
	W float64
	H float64
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{X: x, Y: y, W: w, H: h}
}

func (r *Rect) Intersects(other *Rect) bool {
	return r.X <= other.maxX() &&
		other.X <= r.maxX() &&
		r.Y <= other.maxY() &&
		other.Y <= r.maxY()
}

func (r *Rect) maxX() float64 {
	return r.X + r.W
}

func (r *Rect) maxY() float64 {
	return r.Y + r.H
}
