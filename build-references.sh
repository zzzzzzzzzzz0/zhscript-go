#!/bin/bash
cd `dirname $0`
gg=$GOPATH
if [ "$gg" == "" ]; then
	gg=/opt2/go
fi
rm $gg/pkg/linux_amd64/github.com/zzzzzzzzzzz0/zhscript-go/*.a
d=$PWD

function st__ {
	stat --printf='%n %y %z\n' $bin
}
function b__ {
	echo
	echo "#### $@"
	cd ../$1-go
	export GOPATH=$PWD:$gg
	echo $GOPATH
	bin=bin/$1
	st__
	/usr/bin/go install -v -gcflags "-N -l" ./...
	st__
	cd "$d"
}
b__ zsp
b__ l