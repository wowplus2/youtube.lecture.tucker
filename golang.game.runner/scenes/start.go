package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"golang.game.runner/actor"
	"golang.game.runner/animation"
	"golang.game.runner/consts"
	"golang.game.runner/font"
	"golang.game.runner/scenes/manager"
)

// StartScene scene of init game scene
type StartScene struct {
	backImg   *ebiten.Image
	animation *animation.Handler
	runner    *actor.Runner
}

// StartUp initialize StartScene
func (s *StartScene) StartUp() {
	s.runner = actor.NewRunner(consts.ScreenWidth/4, consts.ScreenHeight/2)

	var err error
	//배경 이미지 로드
	s.backImg, _, err = ebitenutil.NewImageFromFile("./assets/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.runner.SetState(actor.Idle)
}

// Update StartScene
func (s *StartScene) Update(screen *ebiten.Image) error {
	// 배경그리기
	screen.DrawImage(s.backImg, nil)

	// Idle Animation
	s.runner.Update(screen)

	//텍스트의 길이를 가져온다.
	width := font.TextWidth(consts.StartSceneText, 2)
	//화면에 텍스트를 그린다.
	font.DrawTextWithShadow(screen, consts.StartSceneText, consts.ScreenWidth/2-width/2, consts.ScreenHeight/2, 2, color.White)

	//마우스 입력(왼클릭) 핸들링
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		//PlayScene 화면 전환
		manager.SetScene(&PlayScene{})
	}

	return nil
}
