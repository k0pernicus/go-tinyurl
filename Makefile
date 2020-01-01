prepare:
	cp config.yaml app/

build:
	cd app; go build -o tinyurl

run: prepare build
	cd app; ./tinyurl

clean:
	cp app; go clean
	cp app; rm config.yaml; rm tinyurl
