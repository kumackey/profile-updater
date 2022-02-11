# Profile Updater

ZennやQiitaの記事、connpassのイベントのリストを取得し、プロフィールを更新します。

# Get Started

前提: [プロフィールの README を管理する - GitHub Docs](https://docs.github.com/ja/account-and-profile/setting-up-and-managing-your-github-profile/customizing-your-profile/managing-your-profile-readme)

`.github/workflows/`に以下のようなYAMLファイルを置きます。
`zenn_user_id`には自分のZennのユーザIDを入れてください。

```yaml:.github/workflows/profile.yml
name: profile updater

on: [ workflow_dispatch ]

jobs:
  profile-updater:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: kumackey/profile-updater@v1
        with:
          zenn_user_id: kumackey
      - name: Commit and push
        run: |
          git config --local user.name "GitHub Actions"
          git config --local user.email "action@github.com"
          git add .
          git commit -m "profile updated" || true
          git push origin main
```

`README.md`内でZennの記事のリストを置きたい箇所に、以下記述を入れてください。

```text:README.md
<!-- profile updater begin: zenn -->
<!-- profile updater end: zenn -->
```

該当アクションを手動実行することで、リストを取得し更新できます。

![](./res/readme_output.png)
![](./res/readme_result.png)

# How to use

## withのパラメータ

| パラメータ名                | 必須か | 説明                            |
|-----------------------|-----|-------------------------------|
| `zenn_user_id`        | NO  | ZennのユーザID                    |
| `zenn_max_articles`   | NO  | Zennの表示したい最大記事数。デフォルトは5       |
| `connpass_nickname`   | NO  | connpassのユーザ名                 |
| `connpass_max_events` | NO  | connpassの表示したい最大イベント数。デフォルトは5 |
| `qiita_user_id`       | NO  | QiitaのユーザID                   |
| `qiita_max_articles`  | NO  | Qiitaの表示したい最大記事数。デフォルトは5      |

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

## qiita

withのパラメータとして`qiita_user_id`を指定してください。 その上で、README.md内に以下記述を追加してください。

```text:README.md
<!-- profile updater begin: qiita -->
<!-- profile updater end: qiita -->
```

## replace statement not found って言われるんやけど

README.md内に`<!-- profile updater begin: ...`の記述がないと起こります。 zennおよびconnpassを参考に、`README.md`に追記してください。

## 定期更新の仕方(おすすめ)

Github Actionsではcron式でのアクション実行に対応しているので、以下のように書くことで1日1回のプロフィール更新をさせることができます。

```
on:
  schedule:
    - cron: '0 0 * * *'
  workflow_dispatch:
```
