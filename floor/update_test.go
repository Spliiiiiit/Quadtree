package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"testing"
)

// Tests de la fonction updateFromFileFloor
func TestUpdateFloorCMode1(t *testing.T) {
	// Simulation d'une configuration avec une taille d'écran de 3x3 avec CameraMode = 1
	configuration.Global.ScreenCenterTileX = 1
	configuration.Global.ScreenCenterTileY = 1
	configuration.Global.CameraMode = 1
	//Création des tableaux 'Content' et 'fullContent'
	var testFloor Floor
	testFloor.Content = make([][]int, 3)
	for i := 0; i < len(testFloor.Content); i++ {
		testFloor.Content[i] = make([]int, 3)
	}
	testFloor.fullContent = [][]int{
		{0, 1, 0, 1},
		{1, 0, 1, 0},
		{2, 1, 2, 1},
		{1, 2, 1, 2},
	}
	// Résultats attendus
	res0 := [][]int{
		{-1, -1, -1},
		{-1, 0, 1},
		{-1, 1, 0},
	}
	res1 := [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{2, 1, 2},
	}
	res2 := [][]int{
		{0, 1, 0},
		{1, 2, 1},
		{2, 1, 2},
	}
	// Tests
	testFloor.updateFromFileFloor(0, 0)
	if !equal(testFloor.Content, res0) {
		t.Error("Erreur dans le résultat de la mise à jour de Content avec les coordonnées (0,0)\nRésultat : ", testFloor.Content, "\nRésultat attendu : ", res0)
	}
	testFloor.updateFromFileFloor(1, 1)
	if !equal(testFloor.Content, res1) {
		t.Error("Erreur dans le résultat de la mise à jour de Content avec les coordonnées (1,1)\nRésultat : ", testFloor.Content, "\nRésultat attendu : ", res1)
	}
	testFloor.updateFromFileFloor(2, 2)
	if !equal(testFloor.Content, res2) {
		t.Error("Erreur dans le résultat de la mise à jour de Content avec les coordonnées (2,2)\nRésultat : ", testFloor.Content, "\nRésultat attendu : ", res2)
	}
}

func TestUpdateFloorCMode0(t *testing.T) {
	// Simulation d'une configuration avec une taille d'écran de 3x3 avec CameraMode = 0
	configuration.Global.ScreenCenterTileX = 1
	configuration.Global.ScreenCenterTileY = 1
	configuration.Global.CameraMode = 0
	//Création des tableaux 'Content' et 'fullContent'
	var testFloor Floor
	testFloor.Content = make([][]int, 3)
	for i := 0; i < len(testFloor.Content); i++ {
		testFloor.Content[i] = make([]int, 3)
	}
	testFloor.fullContent = [][]int{
		{0, 1, 0, 1},
		{1, 0, 1, 0},
		{2, 1, 2, 1},
		{1, 2, 1, 2},
	}
	// Résultats attendus
	res0 := [][]int{
		{-1, -1, -1},
		{-1, 0, 1},
		{-1, 1, 0},
	}
	res1 := [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{2, 1, 2},
	}
	res2 := [][]int{
		{0, 1, 0},
		{1, 2, 1},
		{2, 1, 2},
	}
	// Tests
	testFloor.updateFromFileFloor(0, 0)
	if !equal(testFloor.Content, res0) {
		t.Error("Erreur dans le résultat de la mise à jour de Content avec les coordonnées (0,0)\nRésultat : ", testFloor.Content, "\nRésultat attendu : ", res0)
	}
	testFloor.updateFromFileFloor(1, 1)
	if !equal(testFloor.Content, res1) {
		t.Error("Erreur dans le résultat de la mise à jour de Content avec les coordonnées (1,1)\nRésultat : ", testFloor.Content, "\nRésultat attendu : ", res1)
	}
	testFloor.updateFromFileFloor(2, 2)
	if !equal(testFloor.Content, res2) {
		t.Error("Erreur dans le résultat de la mise à jour de Content avec les coordonnées (2,2)\nRésultat : ", testFloor.Content, "\nRésultat attendu : ", res2)
	}
}
