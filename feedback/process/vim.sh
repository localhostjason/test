#!/usr/bin/env bash
vim -e /etc/httpd/conf/httpd.conf<<-!
:1,\$s/^/a/g
:wq
!
