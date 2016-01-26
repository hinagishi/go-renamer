#!/bin/bash

dir=`mktemp -d XXXX`
N=0
while [ $N -lt 10 ];
do
    mktemp $dir/XXXX.txt
    N=`expr $N + 1`
done


