#!/bin/bash
if [ $# -eq 0 ] 
then
echo "参数错"
exit
fi
if [ $# -eq 1 ] 
then
svc_like=$1
elif [ $# -eq 2 ] 
then
svc_like=$1
namespace="data-platform-"$2
kubens $namespace
fi
podname=`kubectl get pod|grep "$svc_like"|awk '{print $1}'|head -1`
kubectl exec -it  $podname  bash
