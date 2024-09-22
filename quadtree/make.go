package quadtree

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	if len(floorContent) != 0 && len(floorContent[0]) != 0 {
		q.Height = len(floorContent)
		q.Width = len(floorContent[0])
	} else {
		q.Height = 0
		q.Width = 0
		q.Root = nil
		return q
	}
	// Vérifier que le terrain est valide
	length := len(floorContent[0])
	for _, line := range floorContent {
		if len(line) != length {
			panic("erreur lors de la construction du terrain, toutes les lignes du terrain doivent avoir la même taille")
		}
	}
	q.Root = createNode(floorContent, 0, 0)
	// Log pour le mode Debug
	if configuration.Global.DebugMode && !configuration.Global.InfiniteMap {
		logQuadtree(q)
	}
	return q
}

// createNode génère les nodes de manière récursive d'un quadtree représentant un terrain
func createNode(zone [][]int, x, y int) *node {
	var n node
	// Définition de la hauteur et de la largeur :
	if len(zone) == 0 {
		panic("Erreur lors de la construction du terrain : une node ne peut pas représenter une zone vide")
	} else {
		n.height = len(zone)
	}
	if len(zone[0]) == 0 {
		panic("Erreur lors de la construction du terrain : une node ne peut pas représenter une zone vide")
	} else {
		n.width = len(zone[0])
	}
	// Définition topLeftX et topLeftY :
	n.topLeftX = x
	n.topLeftY = y
	// Définition du content :
	// On teste si la zone ne contient qu'un seul type de terrain
	floor := zone[0][0]
	sameFloor := true
	for _, line := range zone {
		for _, col := range line {
			if col != floor {
				sameFloor = false
				break
			}
		}
		if !sameFloor {
			break
		}
	}
	if sameFloor {
		n.content = floor
		n.topLeftNode = nil
		n.topRightNode = nil
		n.bottomLeftNode = nil
		n.bottomRightNode = nil
		return &n
	}
	// Définition des 4 zones et appel récursif
	// variables h/vsingle permettent de gérer les terrains impairs
	verticalMiddle := len(zone[0]) / 2 // X middle
	horizontalMiddle := len(zone) / 2  // Y middle
	vsingle := false
	if verticalMiddle == 0 {
		verticalMiddle = 1
		vsingle = true
	}
	hsingle := false
	if horizontalMiddle == 0 {
		horizontalMiddle = 1
		hsingle = true
	}
	var topLeftZ [][]int
	var topRightZ [][]int
	for i := 0; i < horizontalMiddle; i++ {
		topLeftZ = append(topLeftZ, zone[i][:verticalMiddle])
		topRightZ = append(topRightZ, zone[i][verticalMiddle:])
	}
	var botLeftZ [][]int
	var botRightZ [][]int
	for i := horizontalMiddle; i < len(zone); i++ {
		botLeftZ = append(botLeftZ, zone[i][:verticalMiddle])
		botRightZ = append(botRightZ, zone[i][verticalMiddle:])
	}
	n.topLeftNode = createNode(topLeftZ, x, y)
	if !vsingle {
		n.topRightNode = createNode(topRightZ, x+verticalMiddle, y)
	}
	if !hsingle {
		n.bottomLeftNode = createNode(botLeftZ, x, y+horizontalMiddle)
	}
	if !vsingle && !hsingle {
		n.bottomRightNode = createNode(botRightZ, x+verticalMiddle, y+horizontalMiddle)
	}
	return &n
}
