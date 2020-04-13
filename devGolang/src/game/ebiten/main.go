package main

import (
	"log"

	"game/ebiten/global"
	"game/ebiten/scenemanager"
	"game/ebiten/scenes"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	var err error

	scenemanager.SetScene(&scenes.StartScene{})
	err = ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "12 Janggi")
	if err != nil {
		log.Fatalf("Ebiten run error : %v", err)
	}
}
