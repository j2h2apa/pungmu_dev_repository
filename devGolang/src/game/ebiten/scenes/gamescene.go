package scenes

import (
	"game/ebiten/global"
	"game/ebiten/scenemanager"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

/*GimulType : int type aliasing */
type GimulType int

/*iota const 설정 시 자동으로 1씩 증가 효과*/
const (
	GimulTypeNone      GimulType = -1 + iota // -1 + 0
	GimulTypeGreenWang                       // 0
	GimulTypeGreenJa
	GimulTypeGreenJang
	GimulTypeGreenSang
	GimulTypeRedWang
	GimulTypeRedJa
	GimulTypeRedJang
	GimulTypeRedSang
	GimulTypeMax
)

/*TeamType :*/
type TeamType int

/**/
const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
)

/*GameScene :*/
type GameScene struct {
	bgimg       *ebiten.Image
	gameover    bool // bool type default value : false
	board       [global.BoardWidth][global.BoardHeight]GimulType
	gimulImgs   [GimulTypeMax]*ebiten.Image
	selectedImg *ebiten.Image
	selected    bool
	selectedX   int
	selectedY   int
	// 현재 턴 팀
	currentTeam TeamType
}

/*GetTeamType : 장기 타입을 보고 해당 팀 리턴*/
func GetTeamType(gimulType GimulType) TeamType {
	if gimulType == GimulTypeGreenWang ||
		gimulType == GimulTypeGreenJang ||
		gimulType == GimulTypeGreenSang ||
		gimulType == GimulTypeGreenJa {
		return TeamGreen
	}
	if gimulType == GimulTypeRedWang ||
		gimulType == GimulTypeRedJang ||
		gimulType == GimulTypeRedSang ||
		gimulType == GimulTypeRedJa {
		return TeamRed
	}

	return TeamNone
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 이동 기능
// 이전 클릭한 위치와 다를 경우 작동
func (g *GameScene) moveGimul(prevX, prevY, tarX, tarY int) {
	// 이동이 가능한 상황인가
	if g.isMoveable(prevX, prevY, tarX, tarY) {
		g.OnDie(g.board[tarX][tarY])
		g.board[prevX][prevY], g.board[tarX][tarY] = GimulTypeNone, g.board[prevX][prevY]
		g.selected = false
		// 말 이동 후 팀 변경
		if g.currentTeam == TeamGreen {
			g.currentTeam = TeamRed
		} else {
			g.currentTeam = TeamGreen
		}
	}
}

func (g *GameScene) isMoveable(prevX, prevY, tarX, tarY int) bool {
	// as is - to be gimul 이 같다면 같은 팀이므로 이동 불가
	if GetTeamType(g.board[prevX][prevY]) == GetTeamType(g.board[tarX][tarY]) {
		return false
	}

	if tarX < 0 || tarY < 0 {
		return false
	}

	if tarX >= global.BoardWidth || tarY >= global.BoardHeight {
		return false
	}

	switch g.board[prevX][prevY] {
	case GimulTypeGreenJa:
		// ja gimul 은 앞으로 한칸만 가능
		return prevX+1 == tarX && prevY == tarY
	case GimulTypeGreenSang, GimulTypeRedSang:
		// x y 각 1 칸씩의 위치에 있는 경우 이동 가능 (즉 인접한 대각선 위치)
		return abs(prevX-tarX) == 1 && abs(prevY-tarY) == 1
	case GimulTypeGreenJang, GimulTypeRedJang:
		// 가로 이동 또는 세로 이동의 합이 1인 경우 즉 곧은 방향으로만 1칸 이동 가능
		return abs(prevX-tarX)+abs(prevY-tarY) == 1
	case GimulTypeGreenWang, GimulTypeRedWang:
		// 가로 변위 또는 세로 변위가 1인 경우는 모두 이동 가능(주변 1칸 모두 이동가능)
		return abs(prevX-tarX) == 1 || abs(prevY-tarY) == 1
	}

	// 이외의 경우는 모두 빈공간
	return false
}

// OnDie calls when
func (g *GameScene) OnDie(gimulType GimulType) {
	if gimulType == GimulTypeGreenWang ||
		gimulType == GimulTypeRedWang {
		g.gameover = true
		scenemanager.SetScene(&GameoverScene{})
	}
}

/*Startup : */
func (g *GameScene) Startup() {
	var err error
	g.bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read bgimg.png error : %v", err)
	}

	g.gimulImgs[GimulTypeGreenWang], _, err = ebitenutil.NewImageFromFile("images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read green_wang error : %v", err)
	}
	g.gimulImgs[GimulTypeGreenJang], _, err = ebitenutil.NewImageFromFile("images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read green_jang error : %v", err)
	}
	g.gimulImgs[GimulTypeGreenJa], _, err = ebitenutil.NewImageFromFile("images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read green_ja error : %v", err)
	}
	g.gimulImgs[GimulTypeGreenSang], _, err = ebitenutil.NewImageFromFile("images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read green_sang error : %v", err)
	}
	g.gimulImgs[GimulTypeRedWang], _, err = ebitenutil.NewImageFromFile("images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read red_wang error : %v", err)
	}
	g.gimulImgs[GimulTypeRedJang], _, err = ebitenutil.NewImageFromFile("images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read red_jang error : %v", err)
	}
	g.gimulImgs[GimulTypeRedJa], _, err = ebitenutil.NewImageFromFile("images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read red_ja error : %v", err)
	}
	g.gimulImgs[GimulTypeRedSang], _, err = ebitenutil.NewImageFromFile("images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read red_sang error : %v", err)
	}
	g.selectedImg, _, err = ebitenutil.NewImageFromFile("images/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read selected error : %v", err)
	}
	/*image load end*/

	// Initialized board
	for i := 0; i < len(g.board); i++ {
		for j := 0; j < len(g.board[i]); j++ {
			g.board[i][j] = GimulTypeNone
		}
	}

	g.board[0][0] = GimulTypeGreenSang
	g.board[0][1] = GimulTypeGreenWang
	g.board[0][2] = GimulTypeGreenJang
	g.board[1][1] = GimulTypeGreenJa

	g.board[3][0] = GimulTypeRedSang
	g.board[3][1] = GimulTypeRedWang
	g.board[3][2] = GimulTypeRedJang
	g.board[2][1] = GimulTypeRedJa

	g.currentTeam = TeamGreen
}

/*Update :*/
func (g *GameScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(g.bgimg, nil)
	// gameover varify
	if g.gameover {
		return nil
	}

	// input handling
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/global.GridWitdh, y/global.GridHeight

		if i >= 0 && i < global.GridWitdh && j >= 0 && j < global.GridHeight {
			if !g.selected {
				// 장기가 있는 자리이고 현재 팀의 장기인가?
				if g.board[i][j] != GimulTypeNone && g.currentTeam == GetTeamType(g.board[i][j]) {
					g.selected = true
					g.selectedX, g.selectedY = i, j
				}
			} else {
				// 같은 위치 확인
				if g.selectedX == i && g.selectedY == j {
					g.selected = false
				} else {
					// 이동
					g.moveGimul(g.selectedX, g.selectedY, i, j)
				}
			}
		}
	}

	// draw gimul
	for i := 0; i < len(g.board); i++ {
		for j := 0; j < len(g.board[i]); j++ {

			var opts *ebiten.DrawImageOptions = &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(global.GimulStartX+global.GridWitdh*i), float64(global.GimulStartY+global.GridHeight*j))

			switch g.board[i][j] {
			case GimulTypeGreenWang:
				// Draw GimulTypeGreenWang
				screen.DrawImage(g.gimulImgs[GimulTypeGreenWang], opts)
			case GimulTypeGreenJang:
				// Draw GimulTypeGreenJang
				screen.DrawImage(g.gimulImgs[GimulTypeGreenJang], opts)
			case GimulTypeGreenJa:
				// Draw GimulTypeGreenJa
				screen.DrawImage(g.gimulImgs[GimulTypeGreenJa], opts)
			case GimulTypeGreenSang:
				// Draw GimulTypeGreenSang
				screen.DrawImage(g.gimulImgs[GimulTypeGreenSang], opts)
			case GimulTypeRedWang:
				// Draw GimulTypeRedWang
				screen.DrawImage(g.gimulImgs[GimulTypeRedWang], opts)
			case GimulTypeRedJang:
				// Draw GimulTypeRedJang
				screen.DrawImage(g.gimulImgs[GimulTypeRedJang], opts)
			case GimulTypeRedJa:
				// Draw GimulTypeRedJa
				screen.DrawImage(g.gimulImgs[GimulTypeRedJa], opts)
			case GimulTypeRedSang:
				// Draw GimulTypeRedSang
				screen.DrawImage(g.gimulImgs[GimulTypeRedSang], opts)
			}
		}
	}
	// 선택한 포지션에 이미지 변경 (mouse left click image)
	if g.selected {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(global.GimulStartX+global.GridWitdh*g.selectedX),
			float64(global.GimulStartY+global.GridHeight*g.selectedY))
		screen.DrawImage(g.selectedImg, opts)
	}

	return nil
}
