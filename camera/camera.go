package camera

// Camera définit les caractéristiques de la caméra.
//
// Les champs sont :
//   - X, Y : la position de la caméra
//   - XInc, YInc : les incréments en X et Y qui sont en cours d'animation.
//   - Moving : l'information de si l'animation est en cours ou pas.
//   - Shift : le nombre de pixels de décalage dans l'animation de la caméra.
//   - step : l'incrément en pixel à réaliser à chaque étape de l'animation.
//   - frameCount : compte le nombre d'images durant l'animation.
type Camera struct {
	X, Y       int
	XInc, YInc int
	Moving     bool
	Shift      float64
	step       float64
	frameCount int
}

// types de caméra disponibles
const (
	Static int = iota
	FollowCharacter
	Smooth
)
