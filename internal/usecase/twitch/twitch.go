package twitch

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"example.com/internal/adapter/twitch"
	"example.com/internal/config"
)

// OAuth ユーザ認証の実行（アクセストークンの取得）
func (c *UseCase) OAuth(ctx context.Context) error {
	state := "some-random-state"
	authURL := c.twitchClient.GetAuthorizationURL(ctx, state)

	fmt.Println("下記URLをブラウザで開いて認証してください: ")
	fmt.Println(authURL)
	fmt.Println()

	fmt.Print("認可コード[ブラウザURLの?code=xxx]を入力してください: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	code := strings.TrimSpace(scanner.Text())
	if err := c.twitchClient.ExchangeCodeForToken(ctx, code); err != nil {
		return err
	}	
	return nil
}

type UseCase struct {
	cfg 		 *config.Config
	twitchClient twitch.Client
}

func New(cfg *config.Config, twitchClient twitch.Client) *UseCase {
	return &UseCase{cfg, twitchClient}
}
