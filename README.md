# Sofos

[![Watch the video](https://i.imgur.com/TuE78so.png)](https://youtu.be/sGA_f3I-WNw)

## Running database

> Once cockroach is installed.

Running cockroack
```bash
cockroach start --insecure  
```

Now cockroach is running on port 8080.

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
CREATE TABLE DomainInformation (
    SSL_grade VARCHAR(5),
    Title VARCHAR(100),
    Is_down BOOLEAN,
    Domain VARCHAR(50)
  );
```

```sql
CREATE TABLE HistoryQueries (
    Domain VARCHAR(50),
    Latest_query TIMESTAMP
  );
```
 
```sql
CREATE TABLE RelatedServers (
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
Now go is running on port 3000.


> In order to rebuild the executable.

```bash

go get -u github.com/go-chi/chi

go get -u github.com/lib/pq

```


## Running Web Server

>  Once Node and npm are installed


```bash

npm install -g @vue/cli

cd vue_app

npm install

npm run serve
```

Now go is running on port 8081.


> The web application will be available at the showed host. A Plugin for bypassing CORS may be necessary. 
