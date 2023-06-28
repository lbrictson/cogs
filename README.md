# cogs

A simple server for executing adhoc scripts or running scripts on a schedule, complete with audit trails and a web interface.

## Running cogs locally

```bash
go run cmd/server/main.go -data=tmp -port=8080
```

## Adding Database Schemas

```bash
go run -mod=mod entgo.io/ent/cmd/ent new NameOfObject
```

Generate the updated models

```bash
go generate ./ent
```