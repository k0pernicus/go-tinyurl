prepare:
	cp config.yaml app/
	docker build -t go-tinyurl-db -f sqlite-Dockerfile .

db:
	docker run --rm -it -v `pwd`/db:/db go-tinyurl-db urls.db

backup:
	docker run --rm -it -v `pwd`/db:/db go-tinyurl-db urls.db .dump >> urls_dump.sql

build:
	cd app; go build -o tinyurl

run: build
	cd app; ./tinyurl

clean:
	cp app; go clean
	cp app; rm config.yaml; rm tinyurl
