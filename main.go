package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var userNameF = flag.String("u", "kz_dev_", "input user name")
var keywordF = flag.String("k", "最新話です!!", "input keyword which target tweet has")
var max = flag.Int("m", 50, "input number of maximum tweets you want")

type Conditions struct {
	UserName string
	Keyword  string
	Max      int
}

func main() {
	/*
		1. コマンドのフラグでツイート検索条件を取得
		2. 条件より、目的の画像を持つツイートを取得
		3. ツイートから目的の画像のリンク一覧を取得
		4. 3のリンク一覧から、画像を自身のファイルに書き込む
	*/
	// 1. コマンドで条件を取得
	flag.Parse()
	con := Conditions{
		UserName: *userNameF,
		Keyword:  *keywordF,
		Max:      *max,
	}

	// 2. 条件より、目的の画像を持つツイートを取得
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: loading .env file")
	}
	client := NewClient(os.Getenv("BEARER_TOKEN"))
	res, err := client.GetTweets(con)
	if err != nil {
		log.Fatal(err)
	}
	// TODO
}
