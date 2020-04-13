package scenes

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"game/ebiten/scenemanager"
)

/*StartScene : */
type StartScene struct {
	startImg *ebiten.Image
}

/*Startup : 이미지 로딩*/
func (s *StartScene) Startup() {
	var err error

	s.startImg, _, err = ebitenutil.NewImageFromFile("./images/start.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read start.png file error : %v", err)
	}
}

/*Update : */
func (s *StartScene) Update(screen *ebiten.Image) error {
	// ebitenutil.DebugPrint(s.startImg, "Start Scene")
	screen.DrawImage(s.startImg, nil)

	// mouse input (click)
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// set game scene
		scenemanager.SetScene(&GameScene{})
	}

	return nil
}
