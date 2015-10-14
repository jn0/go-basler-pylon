# go-basler-pylon
A minimal set of calls on top of the Basler Pylon4 C++ SDK for linux.

## Compiling

Here is a working install.sh for a project using the modules.

    #! /bin/bash
    
    export GOPATH=`pwd`
    export PYLON_PATH="${PYLON_PATH:-/opt/pylon4}"
    export GENICAM_PATH="${GENICAM_PATH:-/opt/pylon4/genicam/library/CPP}"
    export GO_BASLER_PYLON="${GO_BASLER_PYLON:-$(pwd)/src/github.com/japettyjohn/go-basler-pylon}"
    
    # Dependant headers
    CGO_CPPFLAGS="-I$PYLON_PATH/include -I$GENICAM_PATH/include"
    
    # Headers in our project
    CGO_CPPFLAGS="${CGO_CPPFLAGS} -include $GO_BASLER_PYLON/pylon/cpp_capture.h -include $GO_BASLER_PYLON/pylon/capture.h" 
    
    # Needed for mutex
    CGO_CPPFLAGS="${CGO_CPPFLAGS} -std=c++11"
    
    # Compiler speed ups
    CGO_CPPFLAGS="${CGO_CPPFLAGS} -O0"
    
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
    
    
    cwd=${PWD##*/}
    go install ${cwd}
