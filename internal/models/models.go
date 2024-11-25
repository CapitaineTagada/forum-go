package models

import (
	"database/sql"
	"forum-go/internal/shared"
	"time"
)

type User struct {
	UserId           string         `db:"user_id"`
	Email            string         `db:"email"`
	Username         string         `db:"username"`
	Password         string         `db:"password"`
	Role             string         `db:"role"`
	CreationDate     time.Time      `db:"creation_date"`
	SessionId        sql.NullString `db:"session_id"`
	SessionExpire    sql.NullTime   `db:"session_expire"`
	Posts            []Post         `db:"-"`
	Activities       []Activity     `db:"-"`
	UnreadActivities int            `db:"-"`
}

type Category struct {
	CategoryId string `db:"category_id"`
	Name       string `db:"name"`
}
type Post_Comment interface {
	GetUserLikes() []UserLike
}
type Post struct {
	PostId                string       `db:"post_id"`
	Title                 string       `db:"title"`
	Content               string       `db:"content"`
	UserID                string       `db:"user_id"`
	CreationDate          time.Time    `db:"creation_date"`
	FormattedCreationDate string       `db:"-"`
	UpdateDate            sql.NullTime `db:"update_date"`
	User                  User         `db:"-"`
	Categories            []Category   `db:"-"`
	Comments              []Comment    `db:"-"`
	NbOfComments          int          `db:"-"`
	UserLikes             []UserLike   `db:"-"`
	Likes                 int          `db:"-"`
	Dislikes              int          `db:"-"`
	HasVoted              int          `db:"-"`
}

type Comment struct {
	CommentId             string     `db:"comment_id"`
	Content               string     `db:"content"`
	CreationDate          time.Time  `db:"creation_date"`
	FormattedCreationDate string     `db:"-"`
	UserID                string     `db:"user_id"`
	PostID                string     `db:"post_id"`
	Username              string     `db:"-"`
	UserLikes             []UserLike `db:"-"`
	Likes                 int        `db:"-"`
	Dislikes              int        `db:"-"`
	HasVoted              int        `db:"-"`
}

type PostCategory struct {
	PostId     string `db:"post_id"`
	CategoryId string `db:"category_id"`
}

type UserLike struct {
	LikeId    string `db:"like_id"`
	UserId    string `db:"user_id"`
	PostId    string `db:"post_id"`
	CommentId string `db:"comment_id"`
	IsLike    bool   `db:"is_like"`
}
type Error struct {
	Message    string
	StatusCode int
}

type Activity struct {
	ActivityId            string    `db:"activity_id"`
	UserId                string    `db:"user_id"`
	ActionUserId          string    `db:"action_user_id"`
	ActionUsername        string    `db:"-"`
	ActionType            string    `db:"action_type"`
	PostId                string    `db:"post_id"`
	CommentId             string    `db:"comment_id"`
	CreationDate          time.Time `db:"creation_date"`
	FormattedCreationDate string    `db:"-"`
	Details               string    `db:"details"`
	IsRead                bool      `db:"is_read"`
}

func NewActivity(userId, actionUserId, actionType, postId, commentId, details string) Activity {
	activity := Activity{
		ActivityId:   shared.ParseUUID(shared.GenerateUUID()),
		UserId:       userId,
		ActionUserId: actionUserId,
		ActionType:   actionType,
		PostId:       postId,
		CommentId:    commentId,
		CreationDate: time.Now(),
		Details:      details,
		IsRead:       false,
	}
	return activity
}

func (post Post) GetUserLikes() []UserLike {
	return post.UserLikes
}

func (comment Comment) GetUserLikes() []UserLike {
	return comment.UserLikes
}

type ActionType string

const (
	POST_LIKED           ActionType = "postLiked"
	POST_DISLIKED        ActionType = "postDisliked"
	COMMENT_LIKED        ActionType = "commentLiked"
	COMMENT_DISLIKED     ActionType = "commentDisliked"
	GET_POST_LIKED       ActionType = "getPostLiked"
	GET_POST_DISLIKED    ActionType = "getPostDisliked"
	GET_COMMENT_LIKED    ActionType = "getCommentLiked"
	GET_COMMENT_DISLIKED ActionType = "getCommentDisliked"
	POST_CREATED         ActionType = "postCreated"
	COMMENT_CREATED      ActionType = "commentCreated"
	GET_COMMENT          ActionType = "getComment"
)
