#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)

if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

export USER_RUNTIME_ROOT=$RUNTIME_ROOT
export USER_LOG_DIR="$RUNTIME_ROOT/log"

if [ ! -d "$USER_LOG_DIR/app" ]; then
    mkdir -p "$USER_LOG_DIR/app"
fi

if [ ! -d "$USER_LOG_DIR/rpc" ]; then
    mkdir -p "$USER_LOG_DIR/rpc"
fi

exec "$CURDIR/bin/user"
