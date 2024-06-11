#!/bin/sh

webpkit beside /usr/local/apache2/htdocs

echo "LoadModule rewrite_module modules/mod_rewrite.so" >> /usr/local/apache2/conf/httpd.conf
sed -i 's/AllowOverride None/AllowOverride All/g' /usr/local/apache2/conf/httpd.conf
# sed -i 's/LogLevel warn/LogLevel warn rewrite:trace3/g' /usr/local/apache2/conf/httpd.conf

httpd-foreground