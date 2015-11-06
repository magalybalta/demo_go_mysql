#!/bin/bash

/usr/bin/mysqld_safe &
sleep 5
mysql -u root -e "CREATE DATABASE demo_go"
mysql -u root demo_go < /tmp/db.sql
