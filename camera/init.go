package camera

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init met en place une caméra.
func (c *Camera) Init() {
	c.step = float64(configuration.Global.TileSize) / float64(configuration.Global.NumCharacterAnimImages*configuration.Global.NumFramePerCharacterAnimImage)
	if configuration.Global.CameraMode == Static {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	} else if configuration.Global.CameraMode == Smooth && configuration.Global.Padding < 1 {
		panic("erreur la propriété de configuration 'Padding' doit être supérieure ou égale à 1")
	}
}
