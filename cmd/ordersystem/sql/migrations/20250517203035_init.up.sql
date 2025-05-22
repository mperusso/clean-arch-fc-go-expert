CREATE TABLE orders (
                        id VARCHAR(255) NOT NULL PRIMARY KEY,
                        price DECIMAL(15, 2) NOT NULL, 
                        tax DECIMAL(15, 2) NOT NULL,
                        final_price DECIMAL(15, 2) NOT NULL
);