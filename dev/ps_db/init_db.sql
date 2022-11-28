CREATE TABLE IF NOT EXISTS TOCs(
        id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        webPage_id TEXT,
        version TEXT NOT NULL,
        TOC_key TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS users(
        id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        fingerprint TEXT
);

CREATE TABLE IF NOT EXISTS consents (
        id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        TOC_id INT,
        user_id INT,
        answer BOOL NOT NULL,
        anonym_fingerprint TEXT,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        expires_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP + interval '5' day,

    CONSTRAINT fk_TOC
        FOREIGN KEY(TOC_id)
            REFERENCES TOCs(id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);