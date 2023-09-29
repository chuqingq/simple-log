#!/bin/bash

(sh ./app.sh 3>&2 2>&1 1>&3-) | ./logrotate -Name logrotate_test -EnableMemory true
