Тестовое задание для backend-стажёра в команду Advertising
===========================================================

Само задание см. [здесь](https://github.com/olteffe/avitoad/blob/master/task.md).

***
Установка и запуск
-------------------

1. Клонируем репозиторий: `$ git clone https://github.com/olteffe/avitoad.git`.
2. Переименовываем `.env.example` в `.env`(при необходимости меняем значения).
3. Собираем и запускаем `$ docker-compose up` или с помощью makefile `$ make docker.run`(docker должен быть предварительно 
установлен). Для запуска через Makefile необходимо предварительно установить `$ go get github.com/golang-migrate/migrate`.
4. Остановка контейнеров `$ docker-compose down` или `$ make docker.stop` соответственно.

***

Зависимости
------------

Использованы следующие внешние библиотеки:

* [gorilla/mux](https://github.com/gorilla/mux) - HTTP маршрутизатор.
* [godotenv](https://github.com/joho/godotenv/) - загружает переменные env из файла .env.
* [go-playground/validator](https://github.com/go-playground/validator) - реализует проверку значений для структур 
и отдельных полей на основе тегов.
* [google/uuid](https://github.com/google/uuid) - генерирует и проверяет UUID на основе RFC4122 и DCE 1.1.
* [jmoiron/sqlx](https://github.com/jmoiron/sqlx) - набор расширений стандартной `database/sql` библиотеки.
* [jackc/pgx](https://github.com/jackc/pgx) - драйвер для PostgreSQL.
* [squirrel](https://github.com/Masterminds/squirrel) - генератор SQL запросов.
* [swagger](https://github.com/swaggo/http-swagger) - генератор документации.
  
***
Принятые решения
-----------------

1. Использована стандартная библиотека net/http.
2. Для пагинации и сортировки выбран способ использования пары uuid и даты создания записи в базе данных. 
Подробности см. [здесь](https://medium.com/easyread/how-to-do-pagination-in-postgres-with-golang-in-4-common-ways-12365b9fb528).
3. Для денежного поля принят целочисленный тип данных.

***

Документация
--------------

Документация доступна при запуске контейнеров по ссылке http://0.0.0.0:5000/swagger/index.html
***

Архитектура приложения
-----------------------

![alt-текст](https://github.com/olteffe/avitoad/blob/master/assets/arch.png "Архитектура приложения")