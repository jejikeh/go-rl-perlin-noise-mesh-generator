package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := 800
	screenHeight := 450

	offsetX := 0

	rl.InitWindow(int32(screenWidth), int32(screenHeight), "raylib [textures] example - procedural images generation")

	perlinNoise := rl.GenImagePerlinNoise(64, 64, offsetX, 0, 0.5)
	perlinNoiseTexture := rl.LoadTextureFromImage(perlinNoise)

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(18.0, 16.0, 18.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0

	mesh := rl.GenMeshHeightmap(*perlinNoise, rl.NewVector3(16, 8, 16)) // Generate heightmap mesh (RAM and VRAM)
	model := rl.LoadModelFromMesh(mesh)                                 // Load model from generated mesh

	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, perlinNoiseTexture) // Set map diffuse texture

	mapPosition := rl.NewVector3(-8.0, 0.0, -8.0)

	// Unload image data (CPU RAM)
	rl.UnloadImage(perlinNoise)

	rl.SetTargetFPS(60)

	i := 0
	for !rl.WindowShouldClose() {
		i += 1

		rl.UpdateCamera(&camera, rl.CameraOrbital) // Update camera with orbital camera mode

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawModel(model, mapPosition, 1.0, rl.Red)

		rl.DrawGrid(20, 1.0)

		rl.EndMode3D()

		rl.DrawRectangle(20, 20, 74, 74, rl.Fade(rl.Red, 0.5))
		rl.DrawRectangleLines(20, 20, 74, 74, rl.Fade(rl.White, 0.5))
		rl.DrawTexture(perlinNoiseTexture, 25, 25, rl.White)

		rl.EndDrawing()

		if i%2 == 0 {
			offsetX += 1
		}

		rl.UnloadTexture(perlinNoiseTexture)
		rl.UnloadModel(model)

		perlinNoise = rl.GenImagePerlinNoise(64, 64, offsetX, 0, 0.5)
		perlinNoiseTexture = rl.LoadTextureFromImage(perlinNoise)

		mesh = rl.GenMeshHeightmap(*perlinNoise, rl.NewVector3(16, 8, 16))        // Generate heightmap mesh (RAM and VRAM)
		model = rl.LoadModelFromMesh(mesh)                                        // Load model from generated mesh
		rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, perlinNoiseTexture) // Set map diffuse texture
		rl.UnloadImage(perlinNoise)
	}

	rl.UnloadTexture(perlinNoiseTexture)
	rl.UnloadModel(model)

	rl.CloseWindow()
}
