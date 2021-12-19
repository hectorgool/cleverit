# Test MS Beer

## Problema

Bender es fanático de las cervezas, y quiere tener un registro de todas las cervezas que prueba y cómo calcular el precio que necesita para comprar una caja de algún tipo especifico de cervezas. Para esto necesita una API REST con esta información que posteriormente compartir con sus amigos.

## Descripción

Se solicita crear un API REST basándonos en la definición que se encuentra en el archivo: https://bitbucket.org/lgaetecl/microservices-test/src/master/openapi.yaml

### Seguir los siguientes pasos para instalar el api

Para ejecutar el api, dentro de la carpeta del proyecto:
```sh
docker-compose up --build
```
Para parar el api, dentro de la carpeta del proyecto:
```sh
docker-compose down
```

## Funcionalidad

Lista todas las cervezas que se encuentran en el sistema
```sh
GET /beers: 
```
Lista un detalle de una cerveza especifica
```sh
GET /beers/{beerID}: 
```
Entrega el valor que cuesta una caja especifica de cerveza dependiendo de los parámetros ingresados, esto quiere decir que multiplique el precio por la cantidad una vez se homologa la moneda a la que se ingresó por parámetro.
```sh
GET /beers/{beerID}/boxprice: 
```

#### Quanty: 
Cantidad de cervezas a comprar (valor por defecto 6) 

#### Currency: 
Tipo de moneda con que desea pagar, para este caso se recomienda que utilice esta API: https://currencylayer.com 

#### Acceso a la base de datos dede una terminal: 
```sh
docker exec -it cleverit_db mysql -u cleverit -pasdf cervezas
```

#### Monedas soportadas
USD,AUD,CAD,PLN,MXN,CHF

## Ejemplos de endpoints:

Lista todas las cervezas que se encuentran en el sistema:
```sh
http://localhost:8080/beers
```

Lista un detalle de una cerveza especifica:
```sh
http://localhost:8080/beers/1
```

Entrega el valor que cuesta una caja especifica de cerveza dependiendo de los parámetros ingresados, esto quiere decir que multiplique el precio por la cantidad una vez se homologa la moneda a la que se ingresó por parámetro.
El parametro cuantity representa el númro de cajas
El parametro currency representa una de las monedas soportadas que son: USD,AUD,CAD,PLN,MXN,CHF

```sh
http://localhost:8080/beers/1/boxprice
http://localhost:8080/beers/1/boxprice?quanty=3
http://localhost:8080/beers/1/boxprice?currency=MXN
http://localhost:8080/beers/1/boxprice?quanty=3&currency=MXN
```

Ejemplo de un Post (para insertar un registro en la base):
```sh
http://localhost:8080/beers
{
   "id":666,
   "name":"Tecate",
   "brewery":"Brewery",
   "country":"USD",
   "price":22.3,
   "currency":"USD"
}
```

*Nota:
La base de datos tiene 2 registros de prueba el la tabla beer_items:

| id | name  | brewery  |country  | price  |currency  |
| :-----: | :-: | :-: |:-: | :-: |:-: |
| 1 | Golden | Kross | Chile | 10.5 | EUR |
| 2 | Negra | Modelo | Mexico1 | 1 | USD |