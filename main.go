package main

import (
	"flag"
	"log"
	"os"

	"tweet-image-downloader/entity"
	"tweet-image-downloader/utility"

	"github.com/joho/godotenv"
)

var userNameF = flag.String("u", "_kz_dev", "input user name")
var keywordF = flag.String("k", "最新話です!", "input keyword which target tweet has")
var maxF = flag.Int("m", 50, "input number of maximum tweets you want")

func main() {
	/*
		1. コマンドのフラグでツイート検索条件を取得
		2. 条件より、目的の画像を持つツイートを取得
		3. ツイートから目的の画像のリンク一覧を取得
		4. 3のリンク一覧から、画像を自身のファイルに書き込む
	*/
	// 1. コマンドで条件を取得
	flag.Parse()
	con := utility.Conditions{
		UserName: *userNameF,
		Keyword:  *keywordF,
		Max:      *maxF,
	}

	// 2. 条件より、目的の画像を持つツイートを取得
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: loading .env file")
	}
	client := utility.NewTwitterClient(os.Getenv("BEARER_TOKEN"))
	res, err := client.GetTweets(con)
	if err != nil {
		log.Fatal(err)
	}
	if len(res.Tweets) == 0 {
		log.Fatal("you didn't get any tweet. check url parameters")
	}
	// 3. 最新ツイートから目的の画像のリンク一覧を取得
	links := make([]string, 0, 4)
	tweet := latestTweet(res.Tweets)
	for _, key := range tweet.Attachments.MediaKeys {
		for _, media := range res.Includes.Media {
			if key == media.MediaKey {
				links = append(links, media.URL)
			}
		}
	}
}

func latestTweet(tweets []entity.Tweet) entity.Tweet {
	var latestTweet entity.Tweet = tweets[0]
	for _, tweet := range tweets {
		if latestTweet.CreatedAt.After(tweet.CreatedAt) {
			latestTweet = tweet
		}
	}
	return latestTweet
}
