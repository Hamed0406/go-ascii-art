# Contributing & Project Roles

## Code Style
- Keep `cmd/main.go` very small (entrypoint only).
- All features must go into `internal/` packages.
- Use **clear comments** for each function (GoDoc style).
- Functions must be reusable (no hard-coded values unless default).

## Project Rules
1. `main.go` should only call package functions, never hold heavy logic.
2. Core image-to-ASCII logic lives in `internal/ascii/`.
3. Utility functions live in `internal/utils/`.
4. Every new feature should have at least **one test** in `/tests`.
5. Use `testdata/` for input images used in tests.

## Testing
- Run tests with:
  ```bash
  go test ./tests/...
