package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"

	"golang.game.grimpuzzle/consts"
	"golang.game.grimpuzzle/scenes"
	"golang.game.grimpuzzle/scenes/manager"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// 시작 scene 호출
	manager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(manager.Update, consts.ScreenWidth, consts.ScreenHeight, 1.0, "그림 퍼즐")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
