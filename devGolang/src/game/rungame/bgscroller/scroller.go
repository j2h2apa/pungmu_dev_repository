package bgscroller

import "github.com/hajimehoshi/ebiten"

// Scroller horizontal background scroller
type Scroller struct {
	bgimg  *ebiten.Image
	speed  int
	frames int
}

// NewScroller create scroller
func NewScroller(bgimg *ebiten.Image, speed int) *Scroller {
	return &Scroller{bgimg: bgimg, speed: speed, frames: 0}
}

// Update update scroller
func (s *Scroller) Update(screen *ebiten.Image) {
	s.frames++

	bgWidth, _ := s.bgimg.Size()
	var opts *ebiten.DrawImageOptions = &ebiten.DrawImageOptions{}
	backX := (s.frames / s.speed) % bgWidth
	opts.GeoM.Translate(float64(-backX), 0)
	screen.DrawImage(s.bgimg, opts)

	opts.GeoM.Translate(float64(bgWidth), 0)
	screen.DrawImage(s.bgimg, opts)
}
