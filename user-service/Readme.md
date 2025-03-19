# USER SERVICE APP

This is a User Service application built with Golang, designed to manage user-related operations. The application is containerized using Docker and provides API documentation via Swagger.


## Getting Started


**Prerequisites**


Make sure you have the following installed:

- [Git](https://git-scm.com/)

- [Docker](https://www.docker.com/)

- [Docker Compose](https://docs.docker.com/compose/install/)

**Installation**

1. Clone the repository:


   ```bash
    git clone <github-repo-url>
    cd user-service
   ```


2. Build and run the service using Docker Compose:


   ```bash
    docker compose --build -d
   ```

3. Check if the service is running:


   ```bash
    docker ps
   ```

**Api Documentation**


The API documentation is available via Swagger. Once the service is running, you can access it at

```bash
http://localhost:8080/swagger/index.html
```
















