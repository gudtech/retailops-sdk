#!/bin/bash
set -e

/usr/local/bin/provision-soa-service sdk_service

exec "$@"