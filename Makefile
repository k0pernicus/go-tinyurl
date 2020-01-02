prepare:
	cp config.yaml app/
	docker build -t go-tinyurl-db -f sqlite-Dockerfile .

run-db:
	docker run --rm -v `pwd`/db:/db go-tinyurl-db urls.db

backup-db:
	docker run --rm -it -v `pwd`/db:/db go-tinyurl-db urls.db .dump >> urls_dump_$(shell date -u +"%Y-%m-%dT%H:%M:%SZ").sql

build:
	cd app; go build -o tinyurl

run: run-db build
	cd app; ./tinyurl

clean:
	cp app; go clean
	cp app; rm config.yaml; rm tinyurl
	rm -rf db/
