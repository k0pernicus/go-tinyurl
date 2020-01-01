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

## Creation

You can create a tiny URL using:

```bash
curl -X POST HOST:PORT/create -H "Content-Type: application/json" -d '{"url":"https://lemonde.fr"}'
```

The service will return a JSON structure than contains some informations.
For URL creation, you will have this kind of structure:

```json
{"status_code":200,"response":{"id":"SRsidXIl","message":"OK"}}
```

The field `id` corresponds of the ID to append to the URL domain.

If you want to append a deadline to the tiny URL (default: **no limit**), like 30 seconds:

```bash
curl -X POST HOST:PORT/create -H "Content-Type: application/json" -d '{"url":"https://lemonde.fr", "dead_in":"30s"}'
```

Also, you can generate a qr-code that points to the tiny URL (default: **false**), using:

```bash
curl -X POST HOST:PORT/create -H "Content-Type: application/json" -d '{"url":"https://lemonde.fr", "gen_qrcode":true}'
```

## Get

To reach the final URL, you can use the `get` route: `http://HOST:PORT/<ID>`.
Please to replace the `<ID>` with the one returned by the service previously (or scan you qr code).
If you use your browser, this will redirect you to the right website.

## Check if tiny URL still exists

To check if an ID exists, you can use the route `exists`:

```bash
curl -X GET HOST:PORT/exists/SRsidXIl
```
