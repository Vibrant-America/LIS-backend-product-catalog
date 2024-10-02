docker run -d --name jaeger -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:1.31.0
#   -p 5775:5775/udp \
# agent port

#   -p 6832:6832/udp \
#   -p 5778:5778 \
# serve frontend
#   -p 14250:14250 \
#   -p 14268:14268 \
#   -p 14269:14269 \
#   -p 9411:9411 \

# need -p 8600:8600/udp for dns server. Also 8301 9302 tcp and udp for LAN and WAN server communicatinos, 8300 tcp for gRPC.
docker run -d --name local_consul -p 8500:8500 -p 8600:8600/udp -v consul_data:/CONSUL/DATA consul

#mysql
docker run --name local-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=abc123456 -d mysql:8.0.31

#redis
docker run -d --name local-reids -p 6379:6379 redis:7.0.5-alpine3.16