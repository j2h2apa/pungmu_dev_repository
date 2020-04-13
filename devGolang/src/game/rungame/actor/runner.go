package actor

import (
	"game/rungame/animation"
	"game/rungame/global"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// RunnerState runner state
type RunnerState int

const (
	// Idle state of idle
	Idle RunnerState = iota
	// Running state of running
	Running
)

// Runner runner actor
type Runner struct {
	X, Y      float64
	state     RunnerState
	animation *animation.Handler
}

// NewRunner make a new runner
func NewRunner(x, y float64) *Runner {
	var r *Runner = &Runner{}

	r.X, r.Y = x, y
	r.state = Idle
	r.animation = animation.New()

	runnerImg, _, err := ebitenutil.NewImageFromFile("images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read runner image error : %v", err)
	}

	var sprites []*ebiten.Image = make([]*ebiten.Image, global.IdleFrames)
	for i := 0; i < global.IdleFrames; i++ {
		sx := global.IdleX + global.FrameWidth*i
		sy := global.IdleY
		sprites[i] = runnerImg.SubImage(image.Rect(sx, sy,
			sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	}
	r.animation.Add("Idle", sprites, global.IdleAnimSpeed)

	sprites = make([]*ebiten.Image, global.RunningFrames)
	for i := 0; i < global.RunningFrames; i++ {
		sx := global.RunningX + global.FrameWidth*i
		sy := global.RunningY
		sprites[i] = runnerImg.SubImage(image.Rect(sx, sy,
			sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	}
	r.animation.Add("Run", sprites, global.RunningAnimSpeed)

	return r
}

// SetState change and play for runner state
func (r *Runner) SetState(state RunnerState) {
	r.state = state
	switch r.state {
	case Idle:
		r.animation.Play("Idle")
	case Running:
		r.animation.Play("Run")
	}
}

// Update Draw runner
func (r *Runner) Update(screen *ebiten.Image) {
	r.animation.Update(screen, r.X, r.Y)
}
