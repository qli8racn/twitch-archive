package cli

import (
	"context"
	"fmt"

	"example.com/internal/usecase/twitch"
)

// Execute 実行関数
func (c *CLI) Execute(ctx context.Context) error {
	// OAuth 認証の実施
	if err := c.twitchUseCase.OAuth(ctx); err != nil {
		panic(err)
	}

	archives, err := c.twitchUseCase.GetArchives(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(archives)

	return nil
}

type CLI struct {
	twitchUseCase *twitch.UseCase
}

func New(twitchUseCase *twitch.UseCase) *CLI {
	return &CLI{twitchUseCase: twitchUseCase}
}
