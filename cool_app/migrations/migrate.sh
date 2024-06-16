#!/bin/sh

while ! (migrate -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable -path /usr/var/migrations up);
do 
        echo "Migration is not successful. Trying again..."
done