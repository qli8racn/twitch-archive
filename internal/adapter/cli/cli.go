package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/qli8racn/twitch-archive/internal/usecase/twitch"
)

type InputParams struct {
	StartDate string `validate:"omitempty,datetime=2006-01-02"`
	EndDate   string `validate:"omitempty,datetime=2006-01-02"`
}

// Execute 実行関数
func (c *CLI) Execute(ctx context.Context, p InputParams) error {
	// 入力値の検証
	if err := c.validator.Struct(p); err != nil {
		return err
	}

	useCaseParams := twitch.GetArchivesParams{}
	if p.StartDate != "" {
		startT, _ := time.Parse("2006-01-02", p.StartDate)
		useCaseParams.StartAt = &startT
	}
	if p.EndDate != "" {
		endT, _ := time.Parse("2006-01-02", p.EndDate)
		endT = time.Date(endT.Year(), endT.Month(), endT.Day(), 23, 59, 59, 0, endT.Location())
		useCaseParams.EndAt = &endT
	}
	if useCaseParams.StartAt != nil && useCaseParams.EndAt != nil {
		if useCaseParams.EndAt.Before(*useCaseParams.StartAt) {
			return errors.New("end date must be after start date")
		}
	}

	// OAuth 認証の実施
	if err := c.twitchUseCase.OAuth(ctx); err != nil {
		return err
	}

	archives, err := c.twitchUseCase.GetArchives(ctx, useCaseParams)
	if err != nil {
		return err
	}
	fmt.Println(archives)

	return nil
}

type CLI struct {
	twitchUseCase *twitch.UseCase
	validator     *validator.Validate
}

func New(twitchUseCase *twitch.UseCase, v *validator.Validate) *CLI {
	return &CLI{twitchUseCase: twitchUseCase, validator: v}
}
