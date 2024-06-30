CREATE TABLE webhooks (
    id VARCHAR(255) PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL Unique,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE endpoints (
    id VARCHAR(255) PRIMARY KEY,
    webhook_id  VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (webhook_id) REFERENCES webhooks(id)
);

CREATE TABLE response (
    id VARCHAR(255) PRIMARY KEY,
    endpoint_id VARCHAR(255) NOT NULL,
    webhook_id VARCHAR(255) NOT NULL,
    body TEXT,
    result TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (endpoint_id) REFERENCES endpoints(id)
);
