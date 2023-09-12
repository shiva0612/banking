CREATE DATABASE banking;

\c banking;

CREATE TABLE customers (
    customer_id int PRIMARY KEY,
    name varchar(100) NOT NULL,
    date_of_birth date NOT NULL,
    city varchar(100) NOT NULL,
    zipcode varchar(10) NOT NULL,
    status boolean NOT NULL DEFAULT true
);


CREATE TABLE accounts (
    account_id serial PRIMARY KEY,
    customer_id int NOT NULL,
    opening_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    account_type varchar(10) NOT NULL,
    amount numeric(10,2) NOT NULL,
    status boolean NOT NULL DEFAULT true,
    FOREIGN KEY (customer_id) REFERENCES customers (customer_id)
);


CREATE TABLE transactions (
    transaction_id serial PRIMARY KEY,
    account_id int NOT NULL,
    amount decimal(10,2) NOT NULL,
    transaction_type varchar(10) NOT NULL,
    transaction_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

CREATE TABLE users (
    username varchar(20) NOT NULL PRIMARY KEY,
    password varchar(20) NOT NULL,
    role varchar(20) NOT NULL,
    customer_id int,
    created_on timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE refresh_token_store(
    username var(20) NOT NULL 
    refresh_token varchar(300) NOT NULL PRIMARY KEY,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    FOREIGN KEY (username) REFERENCES users (username)
)

INSERT INTO customers (customer_id,name, date_of_birth, city, zipcode, status) VALUES
    (100,'Krishna', '1985-09-05', 'Chennai', '600001', true),
    (101,'Lakshmi', '1990-03-15', 'Bangalore', '560001', true),
    (102,'Suresh', '1982-11-20', 'Hyderabad', '500001', true),
    (103,'Meena', '1995-07-10', 'Kochi', '682001', true),
    (104,'Rajesh', '1988-12-25', 'Trivandrum', '695001', true);


INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES
    (100, '2020-08-22 10:20:06', 'saving', 10500.50, true),
    (100, '2020-08-09 10:27:22', 'checking', 7200.75, true),
    (101, '2020-08-09 10:35:22', 'saving', 12500.00, true),
    (101, '2020-08-09 10:38:22', 'checking', 8500.25, true);

INSERT INTO users (username,password,role,customer_id) VALUES
    ('krishna_1','krishna_psw','user',100),
    ('laxmi_1','laxmi_psw','user',101),
    ('suresh_1','suresh_psw','user',102),
    ('meena_1','meena_psw','user',103),
    ('rajesh_1','rajesh_psw','user',104),
    ('admin_1','admin_psw','admin',NULL)


