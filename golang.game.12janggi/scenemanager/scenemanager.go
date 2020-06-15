package scenemanager

import (
	"github.com/hajimehoshi/ebiten"
)

// Scene interface
type Scene interface {
	StartUp()
	Update(*ebiten.Image) error
}

type scenemanager struct {
	currentScene Scene //현재 scene의 instance
}

var manager *scenemanager

func init() {
	manager = &scenemanager{}
}

func Update(screen *ebiten.Image) error {
	if manager.currentScene != nil {
		return manager.currentScene.Update(screen)
	}

	return nil
}

func SetScene(scene Scene) {
	manager.currentScene = scene
	scene.StartUp()
}
