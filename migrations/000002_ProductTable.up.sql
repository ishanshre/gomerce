CREATE TABLE "products" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(10000),
    brand VARCHAR(50),
    sku VARCHAR(50),
    stock BOOLEAN DEFAULT true,
    image VARCHAR(500),
    price INT DEFAULT 0,
    discounted_price INT,
    category_id INT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
