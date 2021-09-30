package utility

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tweet-image-downloader/entity"
)

type TwitterClient struct {
	token      string
	httpClient HttpClient
}

var baseURL string = "https://api.twitter.com/2/tweets/search/recent"

func NewTwitterClient(token string, hClient HttpClient) *TwitterClient {
	return &TwitterClient{
		token:      token,
		httpClient: hClient,
	}
}

func (t *TwitterClient) GetTweets(con Conditions) (*entity.TweetResponse, error) {
	// リクエストの作成
	req, err := t.createRequest(con)
	if err != nil {
		return nil, err
	}

	// APIエンドポイントにリクエストを投げる
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, entity.ErrorIsnotIdealStatusCode(resp.StatusCode)
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

	params, err := NewParamBuilder().
		Query(queryFields).
		MaxResults(con.Max).
		Expansions(expansionFields).
		MediaFields(mediaFields).
		TweetFields(tweetFields).Build()
	if err != nil {
		return nil, err
	}

	// リクエストの作成
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", baseURL, params), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))
	return req, nil
}

type Conditions struct {
	UserName string
	Keyword  string
	Max      int
}

func (c Conditions) ValidateMaxFieldValue() bool {
	if 10 <= c.Max && c.Max <= 100 {
		return true
	}
	return false
}
