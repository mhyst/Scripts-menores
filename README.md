# Scripts-menores
Aquí iré añadiendo pequeños scripts que uso frecuentemente pero que no merezca la pena crear para ellos un repositorio aparte.

## Scripts que me facilitan trabajar en scripts

### prod

Cada uno tiene sus manías a la hora de trabajar en nuevos scripts. Generalmente hay una carpeta que contiene otra carpeta por cada script que hayas escrito. En mi caso esa carpeta se llama scripts y está en ~/bin. Es decir, en la carpeta bin de mi directorio personal. Pues bien, se me ocurren cambios de vez en cuando y muchas veces olvido hacer copia de los scripts que he modificado a la carpeta ~/bin, que es la única que se encuentra en el PATH. Cuando hago uso de aquel script y veo que no funciona como creía, es cuando me acuerdo de los cambios que hice y de si lo puse en producción en la carpeta ~/bin. Pa aclararlo normalmente hago uso de diff y si veo que hay diferencias, entonces sobreescribo la versión del script de ~/bin con la de su carpeta de ~/bin/scripts. Pero claro, si tienes en rueda unos pocos scripts que usas a diario, la cosa se torna tediosa. Fue en ese momento cuando pensé automatizar el proceso. Y para eso es este script. Para que funcione te tienes que situar en la carpeta donde esté el archivo maestro (en mi caso en ~/bin/scripts/carpeta-del-script-que-sea). Una vez en el directorio deseado, escribiremos prod, espacio y el nombre del archivo que contiene el script. Al darle a intro, prod, realiza algunas comprobaciones y en base a ellas decide si tiene que poner el script en producción o no. Además, avisa al usuario de su decisión y realiza todas las operaciones necesarias. Una maravilla, aunque es realmente sencillo.

### go

Otra cosa que me incomodaba era tener que marcar toda la ruta en el comando cd cuando quera trabajar en, o consultar, un script cualquiera. Puesto que todos los archivos maestros de mis scripts están ubicados en ~/bin/scripts/*, dependiendo del nombre del script, tenía que escribir rutas de a partir de once o doce caracteres. Para evitar eso, escribí "go". Si quieres ir a la carpeta de tu script, no tienes mas que poner "go" y el nombre del script al que quieras ir. Y "go" se encarga de llevarte. Debes tener cuidado en una cosa. En mi caso, procuro que las carpetas que contienen los scripts se llamen igual que el script al que contienen. De esa forma nunca olvido el nombre. "go" no lleva la cuenta de los scripts que creas, sino que te lleva a ~/bin/scripts/la-carpeta-del-script-que-buscas. Si el nombre que le proporcionas no corresponde con una carpeta existente dentro de ~/bin/scripts, te avisará del problema y no te llevará a ninguna parte. Aún así, lo bueno de go es que puedes invocarlo (si lo pones en una carpeta como ~/bin que está en el PATH) desde cualquier ubicación de cualquiera de tus discos físicos o de red. Yo lo invoco habitualmente para ir a un script desde la ubicación de otro.
