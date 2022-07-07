#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)

if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

export COMMENT_RUNTIME_ROOT=$RUNTIME_ROOT
export COMMENT_LOG_DIR="$RUNTIME_ROOT/log"

if [ ! -d "$COMMENT_LOG_DIR/app" ]; then
    mkdir -p "$COMMENT_LOG_DIR/app"
fi

if [ ! -d "$COMMENT_LOG_DIR/rpc" ]; then
    mkdir -p "$COMMENT_LOG_DIR/rpc"
fi

exec "$CURDIR/bin/comment"
