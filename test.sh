#!/bin/bash

a=$(cat /tmp/files/hello.txt)
b=sdd
if [ "$a" = "$b" ]
then
echo "TEST1: SUCCESS"
else
echo "TEST2: FAIL"
fi