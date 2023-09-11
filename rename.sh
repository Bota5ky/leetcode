#!/bin/bash

find "." -type f -name "*.go" -execdir mv {} solution.go \;
