genProtoc: 
	protoc --go_out=. --go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	**/proto/*.proto

genGql:
	cd graphql; go get github.com/99designs/gqlgen \
	&& go run github.com/99designs/gqlgen

upUser:
	docker compose --profile user up --build -d --remove-orphans 

upPet:
	docker compose  --profile pet up --build -d --remove-orphans

up:
	docker compose  --profile all up --build -d --remove-orphans

down:
	docker compose --profile all down