package scenes

import (
	"game/puzzle/font"
	"game/puzzle/global"
	"game/puzzle/scenemanager"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

/*StartScene : */
type StartScene struct {
	startImg *ebiten.Image
}

/*Startup :*/
func (s *StartScene) Startup() {
	println("StartScene Startup")

	var err error

	s.startImg, _, err = ebitenutil.NewImageFromFile("./images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read monalisa.png error : %v", err)
	}
}

/*Update : draw image and control mouse button event*/
func (s *StartScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(s.startImg, nil)

	// get font size
	var width int = font.TextWidth(global.StartSceneText, 2)
	// font draw on screen center
	font.DrawTextWithShadow(screen, global.StartSceneText, global.ScreenWidth/2-width/2, global.ScreenHeight/2,
		2, color.Black)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		println("StartScene Update")
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}
