# Introduction to RESTful API

#### Outline 
* Dapat mengerti apa itu API
* Penjelasan tentang API
* REST
* JSON
* HTTP Response Code
* Open API
* Introduction to Postman
* Package net/http 

## API
API adalah sekumpulan fungsi dan prosedur yang memungkinkan pembuatan aplikasi yang mengakses fitur atau data suatu sistem operasi, aplikasi, atau layanan lainnya

* Cara API Bekerja

```
Client request to server
and
server response to Client
```


## Rest (Representational State Transfer)
```
Client request to server
and
server response to Client
```

```
Use : HTTP Protocol
example : https://www.instagram.com/graphql/query
```
Request & Response format :
* JSON
* XML
* SOAP

HTTP request method :
* GET
* POST 
* PUT 
* DELETE 
* HEAD 
* OPTION
* PATCH

## JSON (JavaScript Object Notation)

Example :
```
"name" : "Aimar Rizki",
"umur" : 21
"student" : true,
"hobi" : [
    "belajar",
    "mengaji"
]
```

HTTPS Response Code :
* 200 : OK
* 201 : Created
* 400 : Bad Request
* 404 : Not Found
* 401 : Unauthorized 
* 405 : Method Not Allowed 
* 500 : Internal Server Error

API Tools :
* Katalon
* Apache JMeter
* SoapUI
* Postman

## Postman
Postman adalah klien HTTP untuk menguji layanan web, tukang pos memudahkan pengujian, pengembangan, dan mendokumentasikan API dengan memungkinkan pengguna dengan cepat menyusun permintaan HTTP yang sederhana dan kompleks

## Package Go "net/http" & "encoding/json"

#### Package net/http
Menyediakan implementasi Klien HTTP dan server

#### Package encoding /json
Berisi banyak fungsi untuk operasi json berikut:
* decode json to object struct
* decode json to map [string] interface {} & interface{}
* decode array json to array object
* encode object to json string
