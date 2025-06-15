package main

import (
	"github.com/qli8racn/twitch-archive/internal/adapter/cli"
	"github.com/qli8racn/twitch-archive/internal/config"
	twitchImpl "github.com/qli8racn/twitch-archive/internal/driver/twitch"
	twitchUseCase "github.com/qli8racn/twitch-archive/internal/usecase/twitch"
	validatorPkg "github.com/qli8racn/twitch-archive/pkg/validator"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

func main() {
	container := build()

	var (
		start string
		end   string
	)
	cmd := &cobra.Command{
		Use: "app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return container.Invoke(func(c *cli.CLI) {
				err := c.Execute(cmd.Context(), cli.InputParams{
					StartDate: start,
					EndDate:   end,
				})
				if err != nil {
					panic(err)
				}
			})
		},
	}
	cmd.Flags().StringVarP(&start, "start", "s", "", "Start Date (YYYY-MM-DD)")
	cmd.Flags().StringVarP(&end, "end", "e", "", "End Date (YYYY-MM-DD)")

	cobra.CheckErr(cmd.Execute())
}

func build() *dig.Container {
	container := dig.New()

	// validator を提供
	container.Provide(validatorPkg.New)

	container.Provide(config.New)
	container.Provide(twitchImpl.New)
	container.Provide(twitchUseCase.New)
	container.Provide(cli.New)

	return container
}
