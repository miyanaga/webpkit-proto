FROM httpd:2.4

ADD ./docker/html/ /usr/local/apache2/htdocs/
ADD ./docker/apache/mirror/.htaccess /usr/local/apache2/htdocs/.htaccess
ADD ./webpkit /usr/local/bin/webpkit
ADD ./docker/apache/mirror/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]