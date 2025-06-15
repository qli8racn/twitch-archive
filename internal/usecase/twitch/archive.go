package twitch

import (
	"context"
	"errors"
	"time"

	"github.com/nicklaw5/helix/v2"
)

type GetArchivesParams struct {
	StartAt *time.Time
	EndAt   *time.Time
}

// GetArchives アーカイブ動画を取得する
func (u *UseCase) GetArchives(ctx context.Context, p GetArchivesParams) ([]helix.Video, error) {
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
	var filteredVideos = make([]helix.Video, 0)
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
			if p.StartAt != nil && p.EndAt != nil {
				// 開始日と終了日が NULL ではない場合
				if createdAt.After(*p.StartAt) && createdAt.Before(*p.EndAt) {
					filteredVideos = append(filteredVideos, video)
				}
			} else if p.StartAt != nil {
				// 開始日のみ NULL ではない場合
				if createdAt.After(*p.StartAt) {
					filteredVideos = append(filteredVideos, video)
				}
			} else if p.EndAt != nil {
				// 終了日のみ NULL ではない場合
				if createdAt.Before(*p.EndAt) {
					filteredVideos = append(filteredVideos, video)
				}
			} else {
				filteredVideos = append(filteredVideos, video)
			}
		}
	}

	return filteredVideos, nil
}
