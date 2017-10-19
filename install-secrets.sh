#!/bin/bash

DIR=$(cd `dirname $0` && pwd)

if [ -z "$DRONE_SERVER" ] || [ -z "$DRONE_TOKEN" ]; then
  echo "Missing drone environment [ DRONE_SERVER, DRONE_TOKEN ]"
  echo
  echo "export DRONE_SERVER={URL}"
  echo "export DRONE_TOKEN={TOKEN}"
  echo
  exit 0
fi

if [ -z "$1" ]; then
  echo "Missing repo parameter. run $0 {target}."
  exit 0
fi

if ! type "drone" &>/dev/null; then
  echo "'drone' command is not available."
  exit 0
fi

TARGET=$1

if drone repo info $TARGET |& grep 'client error 404' &>/dev/null; then
  echo "'$TARGET' repo is not available."
  exit 0
fi

if [ ! -d $DIR/secrets ]; then
  echo "'secrets' directory does not exist."
  exit 0
fi

read -p "Are you sure to import secrets to $TARGET? [y/n]" -n 1 -r
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
  [[ "$0" = "$BASH_SOURCE" ]] && exit 1 || return 1
fi
echo

for f in $DIR/secrets/*
do
  if [ ! -f $f ] || [ -d $f ]; then
    continue
  fi
  SECRET_NAME=`basename $f`
  SECRET_FILE="@$f"
  # remove existing secret
  if drone secret rm --repository $TARGET --name $SECRET_NAME &>/dev/null; then
    echo "update $SECRET_NAME..."
  else
    echo "import $SECRET_NAME..."
  fi
  # add secret
  drone secret add --repository $TARGET --name $SECRET_NAME --value $SECRET_FILE
done


