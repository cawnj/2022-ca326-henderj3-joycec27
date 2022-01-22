# Sonic

## Documentation/Help

### .env
- Store sensitive vars in a .env file to be loaded by make/go
  - e.g. POSTGRES_USER, POSTGRES_PASSWORD, etc.

### golang-migrate
- Install with:
  - `brew install golang-migrate`
- Create migrations with:
  - `migrate create -ext sql -dir db/migrations -seq TABLE_NAME`
