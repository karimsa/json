#!/bin/sh -e

alias json="$(dirname $0)/../bin/json"
jsonFile="$(dirname $0)/test.json"

function at () {
  echo "$(eval echo "\${${1}_${2}${3}}")"
}

## without root prefix

eval $(json < "${jsonFile}")

echo "array is of size: $(at tags length)"

for ((i=0;i<$(at tags length);++i)); do
  echo "arr[${i}] => $(at tags $i)"
done

## with a root prefix

eval $(json obj < "${jsonFile}")

echo ""
echo "array is of size: $(at obj_tags length)"

for ((i=0;i<$(at obj_tags length);++i)); do
  echo "arr[${i}] => $(at obj_tags $i)"
done

## when array elements are objects

eval $(json < "${jsonFile}")

echo ""
echo "array of objects is of size: $(at friends length)"

for ((i=0;i<$(at friends length);++i)); do
  echo "arr[${i}] => {"
  echo "  id: $(at friends $i _id),"
  echo "  name: $(at friends $i _name)"
  echo "}"
done
