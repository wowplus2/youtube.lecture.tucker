package scenes

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"golang.game.12janggi/scenemanager"
)

type GameOverScene struct {
	gameoverImg *ebiten.Image
}

func (s *GameOverScene) StartUp() {
	var err error
	// 시작이미지 로드
	s.gameoverImg, _, err = ebitenutil.NewImageFromFile("./assets/gameover.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
}

func (s *GameOverScene) Update(screen *ebiten.Image) error {
	//시작화면을 화면에 그린다.
	screen.DrawImage(s.gameoverImg, nil)
	//ebitenutil.DebugPrint(screen, "Game Over")

	// 마우스 입력(왼클릭) 핸들링
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// 게임 플레이 scene 호출
		scenemanager.SetScene(&StartScene{})
	}

	return nil
}
