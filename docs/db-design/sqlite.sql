CREATE TABLE stations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE transactions (
    id TEXT PRIMARY KEY,
    card_id TEXT NOT NULL,
    checkin_station_id INTEGER,
    checkout_station_id INTEGER,
    checkin_time DATETIME,
    checkout_time DATETIME,
    fare INTEGER,
    balance_before INTEGER,
    balance_after INTEGER,
    is_synced BOOLEAN DEFAULT 0,
    synced_at DATETIME
);

CREATE TABLE checkin_logs (
    id TEXT PRIMARY KEY,
    card_id TEXT NOT NULL,
    station_id INTEGER,
    checkin_time DATETIME,
    is_valid BOOLEAN,
    error_message TEXT
);
