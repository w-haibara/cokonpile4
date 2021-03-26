package main

import (
	"fmt"
	"os"
	"sample/sanmoku"
)

func main() {
	sm := sanmoku.Init()

	for {
		p1(sm)
		fmt.Println()

		p2(sm)
		fmt.Println()
	}
}

func p1(sm sanmoku.Game) {
	var (
		x int
		y int
	)

	for {
		fmt.Print("Player1 (x y) : ")
		if _, err := fmt.Scanln(&x, &y); err != nil {
			fmt.Println("\n[error:", err, "]")
			continue
		}

		if err := sm.Set(&sm.P1, x, y); err != nil {
			fmt.Println("[error:", err, "]")
			continue
		}

		break
	}

	printBoad(sm.Boad)

	if sm.P1.IsWinner() {
		fmt.Println("[ YOU WIN ]")
		os.Exit(0)
	}
}

func p2(sm sanmoku.Game) {
	x, y, err := sm.AutoSet(&sm.P2)
	if err != nil {
		fmt.Println("[error:", err, "]")
		os.Exit(1)
	}

	fmt.Println("Player2 (x y) :", x, y)
	printBoad(sm.Boad)

	if sm.P2.IsWinner() {
		fmt.Println("[ YOU LOSE ]")
		os.Exit(0)
	}
}

func printBoad(boad [][]string) {
	fmt.Println("*---*---*---*")
	for i := 0; i < len(boad); i++ {
		fmt.Print("|")
		for j := 0; j < len(boad[i]); j++ {
			fmt.Printf(" %v |", boad[i][j])
		}
		fmt.Println("\n*---*---*---*")
	}
}
