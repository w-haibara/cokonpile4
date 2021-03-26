package sanmoku

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Len  int
	P1   Player
	P2   Player
	Boad [][]string
}

type Player struct {
	ID   string
	Mark string
	Boad [][]bool
}

func initPlayerBoad() [][]bool {
	return [][]bool{
		[]bool{false, false, false},
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
}

func Init() Game {
	return Game{
		Len: 3,
		P1: Player{
			ID:   "P1",
			Mark: "o",
			Boad: initPlayerBoad(),
		},
		P2: Player{
			ID:   "P2",
			Mark: "x",
			Boad: initPlayerBoad(),
		},
		Boad: [][]string{
			[]string{" ", " ", " "},
			[]string{" ", " ", " "},
			[]string{" ", " ", " "},
		},
	}
}

func (p Player) IsWinner() bool {
	//vertical, horizontal, and diagonal
	d1 := true
	d2 := true
	for i := 0; i < len(p.Boad); i++ {
		v := true
		h := true
		for j := 0; j < len(p.Boad[i]); j++ {
			v = v && p.Boad[i][j]
			h = h && p.Boad[j][i]
		}

		d1 = d1 && p.Boad[i][i]
		d2 = d2 && p.Boad[i][len(p.Boad)-1-i]

		if v || h {
			return true
		}
	}

	if d1 || d2 {
		return true
	}

	return false
}

func (g Game) IsOccupied(x, y int) bool {
	return g.Boad[x][y] == g.P1.Mark || g.Boad[x][y] == g.P2.Mark
}

func (g Game) Set(p *Player, x, y int) error {
	if x < 0 || y < 0 || x >= g.Len || y >= g.Len {
		return fmt.Errorf("(%d %d) is invalid value", x, y)
	}
	if g.IsOccupied(x, y) {
		return fmt.Errorf("boad[%d][%d] is not empty", x, y)
	}

	p.Boad[x][y] = true
	g.Boad[x][y] = p.Mark

	return nil
}

func (g Game) Unset(p *Player, x, y int) error {
	if x < 0 || y < 0 || x >= g.Len || y >= g.Len {
		return fmt.Errorf("(%d %d) is invalid value", x, y)
	}
	if !g.IsOccupied(x, y) {
		return fmt.Errorf("boad[%d][%d] is empty", x, y)
	}

	p.Boad[x][y] = false
	g.Boad[x][y] = " "

	return nil
}

func (g Game) AutoSet(p *Player) (int, int, error) {
	var (
		x       int
		y       int
		empties [][]int
	)

	for i := 0; i < g.Len; i++ {
		for j := 0; j < g.Len; j++ {
			if !g.IsOccupied(i, j) {
				empties = append(empties, []int{i, j})
			}
		}
	}

	enemy := g.P1
	if p.ID == g.P1.ID {
		enemy = g.P2
	}
	for _, v := range empties {
		g.Set(&enemy, v[0], v[1])
		win := enemy.IsWinner()
		g.Unset(&enemy, v[0], v[1])
		if win {
			x = v[0]
			y = v[1]
			return x, y, g.Set(p, x, y)
		}
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(g.Len)

	x = empties[r][0]
	y = empties[r][1]

	return x, y, g.Set(p, x, y)
}
