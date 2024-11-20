package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
	"time"
)

const (
	playState = iota
	pointState
)

func main() {

	/*go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			fmt.Printf("TotalAlloc: %d\nHeapAlloc: %d\nHeapInuse: %d\n\n\n", m.TotalAlloc, m.HeapAlloc, m.HeapInuse)
			time.Sleep(1 * time.Second)
		}
	}()*/

	var screenWidth int32 = 1280
	var screenHeight int32 = 800
	Green := rl.Color{R: 38, G: 185, B: 154, A: 255}
	DarkGreen := rl.Color{R: 20, G: 160, B: 133, A: 255}
	LightGreen := rl.Color{R: 129, G: 204, B: 184, A: 255}
	Yellow := rl.Color{R: 243, G: 213, B: 91, A: 255}

	rl.InitWindow(screenWidth, screenHeight, "Pong - Raylib")
	rl.SetTargetFPS(60)

	rl.InitAudioDevice()
	paddleSound := rl.LoadSound("resources/paddle.wav")
	wallSound := rl.LoadSound("resources/wall.wav")
	scoreSound := rl.LoadSound("resources/score.wav")

	ball := NewBall(screenWidth/2, screenHeight/2, 20, Yellow)
	ball.Reset()
	cpuPaddle := NewAIPaddle(10, screenHeight/2-60, 25, 120, rl.White)
	playerPaddle := NewPlayerPaddle(screenWidth-45, screenHeight/2-60, 35, 120, rl.White)

	var cpuScore int32 = 0
	var playerScore int32 = 0
	tickCounter := 3
	paddleCollision := false
	wallCollision := false
	gameState := playState
	pointScored := false
	firstRun := true

	for !rl.WindowShouldClose() {

		if gameState == playState {
			if firstRun {
				// Load initial countdown
				pointScored = true
				gameState = pointState
				firstRun = false
			}
			rl.BeginDrawing()
			rl.ClearBackground(DarkGreen)
			rl.DrawRectangle(screenWidth/2, 0, screenWidth/2, screenHeight, Green)
			rl.DrawCircle(screenWidth/2, screenHeight/2, 150, LightGreen)

			// Check collision with wall
			wallCollision = ball.Update()
			if wallCollision {
				rl.PlaySound(wallSound)
			}
			playerPaddle.Update()
			cpuPaddle.Update(ball.y)

			// Check collision with paddles
			paddleCollision = checkPaddleCollision(ball, playerPaddle, cpuPaddle)
			if paddleCollision {
				rl.PlaySound(paddleSound)
				ball.Collision()
			}

			// Check if Score
			checkIfScore := ball.CheckIfScore()
			if checkIfScore != "" {
				setScore(checkIfScore, &cpuScore, &playerScore)
				rl.PlaySound(scoreSound)
				playerPaddle.Reset()
				cpuPaddle.Reset()
				ball.Reset()
				pointScored = true
			}

			// Draw
			rl.DrawLine(screenWidth/2, 0, screenWidth/2, screenHeight, rl.White)
			ball.Draw()
			cpuPaddle.Draw()
			playerPaddle.Draw()
			drawScoreBoard(cpuScore, playerScore, screenWidth)

			if pointScored {
				gameState = pointState
			}
			rl.EndDrawing()
		}

		if gameState == pointState {
			//Prepare window
			rl.BeginDrawing()
			rl.ClearBackground(DarkGreen)
			rl.DrawRectangle(screenWidth/2, 0, screenWidth/2, screenHeight, Green)
			rl.DrawCircle(screenWidth/2, screenHeight/2, 150, LightGreen)

			// Countdown ticker
			timeTicker := time.NewTicker(1 * time.Second)
			select {
			case <-timeTicker.C:
				{
					rl.DrawText(strconv.Itoa(tickCounter), screenWidth/2-rl.MeasureText(strconv.Itoa(tickCounter), 200)/2, 100, 200, rl.Black)

					if tickCounter == 0 {
						timeTicker.Stop()
						gameState = playState
						pointScored = false
						tickCounter = 3
					}
					if pointScored {
						tickCounter--
					}
				}
			}
			//DRAW
			ball.Draw()
			cpuPaddle.Draw()
			playerPaddle.Draw()
			rl.DrawLine(screenWidth/2, 0, screenWidth/2, screenHeight, rl.White)
			drawScoreBoard(cpuScore, playerScore, screenWidth)
			rl.EndDrawing()
		}
	}
	defer rl.UnloadSound(paddleSound)
	defer rl.UnloadSound(wallSound)
	defer rl.UnloadSound(scoreSound)
	defer rl.CloseAudioDevice()
	defer rl.CloseWindow()
}

func drawScoreBoard(cpuScore int32, playerScore int32, screenWidth int32) {
	rl.DrawText(strconv.Itoa(int(cpuScore)), screenWidth/4-20, 20, 80, rl.White)
	rl.DrawText(strconv.Itoa(int(playerScore)), 3*screenWidth/4-20, 20, 80, rl.White)
}

func setScore(whoScored string, cpuScore *int32, playerScore *int32) {
	if whoScored == "cpu" {
		*cpuScore++
	} else if whoScored == "player" {
		*playerScore++
	}
}

func checkPaddleCollision(ball *Ball, playerPaddle *PlayerPaddle, cpuPaddle *CPUPaddle) bool {
	if rl.CheckCollisionCircleRec(rl.Vector2{float32(ball.x), float32(ball.y)}, ball.radius, rl.Rectangle{float32(playerPaddle.x), float32(playerPaddle.y), float32(playerPaddle.width), float32(playerPaddle.height)}) {
		return true
	}

	if rl.CheckCollisionCircleRec(rl.Vector2{float32(ball.x), float32(ball.y)}, ball.radius, rl.Rectangle{float32(cpuPaddle.x), float32(cpuPaddle.y), float32(cpuPaddle.width), float32(cpuPaddle.height)}) {
		return true
	}
	return false
}
