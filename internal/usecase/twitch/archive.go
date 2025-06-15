package twitch

import (
	"context"
	"errors"
	"time"

	"github.com/nicklaw5/helix/v2"
)

// GetArchives アーカイブ動画を取得する
func (u *UseCase) GetArchives(ctx context.Context) ([]helix.Video, error) {
	// ログインユーザの取得
	user, err := u.twitchClient.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("failed to get logged in user info")
	}

	// フォロー中配信者取得
	follows, err := u.twitchClient.GetFollowedChannels(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	// 昨日分のアーカイブ動画を取得
	var (
		now 			= time.Now().UTC()
		yesterdayStart 	= time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.UTC)
		yesterdayEnd 	= yesterdayStart.Add(24 * time.Hour)
		filteredVideos 	= make([]helix.Video, 0)
	)
	for _, followed := range follows {
		videos, err := u.twitchClient.GetArchiveVideos(ctx, followed.BroadcasterID)
		if err != nil {
			return nil, err
		}

		for _, video := range videos {
			createdAt, err := time.Parse(time.RFC3339, video.CreatedAt)
			if err != nil {
				continue
			}
			if createdAt.After(yesterdayStart) && createdAt.Before(yesterdayEnd) {
				filteredVideos = append(filteredVideos, video)
			}
		}
	}

	return filteredVideos, nil
}
