package scenes

import (
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"golang.game.grimpuzzle/consts"
)

// PlayScene scene of game
type PlayScene struct {
	playImg        *ebiten.Image
	subImgs        [consts.PuzzleColumns * consts.PuzzleRows]*ebiten.Image
	boards         [consts.PuzzleColumns][consts.PuzzleRows]int
	blankX, blankY int
}

// StartUp initialize PlayScene
func (s *PlayScene) StartUp() {
	var err error
	// 시작이미지 로드
	s.playImg, _, err = ebitenutil.NewImageFromFile("./assets/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	w := consts.ScreenWidth / consts.PuzzleColumns
	h := consts.ScreenHeight / consts.PuzzleRows
	// 이미지를 컬럼갯수 만큼 쪼갠다.
	for i := 0; i < consts.PuzzleColumns; i++ {
		for j := 0; j < consts.PuzzleRows; j++ {
			s.subImgs[j*consts.PuzzleColumns+i] = s.playImg.SubImage(image.Rect(i*w, j*h, i*w+w, j*h+h)).(*ebiten.Image)
		}
	}

	// 컬럼갯수 만큼 배열을 만들어준다.
	arr := make([]int, consts.PuzzleColumns*consts.PuzzleRows-1)
	idx := 0
	for i := 0; i < consts.PuzzleColumns; i++ {
		for j := 0; j < consts.PuzzleRows; j++ {
			if i == consts.PuzzleColumns-1 && j == consts.PuzzleRows-1 {
				continue
			}
			arr[j*consts.PuzzleColumns+i] = idx
			idx++
		}
	}

	// 비어있는 칸 정의
	s.blankX = consts.PuzzleColumns - 1
	s.blankY = consts.PuzzleRows - 1

	// 순차컬럼을 랜덤컬럼으로 배치한다.
	for i := 0; i < consts.PuzzleColumns; i++ {
		for j := 0; j < consts.PuzzleRows; j++ {
			if i == s.blankX && j == s.blankY {
				s.boards[i][j] = -1
				continue
			}
			idx := rand.Intn(len(arr))
			s.boards[i][j] = arr[idx]
			// 랜덤으로 선택된 idx를 제외하고 배열을 다시 만든다.
			arr = append(arr[:idx], arr[idx+1:]...)
		}
	}
}

// Update PlayScene
func (s *PlayScene) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) { // 위쪽 방향키를 눌렀다 뗐을 경우
		if s.blankY > 0 {
			s.boards[s.blankX][s.blankY] = s.boards[s.blankX][s.blankY-1]
			s.boards[s.blankX][s.blankY-1] = -1
			s.blankY--
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyDown) { // 아래쪽 방향키를 눌렀다 뗐을 경우
		if s.blankY < consts.PuzzleRows-1 {
			s.boards[s.blankX][s.blankY] = s.boards[s.blankX][s.blankY+1]
			s.boards[s.blankX][s.blankY+1] = -1
			s.blankY++
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyLeft) { // 왼쪽 방향키를 눌렀다 뗐을 경우
		if s.blankX > 0 {
			s.boards[s.blankX][s.blankY] = s.boards[s.blankX-1][s.blankY]
			s.boards[s.blankX-1][s.blankY] = -1
			s.blankX--
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyRight) { // 오른쪽 방향키를 눌렀다 뗐을 경우
		if s.blankX < consts.PuzzleColumns-1 {
			s.boards[s.blankX][s.blankY] = s.boards[s.blankX+1][s.blankY]
			s.boards[s.blankX+1][s.blankY] = -1
			s.blankX++
		}
	}

	w := consts.ScreenWidth / consts.PuzzleColumns
	h := consts.ScreenHeight / consts.PuzzleRows

	for i := 0; i < consts.PuzzleColumns; i++ {
		for j := 0; j < consts.PuzzleRows; j++ {
			if s.boards[i][j] == -1 {
				continue
			}
			x := i * w
			y := j * h

			// The previous empty option struct
			opts := &ebiten.DrawImageOptions{}
			// Add the Translate effect to the option struct.
			opts.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(s.subImgs[s.boards[i][j]], opts)
		}
	}

	return nil
}
