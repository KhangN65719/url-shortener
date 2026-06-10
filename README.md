# URL Shortener

A simple URL shortening service written in Go. Submit a long URL and get back a short code. Visiting the short URL redirects to the original.

## Tech Stack

- Go 1.25
- SQLite
- Standard library HTTP server

## Requirements

- Go 1.25 or higher

## Running the Server

Build and start the server:

```bash
make run-server
```

The server runs on port 6767.

## Endpoints

### Shorten a URL

```
POST /shorten
```

Request body:

```json
{
  "url": "https://example.com"
}
```

Response:

```
localhost:6767/abc123
```

### Redirect

```
GET /{code}
```

Redirects to the original URL associated with the given short code. Returns 404 if the code does not exist.
