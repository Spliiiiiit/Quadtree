package quadtree

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"testing"
)

// Tests de la fonction GetContent
func TestGetContent1(t *testing.T) {
	// Test avec un terrain de taille 4x4 avec camera mode 1 et taille d'écran 3x3
	configuration.Global.NumTileX = 3
	configuration.Global.NumTileY = 3
	configuration.Global.ScreenCenterTileX = 1
	configuration.Global.ScreenCenterTileY = 1
	configuration.Global.CameraMode = 1
	var q Quadtree
	var nodeTl node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           2,
		height:          2,
		content:         1,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTrTl node = node{
		topLeftX:        2,
		topLeftY:        0,
		width:           1,
		height:          1,
		content:         3,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTrTr node = node{
		topLeftX:        3,
		topLeftY:        0,
		width:           1,
		height:          1,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTrBl node = node{
		topLeftX:        2,
		topLeftY:        1,
		width:           1,
		height:          1,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTrBr node = node{
		topLeftX:        3,
		topLeftY:        1,
		width:           1,
		height:          1,
		content:         3,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTr node = node{
		topLeftX:        2,
		topLeftY:        0,
		width:           2,
		height:          2,
		content:         0,
		topLeftNode:     &nodeTrTl,
		topRightNode:    &nodeTrTr,
		bottomLeftNode:  &nodeTrBl,
		bottomRightNode: &nodeTrBr,
	}
	var nodeBl node = node{
		topLeftX:        0,
		topLeftY:        2,
		width:           2,
		height:          2,
		content:         0,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeBr node = node{
		topLeftX:        2,
		topLeftY:        2,
		width:           2,
		height:          2,
		content:         2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var root node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           4,
		height:          4,
		content:         0,
		topLeftNode:     &nodeTl,
		topRightNode:    &nodeTr,
		bottomLeftNode:  &nodeBl,
		bottomRightNode: &nodeBr,
	}
	q.Width = 4
	q.Height = 4
	q.Root = &root
	contentHolder := make([][]int, 3)
	for i := 0; i < len(contentHolder); i++ {
		contentHolder[i] = make([]int, 3)
	}
	// Résultats attendus
	res0 := [][]int{
		{-1, -1, -1},
		{-1, 1, 1},
		{-1, 1, 1},
	}
	res1 := [][]int{
		{1, 1, 3},
		{1, 1, 4},
		{0, 0, 2},
	}
	res2 := [][]int{
		{4, 3, -1},
		{2, 2, -1},
		{2, 2, -1},
	}
	q.GetContent(-1, -1, contentHolder, generateFloor)
	if !equalTab(contentHolder, res0) {
		t.Error("Erreur dans l'affichage du terrain avec les coordonnées -1,-1\nRésultat : ", contentHolder, "\nRésultat attendu : ", res0)
	}
	q.GetContent(0, 0, contentHolder, generateFloor)
	if !equalTab(contentHolder, res1) {
		t.Error("Erreur dans l'affichage du terrain avec les coordonnées 0,0\nRésultat : ", contentHolder, "\nRésultat attendu : ", res1)
	}
	q.GetContent(2, 1, contentHolder, generateFloor)
	if !equalTab(contentHolder, res2) {
		t.Error("Erreur dans l'affichage du terrain avec les coordonnées 2,1\nRésultat : ", contentHolder, "\nRésultat attendu : ", res2)
	}
}

func TestGetContent0(t *testing.T) {
	// Test avec un terrain de taille 4x4 avec camera mode 0 et taille d'écran 3x3
	configuration.Global.NumTileX = 3
	configuration.Global.NumTileY = 3
	configuration.Global.ScreenCenterTileX = 1
	configuration.Global.ScreenCenterTileY = 1
	configuration.Global.CameraMode = 0
	var q Quadtree
	var nodeTl node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           2,
		height:          2,
		content:         1,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTrTl node = node{
		topLeftX:        2,
		topLeftY:        0,
		width:           1,
		height:          1,
		content:         3,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTrTr node = node{
		topLeftX:        3,
		topLeftY:        0,
		width:           1,
		height:          1,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTrBl node = node{
		topLeftX:        2,
		topLeftY:        1,
		width:           1,
		height:          1,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTrBr node = node{
		topLeftX:        3,
		topLeftY:        1,
		width:           1,
		height:          1,
		content:         3,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTr node = node{
		topLeftX:        2,
		topLeftY:        0,
		width:           2,
		height:          2,
		content:         0,
		topLeftNode:     &nodeTrTl,
		topRightNode:    &nodeTrTr,
		bottomLeftNode:  &nodeTrBl,
		bottomRightNode: &nodeTrBr,
	}
	var nodeBl node = node{
		topLeftX:        0,
		topLeftY:        2,
		width:           2,
		height:          2,
		content:         0,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeBr node = node{
		topLeftX:        2,
		topLeftY:        2,
		width:           2,
		height:          2,
		content:         2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var root node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           4,
		height:          4,
		content:         0,
		topLeftNode:     &nodeTl,
		topRightNode:    &nodeTr,
		bottomLeftNode:  &nodeBl,
		bottomRightNode: &nodeBr,
	}
	q.Width = 4
	q.Height = 4
	q.Root = &root
	contentHolder := make([][]int, 3)
	for i := 0; i < len(contentHolder); i++ {
		contentHolder[i] = make([]int, 3)
	}
	// Résultats attendus
	res0 := [][]int{
		{-1, -1, -1},
		{-1, 1, 1},
		{-1, 1, 1},
	}
	res1 := [][]int{
		{1, 1, 3},
		{1, 1, 4},
		{0, 0, 2},
	}
	res2 := [][]int{
		{4, 3, -1},
		{2, 2, -1},
		{2, 2, -1},
	}
	q.GetContent(-1, -1, contentHolder, generateFloor)
	if !equalTab(contentHolder, res0) {
		t.Error("Erreur dans l'affichage du terrain avec les coordonnées -1,-1\nRésultat : ", contentHolder, "\nRésultat attendu : ", res0)
	}
	q.GetContent(0, 0, contentHolder, generateFloor)
	if !equalTab(contentHolder, res1) {
		t.Error("Erreur dans l'affichage du terrain avec les coordonnées 0,0\nRésultat : ", contentHolder, "\nRésultat attendu : ", res1)
	}
	q.GetContent(2, 1, contentHolder, generateFloor)
	if !equalTab(contentHolder, res2) {
		t.Error("Erreur dans l'affichage du terrain avec les coordonnées 2,1\nRésultat : ", contentHolder, "\nRésultat attendu : ", res2)
	}
}

// equalTab permet de vérifier si deux tableaux sont exactement les mêmes
func equalTab(res, sol [][]int) bool {
	if len(res) != len(sol) {
		return false
	}
	for i := 0; i < len(res); i++ {
		if len(res[i]) != len(sol[i]) {
			return false
		}
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] != sol[i][j] {
				return false
			}
		}
	}
	return true
}
