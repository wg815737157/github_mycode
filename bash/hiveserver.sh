#!/usr/bin/expect -f
set user wanggang9
set host 39.107.99.64
set port 2222
set password Wg630632!?
set timeout -1
#set totp totp
#set server_name Jumpserver
#set cmd "totp Jumpserver"
#spawn totp $server_name
#expect eof
#catch wait result
set google_code [exec totp Jumpserver]
spawn ssh $user@$host -p$port
expect "*password:"
send "$password\r"
expect "*auth]:"
send "$google_code\r"
expect "Opt>"
send "p\r"
expect "Opt>"
send "1\r"
interact
