package command

import (
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/codegangsta/cli"
)


type Command interface {
	Configs() CommandConfigs
	Metadata() CommandMetadata
	GetRequirements(requirementsFactory requirements.Factory, context *cli.Context) (reqs []requirements.Requirement, err error)
	Run(context *cli.Context)
}
