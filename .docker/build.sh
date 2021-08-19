#!/bin/bash

[ -z "$1" ] && echo "need project name to proceed." && exit 1 || PROJECT=$1
[ -z "$2" ] && RELEASE_VERSION="latest" || RELEASE_VERSION="$2"
[ -z "$3" ] || INSTALL_VERSION="$3"
set -ex

# prepare debian package so that it can be build
# we assume a formalized directory setup `/usr/local/src/$PROJECT/$PROJECT`
# within /usr/local/src/$PROJECT is containing a debian dir containing the 
# necessary files to build a deb package.
BASE="/usr/local/src/$PROJECT"
DEBIAN_LINK="../debian"
[ -d $BASE/$RELEASE_VERSION/debian ] && DEBIAN_LINK="../$RELEASE_VERSION/debian"
cd $BASE/$PROJECT && ln -s $DEBIAN_LINK .
# in python use find __version__.py
# otherwise we would have to install package and than run pontos-version show
[ -z "$INSTALL_VERSION" ] && \
	INSTALL_VERSION=$(pontos-version show || find . -name "__version__.py" | xargs -I {} grep "__version__ =" {} | sed 's/.*"\(.*\)"/\1/')
MAJOR=$(echo $INSTALL_VERSION | sed 's/\..*//')
FULL="$INSTALL_VERSION~git~$(git rev-parse --short HEAD || echo "unknown" )-$INSTALL_VERSION"
CHANGELOG_DATE=$(date +"%a, %d %b %Y %T %z")
find ./debian/ -type f | xargs -I {} sed -i "s/{{major_version}}/$MAJOR/g" {}
find ./debian/ -type f | xargs -I {} sed -i "s/{{full_version}}/$FULL/g" {}
find ./debian/ -type f | xargs -I {} sed -i "s/{{date}}/$CHANGELOG_DATE/g" {}
# rename files to the correct version
for f in `find ./debian/ -name *{{major_version}}.*`; do
	mv $f ${f//\{\{major_version\}\}/$MAJOR}
done

# build package
PYBUILD_SYSTEM=distutils
PKG_NAME=$(cat debian/changelog | head -n 1 | sed 's/ .*//')
apt-get update
mk-build-deps -irt 'apt-get --no-install-recommends -y'
dpkg-buildpackage -b
apt-get remove -y --autoremove $PKG_NAME-build-deps

# build local repository
cd /usr/local/src
mkdir -p packages/pkg
mv $BASE/*.deb packages/
cd packages
dpkg-scanpackages . /dev/null | gzip -9c > pkg/Packages.gz
cd /usr/local/src
echo "deb [trusted=yes] file:/usr/local/src/packages pkg/" > docker.list
