package bgscroller

import (
	"github.com/hajimehoshi/ebiten"
	"golang.game.runner/consts"
)

// Scroller horizontal background scroller
type Scroller struct {
	bgImg  *ebiten.Image
	speed  int
	frames int
}

// New make new background scroll image
func New(bgimg *ebiten.Image, speed int) *Scroller {
	return &Scroller{bgImg: bgimg, speed: speed, frames: 0}
}

// Update horizontal backgroun animation
func (s *Scroller) Update(screen *ebiten.Image) {
	s.frames++

	// 움직이는 배경그리기
	orgBackX, _ := s.bgImg.Size()
	runBackX := (s.frames / consts.BgScrollSpeed) % orgBackX

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(-runBackX), 0)
	screen.DrawImage(s.bgImg, opt)

	opt.GeoM.Translate(float64(orgBackX), 0)
	screen.DrawImage(s.bgImg, opt)
}
