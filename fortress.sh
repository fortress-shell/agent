#!/usr/bin/env bash

export LIBVIRT_URL=
export KAFKA_URL=

cp /var/lib/libvirt/images/ubuntu.qcow2 $NOMAD_TASK_DIR
export VM_PATH=$NOMAD_TASK_DIR/ubuntu.qcow2
fortress
