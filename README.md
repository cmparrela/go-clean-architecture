# API in Go using clean architecture

### Stack
- Mysql (https://www.mysql.com/)
- Go (https://go.dev/)
- Fiber (https://docs.gofiber.io/)
- GORM (https://gorm.io/)

### TODO
- Trabalhar com inversao de dependencia
- Analisar a clean architecture
- Colocar swagger
- Melhorar docker
- Melhorar a doc para iniciar o projeto
- Fazer uma parte via cmd com a lib do cobra
- Adicionar mais dados na tabela de usuários, hoje so tem o nome do usuário
- Criar a entitidade Book e vincular com o User

--------
# API Documentation

## Indices
* [User](#user)
  * [Create](#1-create)
  * [Delete](#2-delete)
  * [List](#3-list)
  * [Find](#4-find)
  * [Update](#5-update)


## User

### 1. Create

***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://localhost:5001/users
```

***Body:***

```js        
{
  "name": "User1"
}
```

### 2. Delete

***Endpoint:***
```bash
Method: DELETE
Type: 
URL: http://localhost:5001/users/2
```

### 3. List

***Endpoint:***
```bash
Method: GET
Type: 
URL: http://localhost:5001/users
```

### 4. Find

***Endpoint:***
```bash
Method: GET
Type: 
URL: http://localhost:5001/users/2
```

### 5. Update

***Endpoint:***
```bash
Method: PUT
Type: URLENCODED
URL: http://localhost:5001/users/1
```

***Body:***

```js        
{
  "name": "NewUserName"
}
```


---
[Back to top](#api-documentation)