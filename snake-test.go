package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 20
	height = 10
)

type snake struct {
	body      []point
	direction point
}

type point struct {
	x, y int
}

func newSnake() *snake {
	return &snake{
		body: []point{{width / 2, height / 2}},
		direction: point{
			x: 1,
			y: 0,
		},
	}
}

func (s *snake) move() {
	head := s.body[len(s.body)-1]
	newHead := point{
		x: head.x + s.direction.x,
		y: head.y + s.direction.y,
	}
	s.body = append(s.body, newHead)
	s.body = s.body[1:]
}

func (s *snake) grow() {
	tail := s.body[0]
	newTail := point{
		x: tail.x - s.direction.x,
		y: tail.y - s.direction.y,
	}
	s.body = append([]point{newTail}, s.body...)
}

func (s *snake) checkCollision() bool {
	head := s.body[len(s.body)-1]
	if head.x < 0 || head.x >= width || head.y < 0 || head.y >= height {
		return true
	}
	for i := 0; i < len(s.body)-1; i++ {
		if s.body[i] == head {
			return true
		}
	}
	return false
}

func (s *snake) checkFood(food point) bool {
	head := s.body[len(s.body)-1]
	if head == food {
		return true
	}
	return false
}

func generateFood(s *snake) point {
	for {
		food := point{
			x: rand.Intn(width),
			y: rand.Intn(height),
		}
		if !contains(s.body, food) {
			return food
		}
	}
}

func contains(points []point, p point) bool {
	for _, point := range points {
		if point == p {
			return true
		}
	}
	return false
}

func printBoard(s *snake, food point) {
	board := make([][]byte, height)
	for i := range board {
		board[i] = make([]byte, width)
		for j := range board[i] {
			board[i][j] = ' '
		}
	}

	for _, p := range s.body {
		board[p.y][p.x] = 'o'
	}

	board[food.y][food.x] = '*'

	fmt.Println("+---------------------+")
	for _, row := range board {
		fmt.Print("| ")
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println("|")
	}
	fmt.Println("+---------------------+")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	s := newSnake()
	food := generateFood(s)

	for {
		printBoard(s, food)

		time.Sleep(200 * time.Millisecond)

		s.move()

		if s.checkCollision() {
			fmt.Println("Game over!")
			return
		}

		if s.checkFood(food) {
			s.grow()
			food = generateFood(s)
		}
	}
}
