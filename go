#!/bin/bash

if [[ $1 == "" ]]; then
    echo "Debe introducir el nombre del script"
    cd ~/bin/scripts
else
    if [[ $1 == "-e" ]]; then
    	echo "Ubicaciones especiales:"
    	dir=""
    	names=( "0-bash" "1-trabajos" "2-exdownloads" "3-pelisd"
	"4-escritos" "5-online")
    	qnames=( "bash" "work" "exdow" "pelisd" "write" "on")
    	folders=( "/home/mhyst/Proyectos/bash" "/home/mhyst/Descargas/Trabajos" "/home/mhyst/Extra/downloads" "/home/mhyst/Descargas/PeliculasD" "/home/mhyst/Proyectos/Escritos" "/home/mhyst/Descargas/PeliculasD/online")
	nfolders=$(( ${#folders[@]} - 1 ))

    	if [[ -n "$2" ]]; then
	    re='^[0-9]+$'
	    if [[ "$2" =~ $re ]]; then
		let dirnum=$(( $2 ))
		if [[ $dirnum > $nfolders ]]; then
		    echo "No existe el directorio especial $dirnum"
		else
		    dir="${folders[$dirnum]}"
		    echo "Ubicación: $dir"
		    cd "${dir}"
		fi
	    else
		let i=0
		for name in "${qnames[@]}"; do
		    if [[ "$name" == "$2" ]]; then
			    dir="${folders[$i]}"
			    echo "Ubicación: $dir"
			    cd "${dir}"
			    break	
		    fi
		    (( i++ ))
		done
		    if [[ $i == ${#qnames[@]} ]]; then
			    echo "Palabra clave no encontrada"
		    fi
	    fi
	else
	    select term in ${names[@]}; do
		if [[ ${term#} > 0 ]]; then
		    data=(${term//-/ })
		    id="${data[0]}"
		    dir="${folders[$id]}"
		    echo "Ubicación: $dir"
		    cd "${dir}"
		    break
		fi
	    done
	fi
    elif [[ "$1" == "-p" ]]; then
	shift
	if [[ -d ~/bin/packages/$1 ]]; then
	    dir="bin/packages/$1"

	    cd
	    cd "${dir}"
	else
	    echo "El directorio $1 no existe"
	fi
    else
	if [[ -d ~/bin/scripts/$1 ]]; then
	    dir="bin/scripts/$1"

	    cd
	    cd "${dir}"
	else
	    echo "El directorio $1 no existe"
	fi
    fi
fi
