package main

//ebiten 라이브러리 이용
import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"golang.game.12janggi/consts"
	"golang.game.12janggi/scenemanager"
	"golang.game.12janggi/scenes"
)

func main() {
	// 시작 scene 호출
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, consts.WindowWidth, consts.WindowHeight, 1.0, "12 장기")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
