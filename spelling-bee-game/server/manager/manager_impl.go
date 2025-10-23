package manager

import (
	"fmt"
	"spelling-bee-game/server/game"
	"sync"
)

var (
	once     sync.Once
	instance *manager
)

type manager struct {
	mutex  sync.RWMutex
	nextId int
	games  map[int]game.Game
}

func (m *manager) Create() (int, game.Game) {
	g := game.New()
	m.mutex.Lock()
	defer m.mutex.Unlock()
	id := m.nextId
	m.games[id] = g
	m.nextId++
	fmt.Println("Created game of id ", id)
	return id, g
}

func (m *manager) GetGame(id int) (game.Game, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	g, ok := m.games[id]
	return g, ok
}

func (m *manager) End(id int) error {
	m.games[id] = nil
	fmt.Println("Ended game of id ", id)
	return nil
}

func GetManager() Manager {
	once.Do(func() {
		instance = &manager{
			games: make(map[int]game.Game),
		}
	})
	return instance
}
