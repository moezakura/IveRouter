@:
	echo "no implements"

build:
	mkdir -p dist/
	make api-generate
	cd server && go build -o ../dist/IveRouter

api-generate:
	openapi-generator generate -i api/packet-traffic-service.v1.yaml -g typescript-axios -o web/api
	openapi-generator generate -i api/packet-traffic-service.v1.yaml -g go-server --git-user-id moezakura --git-repo-id IveRouter/server/api -o server/api