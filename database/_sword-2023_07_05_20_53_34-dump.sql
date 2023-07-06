-- MySQL dump 10.13  Distrib 8.0.30, for macos12.4 (x86_64)
--
-- Host: 127.0.0.1    Database: sword
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tasks`
--

DROP TABLE IF EXISTS `tasks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tasks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `summary` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('pending','finished','in-progress','blocked') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'pending',
  `team_id` bigint DEFAULT NULL,
  `assigned_technician_id` bigint DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `deleted_at` datetime DEFAULT NULL,
  `finished_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tasks_teams_id_fk` (`team_id`),
  KEY `tasks_users_id_fk` (`assigned_technician_id`),
  CONSTRAINT `tasks_teams_id_fk` FOREIGN KEY (`team_id`) REFERENCES `teams` (`id`),
  CONSTRAINT `tasks_users_id_fk` FOREIGN KEY (`assigned_technician_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tasks`
--

INSERT INTO `tasks` VALUES (1,'dadada','xxxasasd','finished',1,1,'2023-07-04 00:32:33',NULL,'2023-07-04 04:58:52'),(2,'dadada','xxxasasd','finished',1,1,'2023-07-04 00:34:51',NULL,'2023-07-04 05:42:40'),(3,'dadada','xxxasasd','finished',1,1,'2023-07-04 00:35:01',NULL,'2023-07-04 05:46:56'),(4,'dadada','xxxasasd','finished',1,NULL,'2023-07-04 00:35:01',NULL,'2023-07-04 05:48:38'),(5,'dadada','xxxasasd','finished',1,2,'2023-07-04 00:35:01',NULL,'2023-07-04 05:49:38'),(6,'dadada','xxxasasd','finished',1,2,'2023-07-04 00:35:01',NULL,'2023-07-04 05:51:21'),(7,'dadada','xxxasasd','finished',2,NULL,'2023-07-04 00:35:01',NULL,'2023-07-04 05:52:51'),(8,'dadada','xxxasasd','pending',2,NULL,'2023-07-04 00:35:01',NULL,NULL),(9,'dadada','xxxasasd','pending',2,NULL,'2023-07-04 00:35:01',NULL,NULL),(10,'dadada','xxxasasd','pending',2,NULL,'2023-07-04 00:35:01',NULL,NULL),(11,'dadada','xxxasasd','finished',2,NULL,'2023-07-04 00:35:01',NULL,'2023-07-04 23:04:15'),(12,'dadada','xxxasasd','finished',1,2,'2023-07-04 00:35:01',NULL,'2023-07-04 23:04:18'),(13,'dadada','xxxasasd','finished',1,NULL,'2023-07-04 00:35:01',NULL,'2023-07-04 23:04:21'),(14,'dadada','xxxasasd','finished',1,2,'2023-07-04 00:35:01',NULL,'2023-07-04 23:05:23'),(15,'dadada','xxxasasd','finished',1,NULL,'2023-07-04 00:35:01',NULL,'2023-07-04 23:06:02'),(16,'dadada','xxxasasd','finished',2,NULL,'2023-07-04 00:35:01',NULL,'2023-07-04 23:07:34'),(17,'dadada','xxxasasd','finished',2,NULL,'2023-07-04 00:35:01',NULL,'2023-07-04 23:08:45'),(18,'dadada','xxxasasd','finished',1,2,'2023-07-04 00:35:01',NULL,'2023-07-04 23:03:41'),(19,'dadada','xxxasasd','finished',1,NULL,'2023-07-04 00:35:01',NULL,'2023-07-04 23:03:09'),(20,'dadada','xxxasasd','finished',1,NULL,'2023-07-04 00:35:01',NULL,'2023-07-04 23:00:03'),(21,'Test title','Test summary','pending',NULL,NULL,'2023-07-05 02:24:56',NULL,NULL),(22,'Test title','Test summary','pending',1,NULL,'2023-07-05 02:26:51',NULL,NULL),(23,'Test title','Test title','pending',1,NULL,'2023-07-05 05:35:31',NULL,NULL);

--
-- Table structure for table `teams`
--

DROP TABLE IF EXISTS `teams`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `teams` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teams`
--

INSERT INTO `teams` VALUES (1,'Alpha Team'),(2,'Beta Team');

--
-- Table structure for table `technicians`
--

DROP TABLE IF EXISTS `technicians`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `technicians` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `team_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `technicians`
--


--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `role` enum('manager','technician') COLLATE utf8mb4_unicode_ci NOT NULL,
  `uuid` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `team_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_pk` (`uuid`),
  KEY `users_teams_id_fk` (`team_id`),
  CONSTRAINT `users_teams_id_fk` FOREIGN KEY (`team_id`) REFERENCES `teams` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

INSERT INTO `users` VALUES (1,'John Doe','manager','af7c1fe6-d669-414e-b066-e9733f0de7a',1),(2,'Elsa Arcadia','manager','08c71152-c552-42e7-b094-f510ff44e9cb',2),(3,'Shelley Adelia','technician','c558a80a-f319-4c10-95d4-4282ef745b4b',1),(4,'Petros Jace','technician','1ad1fccc-d279-46a0-8980-1d91afd6ba67',1),(5,'Gutierre Szilveszter ','technician','5108babc-bf35-44d5-a9ba-de08badfa80a',2),(6,'Hadi Sevastyan','technician','2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb',2);
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-05 20:53:34
