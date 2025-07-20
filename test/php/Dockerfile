FROM php:7.4-cli

# Install composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

WORKDIR /workspace

COPY composer.json .

RUN DEBIAN_FRONTEND=noninteractive apt update -y && \
    DEBIAN_FRONTEND=noninteractive apt install -y git zip unzip libzip-dev && \
    docker-php-ext-install zip
