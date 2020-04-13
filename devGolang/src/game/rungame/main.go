package main

import (
	"game/rungame/global"
	"game/rungame/scenemanager"
	"game/rungame/scenes"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scenemanager.SetScene(&scenes.StartScene{})

	var err error = ebiten.Run(scenemanager.Update, global.ScreenWidth,
		global.ScreenHeight, 2.0, global.GameTitleText)

	if err != nil {
		log.Fatalf("RunGame ebiten run error : %v", err)
	}
}
