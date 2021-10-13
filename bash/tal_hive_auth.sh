#!/bin/bash
kubens jituan-bigdata-tdata
podname=`kubectl get pod|grep "hive-auth"|awk '{print $1}'|head -1`
kubectl logs -f $podname
