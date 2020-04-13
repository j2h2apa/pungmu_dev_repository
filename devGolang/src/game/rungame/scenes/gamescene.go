package scenes

import (
	"game/rungame/actor"
	"game/rungame/bgscroller"
	"game/rungame/global"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

/*GameScene :*/
type GameScene struct {
	bgscroller *bgscroller.Scroller
	runner     *actor.Runner
}

/*Startup : */
func (g *GameScene) Startup() {
	frameCount = 0

	g.runner = actor.NewRunner(0, global.ScreenHeight/2)

	backImg, _, err := ebitenutil.NewImageFromFile("images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("StartScene read background image error : %v", err)
	}

	g.bgscroller = bgscroller.NewScroller(backImg, global.BGScrollSpeed)
	g.runner.SetState(actor.Running)
}

/*Update :*/
func (g *GameScene) Update(screen *ebiten.Image) error {

	g.bgscroller.Update(screen)
	g.runner.Update(screen)

	return nil
}
