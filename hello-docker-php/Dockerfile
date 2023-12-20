ARG ALPINE_VERSION=3.18
FROM alpine:${ALPINE_VERSION} as builder

WORKDIR /app


RUN apk add --no-cache \
  curl \
  nginx \
  php82 \
  php82-ctype \
  php82-curl \
  php82-dom \
  php82-fileinfo \
  php82-fpm \
  php82-gd \
  php82-intl \
  php82-mbstring \
  php82-mysqli \
  php82-opcache \
  php82-openssl \
  php82-phar \
  php82-session \
  php82-tokenizer \
  php82-xml \
  php82-xmlreader \
  php82-xmlwriter

RUN ln -s /usr/bin/php82 /usr/bin/php

COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer

COPY src/ /app

RUN composer install \
    --ignore-platform-reqs \
    --no-interaction \
    --no-plugins \
    --no-scripts \
    --prefer-dist


FROM alpine:${ALPINE_VERSION}
WORKDIR /var/www/html


RUN apk add --no-cache \
  curl \
  nginx \
  php82 \
  php82-fpm \
  supervisor

COPY config/nginx.conf /etc/nginx/nginx.conf
COPY config/conf.d /etc/nginx/conf.d/
COPY config/fpm-pool.conf /etc/php82/php-fpm.d/www.conf
COPY config/php.ini /etc/php82/conf.d/custom.ini
COPY config/supervisord.conf /etc/supervisor/conf.d/supervisord.conf
RUN chown -R nobody.nobody /var/www/html /run /var/lib/nginx /var/log/nginx
RUN ln -s /usr/bin/php82 /usr/bin/php


USER nobody

COPY --from=builder /app /var/www/html/

EXPOSE 8080

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
HEALTHCHECK --timeout=10s CMD curl --silent --fail http://127.0.0.1:8080/fpm-ping
