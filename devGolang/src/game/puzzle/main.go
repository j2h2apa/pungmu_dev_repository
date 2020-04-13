package main

import (
	"game/puzzle/global"
	"game/puzzle/scenemanager"
	"game/puzzle/scenes"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scenemanager.SetScene(&scenes.StartScene{})

	var err error = ebiten.Run(scenemanager.Update, global.ScreenWidth,
		global.ScreenHeight, 1.0, "Puzzle")

	if err != nil {
		log.Fatalf("Puzzle ebiten run error : %v", err)
	}
}
