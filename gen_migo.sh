#!/bin/bash


# Not the greatest script, but it works
ls -1 */main.go |
	while read file; do
		dir="${file%%\/*}";
		cd "$dir";
		~/Code/go/bin/migoinfer -log main.migo.log main.go > main.migo 2> main.migo.stderr;
		cd ..;
	done
