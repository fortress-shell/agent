#!/usr/bin/env bash

#cp /var/lib/libvirt/images/ubuntu.qcow2 $PWD
export NOMAD_META_REPOSITORY_URL=git@github.com:mikefaraponov/fortress.git
export NOMAD_META_BRANCH=master
export NOMAD_META_COMMIT=238cb63efdb25dd228689b02ab46255c2f3c5ff5
export NOMAD_DC=dc1
export NOMAD_META_BUILD_ID=100
export NOMAD_META_USER_ID=1
export JOB_ID=$(uuidgen)
export LIBVIRT_URL=qemu+ssh://root@fortress-4/system?socket=/var/run/libvirt/libvirt-sock
export KAFKA_URL=192.168.2.3:9092
export VM_PATH=/home/linux/TaskDie/ubuntu.qcow2
export PAYLOAD_PATH=$PWD/payload.yml
export NOMAD_META_SSH_KEY=$(cat ./gh_rsa)
export DISK_PATH=/home/linux/TaskDie/user-data.img
