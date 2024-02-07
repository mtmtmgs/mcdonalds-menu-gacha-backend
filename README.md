# mcdonalds-menu-gacha-backend

某ハンバーガー店のメニューを予算内でランダムにリストで返却

URL: https://hm-mtmtmgs.net/v1/menu-gacha?budget=1000

## 目次

- [mcdonalds-menu-gacha-backend](#mcdonalds-menu-gacha-backend)
  - [目次](#目次)
  - [システム構成図](#システム構成図)
  - [技術構成](#技術構成)
  - [ディレクトリ構造](#ディレクトリ構造)
  - [API](#api)
    - [ユーザ](#ユーザ)
    - [メニュー](#メニュー)
  - [環境構築](#環境構築)
  - [ブランチ運用](#ブランチ運用)

## システム構成図

![menu-gacha-system-architecture drawio](https://github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/assets/150935913/e7fe37ac-1d86-4483-8950-493dff68eaf7)

## 技術構成

| カテゴリ | 技術スタック                                                            |
| -------- | ----------------------------------------------------------------------- |
| 言語     | Go, [echo](https://github.com/labstack/echo)                            |
| DB       | PostgreSQL                                                              |
| ORM      | [bun](https://github.com/uptrace/bun)                                   |
| インフラ | Cloud Build, Cloud Run, Cloud SQL                                       |
| CI/CD    | GitHub Actions                                                          |
| テスト   | ~~go test~~, [golangci-lint](https://github.com/golangci/golangci-lint) |
| 開発環境 | Docker                                                                  |
| その他   | JWT 認証, [validator](https://github.com/go-playground/validator)       |

## ディレクトリ構造

```sh
mcdonalds-menu-gacha-backend/
    |-- src/
    |    |-- batch/
    |    |    |-- scripts/ # バッチスクリプト群
    |    |    |-- batch.go # バッチのエントリ
    |    |
    |    |-- consts/ # アプリケーション定数群
    |    |-- controllers/ # コントローラー群, ハンドラ
    |    |    |-- requests/ # usecaseの引数の型, DTO, バリデーション
    |    |    |-- responses/ # usecaseの返り値の型, DTO
    |    |
    |    |-- db/
    |    |    |-- migrate/
    |    |    |    |-- migrations/ # マイグレーションファイル群
    |    |    |    |-- migrate.go # マイグレーションのエントリ
    |    |    |
    |    |    |-- seed/
    |    |    |    |-- csv/ # シードデータ
    |    |    |    |-- seed.go # シードデータのエントリ
    |    |    |
    |    |    |-- db.go # DBクライアント生成
    |    |
    |    |-- domains
    |    |    |-- entities/ # エンティティ群, 主にインスタンス生成
    |    |    |-- models/ # モデル群, 主にDBカラム定義
    |    |    |-- services/ # ドメインサービス群, ビジネスロジック
    |    |    |-- values/ # 値オブジェクト群, ビジネスロジック
    |    |
    |    |-- env/env.go # devとprodの設定
    |    |-- repositories/ # リポジトリ群, DBデータ操作
    |    |-- router/
    |    |    |-- middleware/ # ミドルウェア群 jwt, validator, etc.
    |    |    |-- router.go # ルーティング
    |    |
    |    |-- usecases/ # ユースケース群, コントローラーから呼び出す
    |    |-- utils/ # ユーティリティ群 hash, time, etc.
    |    |-- .dockerignore
    |    |-- .env.sample # アプリケーションの.env
    |    |-- Dockerfile
    |    |-- go.mod
    |    |-- go.sum
    |    |-- main.go # アプリケーションのエントリ
    |
    |-- .github/workflows/main.yml # lint, テスト, デプロイ
    |-- .env.sample # docker-compose.ymlの.env
    |-- .gitignore
    |-- Makefile # 開発用のコマンドリスト
    |-- README.md
    |-- docker-compose.yml
```

## API

### ユーザ

|    機能    | メソッド | URI        | JWT 認証 | 備考             |
| :--------: | :------: | ---------- | :------: | ---------------- |
|  会員登録  |   POST   | /v1/signup |          | メールは飛ばない |
|  ログイン  |   POST   | /v1/login  |          |                  |
| ユーザ情報 |   GET    | /v1/user   |    有    |                  |

### メニュー

|      機能      | メソッド | URI            | JWT 認証 | 備考                                 |
| :------------: | :------: | -------------- | :------: | ------------------------------------ |
| メニューガチャ |   GET    | /v1/menu-gacha |          | ?budget=1500 予算、デフォルトは 1000 |
| メニューリスト |   GET    | /v1/menus      |    有    |                                      |

## 環境構築

1. 環境変数

```txt
.env.sample をコピーして.env を追加
※WEBアプリとDBのタイムゾーンはUTC
```

2. docker-compose 起動

```sh
docker compose up -d
```

3. マイグレーション管理テーブル作成

```sh
make db-migrate-init
```

4. マイグレーション実行

```sh
make db-migrate-up
```

1. ~~バッチ実行、メニューデータを取得して DB 登録~~

```sh
make batch-exec
```

6. シードデータ DB 登録、メニューデータのシード、バッチではなく基本はこちらを使う

```sh
make db-seed
```

7. WEB アプリ起動

```sh
make web-launch
```

## ブランチ運用

GitLab flow にならう

| ブランチ名 | 役割             | 派生元 | マージ先 |
| ---------- | ---------------- | ------ | -------- |
| main       | 公開可能なもの   |        | prod     |
| prod       | リリースするもの |        |          |
| feature/\* | 新規開発のもの   | main   | main     |
| fix/\*     | バグ修正のもの   | main   | main     |

※ prod ブランチ push 時に GitHub Actions で検知し本番へ公開
