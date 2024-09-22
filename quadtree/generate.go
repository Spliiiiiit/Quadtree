package quadtree

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// expand modifie le quadtree actuel pour la root soit un enfant
// d'un nouveau quadtree
func (q *Quadtree) expand(x, y int) {
	w, h := q.Width, q.Height
	// Créer un quadtree vide et lui ajouter la root du quadtree actuel
	var parent Quadtree
	var oldRoot node = *q.Root
	parent.Width, parent.Height = 2*w, 2*h
	parentRoot := node{
		topLeftX: oldRoot.topLeftX,
		topLeftY: oldRoot.topLeftY,
		width:    parent.Width,
		height:   parent.Height,
		content:  0,
	}
	// Placer la root
	if x < oldRoot.topLeftX && y > (oldRoot.topLeftY+(h-1)) {
		// expand top right
		parentRoot.topRightNode = &oldRoot
		parentRoot.topLeftX -= w
	} else if x > (oldRoot.topLeftX+(w-1)) && y < oldRoot.topLeftY {
		// expand bottom left
		parentRoot.bottomLeftNode = &oldRoot
		parentRoot.topLeftY -= h
	} else if x < oldRoot.topLeftX || y < oldRoot.topLeftY {
		// expand bottom right
		parentRoot.bottomRightNode = &oldRoot
		parentRoot.topLeftX -= w
		parentRoot.topLeftY -= h
	} else if x > (oldRoot.topLeftX+(w-1)) || y > (oldRoot.topLeftY+(h-1)) {
		// expand top left
		parentRoot.topLeftNode = &oldRoot
	} else {
		panic("erreur la zone à générer n'est pas en dehors du quadtree actuel")
	}
	// Créer les nodes vides
	if parentRoot.topLeftNode == nil {
		parentRoot.topLeftNode = createEmptyNode(parentRoot.topLeftX, parentRoot.topLeftY, w, h)
	}
	if parentRoot.topRightNode == nil {
		parentRoot.topRightNode = createEmptyNode(parentRoot.topLeftX+w, parentRoot.topLeftY, w, h)
	}
	if parentRoot.bottomLeftNode == nil {
		parentRoot.bottomLeftNode = createEmptyNode(parentRoot.topLeftX, parentRoot.topLeftY+h, w, h)
	}
	if parentRoot.bottomRightNode == nil {
		parentRoot.bottomRightNode = createEmptyNode(parentRoot.topLeftX+w, parentRoot.topLeftY+h, w, h)
	}
	parent.Root = &parentRoot
	// Remplacer le quadtree actuel par celui modifié
	*q = parent
}

// fill remplace la node par un terrain généré aléatoirement ou par un quadtree avec
// des zones non-générées, renvoie true si le terrain a pu être généré
func (n *node) fill(x, y int, gen generateMap) bool {
	w := configuration.Global.RandomMap.Width
	h := configuration.Global.RandomMap.Height
	if n == nil {
		return false
	} else if (x < n.topLeftX || x > (n.topLeftX+(n.width-1))) || (y < n.topLeftY || y > (n.topLeftY+(n.height-1))) {
		return false
	} else if n.content != -2 {
		return false
	}
	if n.width > w && n.height > h {
		// création des 4 nodes vides (-2) qui seront les enfants
		// de la node actuelle
		n.content = 0
		if n.width%2 != 0 || n.height%2 != 0 {
			return false
		}
		n.topLeftNode = createEmptyNode(n.topLeftX, n.topLeftY, n.width/2, n.height/2)
		n.topRightNode = createEmptyNode(n.topLeftX+n.width/2, n.topLeftY, n.width/2, n.height/2)
		n.bottomLeftNode = createEmptyNode(n.topLeftX, n.topLeftY+n.height/2, n.width/2, n.height/2)
		n.bottomRightNode = createEmptyNode(n.topLeftX+n.width/2, n.topLeftY+n.height/2, n.width/2, n.height/2)
	} else if n.width == w && n.height == h {
		// la node va être remplacée par un quadtree représentant
		// un terrain aléatoire
		q := MakeFromArray(gen(w, h))
		q.Root.changeCoordinates(n.topLeftX, n.topLeftY)
		*n = *q.Root
	} else {
		return false
	}
	return true
}

// Fonctions utilitaires

// createEmptyNode créer une node vide avec un contenu non généré
func createEmptyNode(topLeftX, topLeftY, w, h int) *node {
	return &node{
		topLeftX:        topLeftX,
		topLeftY:        topLeftY,
		width:           w,
		height:          h,
		content:         -2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
}

// changeCoordinates incrémente les coordonnées du node et de ses enfants
func (n *node) changeCoordinates(xInc, yInc int) {
	if n == nil {
		panic("erreur lors du changement des coordonnées, la node est vide")
	}
	n.topLeftX += xInc
	n.topLeftY += yInc
	children := []*node{
		n.topLeftNode,
		n.topRightNode,
		n.bottomLeftNode,
		n.bottomRightNode,
	}
	for _, child := range children {
		if child != nil {
			child.changeCoordinates(xInc, yInc)
		}
	}
}
