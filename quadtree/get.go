package quadtree

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// generateMap est une fonction utilisée pour générer des terrains aléatoires
// d'une taille donnée
type generateMap func(int, int) [][]int

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du quadtree q.
func (q *Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int, gen generateMap) {
	if len(contentHolder) == 0 {
		return
	}
	for i := 0; i < len(contentHolder); i++ {
		for j := 0; j < len(contentHolder[i]); j++ {
			x := topLeftX + j
			y := topLeftY + i
			// Détecter si le bloc est en dehors de la map générée
			if (x < q.Root.topLeftX || x > (q.Root.topLeftX+(q.Root.width-1))) || (y < q.Root.topLeftY || y > (q.Root.topLeftY+(q.Root.height-1))) {
				if configuration.Global.InfiniteMap {
					q.expand(x, y)
				} else {
					contentHolder[i][j] = -1
					continue
				}
			}
			contentHolder[i][j] = q.Root.FindContent(x, y, gen)
		}
	}
}

// FindContent renvoie le contenu présent aux coordonnées indiquées en partant
// d'une racine d'un quadtree
func (n *node) FindContent(x, y int, gen generateMap) (content int) {
	if n == nil {
		panic("erreur lors de la lecture du quadtree, la node est vide")
	}
	if (x < n.topLeftX || x > (n.topLeftX+(n.width-1))) || (y < n.topLeftY || y > (n.topLeftY+(n.height-1))) {
		return -1
	}
	children := []*node{
		n.topLeftNode,
		n.topRightNode,
		n.bottomLeftNode,
		n.bottomRightNode,
	}
	endNode := true
	for _, child := range children {
		if child != nil {
			content = child.FindContent(x, y, gen)
			if content == -2 && configuration.Global.InfiniteMap {
				err := child.fill(x, y, gen)
				if !err {
					panic("erreur lors de la génération de la carte")
				}
				return child.FindContent(x, y, gen)
			} else if content != -1 {
				return content
			}
			endNode = false
		}
	}
	if endNode {
		return n.content
	}
	return -1
}
