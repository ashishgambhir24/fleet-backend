#!/bin/bash

rm -rf /tmp/fms-*

listVar="customer_service truck_service"
for i in $listVar; do
    go run "$i"/main.go &
done