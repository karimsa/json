#!/bin/sh -e

alias json="$(dirname $0)/../bin/json"
jsonFile="$(dirname $0)/test.json"

function at () {
  name="$1"
  index="$2"

  echo "$(eval echo "\${${name}_${index}}")"
}

eval $(json < "${jsonFile}")

echo "array is of size: $(at tags length)"

for ((i=0;i<$(at tags length);++i)); do
  echo "arr[${i}] => $(at tags $i)"
done
