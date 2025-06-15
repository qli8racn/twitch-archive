package twitch

import (
	"context"

	"github.com/nicklaw5/helix/v2"
)

type Client interface {
	GetAuthorizationURL(context.Context, string) string
	ExchangeCodeForToken(context.Context, string) error
    GetAuthenticatedUser(context.Context) (*helix.User, error)
	GetFollowedChannels(context.Context, string) ([]helix.FollowedChannel, error)
	GetArchiveVideos(context.Context, string) ([]helix.Video, error) 
}
