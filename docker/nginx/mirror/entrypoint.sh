#!/bin/sh

webpkit mirror /usr/share/nginx/html /usr/share/nginx/html/.webpkit

# 本来不要だが現在のwebpkitのバグで読み込み権限追加が必要
chmod -R go+r /usr/share/nginx/html/.webpkit

nginx -g 'daemon off;'