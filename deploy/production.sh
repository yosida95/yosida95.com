#!/usr/bin/env bash

LOGIN_NAME="rs2p"
HOSTNAME="8b66b542-8b34-476b-808a-68d85be08243.gehirn.ne.jp"
PORT=22329
TARGET="/var/www/yosida95.com/html/"

ROOT=$(cd "$(dirname "$0")"; cd ../; pwd)

rsync --delete --exclude ".DS_Store" -pthrvz  --rsh="ssh  -p ${PORT}" "${ROOT}/output/" "${LOGIN_NAME}@${HOSTNAME}:${TARGET}"
