# credit_certificate

micro service and blockchain

## goctl命令

生成api文件
goctl api go -api *.api -dir ../ --style=goZero

一键生成Dockerfile
goctl docker -go student.go

一键生成k8s部署文件
goctl kube deploy -name course-api -namespace credit-certifitate -image course-api:v1.0 -o course-api.yaml -port 1001

生成rpc代码
goctl rpc
