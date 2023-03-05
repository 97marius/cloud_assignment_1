# Overview

In this assignment i have developed a REST web application in Golang that provides the client to retrieve information about universities. The REST web services uses in this project are:
 * http://universities.hipolabs.com/
 * https://restcountries.com/


## Deploying the service
 * Open in IDE of choice and run main.go

## Using the service
Use an HTTP client of choice (e.g., Browser, Postman) to point requests to the server URL. GET requests will return the intended structure for either endpoint.



This web service will have three resource root paths: 
```
/unisearcher/v1/uniinfo/
/unisearcher/v1/neighbourunis/
/unisearcher/v1/diag/
```
This service has the default port set to 8080. You access the service via http://localhost:8080/

Assuming the web service should run on localhost, port 8080, the resource root paths would look like this:
```
http://localhost:8080/unisearcher/v1/uniinfo/
http://localhost:8080/unisearcher/v1/neighbourunis/
http://localhost:8080/unisearcher/v1/diag/
```

### Using /unisearcher/v1/uniinfo/
Simply extend the path with the full or partial name of an university

Such as: /unisearcher/v1/uniinfo/<partial or complete university name>

 * Examples: 
   * /unisearcher/v1/uniinfo/norwegian%20university%20of%20science%20and%20technology
   * /unisearcher/v1/uniinfo/agricultural

### Using /unisearcher/v1/neighbourunis/
Extend the path with the name of a country, followed by an optional part where you enter a full or partial name of an university

Such as: /unisearcher/v1/neighbourunis/<country>/<partial or complete university name>

 * Examples:
   * /unisearcher/v1/neighbourunis/norway
   * /unisearcher/v1/neighbourunis/norway/agricultural
   * /unisearcher/v1/neighbourunis/sweden/norwegian%20university%20of%20science%20and%20technology

### Using /unisearcher/v1/diag/
Just go to /unisearcher/v1/diag/ and you will see the diagnostics