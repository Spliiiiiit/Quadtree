package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (f Floor) Draw(screen *ebiten.Image, shift float64, xInc, yInc int) {
	p := configuration.Global.Padding
	for y := range f.Content {
		for x := range f.Content[y] {
			if f.Content[y][x] != -1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64((x-p)*configuration.Global.TileSize)+shift*float64(-xInc), float64((y-p)*configuration.Global.TileSize)+shift*float64(-yInc))

				block := f.Content[y][x]
				shiftX := block * configuration.Global.TileSize
				portal := false
				// Le bloc de portail est représenté par le chiffre 7 et le bloc d'eau par le chiffre 4
				if block == 4 && configuration.Global.AnimatedWater.Enabled {
					shiftX = (f.waterAnim.animationStep + 5) * configuration.Global.TileSize
				} else if block == 7 && configuration.Global.Teleporter.Enabled {
					shiftX = f.portalAnim.animationStep * configuration.Global.TileSize
					portal = true
				} else if block > 4 {
					continue
				}
				if portal {
					screen.DrawImage(assets.TeleporterImage.SubImage(
						image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
					).(*ebiten.Image), op)
				} else {
					screen.DrawImage(assets.FloorImage.SubImage(
						image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
					).(*ebiten.Image), op)
				}
			}
		}
	}
}
