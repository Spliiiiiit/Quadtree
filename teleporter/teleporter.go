package teleporter

// Teleporter est le type permettant de représenter les
// données du téléporteur.
//
// - x1, y1, x2, y2 : sont les coordonnées des deux portails.
// - p1, p2 : indiquent si le portail 1 et 2 sont placés.
// - lastX, lastY : indique les dernières coordonnées du personnage.
type Teleporter struct {
	x1, y1, x2, y2 int
	p1, p2         bool
	lastX, lastY   int
	canTeleport    bool
}
