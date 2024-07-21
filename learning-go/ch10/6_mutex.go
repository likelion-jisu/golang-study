package main

import "sync"

func main() {

}

type MutextScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func NewMutextScoreboardManager() *MutextScoreboardManager {
	return &MutextScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *MutextScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutextScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}
