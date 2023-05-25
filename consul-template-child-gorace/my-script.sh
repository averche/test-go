#!/bin/bash

trap 'echo "received SIGTERM; ignoring it"' TERM

for ((n = 20; n; n--)); do
	echo "sleeping for ${n}s"
	sleep 1
done

echo 'done'
exit 1
