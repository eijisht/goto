#!/bin/bash

PROJECT="$($HOME/goto/yamlparse "$@")"


if [[ "$#" -eq 0 ]]; then
	echo "$PROJECT"
	exit 0 
fi

if [ $? -eq 0 ]; then
	if [ "$PROJECT" = "Project already exists" ]; then
		echo "$PROJECT"
		exit 0
	fi
	if [ "$1" = "add" ]; then
		echo "Successfully added $2: $3"
		exit 0
	fi
	cd "$PROJECT"
	nvim .
fi


