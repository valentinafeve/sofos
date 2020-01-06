CREATE TABLE Domain_info (
    SSL_grade VARCHAR(5),
    Title VARCHAR(100),
    Is_down BOOLEAN,
    Domain VARCHAR(50)
  );

CREATE TABLE History_queries (
    Domain VARCHAR(50),
    Latest_query TIMESTAMP
  );

CREATE TABLE Related_servers (
    Domain VARCHAR(50),
    Server VARCHAR(50)
  );
