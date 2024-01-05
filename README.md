# Дефолтный каркас для REST API монолитных приложений на чистой архитектуре с использованием go gin
Делал для себя, чтобы быстрее писать код на хакатонах
* handler - эндпоинты и хэндлеры, middleware слой также тут
* service - основная бизнес логика
* repository - слой для взаимодействия с базой данных
* ent - основные сущности приложения
* docker-compose файл для запуска трех контейнеров, бд(postgres), api, и minio(минио не обязателен, можно просто закоментить и все будет работать отлично)
* Dockerfile - файл для api отдельно
* docs - часть со swagger, если не надо, то просто удалите нахер
* .env - для переменных окружения, в целом удобнее прописывать их в compose файле
  
--- 
* jwt авторизация написана
* написана регистрация
  
---

* для работы с minio используйте дефолтный aws3
  
---


# Дефолтные зависимости
go get github.com/gin-gonic/gin

go get github.com/dgrijalva/jwt-go

go get github.com/jmoiron/sqlx

go get github.com/lib/pq

go get github.com/joho/godotenv


### swagger
go get github.com/swaggo

github.com/swaggo/swag/cmd/swag

github.com/swaggo/files

github.com/swaggo/gin-swagger
