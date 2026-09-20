# go-hello-world-app

Minimal Go HTTP server on stdlib + `lib/pq` that connects to PostgreSQL, runs a schema migration on startup, and serves a health-check endpoint.

## Zerops service facts

- HTTP port: `8080`
- Siblings: `db` (PostgreSQL) — env: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`
- Runtime base: `go@1.22`

## Zerops dev

`setup: dev` idles on `zsc noop --silent`; the agent starts the dev server.

- Dev command: `go run .`
- In-container rebuild without deploy: `go build -o app .`

**All platform operations (start/stop/status/logs of the dev server, deploy, env / scaling / storage / domains) go through the Zerops development workflow via `zcp` MCP tools. Don't shell out to `zcli`.**

## Notes

- `HOME=/home/zerops` is set in dev runtime so Go can locate `GOPATH` (`/home/zerops/go`) and `GOCACHE` (`/home/zerops/.cache/go-build`).
- Migration runs once per deploy via `go run cmd/migrate/main.go` in `initCommands` — schema is ready on SSH.
- Prod build sets `CGO_ENABLED=0` for a fully static binary; `lib/pq` is pure Go so this is safe.
