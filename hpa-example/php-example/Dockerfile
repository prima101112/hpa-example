FROM php:7.2-apache

RUN apt-get update
RUN apt-get install -y libpq-dev \
    && docker-php-ext-configure pgsql -with-pgsql=/usr/local/pgsql \
    && docker-php-ext-install pdo pdo_pgsql pgsql

COPY config/driver.ini $PHP_INI_DIR/conf.d/

RUN apt-get install curl

VOLUME /var/www/html/
EXPOSE 80