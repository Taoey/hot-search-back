#!/bin/bash
echo  "Stopping Server... "
# kill ftp
pid=`ps -o pid,command ax | grep "iris-cil-server" | awk '!/awk/ && !/grep/ {print $1}'`;
if [ "${pid}" != "" ]; then
     kill  ${pid};
fi


pid=`ps -o pid,command ax | grep smartrtb-server | awk '!/awk/ && !/grep/ {print $1}'`;
if [ "${pid}" != "" ]; then
     kill -2 ${pid};
fi
echo  "Server Stopped"
