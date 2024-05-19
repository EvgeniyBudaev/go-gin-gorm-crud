Инициализация зависимостей
```
go mod init github.com/EvgeniyBudaev/go-gin-gorm-crud
```

Сборка
```
go build -v ./cmd/
```

Удаление неиспользуемых зависимостей
```
go mod tidy -v
```

Gin
```
go get -u github.com/gin-gonic/gin
```

Gorm
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

UUID
```
go get -u github.com/google/uuid
```

Validator
```
go get -u github.com/go-playground/validator/v10
```
