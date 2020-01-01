build:
	cp config.yaml app/; cd app; go build -o tinyurl

run: build
	cd app; ./tinyurl