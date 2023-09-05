CREATE TABLE IF NOT EXISTS ecommerece.transaction (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    customer_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INT NOT NULL,
    total_price FLOAT NOT NULL
    );