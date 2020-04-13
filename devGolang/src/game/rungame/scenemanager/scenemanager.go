package scenemanager

import "github.com/hajimehoshi/ebiten"

/*Scene : scene of game*/
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

// package initialize
func init() {
	manager = &scenemanager{}
}

/*SetScene : */
func SetScene(scene Scene) {
	manager.currentScene = scene
	manager.currentScene.Startup()
}

/*Update : */
func Update(screen *ebiten.Image) error {
	if manager.currentScene != nil {
		return manager.currentScene.Update(screen)
	}
	return nil
}
