# GOData

Spesification :

- Docker 
- Go
- Mysql

## INSTALLATION PROCCES ##

1. Create the databases 
        - CREATE DATABASE `GOData` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
        AND THE TABLE ,

CREATE TABLE `posts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `content` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

2. Run the Docker-Compose using
        docker-compose up
3. Test your connection to mysql at docker , host localhost port 3308
4. Run the app
        go run main.go


#API ENDPOINT 
CREATE : 
        /POST
        {
                "title" : "Your Title",
                "Content" : "Your Content"
        }
