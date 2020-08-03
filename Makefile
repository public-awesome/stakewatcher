.PHONY: build test run pkger


generate-models:
	sh contrib/db/generate.sh

install-pkger:
	go install github.com/markbates/pkger/cmd/pkger

pkger:
	pkger -o cmd/stakewatcher
	
build: pkger
	go build -o build/stakewatcher github.com/public-awesome/stakewatcher/cmd/stakewatcher

build-linux: pkger
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/stakewatcher github.com/public-awesome/stakewatcher/cmd/stakewatcher

build-docker:
	docker build -t publicawesome/stakewatcher .

migrate: pkger
	go run github.com/public-awesome/stakewatcher/cmd/stakewatcher migrate
run:
	go run github.com/public-awesome/stakewatcher/cmd/stakewatcher

start:
	./build/stakewatcher
	
fake-post:
	stakecli tx curating post  1 $(POST_ID) "post body"  --from validator --keyring-backend test --trust-node --chain-id $(shell stakecli status | jq '.node_info.network') -b block -y

	
fake-upvote:
	stakecli tx curating upvote 1 $(POST_ID) 1  --from validator --keyring-backend test --trust-node --chain-id $(shell stakecli status | jq '.node_info.network') -b block -y


install-tools: install-sqlboiler install-sqlmigrate


install-sqlboiler:
	VERSION=v4.2.0 ./contrib/dev/install-sqlboiler.sh

install-sqlmigrate:
	go install github.com/rubenv/sql-migrate/sql-migrate
