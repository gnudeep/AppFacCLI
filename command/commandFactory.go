package command

type concreteFactory struct {
	cmdsByName map[string]command.Command
}

func NewFactory() (factory concreteFactory) {
	factory.cmdsByName = make(map[string]command.Command)
	factory.cmdsByName["login"] = service.NewCreateService(ui, config, repoLocator.GetServiceRepository(), serviceBuilder)

}
