package grset

func (t *total) add() {
	t.mu.Lock()
	t.n++
	t.mu.Unlock()
}

func (t *total) done() {
	t.mu.Lock()
	t.n--
	t.mu.Unlock()
}

func (t *total) load() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.n
}
