# go user api

<br>

## Setup

`set-up`コマンドを実行すると`./env`ディレクトリに`.env`ファイルが作成されます。

自分で作成してもらっても問題ないです🙆

※ `.env`ファイルが存在しないとdocker環境が起動しません。

```
make set-up
```

<br>

コマンドが正常終了したのを確認したら、`up`コマンドでdocker環境を起動しましょう。

```
make up
```

<br>

## Directry

```
.
├── Makefile
├── README.md
├── cmd
│   └── api
│       └── server
│           └── main.go # ここからserverを起動します
├── docker
│   ├── dev # dev用のdockerファイルがまとめてあります
│   │   ├── api
│   │   └── mysql
│   │       └── initdb.d # dbの初期化時に実行するsqlです
│   └── prod # prod用のdockerファイルがまとめてあります
│       └── api
├── docker-compose.yml
├── go.mod
├── go.sum
├── pkg
│   ├── apperror # アプリケーション独自のエラーを定義してます
│   ├── config # init関数でサーバーの初期化時に環境変数を取得します
│   ├── db
│   │   └── mysql # mysqlに接続します
│   ├── domain # アプリケーションのdomainです
│   │   ├── model # アプリケーションのビジネスロジックがあります
│   │   └── repository # repositoryのインターフェースをここで定義します
│   ├── infra
│   │   ├── dao # repositoryの具体的な実装です
│   │   └── entity # mysqlとのやり取りをする際のオブジェクトです
│   ├── logger # アプリケーションで統一した形式のloggerを定義してます
│   ├── presenter
│   │   └── http
│   │       ├── adapter # handlerが定義してあります
│   │       │   └── user
│   │       │       ├── handler.go # userのhandlerが定義してあります。またここでrepositoryとusecaseの依存を注入しています
│   │       │       ├── request.go # request bodyのオブジェクトを定義してあります
│   │       │       ├── router.go # routerが定義してあります。ここでどのハンドラーとどのパスを紐づけるかを設定しています
│   │       │       └── view.go # usecaseから受け取ったオブジェクトをjson形式で扱えるように変更します
│   │       ├── middleware
│   │       ├── response # 共通で利用できるレスポンスの定義をしています
│   │       └── router # repositoryを各ハンドラに注入しつつパスの定義も行います
│   ├── registry # dbへの依存を注入したrepositoryのオブジェクトを作成します
│   └── usecase # アプリケーションのユースケースがまとめてあります
│       └── user
│           ├── input.go # presenterの層からusecaseを呼び出す際の引数に利用するinputの定義があります
│           ├── output.go # presenterの層からusecaseを呼び出したときの返り値に利用するoutputの定義があります
│           └── usecase.go # ユースケースが定義されています
├── scripts
```

<br>

## Tips

`make` or `make help`コマンドで`make`コマンドの一覧を確認できます。

```
help                 helpを表示する
up                   docker環境を立ち上げる
down                 docker環境を閉じる
re                   volumesを削除してdocker環境を立ち上げる
log                  docker環境のログを標示する
go-fmt               goのコードを整形します
```

<br>

