FROM mysql:latest
MAINTAINER Mario Inga <mario21ic@gmail.com>

ADD db.sql /tmp/db.sql
ADD init_db.sh /tmp/init_db.sh
RUN /tmp/init_db.sh
