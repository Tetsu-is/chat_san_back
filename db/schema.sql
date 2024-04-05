CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    body TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
    updated_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
    CHECK (body <> '')
);

CREATE TRIGGER IF NOT EXISTS trigger_messages_updated_at AFTER UPDATE ON messages
BEGIN
    UPDATE messages SET updated_at = DATETIME('now') WHERE id = NEW.id;
END;

/* user_id INTEGER NOT NULL,
user_name TEXT NOT NULL, */
