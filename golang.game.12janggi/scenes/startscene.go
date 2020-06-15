package scenes

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"golang.game.12janggi/scenemanager"
)

type StartScene struct {
	startImg *ebiten.Image
}

func (s *StartScene) StartUp() {
	var err error
	// 시작이미지 로드
	s.startImg, _, err = ebitenutil.NewImageFromFile("./assets/start.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
}

func (s *StartScene) Update(screen *ebiten.Image) error {
	//시작화면을 화면에 그린다.
	screen.DrawImage(s.startImg, nil)
	//ebitenutil.DebugPrint(screen, "Start Scene")

	// 마우스 입력(왼클릭) 핸들링
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// GameScene로 화면 전환
		scenemanager.SetScene(&GameScene{})
	}

	return nil
}
