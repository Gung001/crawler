### 运行模拟详情网站
go run mockserver/main.go

### 启动ES
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.4.2

### GO安装ES
go get github.com/olivere/elastic/v7
