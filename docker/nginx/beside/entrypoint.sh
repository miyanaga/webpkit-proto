#!/bin/sh

webpkit beside /usr/share/nginx/html

# 本来不要だが現在のwebpkitのバグで読み込み権限追加が必要
chmod -R go+r /usr/share/nginx/html

nginx -g 'daemon off;'