#!/usr/bin/env bash

wget https://github.com/zinclabs/zinc/releases/download/v0.3.6/zinc_0.3.6_Linux_x86_64.tar.gz
tar -xf zinc_0.3.6_Linux_x86_64.tar.gz

unzip ./data_example/bd_data.zip

mkdir data

setsid sh -c 'ZINC_FIRST_ADMIN_USER=admin ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 ./zinc & ./interfaz & wait'
