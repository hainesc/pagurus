package app

type Command struct {
	Run func() error
}

func (c *Command) Execute() error {
	return c.Run()
}
