# Global Authetication Web
Este é o front-end da API de Autenticação Global, construído com React e TypeScript. Ele permite que os usuários se cadastrem, façam login e acessem uma página que exibe uma mensagem vinda de um endpoint protegido por token JWT.

## Funcionalidades
**Login:** Usuários podem fazer login com seu nome de usuário e senha.

**Registro:** Usuários podem se cadastrar com um nome de usuário e senha.

**Home:** Após o login, usuários acessam uma página que exibe uma mensagem de boas-vindas.

## Estrutura do Projeto

```bash
.
global-authetication-web
├── src
│   └── components
│   │   └── Login.tsx
│   │   └── Register.tsx
│   │   └── Home.tsx
│   └──App.tsx
│   └── index.tsx
│   └──styles.css
├── Dockerfile
├── package.json
├── package-lock.json
└── tsconfig.json
```

## Tecnologias Utilizadas
**React:** Biblioteca para construir interfaces de usuário.

**TypeScript:** Superset do JavaScript que adiciona tipagem estática.

**Axios:** Biblioteca para fazer requisições HTTP.

**React Router:** Biblioteca para gerenciamento de rotas.

## Instalação
### Pré-requisitos
Node.js (v20.17 ou superior)

Docker e Docker Compose (opcional, para execução em contêiner)

### Passo a Passo
- Clone o repositório:
```bash
git clone https://github.com/angelicalombas/global-authentication
```

- Navegue até a pasta da aplicação Web:
```bash
cd global-authentication-web
```

- Instale as dependências:
```bash
npm install
```

- Para executar com Docker, utilize:

```bash
docker-compose up --build
```

- Acesse a página inicial

Acesse o aplicativo no navegador em http://localhost:3000.

## Uso
**Página de Login:** Insira seu nome de usuário e senha. Se você ainda não tem uma conta, clique em "Cadastrar novo usuário".

**Página de Registro:** Insira seu nome de usuário e senha e clique em "Cadastrar". O username deve ter entre 3 e 20 caracteres e a senha deve ter no mínimo 8 caracteres. Após o registro, você será redirecionado para a página de Login.

**Página Home:** Após o login bem-sucedido, você será redirecionado para a página Home, onde verá uma mensagem de boas-vindas.