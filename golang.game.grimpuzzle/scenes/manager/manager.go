package manager

import (
	"github.com/hajimehoshi/ebiten"
)

// Scene interface
type Scene interface {
	StartUp()
	Update(*ebiten.Image) error
}

type manager struct {
	currentScene Scene //현재 scene의 instance
}

var mgr *manager

func init() {
	mgr = &manager{}
}

func Update(screen *ebiten.Image) error {
	if mgr.currentScene != nil {
		return mgr.currentScene.Update(screen)
	}

	return nil
}

func SetScene(scene Scene) {
	mgr.currentScene = scene
	scene.StartUp()
}
