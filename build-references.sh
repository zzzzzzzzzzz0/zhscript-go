#!/bin/bash
cd `dirname $0`
gg=$GOPATH
if [ "$gg" == "" ]; then
	gg=/opt2/go
fi
if [ "$TERM_PROGRAM" == Apple_Terminal ]; then
	fmt1=-f
	fmt2='%N %Sm'
	x_amd64=darwin_amd64
else
	fmt1=--printf
	fmt2='%n %y\n'
	x_amd64=linux_amd64
fi
echo $fmt
rm $gg/pkg/$x_amd64/github.com/zzzzzzzzzzz0/zhscript-go/*.a
d=$PWD

function st__ {
	stat $fmt1 "$fmt2" $bin
}
function b__ {
	echo
	echo "#### $@"
	cd ../$1-go
	export GOPATH=$PWD:$gg
	echo $GOPATH
	bin=bin/$1
	st__
	go install -v -gcflags "-N -l" ./...
	st__
	cd "$d"
}
b__ zsp
b__ l