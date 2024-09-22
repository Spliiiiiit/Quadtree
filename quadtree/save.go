package quadtree

import (
	"fmt"
	"os"
	"strconv"
)

// SaveMapFromQuadtree sauvegarde une carte représentée sous forme d'un quadtree dans un fichier
func SaveMapFromQuadtree(mapFile *os.File, handleError func(err error), q Quadtree, gen generateMap) {
	for y := 0; y < q.Height; y++ {
		var line string
		for x := 0; x < q.Width; x++ {
			line += strconv.Itoa(q.Root.FindContent(x, y, gen))
		}
		if y == q.Height-1 {
			_, err := fmt.Fprint(mapFile, line)
			handleError(err)
			return
		}
		_, err := fmt.Fprintln(mapFile, line)
		handleError(err)
	}
}
