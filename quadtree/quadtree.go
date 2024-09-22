package quadtree

import (
	"log"
	"os"
)

// Quadtree est la structure de données pour les arbres
// quaternaires. Les champs non exportés sont :
//   - Width, Height : la taille en cases de la zone représentée
//     par l'arbre.
//   - Root : le nœud qui est la racine de l'arbre.
type Quadtree struct {
	Width, Height int
	Root          *node
}

// Node représente un nœud d'arbre quaternaire. Les champs sont :
//   - topLeftX, topLeftY : les coordonnées (en cases) de la case
//     située en haut à gauche de la zone du terrain représentée
//     par ce nœud.
//   - width, height : la taille en cases de la zone représentée
//     par ce nœud.
//   - content : le type de terrain de la zone représentée par ce
//     nœud (seulement s'il s'agit d'une feuille).
//   - xxxNode : Une représentation de la partie xxx de la zone
//     représentée par ce nœud, différent de nil si et seulement
//     si le nœud actuel n'est pas une feuille.
type node struct {
	topLeftX, topLeftY int
	width, height      int
	content            int
	topLeftNode        *node
	topRightNode       *node
	bottomLeftNode     *node
	bottomRightNode    *node
}

// logQuadtree fonction utile pour le débogage, permet d'afficher
// un quadtree avec tous ses nœuds sur la sortie standard
func logQuadtree(q Quadtree) {
	logQuad := log.New(os.Stdout, "Quadtree: ", log.Lshortfile|log.Lmsgprefix)
	logQuad.Println("--- Affichage du quadtree ---")
	logQuad.Println("Taille : ", q.Width, "x", q.Height)
	logQuad.Println("Format : (X,Y), (W,H), Content")
	if q.Root == nil {
		logQuad.Println("Le quadtree est vide")
	} else {
		logQuadtreeNode(q.Root, "R")
	}
	logQuad.Println("-----------------------------")
}

// logQuadtreeNode fonction qui permet d'afficher les informations
// d'un nœud d'un quadtree et de tous ses autres nœuds enfants
func logQuadtreeNode(n *node, index string) {
	logNode := log.New(os.Stdout, "Noeud ", log.Lshortfile|log.Lmsgprefix)
	logNode.Println(index, ": (", n.topLeftX, n.topLeftY, "), (", n.width, n.height, "), ", n.content)
	if n.topLeftNode != nil {
		logQuadtreeNode(n.topLeftNode, index+".tl")
	}
	if n.topRightNode != nil {
		logQuadtreeNode(n.topRightNode, index+".tr")
	}
	if n.bottomLeftNode != nil {
		logQuadtreeNode(n.bottomLeftNode, index+".bl")
	}
	if n.bottomRightNode != nil {
		logQuadtreeNode(n.bottomRightNode, index+".br")
	}
}
