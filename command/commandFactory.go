package command

type concreteFactory struct {
	cmdsByName map[string]Command
}

func NewFactory() (factory concreteFactory) {
	factory.cmdsByName = make(map[string]Command)
	//factory.cmdsByName["login"]
	//command for triggering a build
	//factory.cmdsByName["getAppInfo"]
	return
}

