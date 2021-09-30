package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tweet-image-downloader/entity"
)

type Client struct {
	token string
}

var baseURL string = "https://api.twitter.com/2/tweets/search/recent"

func NewClient(token string) *Client {
	return &Client{
		token: token,
	}
}

func (c *Client) GetTweets(con Conditions) (*entity.TweetResponse, error) {
	// リクエストの作成
	req, err := c.createRequest(con)
	if err != nil {
		return nil, err
	}

	// APIエンドポイントにリクエストを投げる
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code is not 200. result is %d", resp.StatusCode)
	}

	// レスポンスの読み取り
	defer resp.Body.Close()
	jsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(entity.TweetResponse)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) createRequest(con Conditions) (*http.Request, error) {
	// apiのurlパラメータの設定
	queryFields := []entity.QueryField{
		entity.QueryFieldHasImages,
		entity.QueryFieldIsRetweet.NOT(),
		entity.QueryFieldFrom(con.UserName),
		entity.QueryFieldKeyword(con.Keyword),
	}
	expansionFields := []entity.ExpansionField{
		entity.ExpasionFieldMediaKeys,
	}
	mediaFields := []entity.MediaField{
		entity.MediaFieldMediaKey,
		entity.MediaFieldURL,
	}

	params := entity.NewParamBuilder().
		Query(queryFields).
		MaxResults(con.Max).
		Expansions(expansionFields).
		MediaFields(mediaFields).String()

	// リクエストの作成
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", baseURL, params), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	return req, nil
}
