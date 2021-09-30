# tweet-image-downloader
このプログラムは推しの漫画家の最新話をGoとTwitterAPIを利用してダウンロードするプログラムです。ぱっとお試し程度に作りましたので問題点などがあれば、Pull Requestをぜひ投げてください!
- ※**コードによるテスト**は最低限の内容しか行っていないので、実際に利用になる際にはご自身でも動作確認をお願いします。(Pull Requestでテストを送ってくださると助かります)
```
- ParameterBuilderの`validateEmptyParamValue`
- httpリクエストやレスポンスをmock化したTwitterClientの`GetTweets`
- TwitterClientの`createResponse`
```
# How to use
1. create `.env` file.
```terminal
$ cp .env.exmaple .env
```
2. set your `Bearer Token` into `.env`
```.env
BEARER_TOKEN="your Bearer Token"
```
3. execute command with some options
```terminal
<!-- オプションの確認 -->
$ go run main.go -h
~~
  -k string
        input keyword which target tweet has (default "最新話です!")
  -m int
        input number of maximum tweets you want between 10 and 100 (default 50)
  -u string
        input user name (default "_kz_dev")
<!-- 好きなようにオプションをつけて実行 -->
$ go run main.go -k=hogehoge -m=100 -u=hogehoge
```