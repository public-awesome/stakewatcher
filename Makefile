.PHONY: build test run pkger


generate-models:
	sqlboiler psql  --struct-tag-casing camel  --wipe

install-pkger:
	go install github.com/markbates/pkger/cmd/pkger

pkger:
	pkger -o cmd/stakewatcher
	
build: pkger
	go build -o build/stakewatcher github.com/public-awesome/stakewatcher/cmd/stakewatcher

migrate: pkger
	go run github.com/public-awesome/stakewatcher/cmd/stakewatcher migrate
run:
	go run github.com/public-awesome/stakewatcher/cmd/stakewatcher

start:
	./build/stakewatcher
	
fake-post:
	stakecli tx stake post stakevaloper1vd333g9t9zemut2nwu39p9hxdelkje7t9ql0tp 1000stake 1 3 "body" 72h  --from validator --keyring-backend test --chain-id $(shell stakecli status | jq '.node_info.network') -b block -y
