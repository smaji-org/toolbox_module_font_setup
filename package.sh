#!/bin/sh

project=font_setup

data=$1

mkdir -p package/${OS}/${ARCH}/${project}
cd package/${OS}/${ARCH}/${project}
ln -s ../../../../build/${OS}/${ARCH}/* ./

if [ "${data#/}" = "$data" ]; then
    ln -s ../../../../$data ./
else
    ln -s $data ./
fi


cd ..
tar chzf ${project}.tgz ${project}
rm -rf ${project}
