# MESSAGE={"userId" => "123e4567-e89b-12d3-a456-426614174000","current_user" => "550e8400-e29b-41d4-a716-446655440000", "body": "Hola mundo!"}
msgCreateUser={"to": "leonardo.lastra@gmail.com","subject": "test","body": "Hola mundo!"}

export TOTAL_POSTS=5
URL_EXTERNAL=${PROTOCOL_EXTERNAL}://${HOST_EXTERNAL}
# Recordar que no se puede usar este token sin agregar "client_credentials" en la tabla hydra_client en grant_types
ACCESS_TOKEN = $(shell \
	curl -s -X POST https://marketdemo.ddns.net/bff/token \
		-d "grant_type=client_credentials" \
		-d "scope=openid offline mediamtx:publish" \
		-d "client_id=$(CLIENT_ID)" \
		-d "client_secret=$(CLIENT_SECRET)" \
	| jq -r '.access_token' | tr -d '\n' \
)
# PROFILE
# prod, dev, test
PROFILE ?= dev

ifeq ($(PROFILE),dev)
    ENVIRONMENTS := $(HOME)/.environments/demo.env
	VOLUME_DIR := /home/leonardo/demo/infra/volumes/dev
else ifeq ($(PROFILE),test)
    ENVIRONMENTS := /home/leonardo/demo/infra/test.env
	TEST_PATTERN := ^TestGetProduct$
	TEST_PACKAGE := integration
	VOLUME_DIR := /home/leonardo/demo/infra/volumes/test
else
    $(error PROFILE inválido: $(PROFILE))
endif

# bff,back,broker,cache,clickhouse,front,hydra,job,media,monitoring,postgres,scrap,solana,ws,e2e
COMPOSE_PROFILES=postgres,cache,hydra,bff,traefik,front,back
# COMPOSE_PROFILES=postgres,cache,hydra,bff,traefik,front,back,monitoring,e2e
# COMPOSE_PROFILES=postgres,cache,broker,hydra,bff,traefik,front,back,monitoring,scrap
# COMPOSE_PROFILES=postgres,broker,monitoring
# COMPOSE_PROFILES=postgres,scrap,broker
# COMPOSE_PROFILES := postgres,bff
# COMPOSE_PROFILES=front,traefik

COMPOSE_DIR := infra/docker/$(PROFILE)
COMPOSE_FILE := $(COMPOSE_DIR)/monitoring.yml:$(COMPOSE_DIR)/hydra.yml:$(COMPOSE_DIR)/rabbitmq.yml:$(COMPOSE_DIR)/mediamtx.yml:$(COMPOSE_DIR)/redis.yml:$(COMPOSE_DIR)/elastic.yml:$(COMPOSE_DIR)/common.yml:$(COMPOSE_DIR)/traefik.yml:$(COMPOSE_DIR)/envoy.yml:$(COMPOSE_DIR)/postgres.yml:$(COMPOSE_DIR)/portainer.yml:$(COMPOSE_DIR)/test.yml:infra/docker/services.yml:$(COMPOSE_DIR)/services.$(PROFILE).yml

export COMPOSE_PROFILES
export COMPOSE_FILE
export VOLUME_DIR
export TEST_PATTERN
export TEST_PACKAGE
# ENVIRONMENTS=/home/leonardo/.environments/demo.env
include ${ENVIRONMENTS}

.PHONY: config sendmsg queue up restart kill down build exec migra scrap delay clean recomender ps tfapply tfdestroy tfssh redis dump trans token test test-unit sshManager sshWorker brokerMigra dumpHydra tfoutput e2e swagger

up:
	@docker compose \
		--env-file ${ENVIRONMENTS} \
		--project-directory ${PWD} \
		--project-name marketplace \
		up -d --force-recreate $(filter-out $@,$(MAKECMDGOALS))

ps:
	@watch docker compose \
		--env-file ${ENVIRONMENTS} \
		--project-directory ${PWD} \
		--project-name marketplace \
		ps -a

down:
	@docker compose --env-file ${ENVIRONMENTS} --project-directory ${PWD} \
		--project-name marketplace down

queue:
	@docker exec broker-job rabbitmqadmin declare queue name=$(QUEUE_NAME) durable=true

delay:
	@docker compose --env-file ${ENVIRONMENTS} exec scrap-worker python -c "from main import run_modelo_spider; run_modelo_spider.delay()"

sendmsg:
	@docker exec broker rabbitmqadmin publish \
  		exchange=$(EVT_USER_REGISTERED_EXCHANGE) \
		routing_key=$(EVT_USER_REGISTERED_EMAIL_QUEUE) \
		payload='$(msgCreateUser)'

tfapply:
	@TF_VAR_ssh_private_key=$(cat ~/.ssh/id_rsa) terraform -chdir=./terraform/ apply -auto-approve

tfdestroy:
	@TF_VAR_ssh_private_key=$(cat ~/.ssh/id_rsa) terraform -chdir=./terraform/ destroy -auto-approve

tfssh:
	ssh root@$(terraform -chdir=./terraform/ output -raw manager_public_ip)

redis:
	@docker compose \
		--project-directory ${PWD} \
		--project-name marketplace \
		 exec -it cache redis-cli

sshManager:
	@ssh root@$(shell terraform -chdir=./terraform/ output -raw manager_public_ip)

sshWorker:
	@ssh root@$(shell terraform -chdir=./terraform/ output -raw worker_public_ip)

tfoutput:
	@terraform -chdir=terraform output

build:
	@docker compose \
		--env-file ${ENVIRONMENTS} \
		--project-directory ${PWD} \
		--project-name marketplace \
		build $(SERVICE) --no-cache

migra:
	@docker run --env-file ${ENVIRONMENTS} --network ${NETWORK_APPLICATION} --rm -v ./flyway/sql:/flyway/sql flyway/flyway:latest-alpine \
	  -url=jdbc:postgresql://${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB} \
	  -user=${POSTGRES_USER} \
	  -password=${POSTGRES_PASSWORD} \
	  -schemas=marketplace,scrap,hydra,marketplacedemo,chat,extension,posts,inventory,scheduling,sales,payment \
	  -cleanDisabled=false \
	  -locations=filesystem:/flyway/sql/migration,filesystem:/flyway/sql/seed \
	  -placeholders.URL_EXTERNAL=${URL_EXTERNAL} \
	  migrate

clean:
	@docker run --env-file ${ENVIRONMENTS} --network ${NETWORK_APPLICATION} --rm -v ./flyway/sql:/flyway/sql flyway/flyway:latest-alpine \
	  -url=jdbc:postgresql://${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB} \
	  -user=${POSTGRES_USER} \
	  -password=${POSTGRES_PASSWORD} \
	  -schemas=marketplace,scrap,hydra,marketplacedemo,chat,extension,posts,inventory,scheduling,sales,payment \
	  -cleanDisabled=false \
	  -locations=filesystem:/flyway/sql/migration,filesystem:/flyway/sql/seed \
	  -placeholders.URL_EXTERNAL=${URL_EXTERNAL} \
	  clean

brokerMigra:
	@docker run --env-file ${ENVIRONMENTS} \
		--network ${NETWORK_APPLICATION}  \
		-v ./infra/docker/config/rabbitmq/init_rabbitmq.sh:/init_rabbitmq.sh \
		-e HOST=broker -e PORT=15672 -e USERNAME=guest -e PASSWORD=guest \
		--entrypoint /init_rabbitmq.sh \
		--rm \
		rabbitmq:${RABBITMQ_VERSION}

scrap:
	@docker exec -it beat-service celery -A main call main.run_modelo_spider

dump:
	@pg_dump \
		--data-only \
		--table=$(SCHEMA).$(TABLE) \
		--column-inserts \
		--rows-per-insert=100000000 \
		--disable-triggers \
		-h localhost \
		-p 5432 \
		-U ${POSTGRES_USER} \
		${POSTGRES_DB} > flyway/sql/seed/R__insert_$(SCHEMA)_$(TABLE).up.sql

dumpHydra:
	@pg_dump \
		-n $(SCHEMA) \
		-h localhost \
		-p 5432 \
		-U ${POSTGRES_USER} \
		${POSTGRES_DB} > flyway/sql/migration/V_0_0_0__$(SCHEMA).up.sql

psql:
	@psql -h localhost -U ${POSTGRES_USER} -d ${POSTGRES_DB} -p 5432

test-unit:
	@$(MAKE) _test-unit ENVIRONMENTS=/home/leonardo/.environments/test.env

_test-unit:
	@docker compose \
		--env-file ${ENVIRONMENTS} \
		--project-directory ${PWD} \
		--project-name marketplace \
		run --rm  --no-deps bff-service \
		go test ./controller -race

_test-integration:
	@docker compose \
		--env-file ${ENVIRONMENTS} \
		--project-directory ${PWD} \
		--project-name marketplace \
		run --rm  --no-deps bff-service \
		go test -run TestGetProduct ./integration

test-integration:
	@$(MAKE) _test-integration ENVIRONMENTS=/home/leonardo/.environments/test.env

_test-front:
	@docker compose \
		--env-file ${ENVIRONMENTS} \
		--project-directory ${PWD} \
		--project-name marketplace \
		run --rm --no-deps front \
		npm run test -- --run

test-frontend:
	@$(MAKE) _test-front ENVIRONMENTS=/home/leonardo/.environments/test.env

e2e:
	@docker exec -it e2e npm run test -- --headed

swagger:
	@docker exec -it bff swag init -g server/docs.go -d . --parseInternal