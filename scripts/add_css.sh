#!/bin/sh

GO_FILE="./mdserver/html/css.go"

echo -e "package html\n\n" > $GO_FILE

for f in ./static/css/*
do
    filename="$(basename -- $f | cut -d '.' -f1)"
    cleancss -o $f $f
    echo -e "const ${filename^^} = \`\n$(cat $f)\`\n" >> $GO_FILE
done
