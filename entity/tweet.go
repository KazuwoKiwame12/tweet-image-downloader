package entity

import (
	"fmt"
	"time"
)

// query parametersの"expasions"に関する設定
type ExpansionField string

const (
	ExpasionFieldMediaKeys ExpansionField = "attachments.media_keys"
)

func (e ExpansionField) String() string {
	return string(e)
}

// query parametersの"media.fields"に関する設定
type MediaField string

const (
	MediaFieldMediaKey MediaField = "media_key"
	MediaFieldURL      MediaField = "url"
)

func (m MediaField) String() string {
	return string(m)
}

// query parametersの"query"に関する設定
type QueryField string

const (
	QueryFieldHasImages QueryField = "has:images"
	QueryFieldIsRetweet QueryField = "is:retweet"
)

func (q QueryField) String() string {
	return string(q)
}

func QueryFieldFrom(user string) QueryField {
	return QueryField(fmt.Sprintf("from:%s", user))
}
func QueryFieldKeyword(keyword string) QueryField {
	return QueryField(fmt.Sprintf("\"%s\"", keyword))
}

func (q QueryField) NOT() QueryField {
	return QueryField(fmt.Sprintf("-%v", q))
}

// query parametersの"tweet.fields"に関する設定
type TweetField string

const (
	TweetFieldAttachments TweetField = "attachments"
	TweetFieldCreatedAt   TweetField = "created_at"
)

func (t TweetField) String() string {
	return string(t)
}

// responseに関する設定
type TweetResponse struct {
	Tweets []struct {
		ID          string    `json:"id"`
		Text        string    `json:"text"`
		CreatedAt   time.Time `json:"created_at"`
		Attachments struct {
			MediaKeys []string `json:"media_keys"`
		} `json:"attachments"`
	} `json:"data"`
	Includes struct {
		Media []struct {
			MediaKey string `json:"media_key"`
			URL      string `json:"url"`
		} `json:"media"`
	} `json:"includes"`
}
