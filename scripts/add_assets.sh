#!/bin/bash

################################################################################
# Compile css
echo "Compiling CSS"
mkdir -vp ./static/css 2>/dev/null || true

command -v sass >/dev/null
if [ $? != 0 ]
then
    command -v sassc >/dev/null
    if [ $? != 0 ]
    then
        echo "Build failed. No sass or sassc in \$PATH"
        exit 1
    else
        SASS=sassc
    fi
else
    SASS=sass
fi

for f in ./assets/sass/*
do
    filename="$(basename -- $f | cut -d '.' -f1)"
    if [[ ! $filename =~ ^_.* ]]
    then
        echo "    Compiling $f"
        $SASS $f ./static/css/$filename.css
    fi
done

################################################################################
# Write to .go file
GO_FILE="./mdserver/gen/genCss.go"

command -v cleancss >/dev/null
if [ $? != 0 ]
then
    CLEANCSS=0
else
    CLEANCSS=1
fi

rm -v $GO_FILE
echo -e "package gen\n\n// Stylesheets\n//This file is autogenerated from assets directory\nconst (" > $GO_FILE

echo "Assembling $GO_FILE"
for f in ./static/css/*.css
do
    echo "    Adding $f"
    filename="$(basename -- $f | cut -d '.' -f1)"
    [ $CLEANCSS != 0 ] && cleancss -o $f $f
    echo -e "\t${filename} = \`$(cat $f)\`\n" >> $GO_FILE
done

echo ")" >> $GO_FILE

rm -rf ./static/css
