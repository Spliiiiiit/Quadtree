package teleporter

// place permet de placer un portail
func (t *Teleporter) place(x, y int) {
	if !t.p1 {
		t.x1, t.y1 = x, y
		t.p1 = true
	} else if !t.p2 {
		t.x2, t.y2 = x, y
		t.p2 = true
	} else {
		t.x1, t.y1 = t.x2, t.y2
		t.x2, t.y2 = x, y
	}
	t.lastX, t.lastY = x, y
	t.canTeleport = false
}
