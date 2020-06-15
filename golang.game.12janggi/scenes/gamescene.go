package scenes

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"golang.game.12janggi/consts"
	"golang.game.12janggi/scenemanager"
)

// GimulType defined type aliasing
type GimulType int

// TeamType defined type aliasing
type TeamType int

const (
	GimulNone GimulType = -1 + iota
	GimulGreenKing
	GimulGreenJa
	GimulGreenJang
	GimulGreenSang
	GimulRedKing
	GimulRedJa
	GimulRedJang
	GimulRedSang
	GimulMax // 이미지 최종 갯수
)
const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
)

type GameScene struct {
	board       [consts.BoardWidth][consts.BoardHeight]GimulType // gimul의 위치좌표
	bgImg       *ebiten.Image                                    // 게임판 이미지
	gimulImgs   [GimulMax]*ebiten.Image                          // gimul 이미지
	selectedImg *ebiten.Image                                    // gimul 선택 이미지
	isSelected  bool                                             // gimul 선택 여부
	selectedX   int                                              // 선택 확인 위치(X) 값
	selectedY   int                                              // 선택 확인 위치(Y) 값
	currentTeam TeamType                                         // 현재 턴 팀
	gameOver    bool                                             // 게임종료 여부
}

func GetTeamType(gimulType GimulType) TeamType {
	if gimulType == GimulGreenJa || gimulType == GimulGreenJang || gimulType == GimulGreenSang || gimulType == GimulGreenKing {
		return TeamGreen
	}

	if gimulType == GimulRedJa || gimulType == GimulRedJang || gimulType == GimulRedSang || gimulType == GimulRedKing {
		return TeamRed
	}

	return TeamNone
}

func (s *GameScene) moveGimul(prevX, prevY, tagX, tagY int) {
	if s.isMovable(prevX, prevY, tagX, tagY) {
		//게임종료 체크
		s.OnGameOver(s.board[tagX][tagY])

		//Gimul위치 이동
		s.board[prevX][prevY], s.board[tagX][tagY] = GimulNone, s.board[prevX][prevY]
		s.isSelected = false

		//턴 넘김
		if s.currentTeam == TeamGreen {
			s.currentTeam = TeamRed
		} else {
			s.currentTeam = TeamGreen
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (s *GameScene) isMovable(prevX, prevY, tagX, tagY int) bool {
	// 같은 편인지 체크
	if GetTeamType(s.board[prevX][prevY]) == GetTeamType(s.board[tagX][tagY]) {
		return false
	}

	// 최소범위를 벗어나는지 체크
	if tagX < 0 || tagY < 0 {
		return false
	}

	// 최대범위를 벗어나는지 체크
	if tagX >= consts.BoardWidth || tagY >= consts.BoardHeight {
		return false
	}

	// gimul 이동 체크
	switch s.board[prevX][prevY] {
	case GimulGreenJa:
		//x측 앞으로 1칸씩만 이동 가능
		return prevY == tagY && prevX+1 == tagX
	case GimulRedJa:
		//x측 앞으로 1칸씩만 이동 가능
		return prevY == tagY && prevX-1 == tagX
	case GimulGreenJang, GimulRedJang:
		//x,y측으로 1칸 이동 가능(4방향으로 1칸씩 이동 가능)
		return abs(prevX-tagX)+abs(prevY-tagY) == 1
	case GimulGreenSang, GimulRedSang:
		//대각선 1칸 이동 가능(x측 1칸, y측 1칸)
		return (abs(prevX-tagX) == 1 && abs(prevY-tagY) == 1)
	case GimulGreenKing, GimulRedKing:
		//x,y측으로 1칸 또는 대각선 1칸 이동가능(8방방향으로 1칸씩 이동 가능)
		return (abs(prevX-tagX) == 1 || abs(prevY-tagY) == 1)
	}

	return false
}

// OnGameOver calls when gimul is died
func (s *GameScene) OnGameOver(gt GimulType) {
	if gt == GimulGreenKing || gt == GimulRedKing {
		s.gameOver = true
		// GameOverScene로 화면 전환
		scenemanager.SetScene(&GameOverScene{})
	}
}

func (s *GameScene) StartUp() {
	s.gameOver = false
	s.currentTeam = TeamGreen

	var err error
	// 배경이미지 로드
	s.bgImg, _, err = ebitenutil.NewImageFromFile("./assets/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	/* 이미지 로딩 ------------------------------------------------------------------------------------------------- */
	// GreenTeam 이미지 로드
	s.gimulImgs[GimulGreenKing], _, err = ebitenutil.NewImageFromFile("./assets/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.gimulImgs[GimulGreenJa], _, err = ebitenutil.NewImageFromFile("./assets/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.gimulImgs[GimulGreenJang], _, err = ebitenutil.NewImageFromFile("./assets/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.gimulImgs[GimulGreenSang], _, err = ebitenutil.NewImageFromFile("./assets/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	// RedTeam 이미지 로드
	s.gimulImgs[GimulRedKing], _, err = ebitenutil.NewImageFromFile("./assets/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.gimulImgs[GimulRedJa], _, err = ebitenutil.NewImageFromFile("./assets/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.gimulImgs[GimulRedJang], _, err = ebitenutil.NewImageFromFile("./assets/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.gimulImgs[GimulRedSang], _, err = ebitenutil.NewImageFromFile("./assets/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.selectedImg, _, err = ebitenutil.NewImageFromFile("./assets/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	/* ------------------------------------------------------------------------------------------------------------- */

	/* Initialze set board --------------------- */
	for i := 0; i < consts.BoardWidth; i++ {
		for j := 0; j < consts.BoardHeight; j++ {
			s.board[i][j] = GimulNone
		}
	}
	s.board[0][0] = GimulGreenSang
	s.board[0][1] = GimulGreenKing
	s.board[0][2] = GimulGreenJang
	s.board[1][1] = GimulGreenJa

	s.board[3][0] = GimulRedSang
	s.board[3][1] = GimulRedKing
	s.board[3][2] = GimulRedJang
	s.board[2][1] = GimulRedJa
	/* ---------------------------------------- */
}

func (s *GameScene) Update(screen *ebiten.Image) error {
	//게임판을 화면에 그린다.
	screen.DrawImage(s.bgImg, nil)

	if s.gameOver {
		return nil
	}

	// 마우스 입력(왼클릭) 핸들링
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()                 // 클릭 했을 때, 현재 마우스의 위치
		i, j := x/consts.GridWidth, y/consts.GridHeight // 보드의 위치(board[i][j])

		if !s.isSelected {
			if s.board[i][j] != GimulNone && s.currentTeam == GetTeamType(s.board[i][j]) {
				s.selectedX, s.selectedY = i, j
				s.isSelected = true
			}
		} else {
			if i == s.selectedX && j == s.selectedY {
				s.isSelected = false
			} else {
				// Move Selected Grid
				s.moveGimul(s.selectedX, s.selectedY, i, j)
			}
		}
	}

	// gimul들을 그려준다.
	for i := 0; i < consts.BoardWidth; i++ {
		for j := 0; j < consts.BoardHeight; j++ {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(consts.GimulStartX+consts.GridWidth*i), float64(consts.GimulStartY+consts.GridHeight*j))

			switch s.board[i][j] {
			case GimulGreenKing:
				// draw GreenKing
				screen.DrawImage(s.gimulImgs[GimulGreenKing], opts)
			case GimulGreenJa:
				// draw GreenJa
				screen.DrawImage(s.gimulImgs[GimulGreenJa], opts)
			case GimulGreenJang:
				// draw GreenJang
				screen.DrawImage(s.gimulImgs[GimulGreenJang], opts)
			case GimulGreenSang:
				// draw GreenSang
				screen.DrawImage(s.gimulImgs[GimulGreenSang], opts)
			case GimulRedKing:
				// draw RedKing
				screen.DrawImage(s.gimulImgs[GimulRedKing], opts)
			case GimulRedJa:
				// draw RedJa
				screen.DrawImage(s.gimulImgs[GimulRedJa], opts)
			case GimulRedJang:
				// draw RedJang
				screen.DrawImage(s.gimulImgs[GimulRedJang], opts)
			case GimulRedSang:
				// draw RedSang
				screen.DrawImage(s.gimulImgs[GimulRedSang], opts)
			}
		}
	}

	if s.isSelected {
		// The previous empty option struct
		opts := &ebiten.DrawImageOptions{}
		// Add the Translate effect to the option struct.
		opts.GeoM.Translate(
			float64(consts.GimulStartX+consts.GridWidth*s.selectedX),
			float64(consts.GimulStartY+consts.GridHeight*s.selectedY),
		)

		screen.DrawImage(s.selectedImg, opts)
	}

	return nil
}
