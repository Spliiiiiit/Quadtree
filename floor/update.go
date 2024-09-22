package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/camera"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(cam camera.Camera) {
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(cam.X, cam.Y)
	case fromFileFloor:
		f.updateFromFileFloor(cam.X, cam.Y)
	case quadTreeFloor:
		f.updateQuadtreeFloor(cam)
	}
	if configuration.Global.AnimatedWater.Enabled {
		f.waterAnim.animationFrameCount++
		if f.waterAnim.animationFrameCount >= configuration.Global.AnimatedWater.NumFramePerAnimImage {
			f.waterAnim.animationFrameCount = 0
			if f.waterAnim.animationStep+1 > configuration.Global.AnimatedWater.NumAnimImages-1 {
				f.waterAnim.animationStep = 0
			} else {
				f.waterAnim.animationStep++
			}
		}
	}
	if configuration.Global.Teleporter.Enabled {
		// Animation du portail
		f.portalAnim.animationFrameCount++
		if f.portalAnim.animationFrameCount >= configuration.Global.Teleporter.NumFramePerAnimImage {
			f.portalAnim.animationFrameCount = 0
			if f.portalAnim.animationStep+1 > configuration.Global.Teleporter.NumAnimImages-1 {
				f.portalAnim.animationStep = 0
			} else {
				f.portalAnim.animationStep++
			}
		}
	}
}

// Le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(camXPos, camYPos int) {
	for y := 0; y < len(f.Content); y++ {
		for x := 0; x < len(f.Content[y]); x++ {
			absCamX := camXPos
			if absCamX < 0 {
				absCamX = -absCamX
			}
			absCamY := camYPos
			if absCamY < 0 {
				absCamY = -absCamY
			}
			f.Content[y][x] = ((x + absCamX%2) + (y + absCamY%2)) % 2
		}
	}
}

// Le sol est récupéré depuis un tableau, qui a été lu dans un fichier
func (f *Floor) updateFromFileFloor(camXPos, camYPos int) {
	for yR := 0; yR < len(f.Content); yR++ {
		for xR := 0; xR < len(f.Content[yR]); xR++ {
			// Calcul des coordonnées absolues à partir des coordonnées
			// relatives dépend de CameraMode pour le centrage
			yA := camYPos + yR - configuration.Global.ScreenCenterTileY - configuration.Global.Padding
			xA := camXPos + xR - configuration.Global.ScreenCenterTileX - configuration.Global.Padding
			// Valeur par défaut d'un terrain vide : -1
			if xA < 0 || yA < 0 || yA >= len(f.fullContent) || xA >= len(f.fullContent[yA]) {
				f.Content[yR][xR] = -1
			} else {
				f.Content[yR][xR] = f.fullContent[yA][xA]
			}
		}
	}
}

// Le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(cam camera.Camera) {
	topLeftX := cam.X - configuration.Global.ScreenCenterTileX - configuration.Global.Padding
	topLeftY := cam.Y - configuration.Global.ScreenCenterTileY - configuration.Global.Padding
	if cam.Moving && configuration.Global.CameraMode == camera.Smooth {
		topLeftX += cam.XInc
		topLeftY += cam.YInc
	} else {
		f.quadtreeContent.GetContent(topLeftX, topLeftY, f.Content, randomFloor)
	}
}
