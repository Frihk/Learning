#!/bin/bash


iNUM=$(head -n 179 the-final-cl-test/mystery/streets/Buckingham_Place | tail -n 1 | grep -o '[0-9]\+')

echo $iNUM

cat "the-final-cl-test/mystery/interviews/interview-${iNUM}"

echo "$main_SUSPECT"

