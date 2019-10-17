#!/bin/bash

PYLON_PATH=/opt/pylon5

# HW libs available: /opt/pylon5/lib64/libpylon_TL_{bcon,camemu,gige,usb}.so
libs=(
	pylonc 
	pylonbase 
	pylon_TL_usb    # FIX here for other HW
	pylonutility 
	GenApi_gcc_v3_1_Basler_pylon 
	GCBase_gcc_v3_1_Basler_pylon
)

CPATH=${PYLON_PATH}/include/
CGO_CPPFLAGS=-I${CPATH}

CGO_LDFLAGS="-L${PYLON_PATH}/lib64 -L."
CGO_LDFLAGS+=" -Wl,--start-group"
for lib in "${libs[@]}"; do CGO_LDFLAGS+=" -l$lib"; done
CGO_LDFLAGS+=' -Wl,--end-group'

cmd=test                    # defaults to test run
if [ -n "$1" ]; then
    case "$1" in
    -*) ;;                  # pass flags directly
    *)  cmd="$1"; shift;;   # get first non-flag word as command
    esac
fi

set -xe
cd pylon
exec env \
    LD_LIBRARY_PATH="${PYLON_PATH}/lib64" \
    PYLON_PATH="${PYLON_PATH}" \
    CGO_CPPFLAGS="${CGO_CPPFLAGS}" \
    CGO_LDFLAGS="${CGO_LDFLAGS}" \
    go "$cmd" "$@"

# EOF #
