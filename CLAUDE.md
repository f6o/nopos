# CLAUDE.md

## Project Overview

**napos** is a Go-based gRPC service for card dealing and game/user management. It includes a gRPC server (Dealer and GameManager services), a CLI client, and a SQLite database via Prisma ORM for persisting users and hand histories.

Module path: `github.com/f6o/napos`

## Repository Structure

```
main.go              # Server entry point (starts gRPC on port 3333)
client/client.go     # CLI client (subcommands: print-freq, add-user, list-users)
server/server.go     # gRPC service implementations (Dealer, GameManager)
hands/               # Protobuf definitions and generated Go code
  hands.proto        # Service and message definitions
  hands.pb.go        # Generated (do not edit)
  hands_grpc.pb.go   # Generated (do not edit)
db/wrapper.go        # Prisma database wrapper
schema.prisma        # Prisma schema (SQLite, User + HandHistory models)
Makefile             # Build targets for proto and prisma
```

## Build & Run Commands

```bash
# Run the server
go run main.go

# Run the client
go run client/client.go <subcommand>

# Regenerate protobuf code after editing hands/hands.proto
make proto

# Push Prisma schema changes to SQLite database
make prisma
```

### Prerequisites

- Go 1.22+
- `protoc` with `protoc-gen-go` and `protoc-gen-go-grpc` plugins (for proto generation)
- Prisma Client Go (`github.com/steebchen/prisma-client-go`)

## Key Dependencies

| Package | Purpose |
|---------|---------|
| `google.golang.org/grpc` | gRPC framework |
| `google.golang.org/protobuf` | Protocol Buffers |
| `github.com/steebchen/prisma-client-go` | Prisma ORM (SQLite) |
| `github.com/joho/godotenv` | .env file loading |
| `golang.org/x/exp` | Experimental stdlib (random seeding) |

## Architecture

- **Server** listens on port `3333` (hardcoded in `main.go`)
- **Dealer service**: `DealCard` (seeded) and `DealRandomCard` RPCs
- **GameManager service**: `AddUser` and `ListUsers` RPCs
- **Database**: SQLite (`dev.db`), schema in `schema.prisma`
- **Generated code**: `hands/*.pb.go` and `db/*_gen.go` are generated — do not edit manually

## Code Conventions

- Standard Go project layout with package-per-directory
- Error handling via `fmt.Errorf` wrapping
- Context passed through for database operations
- No linter configuration; follow standard `gofmt` formatting
- No tests exist yet — use Go standard `*_test.go` convention with `testing` package

## Things to Know

- The `db/` directory has generated Prisma files ignored via `.gitignore` (`*_gen.go`)
- After changing `schema.prisma`, run `make prisma` to regenerate the client and push schema
- After changing `hands/hands.proto`, run `make proto` to regenerate gRPC stubs
- The SQLite database file `dev.db` is local and not committed
