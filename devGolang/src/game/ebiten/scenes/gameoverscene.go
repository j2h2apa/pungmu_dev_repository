package scenes

import (
	"game/ebiten/scenemanager"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

/*GameoverScene :*/
type GameoverScene struct {
	gameoverImg *ebiten.Image
}

/*Startup :*/
func (g *GameoverScene) Startup() {
	var err error
	g.gameoverImg, _, err = ebitenutil.NewImageFromFile("./images/gameover.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read gameover.png file error : %v", err)
	}
}

/*Update :*/
func (g *GameoverScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(g.gameoverImg, nil)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&StartScene{})
	}

	return nil
}
