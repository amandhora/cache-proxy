#!/bin/bash
for ((i=1;i<=10;i++));
do
    curl "localhost:8080/proxy?key=1"
    curl "localhost:8080/proxy?key=2"
    curl "localhost:8080/proxy?key=3"
    curl "localhost:8080/proxy?key=4"
done
