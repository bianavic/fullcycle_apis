# fullcycle_apis


#### Packages

```plaintext
📁 fullcycle_apis
├── 📁 api - especificacoes da api
├── 📁 cmd - o executavel (build, run, main)
│   └── 📁 server - pasta com o nome da aplicacao
│       ├── 📄 .env
│       ├── 📄 main.go  
│       └── 📄 test.db
├── 📁 configs - como projeto inicia (variaveis de ambiente, como subir o sistema) - templates, arquivos go que faz boot
│   └── 📄 config.go
├── 📁 internal - o coracao da app, nao deve ser compartilhado
│   ├── 📁 dto
│   │   └── 📄 dto.go
│   ├── 📁 entity
│   │   ├── 📄 product.go
│   │   ├── 📄 product_test.go  
│   │   ├── 📄 user.go
│   │   └── 📄 user_test.go
│   ├── 📁 infra
│   │   ├── 📁 database
│   │   │   ├── 📄 :memory
│   │   │   ├── 📄 interfaces.go  
│   │   │   ├── 📄 product_db.go
│   │   │   ├── 📄 product_db_test.go
│   │   │   ├── 📄 user_db.go
│   │   │   └── 📄 user_db_test.go
│   │   └── 📁 webserver
│   │       └── 📁 handlers
│   │           └── 📄 product_handlers.go
├── 📁 pkg - libs para serem compartilhadas (ex: autenticacao)
│   ├── 📁 entity
│   │   └── 📄 id.go
├── 📁 test - arquivos adicionais que ajudam (docs, exemplos, stubs, arquivos http, postman...)
|   └── 📄 product.http
├── 📄 .gitignore
├── 📁 go.mod
|   └── 📄 go.sum
└── 📄 README.md
```

### How To

Access project directory and run 
```shell
cd cmd/server 
go run main.go
```

#### Create Product
```shell
curl -H 'Content-Type: application/json' \
      -d '{ "name": "My Product","price": 100 }' \
      -X POST \
      http://localhost:8080/products
```

#### Database
Access sqlite db 
```shell
sqlite3 cmd/server/test.db
```
View products
```shell
sqlite> select * from products;
```

Tools:
- go-chi
- gorm.io
- viper
- google/uuid
- swag

Fontes:
https://github.com/golang-standards/project-layout