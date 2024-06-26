#!/usr/bin/env sh

set -euo pipefail

printf "Can you guess the number? > "

read i

if printf $i | grep -e [^0-9]; then
    printf "bye hacker!"
    exit 1
fi

r=$(head -c512 /dev/urandom | tr -dc 0-9)

if [[ $r == $i ]]; then
    printf "How did you know?!"
    cat flag.txt
else
    printf "Nope. It was $r."
fi
