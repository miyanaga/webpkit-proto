FROM nginx:1

ADD ./docker/html/ /usr/share/nginx/html/
ADD ./docker/nginx/beside/default.conf /etc/nginx/conf.d/default.conf
ADD ./webpkit /usr/local/bin/webpkit
ADD ./docker/nginx/beside/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]