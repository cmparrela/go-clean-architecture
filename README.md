# API in Go using Clean Architecture
API RESTFul to exemplify the use of clean architecture using go language

Fluxo 
- handler recebe o dado e cria o dto
- servico recebe o dto e cria entidade
- repositorio recebe entidade e faz o que tem que fazer
- repositorio busca os dados de algum lugar
- devolve o dado bruto
- servico converte para entidade
- devolve pro controller uma lista de entidades

Camadas

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
- Analisar a clean architecture 
  - As entidades estão sendo criadas no handler, seria melhor criar elas no use cases
  - Presenter
- Colocar swagger
- Fazer uma parte via cmd com a lib do cobra
- Criar repository armazenando em outro local sem ser o banco para exemplificar o repository pattern
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