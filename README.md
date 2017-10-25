# widgets-spa

This is an attempt to create a simple API, written in GOlang, for the Widgets Single Page App Demo API available at https://github.com/RedVentures/widgets-spa. It also uses Docker containers for MySQL and GOlang in order to run the API. It is worth noting it was my first time using the Go programming language.



#### Some Docker Commands Used - MySQL
1) ```sudo docker run --name mysqldb -e MYSQL_ROOT_PASSWORD=secret -d mysql:5.7``` to create and run MySQL container. As seen, the name assigned to the container was "mysqldb" and password was "secret"

2) ```sudo docker inspect mysqldb | grep IPAddr ``` to get the IP Address of the MySQL server. 
An example return of this command could be 
```
"SecondaryIPAddresses": null,
   "IPAddress": "172.17.0.2",
          "IPAddress": "172.17.0.2", 
 ```
 showing that the IP Address of the MySQL server **172.17.0.2**
 
 3) ```mysql -uroot -psecret -h 172.17.0.2 -P 3306 ``` to access the MySQL server from the physical host. To verify if there is any running container, use ``` sudo docker ps ```

#### Some Docker Commands Used - Go
1) With MySQL up and running, we need a Dockerfile to for building a Docker image to install and run our Go server in a Docker container. To build the Docker image, we run the following command: ```sudo docker build -t widgets-spa-master-pablo . ``` It creates the image from widgets-spa-master-pablo directory that buidls Docker image

2) ``` sudo docker run -d -p 8000:8080 widgets-spa-master-pablo ``` to run the Docker image. As shown, it makes the container get accessed at port 8000. The container's port is 8080 in this case.

3) Open the browser and goes to ```http://localhost:8000``` to see it running
