package main

import rl "github.com/gen2brain/raylib-go/raylib"

type PlayerPaddle struct {
	x, y, width, height, speed int32
	color                      rl.Color
}

type CPUPaddle struct {
	x, y, width, height, speed int32
	color                      rl.Color
}

func NewPlayerPaddle(x int32, y int32, width, height int32, color rl.Color) *PlayerPaddle {
	return &PlayerPaddle{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		color:  color,
		speed:  6,
	}
}

func (p *PlayerPaddle) Update() {
	if rl.IsKeyDown(rl.KeyUp) {
		p.y -= p.speed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		p.y += p.speed
	}
	if p.y <= 0 {
		p.y = 0
	}
	if p.y+p.height >= int32(rl.GetScreenHeight()) {
		p.y = int32(rl.GetScreenHeight()) - p.height
	}
}

func (p *PlayerPaddle) Reset() {
	p.y = int32(rl.GetScreenHeight())/2 - 60
}
func (p *PlayerPaddle) Draw() {
	rl.DrawRectangleRounded(rl.Rectangle{X: float32(p.x), Y: float32(p.y), Width: float32(p.width), Height: float32(p.height)}, 0.8, 0, rl.RayWhite)
}

func NewAIPaddle(x int32, y int32, width, height int32, color rl.Color) *CPUPaddle {
	return &CPUPaddle{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		color:  color,
		speed:  6,
	}
}

func (p *CPUPaddle) Update(ballY int32) {
	if p.y+p.height/2 > ballY {
		p.y -= p.speed
	}
	if p.y+p.height/2 <= ballY {
		p.y += p.speed
	}

	if p.y <= 0 {
		p.y = 0
	}
	if p.y+p.height >= int32(rl.GetScreenHeight()) {
		p.y = int32(rl.GetScreenHeight()) - p.height
	}
}

func (p *CPUPaddle) Draw() {
	rl.DrawRectangleRounded(rl.Rectangle{X: float32(p.x), Y: float32(p.y), Width: float32(p.width), Height: float32(p.height)}, 0.8, 0, rl.RayWhite)
}
func (p *CPUPaddle) Reset() {
	p.y = int32(rl.GetScreenHeight())/2 - 60
}
