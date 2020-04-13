package scenes

import (
	"game/rungame/actor"
	"game/rungame/font"
	"game/rungame/global"
	"game/rungame/scenemanager"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

/*StartScene :*/
type StartScene struct {
	backImg *ebiten.Image
	runner  *actor.Runner
}

/*Startup :*/
func (s *StartScene) Startup() {

	s.runner = actor.NewRunner(0, global.ScreenHeight/2)

	var err error
	s.backImg, _, err = ebitenutil.NewImageFromFile("images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("StartScene read background image error : %v", err)
	}
	s.runner.SetState(actor.Idle)
}

// frameCount count of current frame
var frameCount = 0

/*Update :*/
func (s *StartScene) Update(screen *ebiten.Image) error {
	frameCount++

	// draw background image
	screen.DrawImage(s.backImg, nil)
	s.runner.Update(screen)

	// text of start screen
	width := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText, global.ScreenWidth/2-width/2,
		global.ScreenHeight/2, 2, color.White)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}
