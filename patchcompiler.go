package jsonpatch

type InvalidPatchError struct {
	Message string
}

type PatchCompiler interface {
	Compile(patch *Patch, options map[string]interface{}) (*CommandList, error)
}

func (e InvalidPatchError) Error() string {
	return e.Message
}

type CommandList struct {
	Commands []*CommandContainer
}

func (cl *CommandList) AddCommands(commands ...*CommandContainer) {
	cl.Commands = append(cl.Commands, commands...)
}

type ContainerCallback func(transaction interface{}, previousResult interface{}) (interface{}, error)

type CommandContainer struct {
	BeforeCallback ContainerCallback
	MainCallback   ContainerCallback
	AfterCallback  ContainerCallback
}

func (c *CommandContainer) ExecuteMain(transaction interface{}, prev interface{}) (interface{}, error) {
	if c.MainCallback != nil {
		return c.MainCallback(transaction, prev)
	}
	return nil, nil
}

func (c *CommandContainer) ExecuteAfter(transaction, prev interface{}) (interface{}, error) {
	if c.AfterCallback != nil {
		return c.AfterCallback(transaction, prev)
	}
	return nil, nil
}

func (c *CommandContainer) ExecuteBefore(transaction interface{}) (interface{}, error) {
	if c.BeforeCallback != nil {
		return c.BeforeCallback(transaction, nil)
	}
	return nil, nil
}
