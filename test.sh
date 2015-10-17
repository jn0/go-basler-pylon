#! /bin/bash

export GOPATH=`pwd`
export PYLON_PATH="${PYLON_PATH:-/opt/pylon4}"
export GENICAM_PATH="${GENICAM_PATH:-$PYLON_PATH/genicam/library/CPP}"
export GENICAM_ROOT_V2_3="${GENICAM_ROOT_V2_3:-$PYLON_PATH/genicam}"

# Dependant headers
CGO_CPPFLAGS="-I$PYLON_PATH/include -I$GENICAM_PATH/include"

# Headers in our project
CGO_CPPFLAGS="${CGO_CPPFLAGS} -include capture.h"

# Compiler speed ups, don't think O0 takes any effect as O2 is already on the options line
CGO_CPPFLAGS="${CGO_CPPFLAGS} -O0"

# Debug options
# CGO_CPPFLAGS="${CGO_CPPFLAGS} -Wall -g -H -v"

# Linking libs
CGO_LDFLAGS=$PYLON_PATH/lib64/*so
CGO_LDFLAGS=${CGO_LDFLAGS}" "$PYLON_PATH/genicam/bin/Linux64_x64/*so
CGO_LDFLAGS=`echo $CGO_LDFLAGS`

LD_LIBRARY_PATH="$PYLON_PATH/lib64:$PYLON_PATH/genicam/bin/Linux64_x64:$PYLON_GENICAM_ROOT/bin/Linux64_x64/GenApi/Generic"


# Wrappers
if [ `which ccache` ]
then
	if [ `which ccache-gcc` ]
	then
		CC="ccache-gcc"
	fi

	if [ `which ccache-g++` ]
	then
		CXX="ccache-g++"
	fi
else
	echo "Faster if you install ccache."
fi

export CC CXX LD_LIBRARY_PATH CGO_CPPFLAGS CGO_LDFLAGS

cd pylon
go test $@