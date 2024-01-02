include .env

# 引数が指定されていない場合のデフォルト値は未定義
## .envと競合しない名前とする
MIKEFILE_ARG ?=

# おためし
sample:
	@echo "sample"

# WEBコンテナ
## sh接続テスト
web-sh-check:
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "uname -a"
## bash接続
web-sh:
	docker exec -it ${WEB_CONTAINER_NAME} bash
## WEBアプリケーション起動
web-launch:
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app && go run ."

# マイグレーション
## 初期化（マイグレーション管理テーブル作成）
db-migrate-init:
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/migrate && go run . db init"
## ステータス確認
db-migrate-status:
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/migrate && go run . db status"
## マイグレーションファイル生成
db-migrate-create:
ifndef MIKEFILE_ARG
	@echo '第2引数の"MIKEFILE_ARG="にマイグレーションファイル名を入れてください. **例**: make db-migrate-create MIKEFILE_ARG=$${file_name}'
else
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/migrate && go run . db create_go ${MIKEFILE_ARG}"
endif
## 実行
db-migrate-up:
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/migrate && go run . db migrate"
## ロールバック
db-migrate-down:
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/migrate && go run . db rollback"

# seedデータ（menus）
db-seed:
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/seed && go run ."

# バッチ実行
batch-exec:
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/batch && go run ."