#!/bin/bash
set -eo pipefail

# You must be at least this tall to ride this ride.
THRESHOLD=70

# Exclude vendor, tmp, and proto-generated files from coverage.
PACKAGES=`go list ./... | grep -E -v 'vendor|tmp|mock|static'`

echo Packages being tested:
echo "$PACKAGES"

# Run go-acc to produce coverage.txt.  Clean up overly verbose go-acc output.
go-acc $PACKAGES 2> /dev/null | grep -v "of statements"

# print the function-by-function coverage report
go tool cover -func=coverage.txt 

# and also grab the last line for comparing to the threshold
COVERAGELINE=`go tool cover -func=coverage.txt | tail -n1`

# Drop everthing before the percentage
SEDVAL=`echo $COVERAGELINE | sed s/[^0-9.]*//g`

# Drop the % sign
COVERAGE=`echo $SEDVAL | tr -d "\012"`

# Use "bc" to compare the numbers, including decimal values.
PASS=`echo $COVERAGE '>=' $THRESHOLD | bc -l`
if [ "$PASS" -ne 1 ]; then
  echo "Coverage of $COVERAGE% is lower than required threshold of $THRESHOLD. :("
  exit 1
else
  echo "Coverage of $COVERAGE% is above required threshold of $THRESHOLD. :)"
fi
