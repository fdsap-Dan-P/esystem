package db

import (
	"context"
	"simplebank/model"

	"github.com/google/uuid"
)

// var QueriesAccount *account.QueriesAccount = account.New(testDB)

type QuerierSocialMedia interface {
	CreateComment(ctx context.Context, arg CommentRequest) (model.Comment, error)
	GetComment(ctx context.Context, uuid uuid.UUID) (CommentInfo, error)
	GetCommentbyUuId(ctx context.Context, uuid uuid.UUID) (CommentInfo, error)
	ListComment(ctx context.Context, arg ListCommentParams) ([]CommentInfo, error)
	UpdateComment(ctx context.Context, arg CommentRequest) (model.Comment, error)
	DeleteComment(ctx context.Context, uuid uuid.UUID) error

	CreateFollower(ctx context.Context, arg FollowerRequest) (model.Follower, error)
	GetFollower(ctx context.Context, userId int64, followedId int64) (FollowerInfo, error)
	ListFollower(ctx context.Context, arg ListFollowerParams) ([]FollowerInfo, error)
	UpdateFollower(ctx context.Context, arg FollowerRequest) (model.Follower, error)
	DeleteFollower(ctx context.Context, userId int64, followedId int64) error

	CreatePost(ctx context.Context, arg PostRequest) (model.Post, error)
	GetPost(ctx context.Context, uuid uuid.UUID) (PostInfo, error)
	GetPostbyUuId(ctx context.Context, uuid uuid.UUID) (PostInfo, error)
	ListPost(ctx context.Context, arg ListPostParams) ([]PostInfo, error)
	UpdatePost(ctx context.Context, arg PostRequest) (model.Post, error)
	DeletePost(ctx context.Context, uuid uuid.UUID) error
}

var _ QuerierSocialMedia = (*QueriesSocialMedia)(nil)
