# Sonic

## Documentation/Help

### .env
- Store sensitive vars in a .env file to be loaded by make/go
  - e.g. POSTGRES_USER, POSTGRES_PASSWORD, etc.

### golang-migrate
- Install:
  - `brew install golang-migrate`
- Create migration files:
  - `migrate create -ext sql -dir db/migrations -seq TABLE_NAME`
- Perform a migration:
  - `migrate -database $(DB_URL) -path db/migrations up`
- Force migration:
  - `migrate -database $(DB_URL) -path db/migrations force ${MIGRATION_NUM}`
