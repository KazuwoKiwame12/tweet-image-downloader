package entity

import (
	"fmt"
	"time"
)

// query parametersの"expasions"に関する設定
type ExpansionField string

type ExpansionFields []ExpansionField

const (
	ExpasionFieldMediaKeys ExpansionField = "attachments.media_keys"
)

func (e ExpansionField) toString() string {
	return string(e)
}

func (es ExpansionFields) ToStringSlice() []string {
	slice := make([]string, len(es))
	for i, e := range es {
		slice[i] = e.toString()
	}
	return slice
}

// query parametersの"media.fields"に関する設定
type MediaField string

type MediaFields []MediaField

const (
	MediaFieldMediaKey MediaField = "media_key"
	MediaFieldURL      MediaField = "url"
)

func (m MediaField) toString() string {
	return string(m)
}

func (ms MediaFields) ToStringSlice() []string {
	slice := make([]string, len(ms))
	for i, m := range ms {
		slice[i] = m.toString()
	}
	return slice
}

// query parametersの"query"に関する設定
type QueryField string

type QueryFields []QueryField

const (
	QueryFieldHasImages QueryField = "has:images"
	QueryFieldIsRetweet QueryField = "is:retweet"
)

func (q QueryField) toString() string {
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

func (qs QueryFields) ToStringSlice() []string {
	slice := make([]string, len(qs))
	for i, q := range qs {
		slice[i] = q.toString()
	}
	return slice
}

// query parametersの"tweet.fields"に関する設定
type TweetField string

type TweetFields []TweetField

const (
	TweetFieldAttachments TweetField = "attachments"
	TweetFieldCreatedAt   TweetField = "created_at"
)

func (t TweetField) toString() string {
	return string(t)
}

func (ts TweetFields) ToStringSlice() []string {
	slice := make([]string, len(ts))
	for i, t := range ts {
		slice[i] = t.toString()
	}
	return slice
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
