package game

import (
	"fmt"

	"image/color"
	"my-game/assets/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	Player           *Player
	Laser            []*Laser
	meteor           []*Meteor
	meteorSpawnTimer *Timer
	score            int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
	}
	player := NewPlayer(g)
	g.Player = player
	return g
}

func (g *Game) Update() error {
	g.Player.Update()

	for _, laser := range g.Laser {
		laser.Update()
	}

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteor = append(g.meteor, m)
	}

	for _, meteor := range g.meteor {
		meteor.Update()
	}

	for _, m := range g.meteor {
		if m.Collider().Intersects(g.Player.Collider()) {
			g.Reset()
			return nil
		}
	}

	for i, m := range g.meteor {

		for j, l := range g.Laser {
			if m.Collider().Intersects(l.Collider()) {
				g.meteor = append(g.meteor[:i], g.meteor[i+1:]...)

				g.Laser = append(g.Laser[:j], g.Laser[j+1:]...)

				g.score += 1
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)

	for _, laser := range g.Laser {
		laser.Draw(screen)
	}

	for _, meteor := range g.meteor {
		meteor.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos : %d", g.score), assets.FontUi, 20, 100, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Shoot(laser *Laser) {
	g.Laser = append(g.Laser, laser)
}

func (g *Game) Reset() {
	g.Player = NewPlayer(g)
	g.Laser = nil
	g.meteor = nil
	g.meteorSpawnTimer.Reset()
	g.score = 0
}
