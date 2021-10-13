#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o codemonkey ./main
user=root
host=***
port=23
password=***
timeout=-1
/usr/bin/expect <<-EOF
# 设置ssh连接的用户名
#set user wanggang
# 设置ssh连接的host地址
#set host 192.168.144.27
# 设置ssh连接的port端口号
#set port 9999
# 设置ssh连接的登录密码
#set password ***
# 设置ssh连接的超时时间
#set timeout -1

spawn rsync  codemonkey $user@$host:/data/projects/codemonkeycore/codemonkeycore
expect "*password:"
# 提交密码
send "$password\r"
expect eof

spawn ssh  $user@$host
expect "*password:"
# 提交密码
send "$password\r"

expect "*#"
send "supervisorctl stop codemonkeycore &&supervisorctl start codemonkeycore \r"
expect eof
