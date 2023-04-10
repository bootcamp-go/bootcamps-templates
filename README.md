# Consideraciones referidas a la Base de Datos

Instalar MySQL: brew install mysql Opción2: -arm64 brew install mysql

Verificar estado de MySQL: Ejecutar el comando 'make build-database' desde el root del proyecto (comando declarado en el archivo Makefile)

Correr el comando de creación de la BD: Chequear con 'mysql.server' status para verificar que el servicio se encuentra inicializado. Caso contrario ejecutar comando 'mysql.server start'

EN NINGÚN CASO DEBERÁN PONER PASSWORD
