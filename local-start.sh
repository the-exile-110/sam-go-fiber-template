echo  "local start server..."

cd app || exit

export REGION=ap-northeast-1

go run main.go