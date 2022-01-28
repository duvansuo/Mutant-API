# Validate mutant 

### Verificación de ADN:

Se debe realizar un post enviando la secuencia de ADN mediante un HTTP POST con el siguiente JSON:

POST → /mutant/
{
“dna”:["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}

En caso de verificar un mutante, devolvera un HTTP 200-OK, en caso contrario un 403-Forbidden

El metodo ademas de verificar si es mutante o no, verifica que la secuencia de ADN sea Valida y verifica si esta en db. si esta, No continua con la verifica y devuelve el resultado anterior. Si no esta en DB verifica y guarda.

### Estadisticas:

Se debe realizar un get 

Será devuelto un JSON con el siguiente formato de ejemplo:

{"count_mutant_dna":1,"count_human_dna":0,"ratio":1}


## Instrucciones de ejecutarlo localmente:

La API esta totalmente desarrollada en GO, como base de datos en mongoDB.

Como intalar GO, https://golang.org/doc/install

Como instalar MongoDB, 

	Windows
	https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/
	Ubuntu
	https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/
	OSx
	https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/

Para descargar el proyecto tiene que hacer un git clone

	git clone https://github.com/duvansuo/Mutant-API.git

En el proyecto se incluyeron las siguientes librerías externas, que son necesario instalarlas en el entorno donde se vaya a realizar la prueba:

	* go.mongodb.org/mongo-driver/mongo
	* github.com/gorilla/mux (enrutador y despachador de URL para Go)

Para poder instalarlos solamente hay que ejecutar las siguientes lineas en tu consola.

    	- go get go.mongodb.org/mongo-driver/mongo
	- go get github.com/gorilla/mux

Luego puedes dirigirte al directorio donde se encuentra el proyecto y ejecutar:

	go run main.go
  
 por ultimo para correr las pruebas se debe ejecutar el siguiente comando 
 
  	go test ./Testing/ -cover





