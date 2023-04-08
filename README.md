# toDoList

### DataBase

In this project a PostgreSQL data base is applied. In order to start data base from docker container it is required to download a PostgreSQL docker image with the command:

```
docker pull postgres
```
To start the daa base run the command:

```
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
```