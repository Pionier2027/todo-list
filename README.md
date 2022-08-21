# todo-list
TODOリストアプリ

## 実行手順

```
# フォルダの作成
mkdir -p mysql/data
mkdir -p phpmyadmin/sessions

# ビルド
docker-compose up -d --build
# 起動
docker exec -it server go run main.go
# データ確認
http://localhost:4040
# 終了
docker-compose down
```
