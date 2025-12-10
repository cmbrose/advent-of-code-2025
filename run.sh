#! /bin/bash

# Parse arguments using getopt
TEMP=$(getopt -o '' --long target:,debug -n 'run.sh' -- "$@")
if [ $? != 0 ]; then
    echo "Error parsing arguments" >&2
    exit 1
fi

eval set -- "$TEMP"

input="input"
debug_arg=""
target=""

while true; do
    case "$1" in
        --target)
            target="$2"
            shift 2
            ;;
        --debug)
            debug_arg="--debug"
            shift
            ;;
        --)
            shift
            break
            ;;
        *)
            echo "Unexpected option: $1" >&2
            exit 1
            ;;
    esac
done

# Handle positional argument (input)
if [ $# -gt 0 ]; then
    input="$1"
fi

# If --target not set, read from `current` file
if [[ -z "$target" ]]; then
    if [[ -f "current" ]]; then
        target=$(cat current)
    else
        echo "Error: --target not specified and current file not found."
        exit 1
    fi
fi

go run "${target}/main.go" --input "${target}/${input}.txt" $debug_arg