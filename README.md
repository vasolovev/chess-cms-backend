<<<<<<< Updated upstream
Проект по учету посещаемости студентами шахматных турниров. Скачивает результаты турниров в Lichess, сохраняет их в базу данных и формирует отчет в формате (.xlsx). Реализованы:
Запросы к Lichess
Сохранение в базу данных
REST API сервер
Выгрузка результатов турниров в формате (.xlsx)

База данных - MongoDB.
Используется шаблон чистой архитектуры.

В разработке:
Telegram бот
Добавление брокера очередей RabbitMQ
=======



docker run --name mongodb -d -p 27017:27017 mongo

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/chess.proto
>>>>>>> Stashed changes
