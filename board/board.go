package board

import (
	"errors"
	"github.com/Wouterbeets/n-puzzle/plog"
	"strconv"
	"strings"
)

const (
	Up = iota
	Down
	Left
	Right
	MAX_SIZE           = 42
	MANHATTAN          = 1
	MANHATTAN_CONFLICT = 2
	MISSPLACEDTILES    = 3
)

type Board struct {
	LastMove int
	Size     int
	Tiles    []int
	BR       int
	BC       int
	HeurFun  func(int, int, int, int, *Board) int
}

func (b *Board) Copy() *Board {
	r := &Board{
		Size:  b.Size,
		BR:    b.BR,
		BC:    b.BC,
		Tiles: make([]int, len(b.Tiles), len(b.Tiles)),
	}
	copy(r.Tiles, b.Tiles)
	r.HeurFun = b.HeurFun
	return r
}

func (b *Board) initiate() {
	b.Tiles = make([]int, b.Size*b.Size, b.Size*b.Size)
}

func New(size int, heur int) *Board {
	b := new(Board)
	b.Size = size
	b.initiate()
	b.SetHeuristic(heur)
	return b
}

func (b *Board) StateString() string {
	str := strconv.Itoa(b.Size) + " "
	for i := 0; i < b.Size*b.Size; i++ {
		str += strconv.Itoa(b.Tiles[i]) + " "
	}
	return str
}

func (b *Board) GoalString() string {
	str := strconv.Itoa(b.Size) + " "
	for i := 0; i < b.Size*b.Size-1; i++ {
		str += strconv.Itoa(i+1) + " "
	}
	str += "0 "
	return str
}

func (b *Board) String() string {
	str := "\n"
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			str += strconv.Itoa(b.Tiles[b.Size*i+j]) + " "
		}
		str = strings.Trim(str, " ")
		str += "\n"
	}
	return str
}

func (b *Board) findBlanc() error {
	for k, v := range b.Tiles {
		if v == 0 {
			b.BR = k / b.Size
			b.BC = k % b.Size
			return nil
		}
	}
	return errors.New("no blanc")
}
func (b *Board) Input(Values []int) error {
	err := b.checkInputLen(Values)
	if err != nil {
		return err
	}
	err = b.checkNumbers(Values)
	if err != nil {
		return err
	}
	b.Tiles = make([]int, len(Values), len(Values))
	copy(b.Tiles, Values)
	err = b.findBlanc()
	if err != nil {
		return err
	}
	if b.CheckBoard() == false {
		return errors.New("Board unsolvable")
	}
	return nil
}

func (b *Board) checkInputLen(Values []int) error {
	if len(Values) != b.Size*b.Size {
		err := errors.New("length of board Values does not match size of board")
		plog.Error.Println(err)
		return err
	}
	plog.Info.Println("length checked")
	return nil
}

func (b *Board) checkNumbers(Values []int) error {
	for i := 0; i < len(Values); i++ {
		found := false
		j := 0
		for j < len(Values) {
			if i == Values[j] {
				found = true
				break
			}
			j++
		}
		if found == false {
			err := errors.New("not all numbers are present in input")
			plog.Error.Println(err)
			return err
		}
	}
	plog.Info.Println("numbers checked")
	return nil
}

//moveMap is initialised in init()
var moveMap map[int]func(*Board) error

func (b *Board) Move(dir int) error {
	if dir < 4 && dir >= 0 {
		err := moveMap[dir](b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Board) moveUp() error {
	if b.BR == b.Size-1 {
		err := errors.New("cannot slide up with the blanc Tile on bottom Row")
		//plog.Warning.Println(err)
		return err
	}
	b.Tiles[b.BR*b.Size+b.BC], b.Tiles[(b.BR+1)*b.Size+b.BC] = b.Tiles[(b.BR+1)*b.Size+b.BC], b.Tiles[b.BR*b.Size+b.BC]
	b.BR++
	return nil
}

func (b *Board) moveDown() error {
	if b.BR == 0 {
		err := errors.New("cannot slide down with the blanc Tile on top Row")
		//plog.Warning.Println(err)
		return err
	}
	b.Tiles[b.BR*b.Size+b.BC], b.Tiles[(b.BR-1)*b.Size+b.BC] = b.Tiles[(b.BR-1)*b.Size+b.BC], b.Tiles[b.BR*b.Size+b.BC]
	b.BR--
	return nil
}

func (b *Board) moveLeft() error {
	if b.BC == b.Size-1 {
		err := errors.New("cannot slide left  with the blanc Tile on right collumn")
		//plog.Warning.Println(err)
		return err
	}
	b.Tiles[b.BR*b.Size+b.BC], b.Tiles[b.BR*b.Size+b.BC+1] = b.Tiles[b.BR*b.Size+b.BC+1], b.Tiles[b.BR*b.Size+b.BC]
	b.BC++
	return nil
}

func (b *Board) moveRight() error {
	if b.BC == 0 {
		err := errors.New("cannot slide right  with the blanc Tile on left collumn")
		return err
	}
	b.Tiles[b.BR*b.Size+b.BC], b.Tiles[b.BR*b.Size+b.BC-1] = b.Tiles[b.BR*b.Size+b.BC-1], b.Tiles[b.BR*b.Size+b.BC]
	b.BC--
	return nil
}

func (b *Board) SetHeuristic(heuristic int) {
	if heuristic == MANHATTAN {
		b.HeurFun = CalcMD
	} else if heuristic == MANHATTAN_CONFLICT {
		b.HeurFun = CalcMd_linearConflict
	} else if heuristic == MISSPLACEDTILES {
		b.HeurFun = OutOfPlace
	} else {
		b.HeurFun = CalcMD
	}
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}
	return num
}

func CalcMD(x, y, fx, fy int, b *Board) int {
	return abs(x-fx) + abs(y-fy)
}

func get_conflict(tab []int, idxValue int, nbrRow int) int {
	sum := 0
	maxConflict := ((idxValue % nbrRow) * nbrRow) + nbrRow
	value := tab[idxValue]
	if value >= maxConflict-nbrRow && value < maxConflict {
		for i := idxValue; i < maxConflict; i++ {
			if tab[i] >= maxConflict-nbrRow && value > tab[i] {
				sum += 2
			}
		}
	}
	return sum
}

func (b *Board) LinearConflict(tab []int, value int, nbrRow int) int {
	sum := 0
	indexGood := value - 1
	indexValue := 0
	for tab[indexValue] != value {
		indexValue++
	}
	if indexGood > indexValue {
		sum = ((indexGood / nbrRow) - (indexValue / nbrRow))
		if (indexGood % nbrRow) > (indexValue % nbrRow) {
			sum += ((indexGood % nbrRow) - (indexValue % nbrRow))
		} else {
			sum += ((indexValue % nbrRow) - (indexGood % nbrRow))
		}
	} else {
		sum = ((indexValue / nbrRow) - (indexGood / nbrRow))
		if (indexGood % nbrRow) > (indexValue % nbrRow) {
			sum += ((indexGood % nbrRow) - (indexValue % nbrRow))
		} else {
			sum += ((indexValue % nbrRow) - (indexGood % nbrRow))
		}
	}
	sum += get_conflict(tab, indexValue, nbrRow)
	return sum
}

func CalcMd_linearConflict(x, y, fx, fy int, b *Board) int {
	sum := 0
	value := 1
	max := b.Size * b.Size
	for value < max {
		sum += b.LinearConflict(b.Tiles, value, b.Size)
		value++
	}
	return sum
}

func OutOfPlace(x, y, fx, fy int, b *Board) int {
	if x == fx && y == fy {
		return 0
	}
	return 1
}

func (b *Board) GetH() int {
	var (
		h, fx, fy int
	)
	for y := 0; y < b.Size; y++ {
		for x := 0; x < b.Size; x++ {
			if b.Tiles[y*b.Size+x] == 0 {
				fx = b.Size - 1
				fy = b.Size - 1
			} else {
				fx = (b.Tiles[y*b.Size+x] - 1) % b.Size
				fy = (b.Tiles[y*b.Size+x] - 1) / b.Size
			}
			h += b.HeurFun(x, y, fx, fy, b)
		}
	}
	return h
}

func (b *Board) GetPosMoves() (moves []int) {
	moves = make([]int, 0, 4)
	if b.BR < b.Size-1 {
		moves = append(moves, 0)
	}
	if b.BR > 0 {
		moves = append(moves, 1)
	}
	if b.BC > 0 {
		moves = append(moves, 2)
	}
	if b.BC < b.Size-1 {
		moves = append(moves, 3)
	}
	return
}

func (b *Board) GetMoves() []*Board {
	ret := make([]*Board, 0, 4)
	moves := b.GetPosMoves()
	for _, v := range moves {
		move := b.Copy()
		err := move.Move(v)
		if err == nil {
			ret = append(ret, move)
		} else {
			//plog.Warning.Println("move", i, "not possible")
		}
	}
	return ret
}

func (b *Board) GetLastMove() int {
	return b.LastMove
}

func (b *Board) get_inversion(y int, x int) int {
	sum := 0
	flag := 1
	j := x

	if b.Tiles[y*b.Size+x] == 0 {
		return sum
	}
	for i := y; i < b.Size; i++ {
		if flag == 0 {
			j = 0
		}
		for j < b.Size {
			if b.Tiles[y*b.Size+x] > b.Tiles[i*b.Size+j] && b.Tiles[i*b.Size+j] > 0 {
				sum++
			}
			j++
		}
		flag = 0
	}
	return sum
}

func (b *Board) get_position_zero() int {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Tiles[i*b.Size+j] == 0 {
				return b.Size - i
			}
		}
	}
	return 1
}

func (b *Board) CheckBoard() bool {
	sum := 0
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			sum += b.get_inversion(i, j)
		}
	}
	if b.Size%2 == 0 {
		row := b.get_position_zero()
		if row%2 == 0 && sum%2 != 0 {
			return true
		} else if row%2 != 0 && sum%2 == 0 {
			return true
		}
	} else {
		if sum%2 == 0 {
			return true
		}
	}
	return false
}

func init() {
	moveMap = make(map[int]func(*Board) error)
	moveMap[Up] = (*Board).moveUp
	moveMap[Down] = (*Board).moveDown
	moveMap[Left] = (*Board).moveLeft
	moveMap[Right] = (*Board).moveRight
}
