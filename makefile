dev:
	docker compose -f dev.docker-compose.yml up --build --force-recreate --renew-anon-volumes
proto:
	mkdir -p mkdir -p ./api/google/api; \
    curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > ./api/google/api/annotations.proto; \
    curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > ./api/google/api/http.proto; \
	protoc -I ./api \
      --go_out ./api --go_opt paths=source_relative \
      --go-grpc_out ./api --go-grpc_opt paths=source_relative \
      --grpc-gateway_out ./api --grpc-gateway_opt paths=source_relative \
      api/chgk-telegram.proto;

