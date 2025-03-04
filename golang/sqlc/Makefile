include .env
export

MIGRATE=atlas migrate

DB_URL := postgres://$(DB_USER):$(DB_PASS)@:$(DB_FORWARD_PORT)/$(DB_NAME)?sslmode=disable
DEV_URL := docker://postgis/postgis:16-3.4-alpine/dev

migrate-status:
	$(MIGRATE) status --url "${DB_URL}" --dir "file://migrations" 

migrate-diff:
	$(MIGRATE) diff --env local

migrate-apply:
	$(MIGRATE) apply --url "${DB_URL}" --allow-dirty --exec-order non-linear

migrate-down:
	$(MIGRATE) down --url "${DB_URL}" --dev-url "${DEV_URL}"

migrate-hash:
	$(MIGRATE) hash

lint-setup:
	python3 -m ensurepip --upgrade
	sudo pip3 install pre-commit
	pre-commit install
	pre-commit autoupdate

.PHONY: migrate-status migrate-diff migrate-apply migrate-down migrate-hash lint-setup