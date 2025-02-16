package main

import (
	"os"

	"github.com/rancher/k3s/pkg/configfilearg"
	"github.com/rancher/rke2/pkg/cli/cmds"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cmds.NewApp()
	app.Commands = []cli.Command{
		cmds.NewServerCommand(),
		cmds.NewAgentCommand(),
		cmds.NewEtcdSnapshotCommand(),
		cmds.NewCertRotateCommand(),
	}

	if err := app.Run(configfilearg.MustParse(os.Args)); err != nil {
		logrus.Fatal(err)
	}
}
