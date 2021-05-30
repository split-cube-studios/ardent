package engine

import "math"

// StateMachine handles states and their
// associated callbacks. States are represented
// as bitmasks, allowing multiple states to be active
// at the same time.
type StateMachine struct {
	state uint64

	callbacks map[uint64]func()
}

// NewStateMachine returns an instantiated StateMachine.
func NewStateMachine() *StateMachine {
	return &StateMachine{
		callbacks: make(map[uint64]func()),
	}
}

// UpdateState marks a given state as active or not.
func (sm *StateMachine) UpdateState(state uint64, active bool) {
	if active {
		sm.state |= state
	} else {
		sm.state &^= state
	}
}

// SetState deactivates all previous states and only
// sets the state specified.
func (sm *StateMachine) SetState(state uint64, active bool) {
	sm.SetAll(false)
	sm.UpdateState(state, active)
}

// SetAll marks all states as active or not.
func (sm *StateMachine) SetAll(active bool) {
	sm.state = 0
	if active {
		sm.state = math.MaxUint64
	}
}

// State returns the full state bitmask.
func (sm *StateMachine) State() uint64 {
	return sm.state
}

// Is indicates whether a given state is active.
func (sm *StateMachine) Is(state uint64) bool {
	return sm.state&state != 0
}

// IsOnly indicates whether a given state is active,
// and that it is the only active state.
func (sm *StateMachine) IsOnly(state uint64) bool {
	return sm.state&^state == 0
}

// SetCallback sets a given callback function for a given state.
func (sm *StateMachine) SetCallback(state uint64, cb func()) {
	sm.callbacks[state] = cb
}

// RemoveCallback removes a callback for a given state.
func (sm *StateMachine) RemoveCallback(state uint64) {
	delete(sm.callbacks, state)
}

// HandleCallbacks runs all callbacks for every active state.
func (sm *StateMachine) HandleCallbacks() {
	for state, cb := range sm.callbacks {
		if sm.Is(state) {
			cb()
		}
	}
}
