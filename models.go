package main

import (
	"time"

	"github.com/trent-howard/go-rss/internal/database"
)

type User struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        int(dbUser.ID),
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    int       `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        int(dbFeed.ID),
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    int(dbFeed.UserID),
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `json:"user_id"`
	FeedID    int       `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        int(dbFollow.ID),
		CreatedAt: dbFollow.CreatedAt,
		UpdatedAt: dbFollow.UpdatedAt,
		UserID:    int(dbFollow.UserID),
		FeedID:    int(dbFollow.FeedID),
	}
}

func databaseFeedFollowsToFeedFollows(dbFollows []database.FeedFollow) []FeedFollow {
	follows := []FeedFollow{}
	for _, feed := range dbFollows {
		follows = append(follows, databaseFeedFollowToFeedFollow(feed))
	}
	return follows
}
