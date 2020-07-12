#!/bin/sh

################################################################################
# Compile css

mkdir -v ./static/css 2>/dev/null || true

command -v sass >/dev/null
if [ $? != 0 ]
then
    command -v sassc >/dev/null
    if [ $? != 0 ]
    then
        echo "Build failed. No sass or sassc in $PATH"
        exit 1
    else
        SASS=sassc
    fi
else
    SASS=sass
fi

$SASS ./sass/fv_common.sass ./static/css/fv_common.css
$SASS ./sass/fv_dark.sass ./static/css/fv_dark.css
$SASS ./sass/fv_light.sass ./static/css/fv_light.css
$SASS ./sass/ghmd.sass ./static/css/ghmd.css
$SASS ./sass/ghmd_dark.sass ./static/css/ghmd_dark.css
$SASS ./sass/ghmd_light.sass ./static/css/ghmd_light.css
$SASS ./sass/fonts.scss ./static/css/fonts.css
$SASS ./sass/style.sass ./static/css/style.css


################################################################################
# Write to .go file
GO_FILE="./mdserver/html/css.go"

command -v cleancss >/dev/null
if [ $? != 0 ]
then
    echo "Build failed. No cleancss in \$PATH"
    echo "Run 'npm install clean-css'"
    exit 1
fi

echo -e "package html\n\n" > $GO_FILE

for f in ./static/css/*
do
    filename="$(basename -- $f | cut -d '.' -f1)"
    cleancss -o $f $f
    echo -e "const ${filename^^} = \`\n$(cat $f)\`\n" >> $GO_FILE
done
