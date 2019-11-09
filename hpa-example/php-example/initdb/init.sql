CREATE TABLE link (
   ID serial PRIMARY KEY,
   url VARCHAR (255) NOT NULL,
   name VARCHAR (255) NOT NULL,
   description VARCHAR (255),
   rel VARCHAR (50)
);

INSERT INTO link (url, name)
VALUES
   ('http://www.postgresqltutorial.com','PostgreSQL Tutorial');