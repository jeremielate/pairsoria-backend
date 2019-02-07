#!/bin/bash
set -e -x

# packages needed to build gosl
sudo apt-get install libopenmpi-dev libhwloc-dev libsuitesparse-dev libmumps-dev gfortran libvtk6-dev python-scipy python-matplotlib dvipng libfftw3-dev libfftw3-mpi-dev libmetis-dev liblapacke-dev libopenblas-dev libhdf5-dev git

# Download gosl submodule
git submodule update --init

# and run build script
cd ./gosl && ./all.bash
