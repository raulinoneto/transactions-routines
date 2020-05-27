USE transactions_routines;
CREATE TABLE accounts
(
    id              INT NOT NULL AUTO_INCREMENT,
    document_number INT NOT NULL UNIQUE,
    is_blocked      INT(1) ZEROFILL,
    PRIMARY KEY (id)
);

CREATE TABLE transactions
(
    id             INT    NOT NULL AUTO_INCREMENT,
    account_id     INT    NOT NULL,
    amount         DOUBLE NOT NULL,
    operation_type INT NOT NULL,
    event_date     datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts (id)
);