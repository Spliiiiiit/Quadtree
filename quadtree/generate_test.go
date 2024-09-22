package quadtree

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"math/rand"
	"testing"
)

// Tests de la fonction expand
func TestExpand1(t *testing.T) {
	// Paramètres pour la fonction
	var rInit node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           3,
		height:          3,
		content:         1,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var qInit Quadtree
	qInit.Width, qInit.Height = 3, 3
	qInit.Root = &rInit
	// Résultat attendu
	var trRes node = node{
		topLeftX:        3,
		topLeftY:        0,
		width:           3,
		height:          3,
		content:         -2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var blRes node = node{
		topLeftX:        0,
		topLeftY:        3,
		width:           3,
		height:          3,
		content:         -2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var brRes node = node{
		topLeftX:        3,
		topLeftY:        3,
		width:           3,
		height:          3,
		content:         -2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var rRes node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           6,
		height:          6,
		content:         0,
		topLeftNode:     &rInit,
		topRightNode:    &trRes,
		bottomLeftNode:  &blRes,
		bottomRightNode: &brRes,
	}
	var qRes Quadtree
	qRes.Width, qRes.Height = 6, 6
	qRes.Root = &rRes

	qInit.expand(4, 4)
	if !equalQuadtree(qInit, qRes) {
		logQuadtree(qInit)
		t.Error("Erreur dans l'agrandissement du quadtree (top left)")
	}
}

func TestExpand2(t *testing.T) {
	// Paramètres pour la fonction
	var blInit node = node{
		topLeftX:        0,
		topLeftY:        1,
		width:           1,
		height:          2,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var rInit node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           3,
		height:          3,
		content:         0,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  &blInit,
		bottomRightNode: nil,
	}
	var qInit Quadtree
	qInit.Width, qInit.Height = 3, 3
	qInit.Root = &rInit
	// Résultat attendu
	var tlRes node = node{
		topLeftX:        -3,
		topLeftY:        -3,
		width:           3,
		height:          3,
		content:         -2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var trRes node = node{
		topLeftX:        0,
		topLeftY:        -3,
		width:           3,
		height:          3,
		content:         -2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var blRes node = node{
		topLeftX:        -3,
		topLeftY:        0,
		width:           3,
		height:          3,
		content:         -2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var brblRes node = node{
		topLeftX:        0,
		topLeftY:        1,
		width:           1,
		height:          2,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var brRes node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           3,
		height:          3,
		content:         0,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  &brblRes,
		bottomRightNode: nil,
	}
	var rRes node = node{
		topLeftX:        -3,
		topLeftY:        -3,
		width:           6,
		height:          6,
		content:         0,
		topLeftNode:     &tlRes,
		topRightNode:    &trRes,
		bottomLeftNode:  &blRes,
		bottomRightNode: &brRes,
	}
	var qRes Quadtree
	qRes.Width, qRes.Height = 6, 6
	qRes.Root = &rRes

	qInit.expand(-1, -1)
	if !equalQuadtree(qInit, qRes) {
		logQuadtree(qInit)
		t.Error("Erreur dans l'agrandissement du quadtree (bottom right)")
	}
}

// Tests de la fonction fill
func TestFill1(t *testing.T) {
	configuration.Global.RandomMap.Width = 2
	configuration.Global.RandomMap.Height = 2
	var brInit node = node{
		topLeftX:        1,
		topLeftY:        1,
		width:           2,
		height:          2,
		content:         -2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}

	err := brInit.fill(2, 2, generateFloor)
	if !err {
		t.Error("erreur la node n'a pas pu être remplacée")
	}
	if brInit.content != 0 {
		logQuadtreeNode(&brInit, "R.br")
		t.Error("erreur le contenu de la node n'a pas été mis à jour")
	}
	if brInit.topLeftNode == nil || brInit.topRightNode == nil || brInit.bottomLeftNode == nil || brInit.bottomRightNode == nil {
		logQuadtreeNode(&brInit, "R.br")
		t.Error("erreur la node à générer n'a pas d'enfant")
	}
}

func TestFill2(t *testing.T) {
	configuration.Global.RandomMap.Width = 2
	configuration.Global.RandomMap.Height = 2
	var brInit node = node{
		topLeftX:        1,
		topLeftY:        1,
		width:           4,
		height:          4,
		content:         -2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}

	err := brInit.fill(2, 2, generateFloor)
	if !err {
		t.Error("erreur la node n'a pas pu être remplacée")
	}
	if brInit.content != 0 {
		logQuadtreeNode(&brInit, "R.br")
		t.Error("erreur le contenu de la node n'a pas été mis à jour")
	}
	children := []*node{
		brInit.topLeftNode,
		brInit.topRightNode,
		brInit.bottomLeftNode,
		brInit.bottomRightNode,
	}
	coordinates := [][]int{
		{1, 1},
		{3, 1},
		{1, 3},
		{3, 3},
	}
	for i, child := range children {
		if child == nil {
			logQuadtreeNode(&brInit, "R.br")
			t.Error("erreur la node à générer n'a pas d'enfant")
		}
		if child.content != -2 {
			logQuadtreeNode(&brInit, "R.br")
			t.Error("erreur les enfants de la node n'ont pas le bon contenu")
		}
		if child.width != 2 || child.height != 2 {
			logQuadtreeNode(&brInit, "R.br")
			t.Error("erreur les enfants de la node n'ont pas la bonne taille")
		}
		if child.topLeftX != coordinates[i][0] || child.topLeftY != coordinates[i][1] {
			logQuadtreeNode(&brInit, "R.br")
			t.Error("erreur le noeud ", i, " n'a pas les bonnes coordonnées")
		}
	}
}

// Fonction utilitaires pour les tests

// generateFloor crée un tableau représentant un terrain aléatoire d'une taille donnée
func generateFloor(w, h int) (floorContent [][]int) {
	floorContent = make([][]int, h)
	for i := 0; i < len(floorContent); i++ {
		floorContent[i] = make([]int, w)
		for j := 0; j < len(floorContent[i]); j++ {
			floorContent[i][j] = rand.Intn(5)
		}
	}
	return floorContent
}
