package jsonpatch

type InvalidPatchError struct {
	Message string
}

type PatchCompiler interface {
	Compile(id string, patch *Patch) (*CommandList, error)
}

func (e InvalidPatchError) Error() string {
	return e.Message
}

type CommandList struct {
	Commands []CommandContainer
}

type ContainerCallback func() error

type CommandContainer interface {
	ExecuteBefore(interface{}) error
	ExecuteMain(interface{}) error
	ExecuteAfter(interface{}) error
}
