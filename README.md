# go user api

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

# 疑問点
- srcディレクトリにusecaseを切り分けた理由
- 値わたししている理由
- なぜusecaseのinterfaceがないのか
  - あったらテストしやすいのでは？