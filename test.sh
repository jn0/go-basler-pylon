#! /bin/bash

export GOPATH=`pwd`
export PYLON_PATH="${PYLON_PATH:-/opt/pylon4}"
export GENICAM_PATH="${GENICAM_PATH:-$PYLON_PATH/genicam/library/CPP}"
export GENICAM_ROOT_V2_3="${GENICAM_ROOT_V2_3:-$PYLON_PATH/genicam}"

# Dependant headers
CGO_CPPFLAGS="-I$PYLON_PATH/include -I$GENICAM_PATH/include"

# Headers in our project
CGO_CPPFLAGS="${CGO_CPPFLAGS} -include cpp_capture.h"

# Needed for mutex
CGO_CPPFLAGS="${CGO_CPPFLAGS} -std=c++0x"

# Compiler speed ups, don't think O0 takes any effect as O2 is already on the options line
CGO_CPPFLAGS="${CGO_CPPFLAGS} -O0 -fno-var-tracking"

# Debug options
#CGO_CPPFLAGS="${CGO_CPPFLAGS} -Wall -g -H -v"

CGO_LDFLAGS=$PYLON_PATH/lib64/*so
CGO_LDFLAGS=${CGO_LDFLAGS}" "$PYLON_PATH/genicam/bin/Linux64_x64/*so
CGO_LDFLAGS=`echo $CGO_LDFLAGS`

export CGO_CPPFLAGS
export CGO_LDFLAGS

# Wrappers
if [ `which ccache` ]
then
	echo "Using ccache to speed this up."
	alias gcc="ccache $(which gcc)"
	alias g++="ccache $(which g++)"
else
	echo "Faster if you install ccache."
fi

export CC="gcc"
export CXX="g++"

export LD_LIBRARY_PATH="$PYLON_PATH/lib64:$PYLON_PATH/genicam/bin/Linux64_x64:$PYLON_GENICAM_ROOT/bin/Linux64_x64/GenApi/Generic"

cd pylon
if ! go test $@
then
	echo "Env variables used:"
	echo $GOPATH
	echo $CGO_CPPFLAGS
	echo $CGO_LDFLAGS
	echo $CC
	echo $CXX
	echo $LD_LIBRARY_PATH
fi