# Sofos

## Executing Go Server

>  Once go is installed

```bash

cd go_chi

./main

```

## Executing Web Server

```bash
cd vue_app

npm install

npm run serve
```


## Setting Database

> Once cockroach is installed


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