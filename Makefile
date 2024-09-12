vault-file=passkey

proto-gen:
	@rm -rf pb/*
	@protoc --proto_path=proto --go_out=proto/gen --go_opt=paths=source_relative \
	--go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
	proto/*.proto

husky:
	@alias husky="go run \"github.com/automation-co/husky\""

account:
	@air --build.cmd "go build -o bin/account cmd/account/main.go" --build.bin "./bin/account"

set-alias:
	@alias vault="docker run -v \"$PWD\":/workdir ghcr.io/goodwaygroup/gwvault"

encrypt:
	@docker run -v "$PWD":/workdir ghcr.io/goodwaygroup/gwvault encrypt --vault-password-file /workdir/${vault-file} main.go

decrypt:
	@docker run -v "$PWD":/workdir ghcr.io/goodwaygroup/gwvault decrypt --vault-password-file ${vault-file} .env

### DATABASE

migration_dir=db/migration

migration-new:
	@read -p "Migration name: " migration_name;
	@goose -dir=${migration_dir} create "$migration_name" sql