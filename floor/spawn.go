package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/camera"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Spawn permet de changer les coordonnées d'apparition du character et de la camera
func (f *Floor) Spawn(x, y, camX, camY *int) {
	if !configuration.Global.WaterWalk {
		switch configuration.Global.FloorKind {
		case fromFileFloor:
			if *y < 0 || *y >= len(f.fullContent) || *x < 0 || *x >= len(f.fullContent[0]) {
				panic("erreur point de spawn invalide")
			}
			content := f.fullContent[*y][*x]
			if configuration.Global.CameraMode == camera.Static && content == 4 {
				*x = 0
				*y = 0
			}
			for content == 4 {
				if *x+1 >= len(f.fullContent[0]) {
					if *y+1 >= len(f.fullContent) {
						panic("erreur aucun point de spawn valide trouvé")
					}
					*x = 0
					*y++
				} else {
					*x++
				}
				content = f.fullContent[*y][*x]
			}
		case quadTreeFloor:
			content := f.quadtreeContent.Root.FindContent(*x, *y, randomFloor)
			if configuration.Global.CameraMode == camera.Static && content == 4 {
				*x = 0
				*y = 0
			}
			for content == 4 {
				if *x+1 > f.quadtreeContent.Width-1 {
					if *y+1 > f.quadtreeContent.Height-1 {
						panic("erreur aucun point de spawn valide trouvé")
					}
					*x = 0
					*y++
				} else {
					*x++
				}
				content = f.quadtreeContent.Root.FindContent(*x, *y, randomFloor)
			}
		}
		*camX = *x
		*camY = *y
	}
}
