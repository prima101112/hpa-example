# Docker

This is walkthrough for docker, from building an image untill run it in docker machine

# Dockerfile

A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image.
Dockerfile should exactly named `Dockerfile` all docs could be read here
https://docs.docker.com/engine/reference/builder/

example Dockerfile
```
FROM alpine
LABEL maintainer="prima.adi@tokopedia.com"
WORKDIR /hpa
COPY hpa .
ENTRYPOINT /hpa/hpa
```

`FROM alpine` define where is the base image, in here we use alpine for the base image. it could be ubuntu or any image in docker registry

`LABEL maintainer="prima.adi@tokopedia.com"` The LABEL instruction adds metadata to an image. A LABEL is a key-value pair.

`WORKDIR /hpa` The WORKDIR instruction sets the working directory for any `RUN, CMD, ENTRYPOINT, COPY` and `ADD` instructions that follow it in the Dockerfile. If the WORKDIR doesn’t exist, it will be created even if it’s not used in any subsequent Dockerfile instruction. in this case is /hpa

`COPY hpa .` execute copy from our local to workdir that defined before

`ENTRYPOINT /hpa/hpa` entrypoint will execute the binary

since we all love php :p. lets build our docker images in php. we will use php with apache because we love to show to browser

```
FROM php:7.2-apache
COPY src/ /var/www/html/
```

thats from the official page docker apache https://hub.docker.com/_/php

`FROM php:7.2-apache` pull php-apache tag 
`COPY src/ /var/www/html/` and copy from src to their folder that will serve request

lets move.

## build the image

after we create the docker file we need to create docker images to our local docker machine.

`docker build . -t repo/imagename:version`

like :

`docker build . -t prima101112/hpa-example:php-example-01`

you will see the image already in list with `docker images` command

```
# prima at Nullp-MacPro.local in ~ [19:32:13]
→ uns-kube-srv215 $ docker images
REPOSITORY                TAG                 IMAGE ID            CREATED             SIZE
prima101112/hpa-example   php-example-01      8c0aaa7814c2        3 minutes ago       410MB
php                       7.2-apache          071b437a2194        2 weeks ago         410MB
prima101112/hpa-example   v0.1                71d0188db4a7        4 months ago        12.1MB
```

yes we success building our images
lets build our php images this image is pretty big.

## running the container

for running the container, there is several usual flag that need to remember. this is the usual command 

```
docker run -d --name myphp -p 8089:80 -it prima101112/hpa-example:php-example-01
```

that docker will run
- `-d` as daemon wthat will imidiat exit and logs will be show with `docker logs myphp`
- `--name myphp` is the name of container of this is not provided docker will create them self
- `-p 8089:80` will make port 8089 in our host go trhogh port 80 in container since php apache open port 80
- `-it` instructs Docker to allocate a pseudo-TTY connected to the container’s stdin; creating an interactive bash shell in the container
- `prima101112/hpa-example:php-example-01` image name, will pull if not exists

lets open our pages

http://localhost:8089/ ; http://localhost:8089/info.php

thats our php container

now what if we want to connect our php with postgress
we need extention.

## adding more command in Dockerfile

lets add more command to our Dockerfile this will install pdo on our container and enable the php extention

```
RUN apt-get update
RUN apt-get install -y libpq-dev \
    && docker-php-ext-configure pgsql -with-pgsql=/usr/local/pgsql \
    && docker-php-ext-install pdo pdo_pgsql pgsql
```

thats will install pdo psql

build again 

`docker build . -t prima101112/hpa-example:php-example-01`

always see the result with `docker images`, with `grep` if you already have so many images

# Docker Compose

Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration. 

lets try our compose.

make docker-compose.yml and start composing

```
version: '3'
services:
  myphp:
    build: .
    ports:
      - "8090:80"
    volumes:
      - ./src:/var/www/html/
  db:
    image: postgres:11
    restart: always
    environment:
      POSTGRES_PASSWORD: example
      POSTGRES_USER: example
    volumes:
      - ./initdb/:/docker-entrypoint-initdb.d/
  phppgadmin:
    image: dockage/phppgadmin:latest
    restart: always
    ports:
      - 8080:80
    environment:
      PHP_PG_ADMIN_SERVER_HOST: db
```

version 3 is the compose api

service is list of service in our compose they will share network and could call each other by service name.

`docker compose up -d` -d is for daemon
wait for image pulling and running container
and we have and fully development env