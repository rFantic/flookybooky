genProtoc:
	protoc --go_out=pb --go-grpc_out=pb \
	--go_opt=paths=import \
	--go-grpc_opt=paths=import \
	**/**/proto/*.proto

genGql:
	cd services/graphql; go get github.com/99designs/gqlgen@v0.17.31 \
	&& go run github.com/99designs/gqlgen

upProfile:
	docker compose --profile $(Profile) up -d --build --remove-orphans

up:
	docker compose --profile graphql up -d --build --remove-orphans

down:
	docker compose --profile graphql down
