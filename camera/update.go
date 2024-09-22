package camera

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour la position de la caméra à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(charMoving bool, charPosX, charPosY, charXInc, charYInc int) {
	switch configuration.Global.CameraMode {
	case Static:
		c.updateStatic()
	case FollowCharacter:
		c.updateFollowCharacter(charPosX, charPosY)
	case Smooth:
		c.updateSmooth(charMoving, charXInc, charYInc)
	}
}

// updateStatic est la mise-à-jour d'une caméra qui reste
// toujours à la position (0,0). Cette fonction ne fait donc
// rien.
func (c *Camera) updateStatic() {}

// updateFollowCharacter est la mise-à-jour d'une caméra qui
// suit toujours le personnage. Elle prend en paramètres deux
// entiers qui indiquent les coordonnées du personnage et place
// la caméra au même endroit.
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY int) {
	c.X = characterPosX
	c.Y = characterPosY
}

// updateSmooth est la mise-à-jour d'une caméra qui
// suit le personnage pixel par pixel.
func (c *Camera) updateSmooth(charMoving bool, charXInc, charYInc int) {
	if charMoving && !c.Moving {
		c.Moving = true
		c.XInc = charXInc
		c.YInc = charYInc
	} else if c.Moving {
		c.frameCount++
		c.Shift += c.step
		if c.frameCount >= configuration.Global.NumCharacterAnimImages*configuration.Global.NumFramePerCharacterAnimImage {
			c.Moving = false
			c.Shift = 0
			c.frameCount = 0
			c.X += c.XInc
			c.Y += c.YInc
			c.XInc, c.YInc = 0, 0
		}
	}
}
