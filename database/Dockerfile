FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD rootpass
ENV MYSQL_DATABASE fullcalendar
ENV MYSQL_USER testuser
ENV MYSQL_PASSWORD testpass
ENV TZ Asia/Tokyo

COPY ./config/my.conf /etc/mysql/conf.d/my.cnf

# docker build -t docker-mysql-8-0:1 -f Dockerfile-8-0 .
# docker run --name docker-mysql-8-0 -d -v $PWD/db-8-0:/var/lib/mysql -p 13306:3306 docker-mysql-8-0:1
# docker exec -it docker-mysql-8-0 bash