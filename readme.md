# Desafio Clean Architecture go

Este projeto foi criado para o desafio de Clean Architecture do curso de go

## Como executar:

**Passo 1:** Faça o clone do projeto no diretório de sua preferência.

```shell
git clone https://github.com/elielbatiston/fullcycle_desafio_clean_arch_curso_go
```

**Passo 2:** Vá para o terminal, acesse o diretório que você clonou o projeto e execute o comando abaixo:

```shell
make migrate
```

**Passo 3:** Acesse o diretorio cmd

```shell
cd cmd
```

**Passo 4:** Acesse o diretorio ordersystemcd

```shell
cd ordersystem
```

**Passo 5:** Suba o servidor com o comando abaixo

```shell
go run main.go wire_gen.go
```

## Como testar:

1. WebServcer

Acesse o diretorio api dentro do VSCode e envie a requisição de POST e depois a de GET

2. GraphQL

Acesse a URL http://localhost:8080 e utilize o playground pra executar

a. Exemplo para criar orders

```
mutation createOrder {
	createOrder(input:{id:"el1",Price:50.0, Tax:10.0}){
    id
    Price
    Tax
    FinalPrice
  }
}
```

b. Consultar as orders criadas

```
query queryOrders {
  ListOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

3. gRPC

a. Digite o comando abaixo

```shell
evans -r repl
```

b. Digite o comando abaixo para criar uma order

```shell
call CreateOrder
```

Preencha os campos solicitados

c. Digite o comando abaixo para listar as orders

```shell
call ListOrders
``` 

## Observaçõs:

1. Este documento leva em consideração que você tem o VSCode instalado em sua máquina com a extensão Rest Client instalada
2. Este documento leva em consideração que você tem o evans instalado em sua máquina
3. Não fazia parte do desafio utilizar Docker
