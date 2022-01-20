<div align="center">
      <h1>Todo List CLI</h1> </br>
</div>

 [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)


Command Line Interface que permite gerenciar uma lista de tarefas

## Comandos & Flags

Comando para adicionar uma task:

~~~shell
./task add "descricao de um todo" -H
~~~

Seguido das flags de prioridade:
- --high ou -H
- --low ou -L
- Sem flag, prioridade normal

Listar todos:

~~~shell
./task list-all
~~~

Listar um todo pelo Id:
~~~shell
./task list "id"
~~~

Atualizar o status de um todo:
~~~shell
./task update "id"
~~~


Deletar um todo:
~~~shell
./task remove "id"
~~~

Deletar todos:
~~~shell
./task delete
~~~


