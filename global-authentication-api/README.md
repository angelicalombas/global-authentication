# API GLOBAL-AUTHENTICATION

Esta API fornece funcionalidades básicas de autenticação, como registro de usuários e login. É construída com Go e utiliza o GORM para interagir com um banco de dados PostgreSQL.

## Funcionalidades

**Registro de Usuário:** Permite que novos usuários se registrem com um nome de usuário e senha.

**Login:** Permite que usuários existentes façam login e recebam um token JWT.

**Home:** Endpoint protegido que exibe uma mensagem de boas-vindas para usuários autenticados.

## Estrutura do Projeto

```shell
.
global-authentication-api
├── config
│   └── connection_db.go
├── controllers
│   └── user_controller_interface.go
│   └── user_controller.go
├── docs
│   └── docs.go
├── models
│   └── user.go
├── repositories
│   └── user_repository.go
├── routes
│   └── routes.go
├── services
│   └── userService.go
├── utils
├── docker-compose.yml
├── main.go
├── Dockerfile
├── go.mod
├── go.sum
│   └── main.go
└── README.md
```

## Como Executar

### Pré-requisitos
Go (1.23.1 ou superior)

Docker e Docker Compose

### Passo a Passo
- Clone o repositório:

```bash
git clone https://github.com/angelicalombas/global-authentication
```

- Navegue até a pasta da API:

```bash
cd global-authentication-api
```

- Crie a imagem e inicie os serviços:

```bash
docker-compose up --build
```

- Aguarde até que a mensagem abaixo seja exibida no terminal, indicando que o serviço foi iniciado:

```bash
Server running on port :8000
```

- Acesse a API:

A API estará rodando em http://localhost:8000

A documentação Swagger estará disponível em http://localhost:8000/swagger/index.html

### Endpoints

**1. Registrar Usuário**

**URL:** /register

**Método:** POST

**Corpo da Requisição:**

```json
{
  "username": "string",
  "password": "string"
}
```

**Respostas:**

201 Created: Usuário criado com sucesso.

400 Bad Request: Falha na validação.

409 Conflict: Nome de usuário já está em uso.

**2. Login**

**URL:** /login

**Método:** POST

**Corpo da Requisição:**

```json
{
  "username": "string",
  "password": "string"
}
```

**Respostas:**

200 OK: Retorna um token JWT.

400 Bad Request: Falha na validação.

401 Unauthorized: Credenciais inválidas.

**3. Home**

**URL:** /home

**Método:** GET

**Cabeçalho:**

**Authorization:** Bearer <token>

**Respostas:**

200 OK: Mensagem de boas-vindas.

### Banco de Dados
A aplicação utiliza o PostgreSQL. As credenciais e o nome do banco de dados estão configurados no arquivo docker-compose.yml.

```yaml
services:
  db:
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: userdb
```

## Rodando os testes

Para rodar os testes, rode o seguinte comando

```bash
  go test ./...
```