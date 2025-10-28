## Tutorial: Instalação e Execução do Projeto Go

Este projeto é uma API REST para gerenciamento de filmes usando Go, Gin e GORM.

### 1. Instalar o Go

Baixe e instale o Go em https://go.dev/dl/
Siga as instruções para seu sistema operacional.

Para verificar a instalação:
```bash
go version
```

### 2. Clonar o projeto

```bash
git clone <url-do-repositorio> add 
cd go
```

### 3. Instalar as dependências

O arquivo `go.mod` lista todas as dependências do projeto (similar ao `package.json` ou `requirements.txt`).

Para instalar as dependências:
```bash
go mod tidy
```

### 4. Configurar o banco de dados

Edite a string de conexão no arquivo `main.go` (variável `dsn`) com seu usuário, senha e nome do banco MySQL.

Exemplo:
```
usuario:senha@tcp(127.0.0.1:3306)/acervo?charset=utf8mb4&parseTime=True&loc=Local
```

### 5. Executar o projeto

```bash
go run main.go
```

O servidor estará disponível em `http://localhost:3000`.

### 6. Testar rotas

Utilize ferramentas como Postman ou Insomnia para testar as rotas da API:
- `GET /movies`
- `GET /movies/:id`
- `POST /movies`
- `PUT /movies/:id`
- `DELETE /movies/:id`

---

#### Observação
O arquivo `go.mod` foi criado para gerenciar as dependências do projeto.


O Go não utiliza arquivos .env nativamente, mas é comum usar a biblioteca github.com/joho/godotenv para carregar variáveis de ambiente de um arquivo .env.


O GORM não possui um comando nativo de "seed" como Prisma ou Sequelize, mas é comum criar uma função personalizada para ler um arquivo (ex: JSON) e inserir os dados na tabela.

Na pasta database foi criada uma função Go que lê o arquivo movies.json e insere os registros na tabela movies usando o GORM. Dessa forma será possível popular a tabela movies.


Para popular a tabela, basta chamar database.SeedMovies("database/movies.json") após conectar ao banco. 

Para executar a função SeedMovies, basta chamá-la em main.go após conectar ao banco. Exemplo:

Logo após a conexão, antes de iniciar o servidor. Assim, os dados serão inseridos toda vez que rodar o projeto. 
```
database.ConnectDatabase()
database.SeedMovies("database/movies.json")
```

mas o mais correto é executar SeedMovies de forma isolada.
Para executar SeedMovies de forma isolada, crie um novo arquivo Go, por exemplo seed.go na raiz do projeto, com o seguinte conteúdo:
```
package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "github.com/learning_new_techs/go/database"
)

func main() {
    // Carrega variáveis do .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar .env")
    }

    // Conecta ao banco
    database.ConnectDatabase()

    // Executa o seed
    database.SeedMovies("database/movies.json")
}
```
Depois, execute no terminal:
```
go run seed.go
```


O GORM possui um recurso de migration automático via o método AutoMigrate. Ele cria e atualiza tabelas conforme os structs do seu código, mas não gera arquivos de migração versionados como o Prisma, Sequelize ou Alembic.

Exemplo de uso (já presente no projeto - vide database\db.go):
```
database.AutoMigrate(&structs.Movie{})
```
Para projetos mais avançados, existem ferramentas externas como golang-migrate para migrations versionadas, mas o padrão do GORM é o AutoMigrate.