#!/bin/sh
set -ex

cd /app
dnu restore
dnx web