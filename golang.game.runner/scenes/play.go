package scenes

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"golang.game.runner/actor"
	"golang.game.runner/animation"
	"golang.game.runner/bgscroller"
	"golang.game.runner/consts"
)

// PlayScene scene of init game scene
type PlayScene struct {
	bgscroller *bgscroller.Scroller
	animation  *animation.Handler
	runner     *actor.Runner
}

// StartUp initialize PlayScene
func (s *PlayScene) StartUp() {
	s.runner = actor.NewRunner(consts.ScreenWidth/4, consts.ScreenHeight/2)

	//배경 이미지 로드
	backImg, _, err := ebitenutil.NewImageFromFile("./assets/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.bgscroller = bgscroller.New(backImg, consts.BgScrollSpeed)
	s.runner.SetState(actor.Running)
}

// Update PlayScene
func (s *PlayScene) Update(screen *ebiten.Image) error {

	s.bgscroller.Update(screen)

	// Running Animation
	s.runner.Update(screen)

	return nil
}
