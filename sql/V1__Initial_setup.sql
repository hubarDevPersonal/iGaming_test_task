
CREATE TABLE players (
                         user_id VARCHAR(255) PRIMARY KEY,
                         user_nick VARCHAR(255) NOT NULL,
                         jp_key VARCHAR(255) NOT NULL
);


CREATE TABLE balance (
                         id VARCHAR(255) PRIMARY KEY,
                         user_id VARCHAR(255) NOT NULL,
                         balance BIGINT NOT NULL,
                         denomination INT NOT NULL,
                         currency VARCHAR(50) NOT NULL,
                         max_win BIGINT NOT NULL,
                         game_session_id VARCHAR(255) NOT NULL,
                         FOREIGN KEY (user_id) REFERENCES players(user_id)
);