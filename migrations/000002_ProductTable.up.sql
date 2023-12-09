CREATE TABLE "products" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(10000),
    brand VARCHAR(50),
    sku VARCHAR(50),
    in_stock BOOLEAN DEFAULT true,
    image VARCHAR(500),
    price NUMERIC(12,2) DEFAULT 0.0,
    discounted_price NUMERIC(12,2) DEFAULT 0.0,
    category_id INT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
