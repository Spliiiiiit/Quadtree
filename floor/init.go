package floor

import (
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	const savePath string = "../saved-maps"
	f.Content = make([][]int, configuration.Global.NumTileY+configuration.Global.Padding*2)
	for y := 0; y < len(f.Content); y++ {
		f.Content[y] = make([]int, configuration.Global.NumTileX+configuration.Global.Padding*2)
	}
	if configuration.Global.RandomMap.Enabled {
		// Vérification de la validité des probabilités
		if configuration.Global.RandomMap.Width < 1 || configuration.Global.RandomMap.Height < 1 {
			panic("erreur la taille de la carte générée aléatoirement doit être supérieure ou égale à 1")
		}
		if len(configuration.Global.RandomMap.BlockProbability) != 5 {
			panic("erreur le tableau BlockProbability doit contenir exactement 5 nombres")
		}
		sum := 0
		for _, p := range configuration.Global.RandomMap.BlockProbability {
			if p < 0 || p > 100 {
				panic("erreur le tableau BlockProbability doit contenir des nombres compris entre 0 et 100")
			}
			sum += p
		}
		if sum != 100 {
			panic("erreur la somme des nombres du tableau BlockProbability doit être égale à 100")
		}
		if configuration.Global.InfiniteMap && configuration.Global.FloorKind != quadTreeFloor {
			panic("erreur le champ floorKind doit être à 2 (quadtree) pour utiliser la génération infinie")
		}
		switch configuration.Global.FloorKind {
		case fromFileFloor:
			f.fullContent = randomFloor(configuration.Global.RandomMap.Width, configuration.Global.RandomMap.Height)
		case quadTreeFloor:
			f.quadtreeContent = quadtree.MakeFromArray(randomFloor(configuration.Global.RandomMap.Width, configuration.Global.RandomMap.Height))
		}
		// Sauvegarde de la map
		if configuration.Global.SaveMap {
			f.saveMap(savePath)
		}
	} else {
		if configuration.Global.InfiniteMap {
			panic("erreur la génération aléatoire (RandomMap) doit être activée pour utiliser la génération infinie")
		}
		if configuration.Global.SaveMap {
			fmt.Println("SaveMap : impossible de sauvegardé la carte.\nLa génération aléatoire (RandomMap) doit être activée pour pouvoir sauvegarder la carte")
		}
		switch configuration.Global.FloorKind {
		case fromFileFloor:
			f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
		case quadTreeFloor:
			f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
		}
	}
}

// Lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	var floorFileContent []byte
	var err error
	// Lit le fichier
	floorFileContent, err = os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// Sépare les lignes
	var lines []string = strings.Split(string(floorFileContent), "\n")
	//Définition de la taille du tableau en fonction du nombre de lignes
	floorContent = make([][]int, len(lines))
	// Pour chaque ligne, on sépare les chiffres
	for i, l := range lines {
		// Création du tableau correspondant à une ligne
		// avec une taille qui dépend du nombre de chiffres contenu dans une ligne
		floorContent[i] = make([]int, len(l))
		for j, c := range l {
			// - 48 correspond à la conversion du code ascii en chiffre
			floorContent[i][j] = int(c) - 48
		}
	}
	// Log pour le mode Debug
	if configuration.Global.DebugMode {
		logFloor := log.New(os.Stdout, "readFloorFromFile: ", log.Lshortfile|log.Lmsgprefix)
		logFloor.Println("floorContent =\n", floorContent)
	}
	return floorContent
}

// randomFloor crée un tableau représentant un terrain aléatoire d'une taille donnée
func randomFloor(w, h int) (floorContent [][]int) {
	blocks := []int{0, 1, 2, 3, 4}
	var probs []int
	for k, block := range blocks {
		for i := 0; i < configuration.Global.RandomMap.BlockProbability[k]; i++ {
			probs = append(probs, block)
		}
	}
	floorContent = make([][]int, h)
	for i := 0; i < len(floorContent); i++ {
		floorContent[i] = make([]int, w)
		for j := 0; j < len(floorContent[i]); j++ {
			floorContent[i][j] = probs[rand.Intn(100)]
		}
	}
	return floorContent
}

// saveMapFromArray sauvegarde une carte représentée sous forme d'un tableau dans un fichier
func saveMapFromArray(mapFile *os.File, handleError func(err error), m [][]int) {
	for y := range m {
		var line string
		for x := range m[y] {
			line += strconv.Itoa(m[y][x])
		}
		if y == len(m)-1 {
			_, err := fmt.Fprint(mapFile, line)
			handleError(err)
			return
		}
		_, err := fmt.Fprintln(mapFile, line)
		handleError(err)
	}
}

func (f Floor) saveMap(directoryPath string) {
	if configuration.Global.InfiniteMap {
		fmt.Println("SaveMap : impossible de sauvegarder le terrain car 'InfiniteMap' est activée")
		return
	}
	handleError := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	var mapFile *os.File
	var err error
	currentTime := time.Now()
	var fileName string = directoryPath + "/" + strconv.Itoa(currentTime.Year()) + fmt.Sprintf("%02d", int(currentTime.Month())) +
		fmt.Sprintf("%02d", currentTime.Day()) + "_" + fmt.Sprintf("%02d", currentTime.Hour()) + fmt.Sprintf("%02d", currentTime.Minute())
	// Création du dossier si nécessaire
	if _, err = os.Stat(directoryPath); os.IsNotExist(err) {
		err = os.MkdirAll(directoryPath, 0700)
		handleError(err)
	}
	// Vérifier si le fichier existe déjà (même date)
	_, exist := os.Stat(fileName)
	i := 0
	for !os.IsNotExist(exist) {
		i++
		_, exist = os.Stat(fileName + "_" + strconv.Itoa(i))
	}
	if i != 0 {
		fileName += "_" + strconv.Itoa(i)
	}
	// Création du fichier
	mapFile, err = os.Create(fileName)
	handleError(err)
	// Sauvegarde de la map
	switch configuration.Global.FloorKind {
	case fromFileFloor:
		saveMapFromArray(mapFile, handleError, f.fullContent)
	case quadTreeFloor:
		quadtree.SaveMapFromQuadtree(mapFile, handleError, f.quadtreeContent, randomFloor)
	}
	err = mapFile.Close()
	handleError(err)
}
