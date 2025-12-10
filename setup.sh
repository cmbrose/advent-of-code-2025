#! /bin/bash

dest="$1"

# If it's a part 2, use the part 1, else use template
src="template"
if [[ "$dest" == *.2 ]]; then
    src="${dest%.2}.1"
fi

echo $dest > current

cp -r $src $dest

cd $dest

code *