include .env
export

MIGRATE=sql-migrate

limit=1
name=default
branch=test-mysql

migrate-status:
	$(MIGRATE) status -env=${branch}

migrate-up:
	$(MIGRATE) up -limit=${limit} -env=${branch}

migrate-down:
	$(MIGRATE) down -limit=${limit} -env=${branch}

migrate-redo:
	$(MIGRATE) redo -env=${branch}

migrate-create:
	$(MIGRATE) new -env=${branch} ${name}

migrate-skip:
	$(MIGRATE) skip -env=${branch}

.PHONY: migrate-status migrate-up migrate-down migrate-redo migrate-create migrate-skip