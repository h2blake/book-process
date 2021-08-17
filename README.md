CREATE USER 'hblake'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON books . * TO 'hblake'@'localhost';



CREATE TABLE `books`.`dogs` (
  `first_name` VARCHAR(20) NOT NULL,
  `last_name` VARCHAR(20) NOT NULL,
  `id` INT NOT NULL,
  `email` VARCHAR(245) NOT NULL,
  PRIMARY KEY (`id`));
