package manager

import "spelling-bee-game/server/game"

type Manager interface {
	Create() (int, game.Game, error)
	GetGame(id int) (game.Game, bool)
	End(id int) error
}
