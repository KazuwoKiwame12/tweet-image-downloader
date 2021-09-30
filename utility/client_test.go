package utility_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
	"tweet-image-downloader/entity"
	mock "tweet-image-downloader/mock/utility"
	"tweet-image-downloader/utility"

	"github.com/golang/mock/gomock"
)

func Test_GetTweets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type input struct {
		con   utility.Conditions
		setup func() *mock.MockHttpClient
	}

	correctCreatedAt, _ := time.Parse(time.RFC3339Nano, "2021-03-19T19:59:10.000Z")
	tests := []struct {
		name   string
		input  input
		want   *entity.TweetResponse
		hasErr bool
	}{
		{
			name: "failed to send request",
			input: input{
				con: utility.Conditions{
					UserName: "sample",
					Keyword:  "sample",
					Max:      50,
				},
				setup: func() *mock.MockHttpClient {
					m := mock.NewMockHttpClient(ctrl)
					m.EXPECT().Do(gomock.Any()).Return(
						nil,
						errors.New("error: failed to send request"),
					)
					return m
				},
			},
			want:   nil,
			hasErr: true,
		},
		{
			name: "failed with StatusCode",
			input: input{
				con: utility.Conditions{
					UserName: "sample",
					Keyword:  "sample",
					Max:      50,
				},
				setup: func() *mock.MockHttpClient {
					m := mock.NewMockHttpClient(ctrl)
					m.EXPECT().Do(gomock.Any()).Return(
						&http.Response{
							StatusCode: http.StatusBadRequest,
						},
						errors.New("error: status code is not 200"),
					)
					return m
				},
			},
			want:   nil,
			hasErr: true,
		},
		{
			name: "failed when json unmarshal",
			input: input{
				con: utility.Conditions{
					UserName: "sample",
					Keyword:  "sample",
					Max:      50,
				},
				setup: func() *mock.MockHttpClient {
					body := `{
						"sample": "sample",
					}`
					m := mock.NewMockHttpClient(ctrl)
					m.EXPECT().Do(gomock.Any()).Return(
						&http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(body)),
						},
						nil,
					)
					return m
				},
			},
			want:   nil,
			hasErr: true,
		},
		{
			name: "success",
			input: input{
				con: utility.Conditions{
					UserName: "sample",
					Keyword:  "sample",
					Max:      50,
				},
				setup: func() *mock.MockHttpClient {
					body := `{
						"data": [
							{
								"id": "1",
								"text": "sample",
								"created_at": "2021-03-19T19:59:10.000Z",
								"attachments": {
									"media_keys": ["100", "200"]
								}
							}
						],
						"includes": {
							"media": [
								{
									"media_key": "100",
									"url": "http://sample.com/sample1.html"
								},
								{
									"media_key": "200",
									"url": "http://sample.com/sample2.html"
								}
							]
						}
					}`
					m := mock.NewMockHttpClient(ctrl)
					m.EXPECT().Do(gomock.Any()).Return(
						&http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(body)),
						},
						nil,
					)
					return m
				},
			},
			want: &entity.TweetResponse{
				Tweets: []entity.Tweet{
					{
						ID:        "1",
						Text:      "sample",
						CreatedAt: correctCreatedAt,
						Attachments: struct {
							MediaKeys []string "json:\"media_keys\""
						}{
							MediaKeys: []string{"100", "200"},
						},
					},
				},
				Includes: struct {
					Media []struct {
						MediaKey string "json:\"media_key\""
						URL      string "json:\"url\""
					} "json:\"media\""
				}{
					Media: []struct {
						MediaKey string "json:\"media_key\""
						URL      string "json:\"url\""
					}{
						{
							MediaKey: "100",
							URL:      "http://sample.com/sample1.html",
						},
						{
							MediaKey: "200",
							URL:      "http://sample.com/sample2.html",
						},
					},
				},
			},
			hasErr: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, test.name), func(t *testing.T) {
			c := utility.NewTwitterClient("sample", test.input.setup())
			tweets, err := c.GetTweets(test.input.con)

			if (err != nil) != test.hasErr {
				t.Errorf("failed to match for error: result is %v, hasErr is %v, err content is %v", (err != nil), test.hasErr, err)
			}
			if !isSameContent(tweets, test.want, t) {
				t.Errorf("failed to match for response: result is %v, want is %v", tweets, test.want)
			}
		})
	}
}

func isSameContent(a, b *entity.TweetResponse, t *testing.T) bool {
	// nilの場合
	if a == nil && b == nil {
		return true
	}
	if a != nil && b == nil {
		return false
	}
	if a == nil && b != nil {
		return false
	}

	// 中身の比較
	if len(a.Tweets) != len(b.Tweets) {
		t.Logf("number of tweets is not matched: result is %v, want is %v", len(a.Tweets), len(b.Tweets))
		return false
	}
	if !reflect.DeepEqual(a.Includes, b.Includes) {
		t.Logf("includes is not matched: result is %v, want is %v", a.Includes, b.Includes)
		return false
	}
	for i := 0; i < len(a.Tweets); i++ {
		if a.Tweets[i].ID != b.Tweets[i].ID {
			t.Logf("tweet id is not matched: result is %v, want is %v", a.Tweets[i].ID, b.Tweets[i].ID)
			return false
		}
		if a.Tweets[i].Text != b.Tweets[i].Text {
			t.Logf("tweet text is not matched: result is %v, want is %v", a.Tweets[i].Text, b.Tweets[i].Text)
			return false
		}
		if !a.Tweets[i].CreatedAt.Equal(b.Tweets[i].CreatedAt) {
			t.Logf("tweet created_at is not matched: result is %v, want is %v", a.Tweets[i].CreatedAt, b.Tweets[i].CreatedAt)
			return false
		}
		if !reflect.DeepEqual(a.Tweets[i].Attachments, b.Tweets[i].Attachments) {
			t.Logf("tweet attatchments is not matched: result is %v, want is %v", a.Tweets[i].Attachments, b.Tweets[i].Attachments)
			return false
		}
	}

	return true
}
