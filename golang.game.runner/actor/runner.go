package actor

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"golang.game.runner/animation"
	"golang.game.runner/consts"
)

// RunnerState runner state
type RunnerState int

const (
	// Idle idle state
	Idle RunnerState = iota
	// Runnig running state
	Running
)

// Runner runner actor
type Runner struct {
	X, Y      float64
	animation *animation.Handler
	state     RunnerState
}

// NewRunner make a new runner
func NewRunner(x, y float64) *Runner {
	r := &Runner{}
	r.X, r.Y = x, y
	r.state = Idle
	r.animation = animation.New()

	// 이미지 로드
	actorImg, _, err := ebitenutil.NewImageFromFile("./assets/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	// make idle animation sprites
	sprites := make([]*ebiten.Image, consts.IdleFrames)
	for i := 0; i < consts.IdleFrames; i++ {
		sx := consts.IdleX + consts.FrameWidth*i
		sy := consts.IdleY

		sprites[i] = actorImg.SubImage(image.Rect(sx, sy, sx+consts.FrameWidth, sy+consts.FrameHeight)).(*ebiten.Image)
	}
	r.animation.Add("Idle", sprites, consts.IdleAnimSpeed)

	// make running animation sprites
	sprites = make([]*ebiten.Image, consts.RunningFrames)
	for i := 0; i < consts.RunningFrames; i++ {
		sx := consts.RunningX + consts.FrameWidth*i
		sy := consts.RunningY

		sprites[i] = actorImg.SubImage(image.Rect(sx, sy, sx+consts.FrameWidth, sy+consts.FrameHeight)).(*ebiten.Image)
	}
	r.animation.Add("Run", sprites, consts.RunningAnimSpeed)

	return r
}

// SetState change runner state
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
