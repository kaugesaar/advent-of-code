#!/bin/bash

DAY=$1

echo "Creating Boilerplate for day ${DAY}"

# Edits 1 to 01, 2 to 02 etc.
day_length=${#DAY}
if [ $day_length == 1 ]
then DISPLAY_DAY="0${DAY}"
else DISPLAY_DAY=$DAY
fi

TARGET=2022/day${DISPLAY_DAY}

# Make a copy of the boilerplate files
cd boilerplate
cp day.txt day${DAY}.txt
cp day_test.txt day${DAY}_test.txt

# Replace the variables
sed -i "" "s/{DAY}/${DAY}/g" day${DAY}.txt
sed -i "" "s/{DISPLAY_DAY}/${DISPLAY_DAY}/g" day${DAY}.txt
sed -i "" "s/{DAY}/${DAY}/g" day${DAY}_test.txt
sed -i "" "s/{DISPLAY_DAY}/${DISPLAY_DAY}/g" day${DAY}_test.txt

# Move them to to target folder and change to .go
cd ..
mkdir -p ${TARGET}
mv boilerplate/day${DAY}.txt ${TARGET}/day${DAY}.go
mv boilerplate/day${DAY}_test.txt ${TARGET}/day${DAY}_test.go

# Download the input data from advent of code
INPUT_URL="https://adventofcode.com/2022/day/${DAY}/input"
TEMP_INPUT="temp-input.txt"
curl "${INPUT_URL}" -H "cookie: session=$AOC_SESSION_COOKIE" -o "${TEMP_INPUT}" 2>/dev/null
mv ${TEMP_INPUT} ${TARGET}/day${DAY}.txt 
type NUL > ${TARGET}/day${DAY}_test.txt 2>/dev/null