genProtoc: 
	protoc --go_out=. --go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	**/proto/*.proto

genGql:
	cd services/graphql; go get github.com/99designs/gqlgen@v0.17.31 \
	&& go run github.com/99designs/gqlgen

upProfile:
	docker compose --profile $(Profile) up --build -d --remove-orphans 

up:
	docker compose --profile graphql up --build -d --remove-orphans

down:
	docker compose --profile graphql down