CREATE TABLE node(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    label VARCHAR(20) UNIQUE NOT NULL,
    base_url VARCHAR(50) NOT NULL,
    is_online BOOLEAN DEFAULT 0 NOT NULL,
    last_heartbeat DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)