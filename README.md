# API in Go using Clean Architecture
API RESTFul to exemplify the use of clean architecture using go language

![clean-archicture](https://miro.medium.com/max/1400/1*O4pMWCi5kZi20SNOR6V33Q.png)

## 📌 Stack
- Mysql (https://www.mysql.com/)
- Go (https://go.dev/)
- Fiber (https://docs.gofiber.io/)
- GORM (https://gorm.io/)

## 🚀 Install
Run docker
```sh
docker-compose up -d
```

To see the log and follow it

```sh
docker-compose logs -f app.devbooks  
```


## 📝 To Do 
- Trabalhar com inversao de dependencia
- Analisar a clean architecture
- Colocar swagger
- Fazer uma parte via cmd com a lib do cobra
- Adicionar mais dados na tabela de usuários, hoje so tem o nome do usuário
- Criar a entitidade Book e vincular com o User
- Adicionar teste unitario


# API Documentation

## User

### 1. Create
```sh
curl --location --request POST 'http://localhost:5001/users/' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "teste"
}'
```

### 2. Delete
```sh
curl --location --request PUT 'http://localhost:5001/users/1' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "teste2"
}'
```

### 4. List
```sh
curl --location --request GET 'http://localhost:5001/users' \
--header 'Content-Type: application/json'
```

### 4. Find
```sh
curl --location --request GET 'http://localhost:5001/users/2' \
--header 'Content-Type: application/json'
```

### 5. Update
```sh
curl --location --request PUT 'http://localhost:5001/users/1' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "teste2"
}'
```

---
[Back to top of documentation](#api-documentation)