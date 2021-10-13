#! /usr/bin/bash
n=`ps -e |grep codemonkeycore|awk '{print $1}'`

echo $n
kill $n
./codemonkeycore >err 2>&1 &
