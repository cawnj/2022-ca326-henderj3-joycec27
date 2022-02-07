# Sonic Server

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
- **NOTE**: Have now switched to using the migrate/migrate docker image instead
  ```
  docker run \
	-v `pwd`/db/migrations:/migrations \
	--network host migrate/migrate \
	-database $(DB_URL) \
	-path /migrations \
	up
  ```

### Creating endpoints
- Add router to `handler.go`
- Create function in `handler` folder
- Create database query function in `db` folder
  - This also requires a valid model/struct to match the query return content

### Keeping master and external branches in check
- `git pull` in master branch
- `git checkout external`
- `git rebase master` and fix and conflicts if any pop up
