package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"

// animation représente les données nécessaires à l'animation du terrain
//
//   - animationStep : l'étape de l'animation en cours d'affichage.
//   - animationFrameCount : le nombre d'appels à update (ou de 1/60 de seconde)
//     qui ont eu lieu depuis la dernière étape d'animation.
type animation struct {
	animationStep       int
	animationFrameCount int
}

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - Content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
//   - waterAnim : structure représentant les données nécessaires à l'animation
//     de l'eau
//   - portalAnim : structure représentant les données nécessaires à l'animation
//     des téléporteurs
type Floor struct {
	Content         [][]int
	fullContent     [][]int
	quadtreeContent quadtree.Quadtree
	waterAnim       animation
	portalAnim      animation
}

// Types d'affichage du terrain disponibles
const (
	gridFloor int = iota
	fromFileFloor
	quadTreeFloor
)
