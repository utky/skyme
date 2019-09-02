#!/bin/bash
if [ $# -ne 1 ]; then
    echo "Please input path of executable"
    exit 1
fi
APP=$1

${APP} test/01_test.scm