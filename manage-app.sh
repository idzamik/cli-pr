#!/bin/bash
APP_NAME="myapp"
case $1 in
  start) docker run -d --name $APP_NAME -p 8080:80 nginx ;;
  stop) docker stop $APP_NAME && docker rm $APP_NAME ;;
  logs) docker logs $APP_NAME ;;
  *) echo "Usage: $0 {start|stop|logs}"; exit 1 ;;
esac
