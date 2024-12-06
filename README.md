# fullcycle_apis

fontes:
https://github.com/golang-standards/project-layout

#### Packages
1- internal - o coracao da app, nao deve ser compartilhado

2- pkg - libs para serem compartilhadas (ex: autenticacao)

3- cmd - onde fica o projeto, o executavel (build, run, main)
- cria-se uma pasta com o nome da aplicacao (ex: server) onde adicona o main

4- configs - como projeto inicia (variaveis de ambiente, como subir o sistema) - templates, arquivos go que faz boot

5- test - arquivos adicionais que ajudam (docs, exemplos, stubs, arquivos http, postman - podem nao ser arquivos .go)

6- api - especificacoes da api