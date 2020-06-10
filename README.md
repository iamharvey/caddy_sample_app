# caddy_sample_app

以下说明based on Caddy [V2.0.0](https://github.com/caddyserver/caddy/releases/tag/v2.0.0)

Caddy默认有一个directive的执行顺序：

root

header

redir

rewrite

uri

try_files

basicauth

request_header

encode

templates

handle

handle_path

route

respond

reverse_proxy

php_fastcgi

file_server

acme_server

如果需要改变这个顺序，可以在glbal option里改变order的设置，或者使用route.

这个顺序很重要。
