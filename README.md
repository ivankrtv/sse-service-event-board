# sse-service-event-board
SSE service to push new events to client on Golang

## Description
Приложение для отправки данных клиенту с помощью SSE. Данные для отправки берутся из очереди 
Rabbit MQ, поэтому необходимо, чтобы приложение имело подключение к очереди, 
без необходимых данных в окружении, приложение будет падать с ошибкой

## Environment
```dotenv
# rabbit
RABBIT_USER=
RABBIT_PASSWORD=
RABBIT_HOST=
RABBIT_PORT=

# SSE
# Эндпоинт к которому будет подключаться клиент для получения данных о событиях
SSE_NEW_EVENT_ROUT="/event/new-event/1"

# APP
APP_PORT=
```