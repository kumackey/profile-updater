# Profile Updater

ZennやQiitaの記事、connpassのイベントのリストを取得し、プロフィールを更新します。

# Get Started



# How to use

## withのパラメータ

| パラメータ名                | 必須か | 説明                            |
|-----------------------|-----|-------------------------------|
| `qiita_user_id`       | NO  | QiitaのユーザID                   |
| `qiita_max_articles`  | NO  | Qiitaの表示したい最大記事数。デフォルトは5      |
| `zenn_user_id`        | NO  | ZennのユーザID                    |
| `zenn_max_articles`   | NO  | Zennの表示したい最大記事数。デフォルトは5       |
| `connpass_nickname`   | NO  | connpassのユーザ名                 |
| `connpass_max_events` | NO  | connpassの表示したい最大イベント数。デフォルトは5 |

## qiita

withのパラメータとして`qiita_user_id`を指定してください。 その上で、README.md内に以下記述を追加してください。

```text:README.md
<!-- profile updater begin: qiita -->
<!-- profile updater end: qiita -->
```

## zenn

withのパラメータとして`zenn_user_id`を指定してください。 その上で、README.md内に以下記述を追加してください。

```text:README.md
<!-- profile updater begin: zenn -->
<!-- profile updater end: zenn -->
```

## connpass

withのパラメータとして`connpass_nickname`を指定してください。 その上で、`README.md`内に以下記述を追加してください。

```text:README.md
<!-- profile updater begin: connpass -->
<!-- profile updater end: connpass -->
```

## 定期更新の仕方(おすすめ)

Github Actionsではcron式でのアクション実行に対応しているので、以下のように書くことで1日1回のプロフィール更新をさせることができます。

```
on:
  schedule:
    - cron: '0 0 * * *'
  workflow_dispatch:
```

## replace statement not found って言われるんやけど

README.md内に`<!-- profile updater begin: ...`の記述がないと起こります。 各説明を参考に、`README.md`に追記してください。

## この機能くれ！

issueかプルリクエストかSNSでのメッセージをお待ちしてます。

## おい、バグってんぞ！

issueかプルリクエストかSNSでのメッセージをお待ちしてます。