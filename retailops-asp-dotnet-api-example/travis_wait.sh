#!/bin/sh
set -ex

while ! echo exit | curl http://localhost:5000; do sleep 10; done