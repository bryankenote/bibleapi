# BibleApi

### Install Depenedencies

- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
- [connectrpc go](https://connectrpc.com/docs/go/getting-started)

### Run

```
make run
```

## Development

### DB

After making changes to the schema or adding or changing new queries, repository methods can be generated using `make sqlc`.

### API Interface

After making changes to the protobuf file, the service interface and clients can be generated using `make pb`.
