package main

import (
	"example.com/internal/adapter/cli"
	"example.com/internal/config"
	twitchImpl "example.com/internal/driver/twitch"
	twitchUseCase "example.com/internal/usecase/twitch"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

func main() {
	container := build()

	cmd := &cobra.Command{
        Use: "app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return container.Invoke(func (c *cli.CLI) {
				c.Execute(cmd.Context())
			})
        },
    }
	cobra.CheckErr(cmd.Execute())
}

func build() *dig.Container {
	container := dig.New()

    container.Provide(config.New)
    container.Provide(twitchImpl.New)
	container.Provide(twitchUseCase.New)
	container.Provide(cli.New)

	return container
}
