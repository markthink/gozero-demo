-- update user set host='%' where user='root';
-- GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';
-- FLUSH PRIVILEGES;

CREATE TABLE `user`
(
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(255) NOT NULL UNIQUE COMMENT 'username',
  `password` varchar(255) NOT NULL COMMENT 'password',
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4