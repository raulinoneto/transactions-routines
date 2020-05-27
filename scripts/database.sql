USE transactions_routines;
CREATE TABLE accounts
(
    id              INT NOT NULL AUTO_INCREMENT,
    document_number INT NOT NULL,
    is_blocked      INT(8) ZEROFILL,
    PRIMARY KEY (id)
);

CREATE TABLE transactions
(
    id             INT    NOT NULL AUTO_INCREMENT,
    account_id     INT    NOT NULL,
    amount         DOUBLE NOT NULL,
    operation_type INT,
    event_date     datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts (id)
);