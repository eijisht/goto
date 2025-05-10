#!/bin/bash

PROJECT="$($HOME/goto/yamlparse "$@")"


if [[ "$#" -eq 0 ]]; then
	echo "$PROJECT"
	return 0 
fi

if [ $? -eq 0 ]; then
	if [ "$PROJECT" = "Project already exists" ]; then
		echo "$PROJECT"
		return 0
	fi
	if [ "$1" = "add" ]; then
		echo "Successfully added $2: $3"
		return 0
	fi
	cd "$PROJECT"
fi


