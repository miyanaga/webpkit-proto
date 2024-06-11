#!/bin/sh

webpkit mirror /usr/local/apache2/htdocs /usr/local/apache2/htdocs/.webpkit

echo "LoadModule rewrite_module modules/mod_rewrite.so" >> /usr/local/apache2/conf/httpd.conf
sed -i 's/AllowOverride None/AllowOverride All/g' /usr/local/apache2/conf/httpd.conf
# sed -i 's/LogLevel warn/LogLevel warn rewrite:trace3/g' /usr/local/apache2/conf/httpd.conf

httpd-foreground