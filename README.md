# CoffeShop

## Introduction

This project implements a coffeshop API using a microservices architecture. The following diagram demonstrates the microservice architechture.

![MicroServices](/assets/architecture.png)

## How to run this project

- Run the currency service

```
cd currency
go run main.go
```

- Test it with grpcCurl

```
grpcurl --plaintext -d '{"Base":"USD","Destination":"INR"}' localhost:9092 currency.Currency.GetRate
```

- Command to install grpcurl `go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest`
- Run the products-api service

```
cd products-api
go run main.go
```

## API Testing Using CUrl


### GET

- `curl localhost:9090/products`
```
[{"id":1,"name":"Masala Chia","description":"Spicy Tea","price":20,"sku":"mas-ala-chi"},{"id":2,"name":"Espresso","description":"Short and strong coffee without milk","price":1.99,"sku":"fjd34"},{"id":3,"name":"Chiya","description":"Fried Tea","price":5,"sku":"chi-yaa-bag"}]
```
- `curl localhost:9090/products/1`
``` 
*   Trying 127.0.0.1:9090...
* Connected to localhost (127.0.0.1) port 9090 (#0)
> GET /products/1 HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.81.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 07 Oct 2022 09:32:29 GMT
< Content-Length: 87
< Content-Type: text/plain; charset=utf-8
<
```
``` 
{"id":1,"name":"Masala Chia","description":"Spicy Tea","price":20,"sku":"mas-ala-chi"}
* Connection #0 to host localhost left intact
```
- `curl -v localhost:9090/products/1?currency=INR`

```
*   Trying 127.0.0.1:9090...
* Connected to localhost (127.0.0.1) port 9090 (#0)
> GET /products/1?currency=INR HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.81.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 07 Oct 2022 09:36:29 GMT
< Content-Length: 92
< Content-Type: text/plain; charset=utf-8
< 
```
```
{"id":1,"name":"Masala Chia","description":"Spicy Tea","price":1621.23,"sku":"mas-ala-chi"}
```

- *Currency Service Response*

```
2022-10-07T15:06:29.540+0530 [INFO]  handle request for GetRate: base=EUR dest=INR
```

### PUT

- `curl -v -X PUT localhost:9090/products/1 -d '{"name":"Masala Chia", "description":"Spicy Tea","price":20,"sku":"mas-ala-chi"}'`

```
*   Trying 127.0.0.1:9090...
* Connected to localhost (127.0.0.1) port 9090 (#0)
> PUT /products/1 HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.81.0
> Accept: */*
> Content-Length: 80
> Content-Type: application/x-www-form-urlencoded
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 07 Oct 2022 09:25:40 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
```
- `curl localhost:9090/products`
```
[{"id":1,"name":"Masala Chia","description":"Spicy Tea","price":20,"sku":"mas-ala-chi"},{"id":2,"name":"Espresso","description":"Short and strong coffee without milk","price":1.99,"sku":"fjd34"},{"id":3,"name":"Chiya","description":"Fried Tea","price":5,"sku":"chi-yaa-bag"}]
```
### POST

- `curl -v localhost:9090/products -d '{"id":4,"name":"Chiya", "description":"Fried Tea","price":5,"sku":"chi-yaa-bag"}'`
```
*   Trying 127.0.0.1:9090...
* Connected to localhost (127.0.0.1) port 9090 (#0)
> POST /products HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.81.0
> Accept: */*
> Content-Length: 80
> Content-Type: application/x-www-form-urlencoded
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 07 Oct 2022 09:17:29 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
```

- `curl localhost:9090/products`
```
[{"id":1,"name":"Latte","description":"Frothy milky coffee","price":2.45,"sku":"abc323"},{"id":2,"name":"Espresso","description":"Short and strong coffee without milk","price":1.99,"sku":"fjd34"},{"id":3,"name":"Chiya","description":"Fried Tea","price":5,"sku":"chi-yaa-bag"}]
```

### DELETE

- `curl -v -X DELETE localhost:9090/products/3`
```
*   Trying 127.0.0.1:9090...
* Connected to localhost (127.0.0.1) port 9090 (#0)
> DELETE /products/3 HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.81.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 07 Oct 2022 09:39:48 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
```
- `curl localhost:9090/products`
```
[{"id":1,"name":"Masala Chia","description":"Spicy Tea","price":20,"sku":"mas-ala-chi"},{"id":2,"name":"Espresso","description":"Short and strong coffee without milk","price":1.99,"sku":"fjd34"}]
```

## API Documentation Using Swagger OPENAPI

- Check out the live documentation at `localhost:9092/docs` after running the project locally

![1](/assets/doc1.png)
![2](/assets/doc2.png)
![3](/assets/doc3.png)
![4](/assets/doc4.png)
![5](/assets/doc5.png)
![6](/assets/doc6.png)
![7](/assets/doc7.png)

- To Generate/Modify the documentation locally, make the required modification. Then run the following:

```
cd product-api
make swagger
```

- Swagger Installation [Refs](https://goswagger.io/install.html)