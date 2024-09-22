package game

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisations, car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.character.Init()
	g.camera.Init()
	g.floor.Init()
	g.floor.Spawn(&g.character.X, &g.character.Y, &g.camera.X, &g.camera.Y)
}
