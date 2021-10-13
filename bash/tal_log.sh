#!/bin/bash
if [ $# -eq 0 ] 
then
echo "参数错"
exit
fi
svc_like=$1
kubens jituan-bigdata-tdata

podname=`kubectl get pod|grep "$svc_like"|awk '{print $1}'|head -1`


kubectl logs -f $podname
