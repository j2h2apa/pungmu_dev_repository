package scenemanager

import "github.com/hajimehoshi/ebiten"

/*Scene :*/
type Scene interface {
	Startup()
	Update(*ebiten.Image) error
}

type scenemanager struct {
	currentScene Scene
}

var (
	manager *scenemanager
)

// package initializition function on import
func init() {
	manager = &scenemanager{}
}

/*SetScene : */
func SetScene(scene Scene) {
	manager.currentScene = scene
	scene.Startup()
}

/*Update :*/
func Update(screen *ebiten.Image) error {
	if manager.currentScene != nil {
		return manager.currentScene.Update(screen)
	}
	return nil
}
