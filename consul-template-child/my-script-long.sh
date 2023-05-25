#!/bin/bash

trap 'echo "Be patient"' TERM

for ((n = 20; n; n--)); do
	echo "sleeping ${n}"
	sleep 1
done

echo "done"
exit 1
