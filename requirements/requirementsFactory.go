package requirements

type Factory interface {
	//Identify the requirement types and add here
}

type Requirement interface {
	Execute() (success bool)
}
