CREATE TABLE `bom` (
    `id` varchar(25) NOT NULL,
    `name` varchar(100) NOT NULL,
    `parent_id` varchar(25),
    `desc` varchar(225),
    `net` float(10, 2) NOT NULL,
    `inv` varchar(4) NOT NULL 
    
    
    FOREIGN KEY (parent_id) REFERENCES bom(id) 
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
