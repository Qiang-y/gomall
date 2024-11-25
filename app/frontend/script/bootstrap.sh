#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=forntend
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}