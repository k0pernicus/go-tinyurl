# tinyurl
Simple REST API for tiny URLs - toy project to help people to learn the go programming language

This tutorial permits to learn go in an easy way, in creating a simple project: a REST API for tiny URLs.

# Clone

> git clone https://github.com/k0pernicus/go-tinyurl

# Build and run

A makefile is available at the root of the project.

To build the project, you just need to run `make build`.

To run the project, you can modify the `config.yaml` file available at the root project path, and run `make run`.

# Usage

You can create a tiny URL using:

```bash
curl -X POST HOST:PORT/create -H "Content-Type: application/json" -d '{"url":"https://lemonde.fr"}'
```

The service will return a JSON structure than contains some informations.
For URL creation, you will have this kind of structure:

```json
{"status_code":200,"response":{"id":"SRsidXIl","message":"OK"}}
```

The field `id` corresponds to the ID to append to the URL domain, like this: `http://HOST:PORT/SRsidXIl`.
If you use your browser, this will redirect you to the right website.

If you want to append a deadline to the tiny URL:

```bash
curl -X POST localhost:8000/create -H "Content-Type: application/json" -d '{"url":"https://lemonde.fr", "dead_in":"30s"}'
```

30 seconds after the creation, the link will not exists anymore.
