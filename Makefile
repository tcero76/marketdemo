ENVIRONMENTS=/home/leonardo/.environments/demo.env
include ${ENVIRONMENTS}

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

.PHONY: config sendmsg queue up kill down build exec migra scrap delay clean recomender ps tfapply tfdestroy tfssh redis dump trans token

up:
	@docker compose \
		--env-file ${ENVIRONMENTS} \
		--project-directory ${PWD} \
		--project-name marketplace \
		up -d $(filter-out $@,$(MAKECMDGOALS))

config:
	@docker compose \
		--env-file ${ENVIRONMENTS} \
		--project-directory ${PWD} \
		--project-name marketplace \
		config

ps:
	@watch docker compose \
		--env-file ${ENVIRONMENTS} \
		--project-directory ${PWD} \
		--project-name marketplace \
		ps -a

down:
	@docker compose --env-file ${ENVIRONMENTS} --project-directory ${PWD} \
		--project-name marketplace down

kill:
	@docker rm -f $(filter-out $@,$(MAKECMDGOALS))

exec:
	@docker exec -it $(filter-out $@,$(MAKECMDGOALS)) /bin/bash

queue:
	@docker exec broker-job rabbitmqadmin declare queue name=$(QUEUE_NAME) durable=true

delay:
	@docker compose --env-file ${ENVIRONMENTS} exec scrap-worker python -c "from main import run_modelo_spider; run_modelo_spider.delay()"

recomender:
	@docker compose --env-file ${ENVIRONMENTS} exec recomender python -c "from main import calculate_recommendations_task; calculate_recommendations_task.delay()"

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
	@docker compose --env-file ${ENVIRONMENTS} exec -it cache redis-cli

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

reset:
	@sudo rm -rf ./infra/volumes/timescaleDB/
	@mkdir -p ./infra/volumes/timescaleDB/data
migra:
	@docker run --env-file ${ENVIRONMENTS} --network ${NETWORK_APPLICATION} --rm -v ./flyway/sql:/flyway/sql flyway/flyway:latest-alpine \
	  -url=jdbc:postgresql://${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB} \
	  -user=${POSTGRES_USER} \
	  -password=${POSTGRES_PASSWORD} \
	  -schemas=marketplace,scrap,hydra,marketplacedemo,chat,extension \
	  -cleanDisabled=false \
	  -locations=filesystem:/flyway/sql/migration,filesystem:/flyway/sql/seed \
	  -placeholders.URL_EXTERNAL=${URL_EXTERNAL} \
	  migrate

clean:
	@docker run --env-file ${ENVIRONMENTS} --network ${NETWORK_APPLICATION} --rm -v ./flyway/sql:/flyway/sql flyway/flyway:latest-alpine \
	  -url=jdbc:postgresql://${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB} \
	  -user=${POSTGRES_USER} \
	  -password=${POSTGRES_PASSWORD} \
	  -schemas=marketplace,scrap,hydra,marketplacedemo,chat,extension \
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

trans:
	@ffmpeg \
		-thread_queue_size 512 -f v4l2 -i /dev/video0 \
		-thread_queue_size 512 -f alsa -i default \
		-c:v libx264 -preset veryfast -tune zerolatency \
		-pix_fmt yuv420p \
		-g 30 \
		-c:a aac -b:a 128k \
		-f flv rtmp://localhost:1935/streams?jwt=$(ACCESS_TOKEN)

webrtc:
	@curl -k -X POST \
		-H "Content-Type: application/sdp" \
		-H "Authorization: Basic $(echo -n 'any:any' | base64)" \
		http://localhost:8889/streams/whip \
		--data-binary @offer.sdp

play:
	@ffplay -headers "Authorization: Bearer $(ACCESS_TOKEN)" http://localhost:8888/streams/index.m3u8
