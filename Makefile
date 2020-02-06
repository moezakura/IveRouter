@:
	echo "no implements"

build:
	mkdir -p dist/
	cd server && go build -o ../dist/IveRouter