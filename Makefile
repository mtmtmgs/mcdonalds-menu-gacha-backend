include .env

# 引数が指定されていない場合のデフォルト値は未定義
## .envと競合しない名前とする
MIKEFILE_ARG ?=

# おためし
sample:
	@echo "sample"

# WEBコンテナ接続
web-ssh-check: # テスト
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "uname -a"
web-ssh: # 接続
	docker exec -it ${WEB_CONTAINER_NAME} bash
web-launch: # WEBアプリケーション起動
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app && go run ."

# マイグレーション
db-migrate-status: # 状態
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/migrate && go run . db status"
db-migrate-create: # マイグレーションファイル生成
ifndef MIKEFILE_ARG
	@echo '第2引数の"MIKEFILE_ARG="にファイル名を入れてください. **例**: make db-migrate-create MIKEFILE_ARG=$${file_name}'
else
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/migrate && go run . db create_go ${MIKEFILE_ARG}"
endif
db-migrate-up: # 実行
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/migrate && go run . db migrate"
db-migrate-down: # ロールバック
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/db/migrate && go run . db rollback"

# バッチ実行
batch-exec:
	docker exec -it ${WEB_CONTAINER_NAME} sh -c "cd /app/batch && go run ."