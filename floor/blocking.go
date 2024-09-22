package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Blocking retourne, étant donné la position du personnage,
// un tableau de booléen indiquant si les cases au-dessus (0),
// à droite (1), au-dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {
	p := configuration.Global.Padding

	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX + p
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY + p

	blocking[0] = relativeYPos <= 0+p || f.Content[relativeYPos-1][relativeXPos] == -1
	blocking[1] = relativeXPos >= configuration.Global.NumTileX-1+p || f.Content[relativeYPos][relativeXPos+1] == -1
	blocking[2] = relativeYPos >= configuration.Global.NumTileY-1+p || f.Content[relativeYPos+1][relativeXPos] == -1
	blocking[3] = relativeXPos <= 0+p || f.Content[relativeYPos][relativeXPos-1] == -1
	// Extension pour empêcher de marcher sur l'eau
	if !configuration.Global.WaterWalk {
		blocking[0] = blocking[0] || f.Content[relativeYPos-1][relativeXPos] == 4
		blocking[1] = blocking[1] || f.Content[relativeYPos][relativeXPos+1] == 4
		blocking[2] = blocking[2] || f.Content[relativeYPos+1][relativeXPos] == 4
		blocking[3] = blocking[3] || f.Content[relativeYPos][relativeXPos-1] == 4
	}
	return blocking
}
