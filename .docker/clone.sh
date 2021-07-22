#!/bin/sh
# clone.sh will clone a specific branch based on input parameter
# the first parameter is the name of the project
# the second parameter is the version
[ -z "$1" ] && (echo "need name to clone, exiting" && exit 1) || NAME=$1
[ -z "$2" ] && BRANCH="main" || BRANCH="$2"

[ -z "$GITHUB_DEFAULT" ] && GITHUB_DEFAULT="git://github.com/greenbone/" 
mkdir -p /usr/local/src/$NAME
GIT_PREFIX="$GITHUB_DEFAULT"
URL="$GIT_PREFIX$NAME$SUFFIX"
echo "cloning $URL of $BRANCH"
cd /usr/local/src/$NAME
git clone --depth 1 --single-branch --branch $BRANCH $URL
