package scenes

import (
	"game/puzzle/global"
	"game/puzzle/scenemanager"
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

/*GameScene :*/
type GameScene struct {
	bgImg *ebiten.Image
	// subImages image array of cropped game image
	subImages [global.PuzzleColumns * global.PuzzleRows]*ebiten.Image
	// status of game board
	board [global.PuzzleColumns][global.PuzzleRows]int
	// save for blank area
	blankX, blankY int
}

/*Startup :*/
func (g *GameScene) Startup() {
	println("GameScene Startup")
	var err error
	g.bgImg, _, err = ebitenutil.NewImageFromFile("./images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read monalisa.png error : %v", err)
	}

	// width of one image
	width := global.ScreenWidth / global.PuzzleColumns
	// heigth of one image
	height := global.ScreenHeight / global.PuzzleRows

	// cropping image
	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			g.subImages[j*global.PuzzleColumns+i] = g.bgImg.SubImage(
				image.Rect(i*width, j*height, i*width+width, j*height+height)).(*ebiten.Image)
		}
	}

	// initialize for blank area
	g.blankX = global.PuzzleColumns - 1
	g.blankY = global.PuzzleRows - 1

	// shuffle of card
	var arr []int = make([]int, global.PuzzleColumns*global.PuzzleRows-1)
	var idx int = 0
	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			if i == global.PuzzleColumns-1 && j == global.PuzzleRows-1 {
				continue
			}
			arr[j*global.PuzzleColumns+i] = idx
			idx++
		}
	}
	// slice
	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			if i == g.blankX && j == g.blankY {
				g.board[i][j] = -1
				continue
			}
			idx = rand.Intn(len(arr))
			g.board[i][j] = arr[idx]
			arr = append(arr[:idx], arr[idx+1:]...)
		}
	}

	// initialize for game board (complete)
	// for i := 0; i < global.PuzzleColumns; i++ {
	// 	for j := 0; j < global.PuzzleRows; j++ {
	// 		if i == g.blankX && j == g.blankY {
	// 			g.board[i][j] = -1
	// 		} else {
	// 			g.board[i][j] = j*global.PuzzleColumns + i
	// 		}
	// 	}
	// }
}

/*Update :*/
func (g *GameScene) Update(screen *ebiten.Image) error {
	// kev event
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		if g.blankY > 0 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX][g.blankY-1]
			g.board[g.blankX][g.blankY-1] = -1
			g.blankY--
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		if g.blankY < global.PuzzleRows-1 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX][g.blankY+1]
			g.board[g.blankX][g.blankY+1] = -1
			g.blankY++
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		if g.blankX > 0 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX-1][g.blankY]
			g.board[g.blankX-1][g.blankY] = -1
			g.blankX--
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		if g.blankX < global.PuzzleColumns-1 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX+1][g.blankY]
			g.board[g.blankX+1][g.blankY] = -1
			g.blankX++
		}
	}

	width := global.ScreenWidth / global.PuzzleColumns
	height := global.ScreenHeight / global.PuzzleRows

	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {

			if g.board[i][j] == -1 {
				continue
			}

			x := i * width
			y := j * height

			// The previous empty option struct
			opts := &ebiten.DrawImageOptions{}
			// Add the Translate effect to the option struct
			opts.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(g.subImages[g.board[i][j]], opts)
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		println("GameScene Update")
		scenemanager.SetScene(&StartScene{})
	}

	return nil
}
