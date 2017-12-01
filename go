#!/bin/bash

if [[ $1 == "" ]]; then
	echo "Debe introducir el nombre del script"
	cd ~/bin/scripts
else
	if [[ $1 == "-e" ]]; then
		echo "Ubicaciones especiales:"
		dir=""
		names=( "0-bash" "1-trabajos" )
		qnames=( "bash" "work" )
		folders=( "/home/mhyst/Proyectos/bash" "/home/mhyst/Descargas/Trabajos" )

		if [[ -n "$2" ]]; then
			re='^[0-9]+$'
			if [[ "$2" =~ $re ]]; then
				let dirnum=$(( $2 + -1 ))
				dir="${folders[$dirnum]}"
				echo "Ubicación: $dir"
				cd "${dir}"				
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
