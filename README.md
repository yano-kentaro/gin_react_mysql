# gin_react_mysql
## initialization
```
git clone https://github.com/yano-kentaro/gin_react_mysql.git
cd gin_react_mysql
touch .env
touch ./backend/.env
cd ./frontend/fullcalendar
npm install
cd ../..
docker-compose up -d --build
# After containers start running, attach backend shell.
cd ./migration
go run main.go up
```

## ./.env
```
MYSQL_ROOT_PASSWORD='{Any Root Password}'
MYSQL_DATABASE='fullcalendar'
MYSQL_USER='{Any Name}'
MYSQL_PASSWORD='{Any Password}'
TZ='{Any Timezone}'
```

## ./backend/.env
```
HTTP_HOST='0.0.0.0'
HTTP_PORT='3000'
MYSQL_USER='{Any Name}'
MYSQL_PASSWORD='{Any Password}'
MYSQL_HOST='database'
MYSQL_PORT=3306
MYSQL_DATABASE='fullcalendar'
DB_DRIVER='mysql'
TZ='{Any Timezone}'
```