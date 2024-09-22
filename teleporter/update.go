package teleporter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour les différents champs du teleporter ainsi que
// les coordonnées du personnage et de la caméra
func (t *Teleporter) Update(charX, charY, camX, camY *int, content [][]int) {
	if !configuration.Global.Teleporter.Enabled {
		return
	}
	if *charX != t.lastX || *charY != t.lastY {
		t.canTeleport = true
	}
	// Placement d'un portail quand la touche T est utilisée
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		t.place(*charX, *charY)
	}
	// Détecter si le personnage marche sur un portail et s'il peut être téléporté
	if *charX == t.x1 && *charY == t.y1 && t.p1 && t.p2 && t.canTeleport {
		t.teleport(charX, charY, camX, camY, t.x2, t.y2)
	} else if *charX == t.x2 && *charY == t.y2 && t.p2 && t.p1 && t.canTeleport {
		t.teleport(charX, charY, camX, camY, t.x1, t.y1)
	}
	// Placer dans content les blocs de portail
	for yR := 0; yR < len(content); yR++ {
		for xR := 0; xR < len(content[yR]); xR++ {
			yA := *camY + yR - configuration.Global.ScreenCenterTileY - configuration.Global.Padding
			xA := *camX + xR - configuration.Global.ScreenCenterTileX - configuration.Global.Padding
			if (yA == t.y1 && xA == t.x1 && t.p1) || (yA == t.y2 && xA == t.x2 && t.p2) {
				content[yR][xR] = 7
			}
		}
	}
}

func (t *Teleporter) teleport(charX, charY, camX, camY *int, x, y int) {
	*charX, *charY = x, y
	*camX, *camY = x, y
	t.lastX, t.lastY = x, y
	t.canTeleport = false
}
