#!/bin/bash

HUGO_VERSION=0.62.2
RELEASE_NAME=hugo_extended_${HUGO_VERSION}_Linux-64bit
FILE_NAME=${RELEASE_NAME}.tar.gz
HUGO_PACKAGE=https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/${FILE_NAME}

TMP_DIR=tmp

pwd

mkdir -p $TMP_DIR
if ! [ -e $TMP_DIR/$FILE_NAME ]; then
    curl $HUGO_PACKAGE -L -s -o $TMP_DIR/$FILE_NAME
fi

tar -zxvf $TMP_DIR/$FILE_NAME -C $TMP_DIR
mv $TMP_DIR/hugo bin/hugo

./bin/hugo --cleanDestinationDir -s web -d ../public
