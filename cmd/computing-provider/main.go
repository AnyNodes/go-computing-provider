package main

import (
	"github.com/lagrangedao/go-computing-provider/build"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	FlagCpRepo = "cp-repo"
)

func main() {
	app := &cli.App{
		Name:                 "computing-provider",
		Usage:                "Swanchain decentralized computing network client",
		EnableBashCompletion: true,
		Version:              build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    FlagCpRepo,
				EnvVars: []string{"CP_PATH"},
				Usage:   "cp repo path",
				Value:   "~/.swan/computing",
			},
		},
		Commands: []*cli.Command{
			initCmd,
			runCmd,
			infoCmd,
			accountCmd,
			taskCmd,
			walletCmd,
			collateralCmd,
			ubiTaskCmd,
		},
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
	}
}
