#!/bin/bash

export WORKER_ID=`uuidgen`
cp /home/linux/Images/xenial.b.qcow2 /home/linux/Images/$WORKER_ID
export VM_PATH/home/linux/Images/$WORKER_ID
agent $@
