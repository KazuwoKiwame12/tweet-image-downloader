package utility

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tweet-image-downloader/entity"
)

type TwitterClient struct {
	token string
}

type Conditions struct {
	UserName string
	Keyword  string
	Max      int
}

var baseURL string = "https://api.twitter.com/2/tweets/search/recent"

func NewTwitterClient(token string) *TwitterClient {
	return &TwitterClient{
		token: token,
	}
}

func (t *TwitterClient) GetTweets(con Conditions) (*entity.TweetResponse, error) {
	// リクエストの作成
	req, err := t.createRequest(con)
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

func (t *TwitterClient) createRequest(con Conditions) (*http.Request, error) {
	// apiのurlパラメータの設定
	queryFields := entity.QueryFields{
		entity.QueryFieldHasImages,
		entity.QueryFieldIsRetweet.NOT(),
		entity.QueryFieldFrom(con.UserName),
	}
	if len(con.Keyword) != 0 {
		queryFields = append(queryFields, entity.QueryFieldKeyword(con.Keyword))
	}
	expansionFields := entity.ExpansionFields{
		entity.ExpasionFieldMediaKeys,
	}
	mediaFields := entity.MediaFields{
		entity.MediaFieldMediaKey,
		entity.MediaFieldURL,
	}
	tweetFields := entity.TweetFields{
		entity.TweetFieldCreatedAt,
		entity.TweetFieldAttachments,
	}

	params := NewParamBuilder().
		Query(queryFields).
		MaxResults(con.Max).
		Expansions(expansionFields).
		MediaFields(mediaFields).
		TweetFields(tweetFields).Build()

	// リクエストの作成
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", baseURL, params), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))
	return req, nil
}
