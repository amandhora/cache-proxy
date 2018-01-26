#!/bin/bash
for ((i=1;i<=100;i++));
do
    echo "localhost:8080/proxy?key=$i"
    curl -v --header "Connection: keep-alive" "localhost:8080/proxy?key=$i"&
done
