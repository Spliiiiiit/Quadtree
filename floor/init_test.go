package floor

import (
	"testing"
)

// Tests de la fonction readFloorFromFile
func TestReadFloor1(t *testing.T) {
	path := "../floor-files/exemple"
	res := readFloorFromFile(path)
	sol := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	if !equal(res, sol) {
		t.Error("Erreur dans le résultat de la lecture du fichier 'floor-files/exemple'\nRésultat : ", res, "\nRésultat attendu : ", sol)
	}
}
func TestReadFloor2(t *testing.T) {
	path := "../floor-files/beaupasbeau"
	res := readFloorFromFile(path)
	sol := [][]int{
		{0, 0, 1, 0, 2, 2, 2, 2},
		{0, 0, 1, 0, 2, 2, 2, 2},
		{0, 0, 1, 0, 2, 2, 2, 2},
		{0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0},
		{2, 2, 0, 0, 0, 1, 0, 0},
		{2, 2, 2, 0, 0, 1, 0, 0},
		{2, 2, 2, 0, 0, 1, 0, 0},
	}
	if !equal(res, sol) {
		t.Error("Erreur dans le résultat de la lecture du fichier 'floor-files/beaupasbeau'\nRésultat : ", res, "\nRésultat attendu : ", sol)
	}
}

// Fonction utile pour les tests, permet de vérifier si deux tableaux de tableaux d'entiers sont identiques
func equal(res, sol [][]int) bool {
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
