#!/bin/sh

webpkit mirror /usr/share/nginx/html /usr/share/nginx/html/.webpkit

nginx -g 'daemon off;'