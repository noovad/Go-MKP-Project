-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabel terminals
CREATE TABLE terminals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    location TEXT NOT NULL,
    status TEXT NOT NULL
);

-- Tabel transactions
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    card_id TEXT NOT NULL,
    terminal_id_in UUID REFERENCES terminals(id),
    terminal_id_out UUID REFERENCES terminals(id),
    checkin_time TIMESTAMP NOT NULL,
    checkout_time TIMESTAMP,
    fare INTEGER,
    balance_before INTEGER,
    balance_after INTEGER
);

-- Tabel users
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
