#!/bin/bash

if [[ $1 == "" ]]; then
	echo "Debe introducir el nombre del script"
	exit 1
fi

if [[ -d ~/bin/scripts/$1 ]]; then
	dir="bin/scripts/$1"
	
	cd
	cd ${dir}
else
	echo "El directorio $1 no existe"
fi

