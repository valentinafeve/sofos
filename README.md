# Sofos

[![Watch the video](https://ibb.co/fdYVbnV)](https://youtu.be/aGuEySRVbCM)

## Running database

> One cockroach is installed.

Running cockroack
```bash
cockroach start --insecure  
```

Now cockroach is running on the port 8080.

Running sql shell
```bash
cockroach sql --insecure 
```

## Setting Database

> Execute in the shell


```sql
CREATE DATABASE sofos;
```

```sql
SET SATABASE = sofos;
```

 
```sql
CREATE TABLE Domain_info (
    SSL_grade VARCHAR(5),
    Title VARCHAR(100),
    Is_down BOOLEAN,
    Domain VARCHAR(50)
  );
```

```sql
CREATE TABLE History_queries (
    Domain VARCHAR(50),
    Latest_query TIMESTAMP
  );
```
 
```sql
CREATE TABLE Related_servers (
    Domain VARCHAR(50),
    Server VARCHAR(50)
  );
``` 

```sql
CREATE TABLE Related_servers (
    Domain VARCHAR(50),
    Server VARCHAR(50)
  );
```

```sql
CREATE USER IF NOT EXISTS sofos_u;
```


```sql
GRANT SELECT ON history_queries TO sofos_u;
GRANT ALL ON history_queries TO sofos_u;
GRANT ALL ON domain_info TO sofos_u;
```

## Running Go Server

>  Once go is installed

```bash

cd go_chi

./main

```
Now go is running on the port 3000.


> In order to rebuild the executable.

```bash

go get -u github.com/go-chi/chi

go get -u github.com/lib/pq

```


## Running Web Server

>  Once Node and npm installed


```bash

npm install -g @vue/cli

cd vue_app

npm install

npm run serve
```

Now go is running on the port 8081.


> The web application will be available at the showed host. 
