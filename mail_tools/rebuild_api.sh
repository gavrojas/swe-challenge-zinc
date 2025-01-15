#!/bin/bash

# Apagar y remover los contenedores de 'api'
echo "Deteniendo y eliminando contenedores..."
sudo docker compose down api

# Limpiar el sistema Docker (contenedores, imágenes y volúmenes no utilizados)
echo "Eliminando recursos no utilizados..."
sudo docker system prune -a -f

# Reconstruir la imagen de 'api' sin utilizar caché
echo "Reconstruyendo la imagen de 'api'..."
sudo docker compose build api --no-cache

# Levantar el contenedor 'api' en segundo plano
echo "Levantando el contenedor 'api'..."
sudo docker compose up api -d

echo "Proceso completado con éxito."

