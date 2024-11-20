package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	x      int32
	y      int32
	radius float32
	color  rl.Color
	speedX int32
	speedY int32
}

func NewBall(x int32, y int32, radius float32, color rl.Color) *Ball {
	return &Ball{
		x:      x,
		y:      y,
		radius: radius,
		color:  color,
		speedX: 7,
		speedY: 7,
	}
}

func (b *Ball) Update() bool {
	b.x += b.speedX
	b.y += b.speedY

	if b.y+int32(b.radius) >= int32(rl.GetScreenHeight()) || b.y-int32(b.radius) <= 0 {
		b.speedY *= -1
		return true
	}
	if b.x+int32(b.radius) >= int32(rl.GetScreenWidth()) || b.x-int32(b.radius) <= 0 {
		b.speedX *= -1
		return true
	}
	return false
}

func (b *Ball) Reset() {
	b.x = int32(rl.GetScreenWidth() / 2)
	b.y = int32(rl.GetScreenHeight() / 2)

	speedChoices := [2]int32{-1, 1}
	b.speedX *= speedChoices[rl.GetRandomValue(0, 1)]
	b.speedY *= speedChoices[rl.GetRandomValue(0, 1)]
}

func (b *Ball) CheckIfScore() string {
	if b.x+int32(b.radius) >= int32(rl.GetScreenWidth()) {
		return "cpu"
	}
	if b.x-int32(b.radius) <= 0 {
		return "player"
	}
	return ""
}

func (b *Ball) Collision() {
	b.speedX *= -1
}
func (b *Ball) Draw() {
	rl.DrawCircle(b.x, b.y, b.radius, b.color)
}
