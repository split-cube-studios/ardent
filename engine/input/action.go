package input

type Action interface {
	CanPerform(State) bool
}

// todo - Maybe this should just be a typed function?
type ActionPerformer interface {
	Perform(State) bool
}

type funcActionPerformer struct {
	perform func(State) bool
}

func (f funcActionPerformer) Perform(s State) bool {
	return f.perform(s)
}

func PerformFunc(f func(State) bool) ActionPerformer {
	return funcActionPerformer{
		perform: f,
	}
}
