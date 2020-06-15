package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"

	"golang.game.runner/consts"
	"golang.game.runner/scenes"
	"golang.game.runner/scenes/manager"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// 시작 scene 호출
	manager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(manager.Update, consts.ScreenWidth, consts.ScreenHeight, 2.0, "런 게임")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
