<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/Toolnado/shorter-api">

# Shorter-api

Shorter-api - это сервис, позволяющий сокращать ссылки, записывать их в базу данных и получать оригинальную ссылку по ее уникальной короткой версии.


## Структура

Сервис состоит из клиента, реализованного в виде утилиты командной строки, grpc-сервера и базы данных postgresql, с реализованными миграциями через утилиту migrate. 


### Docker 

Сервис упокован в docker-контейнеры и собран с помощью docker-compose.
 - shorter-api_client
 - shorter-api_server
 - postgres
 - migrate

Образ клиента доступен в репозитории docker hub  по адресу: /77187719/shorter-api_client.

Образ сервера доступен в репозитории docker hub  по адресу: /77187719/shorter-api_server.

###  Unit-тесты

Реализованный функционал покрыт Unit-тестами.
