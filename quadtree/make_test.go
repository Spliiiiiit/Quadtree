package quadtree

import (
	"testing"
)

// Tests de la fonction MakeFromArray
func TestMakeFromArray1(t *testing.T) {
	floorContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	res := MakeFromArray(floorContent)
	var sol Quadtree
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
	sol.Width = 4
	sol.Height = 4
	sol.Root = &root
	if !equalQuadtree(res, sol) {
		t.Error("Erreur dans la construction du quadtree")
	}
}

func TestMakeFromArray2(t *testing.T) {
	floorContent := [][]int{
		{1, 1, 2, 2},
		{3, 3, 4, 4},
	}
	res := MakeFromArray(floorContent)
	var sol Quadtree
	var nodeTl node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           2,
		height:          1,
		content:         1,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTr node = node{
		topLeftX:        2,
		topLeftY:        0,
		width:           2,
		height:          1,
		content:         2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var nodeBl node = node{
		topLeftX:        0,
		topLeftY:        1,
		width:           2,
		height:          1,
		content:         3,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeBr node = node{
		topLeftX:        2,
		topLeftY:        1,
		width:           2,
		height:          1,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var root node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           4,
		height:          2,
		content:         0,
		topLeftNode:     &nodeTl,
		topRightNode:    &nodeTr,
		bottomLeftNode:  &nodeBl,
		bottomRightNode: &nodeBr,
	}
	sol.Width = 4
	sol.Height = 2
	sol.Root = &root
	if !equalQuadtree(res, sol) {
		logQuadtree(res)
		t.Error("Erreur dans la construction d'un quadtree qui représente un terrain rectangulaire")
	}
}

func TestMakeFromArray3(t *testing.T) {
	floorContent := [][]int{
		{1, 1, 2, 2, 2},
		{3, 3, 4, 4, 4},
		{3, 3, 4, 4, 4},
	}
	res := MakeFromArray(floorContent)
	var sol Quadtree
	var nodeTl node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           2,
		height:          1,
		content:         1,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTr node = node{
		topLeftX:        2,
		topLeftY:        0,
		width:           3,
		height:          1,
		content:         2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var nodeBl node = node{
		topLeftX:        0,
		topLeftY:        1,
		width:           2,
		height:          2,
		content:         3,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeBr node = node{
		topLeftX:        2,
		topLeftY:        1,
		width:           3,
		height:          2,
		content:         4,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var root node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           5,
		height:          3,
		content:         0,
		topLeftNode:     &nodeTl,
		topRightNode:    &nodeTr,
		bottomLeftNode:  &nodeBl,
		bottomRightNode: &nodeBr,
	}
	sol.Width = 5
	sol.Height = 3
	sol.Root = &root
	if !equalQuadtree(res, sol) {
		logQuadtree(res)
		t.Error("Erreur dans la construction d'un quadtree qui représente un terrain avec une largeur et une hauteur impaires")
	}
}

func TestMakeFromArray4(t *testing.T) {
	floorContent := [][]int{
		{1, 1, 2, 2, 2},
	}
	res := MakeFromArray(floorContent)
	var sol Quadtree
	var nodeTl node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           2,
		height:          1,
		content:         1,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeTr node = node{
		topLeftX:        2,
		topLeftY:        0,
		width:           3,
		height:          1,
		content:         2,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	var root node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           5,
		height:          1,
		content:         0,
		topLeftNode:     &nodeTl,
		topRightNode:    &nodeTr,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
	sol.Width = 5
	sol.Height = 1
	sol.Root = &root
	if !equalQuadtree(res, sol) {
		logQuadtree(res)
		t.Error("Erreur dans la construction d'un quadtree qui représente un terrain avec une hauteur de 1")
	}
}

func TestMakeFromArray5(t *testing.T) {
	floorContent := [][]int{
		{1},
		{2},
		{2},
	}
	res := MakeFromArray(floorContent)
	var sol Quadtree
	var nodeTl node = node{
		topLeftX:        0,
		topLeftY:        0,
		width:           1,
		height:          1,
		content:         1,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomRightNode: nil,
		bottomLeftNode:  nil,
	}
	var nodeBl node = node{
		topLeftX:        0,
		topLeftY:        1,
		width:           1,
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
		width:           1,
		height:          3,
		content:         0,
		topLeftNode:     &nodeTl,
		topRightNode:    nil,
		bottomLeftNode:  &nodeBl,
		bottomRightNode: nil,
	}
	sol.Width = 1
	sol.Height = 3
	sol.Root = &root
	if !equalQuadtree(res, sol) {
		logQuadtree(res)
		t.Error("Erreur dans la construction d'un quadtree qui représente un terrain avec une largeur de 1")
	}
}

// equalQuadtree permet de vérifier si les deux Quadtree passés en
// argument sont exactement les mêmes
func equalQuadtree(res, sol Quadtree) bool {
	if res.Width != sol.Width || res.Height != sol.Height {
		return false
	}
	if equalNodes(*res.Root, *sol.Root) {
		return true
	}
	return false
}

// equalNodes permet de vérifier si les deux node passés en
// argument sont exactement les mêmes
func equalNodes(res, sol node) bool {
	if res.width != sol.width || res.height != sol.height || res.content != sol.content || res.topLeftX != sol.topLeftX || res.topLeftY != sol.topLeftY {
		return false
	}
	if (res.topLeftNode == nil || sol.topLeftNode == nil) && res.topLeftNode != sol.topLeftNode {
		return false
	} else if res.topLeftNode != nil && sol.topLeftNode != nil && !equalNodes(*res.topLeftNode, *sol.topLeftNode) {
		return false
	}
	if (res.topRightNode == nil || sol.topRightNode == nil) && res.topRightNode != sol.topRightNode {
		return false
	} else if res.topRightNode != nil && sol.topRightNode != nil && !equalNodes(*res.topRightNode, *sol.topRightNode) {
		return false
	}
	if (res.bottomLeftNode == nil || sol.bottomLeftNode == nil) && res.bottomLeftNode != sol.bottomLeftNode {
		return false
	} else if res.bottomLeftNode != nil && sol.bottomLeftNode != nil && !equalNodes(*res.bottomLeftNode, *sol.bottomLeftNode) {
		return false
	}
	if (res.bottomRightNode == nil || sol.bottomRightNode == nil) && res.bottomRightNode != sol.bottomRightNode {
		return false
	} else if res.bottomRightNode != nil && sol.bottomRightNode != nil && !equalNodes(*res.bottomRightNode, *sol.bottomRightNode) {
		return false
	}
	return true
}
