# SolarSystem
Solar System Challenge

Estructura:
/cmd 
Comandos para ejecutar la simulacion

/pkg
Código de la simulación

Techs:

Golang
Gin-Gonic 

Endpoint expuesto:

/climateForDay/{nroDia}
Recibe un día > 0 y devuelve las condiciones climaticas para dicho día. 

Ej:
URL: http://localhost:8080/v1/climateForDay/{nroDia}

Correr local:
Desde el archivo cmd/main.go se pueden ejecutar los métodos deseados.