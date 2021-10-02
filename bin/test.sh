#!/bin/bash

# go build
hugo --cleanDestinationDir --baseUrl="http://localhost:8080/" -s web -d ../public -w
