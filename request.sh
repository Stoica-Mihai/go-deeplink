#!/bin/sh

if [ $# -eq 0 ]; then
    echo "Usage: ./request.sh <command> <arg>"
    echo "Examples:"
    echo "- ./request.sh index - calls the index endpoint which returns the homepage"
    echo "- ./request.sh delete 123 - calls the delete endpoint which removes the deeplink association"
    echo "- ./request.sh details 123 - calls the details endpoint which returns details about the deeplink"
    echo "- ./request.sh unhandled - calls a endpoint which is not implemented"
    exit 1
fi

homeAddr="http://0.0.0.0:9000"

case "$1" in
    "create")
        curl -X POST $homeAddr/create
	;;
    "delete")
        if [ -z $2 ]; then
            echo "Delete requires a second argument which is the link value."
            echo "Example: ./request.sh delete 123, where "123" is the deeplink value"
	    exit 1
        fi

        curl -X DELETE $homeAddr/link/$2
	;;
    "details")
        if [ -z $2 ]; then
            echo "Details requires a second argument which is the link value that is inspected."
	    echo "Example: ./request details 123, where "123" is the deeplink value"
	    exit 1
	fi

        curl -X GET $homeAddr/link/$1 
	;;
    "unhandled")
        curl -X GET $homeAddr/random/unhandled/link
	exit 0
esac
