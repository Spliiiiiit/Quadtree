package configuration

import (
	"encoding/json"
	"log"
	"os"
)

// randomMap définit les propriétés de la génération aléatoire de terrain
//
// Les champs sont :
//   - Enabled : indique si on utilise la génération aléatoire.
//   - Width : la largeur du terrain à générer.
//   - Height : la hauteur du terrain à générer.
//   - BlockProbability : un tableau contenant la probabilité
//     d'apparition de chaque bloc.
type randomMap struct {
	Enabled          bool
	Width            int
	Height           int
	BlockProbability []int
}

// animatedBlock définit les propriétés de l'animation d'un bloc
//
// Les champs sont :
//   - Enabled : indique si on utilise l'extension.
//   - NumAnimImages : le nombre d'images de l'animation du bloc.
//   - NumFramePerAnimImage : le nombre de 1/60 de seconde qui ont
//     lieu entre deux images de l'animation du bloc.
type animatedBlock struct {
	Enabled              bool
	NumAnimImages        int
	NumFramePerAnimImage int
}

// Configuration définit les élèments de la configuration
// du jeu. Pour ajouter un élèment de configuration, il
// suffit d'ajouter un champ dans cette structure.
//
// Les champs directement lus dans le fichier de configuration sont :
//   - DebugMode : indique si on est en mode debug ou pas
//   - NumTileX, NumTileY : les nombres de cases affichées à l'écran
//     en largeur et hauteur.
//   - Padding : le nombre de blocs à générer autour de la caméra qui ne
//     seront pas affichés.
//   - TileSize : la taille en pixels du côté d'une case.
//   - NumCharacterAnimImages : le nombre d'images de l'animation du
//     personnage.
//   - NumFramePerCharacterAnimImage : le nombre d'appels à update (ou
//     de 1/60 de seconde) qui ont lieu entre deux images de l'animation
//     du personnage.
//   - NumTileForDebug : le nombre de cases à ajouter à droite de l'écran
//     pour afficher les informations de debug
//   - CameraMode : le type de caméra à utiliser (0 pour une caméra fixe
//     et 1 pour une caméra qui suit le personnage, 2 pour une caméra fluide).
//   - FloorKind : détermine la méthode à utiliser pour afficher le terrain
//     (quadrillage, lecture dans un fichier, quadtree, etc)
//   - FloorFile : le chemin d'un fichier où lire les informations sur le
//     terrain si nécessaire
//   - WaterWalk : indique si le personnage peut marcher sur l'eau
//   - RandomMap : indique si le terrain est généré aléatoirement et ses
//     propriétiés
//   - InfiniteMap : indique si le terrain est généré au fur et à mesure
//     de l'exploration
//   - SaveMap : indique si la carte générée aléatoirement doit être
//     sauvegardée dans un fichier (dans le dossier "saved-maps").
//   - AnimatedWater : indique si l'eau est animée et ses propriétés
//   - Teleporter : indique s'il est possible de placer des téléporteurs
//
// Les champs calculés à partir des précédents sont :
//   - ScreenWidth, ScreenHeight : la largeur et la hauteur de l'écran
//     en pixels (hors zone d'affichage pour le debug)
//   - ScreenCenterTileX, ScreenCenterTileY : les coordonnées de la case
//     au centre de l'écran, où sera placée la caméra.
type Configuration struct {
	DebugMode                     bool
	NumTileX, NumTileY            int
	Padding                       int
	TileSize                      int
	NumCharacterAnimImages        int
	NumFramePerCharacterAnimImage int
	NumTileForDebug               int
	CameraMode                    int
	FloorKind                     int
	FloorFile                     string
	WaterWalk                     bool
	RandomMap                     randomMap
	InfiniteMap                   bool
	SaveMap                       bool
	AnimatedWater                 animatedBlock
	Teleporter                    animatedBlock

	ScreenWidth, ScreenHeight            int `json:"-"`
	ScreenCenterTileX, ScreenCenterTileY int `json:"-"`
}

// Global est la variable qui contient la configuration
// du jeu. Sa valeur est fixée à partir de la lecture d'un
// fichier de configuration par la fonction Load. C'est
// cette variable qu'il faut lire (configuration.Global)
// pour accéder à la configuration depuis d'autres paquets.
var Global Configuration

// Load se charge de lire un fichier de configuration, de
// remplir les champs obtenus par simple lecture, puis
// d'appeler la fonction qui remplit les champs calculés.
func Load(configurationFileName string) {
	content, err := os.ReadFile(configurationFileName)
	if err != nil {
		log.Fatal("Error while opening configuration file: ", err)
	}

	err = json.Unmarshal(content, &Global)
	if err != nil {
		log.Fatal("Error while reading configuration file: ", err)
	}

	setComputedFields()
}

// setComputedFields se charge de remplir les champs calculés
// de la configuration à partir des autres champs.
func setComputedFields() {
	Global.ScreenWidth = Global.NumTileX * Global.TileSize
	Global.ScreenHeight = Global.NumTileY * Global.TileSize
	Global.ScreenCenterTileX = Global.NumTileX / 2
	Global.ScreenCenterTileY = Global.NumTileY / 2
}
