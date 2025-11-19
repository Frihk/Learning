#!/bin/bash


INTERVIEW_PATH=$(grep -i "Ms. Church" the-final-cl-test/mystery/interviews/*)


INTERVIEW_NUMBER=$(echo "$INTERVIEW_PATH" | grep -oE '[0-9]+')


export INTERVIEW_NUMBER



echo "$INTERVIEW_NUMBER"



cat "$INTERVIEW_PATH"


echo "Dartey Henv"
