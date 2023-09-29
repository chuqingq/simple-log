#!/bin/bash

while true; do
    d="$(date)"
    echo "${d} this is an output"
    echo "${d} is an error" >&2
done
