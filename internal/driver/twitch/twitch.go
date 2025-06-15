package twitch

import (
	"context"

	"github.com/nicklaw5/helix/v2"
	"github.com/qli8racn/twitch-archive/internal/adapter/twitch"
	"github.com/qli8racn/twitch-archive/internal/config"
)

// GetAuthorizationURL 認可 URL を取得する
func (c *Client) GetAuthorizationURL(ctx context.Context, state string) string {
	return c.sdk.GetAuthorizationURL(&helix.AuthorizationURLParams{
		ResponseType: "code",
		Scopes:       []string{"user:read:follows"},
		State:        state,
	})
}

// ExchangeCodeForToken 認可コードをアクセストークンに変換
func (c *Client) ExchangeCodeForToken(ctx context.Context, code string) error {
	tokenResp, err := c.sdk.RequestUserAccessToken(code)
	if err != nil {
		return err
	}
	c.sdk.SetUserAccessToken(tokenResp.Data.AccessToken)
	return nil
}

// GetUserInfo 認証済みユーザの情報を取得する
func (c *Client) GetAuthenticatedUser(ctx context.Context) (*helix.User, error) {
	resp, err := c.sdk.GetUsers(&helix.UsersParams{})
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Users) == 0 {
		return nil, nil
	}
	return &resp.Data.Users[0], nil
}

// GetFollowedChannels 認証済みユーザのフォロー済み配信者を取得する
func (c *Client) GetFollowedChannels(ctx context.Context, userID string) ([]helix.FollowedChannel, error) {
	resp, err := c.sdk.GetFollowedChannels(&helix.GetFollowedChannelParams{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data.FollowedChannels, nil
}

// GetArchiveVideos アーカイブビデオを取得する
func (c *Client) GetArchiveVideos(ctx context.Context, userID string) ([]helix.Video, error) {
	resp, err := c.sdk.GetVideos(&helix.VideosParams{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data.Videos, nil
}

type Client struct {
	sdk *helix.Client
}

func New(cfg *config.Config) (twitch.Client, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:     cfg.Twitch.ClientID,
		ClientSecret: cfg.Twitch.ClientSecret,
		RedirectURI:  cfg.Twitch.RedirectURI,
	})
	if err != nil {
		return nil, err
	}

	resp, err := client.RequestAppAccessToken([]string{})
	if err != nil {
		return nil, err
	}
	client.SetAppAccessToken(resp.Data.AccessToken)

	return &Client{sdk: client}, nil
}
