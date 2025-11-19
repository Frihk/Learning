#!/bin/bash


iNUM=$(head -n 179 streets/Buckingham_Place | tail -n 1 | grep -o '[0-9]\+')

echo $iNUM

cat "interviews/interview-${iNUM}"

echo "$main_SUSPECT"

