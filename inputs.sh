#!/bin/sh

failure() {
    printf 1>&2 "%s\n" "$1"
    exit 1
}

SESSION_COOKIE=$1
[ "$SESSION_COOKIE" = "" ] && printf 1>&2 "use your session cookie as argument\n"

for day in $(seq 1 25)
do
    OUTPUT_FILE="./src/pkg/day$day/day$day.txt"
    [ -f "$OUTPUT_FILE" ] && continue

    CURRENT_INPUT=$(curl -fs --cookie "$SESSION_COOKIE" "https://adventofcode.com/2023/day/$day/input")
    [ $? != 0 ] && failure "an error occurred when getting input for day $day"

    mkdir -p "$(dirname $OUTPUT_FILE)"
    printf "%s" "$CURRENT_INPUT" > "$OUTPUT_FILE"
done
