#!/bin/bash

dir=`mktemp -d XXXX`
os=`uname`
N=0
while [ $N -lt 10 ];
do
    if [ $os = "Darwin" ];
    then
        mktemp $dir/XXXXX
    else
        mktemp $dir/XXXX.txt
    fi
    N=`expr $N + 1`
done

mktemp $dir/.XXXX
mktemp $dir/.XXXX

