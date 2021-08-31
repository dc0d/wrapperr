#!/bin/sh

while true
do
    watchman-wait -p "**/*.go" -- .
    clear
    make
done
