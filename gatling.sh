#!/bin/bash

for (( i = 0; i < $1; ++i )); do
    make rerun-gatling
done