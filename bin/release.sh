#!/bin/bash
go build
hugo --cleanDestinationDir -s web -d ../public
