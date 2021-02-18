BEGIN;

CREATE TABLE IF NOT EXISTS COMPANIES(

    ID TEXT NOT NULL PRIMARY KEY,
    NAME TEXT NOT NULL,
    COUNTRY TEXT,
    WEBSITE TEXT,
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UPDATED_AT TIMESTAMPTZ,
    IS_ACTIVE BOOLEAN DEFAULT TRUE,
    UNIQUE (NAME)

);

CREATE TABLE IF NOT EXISTS ROLES(

    ID TEXT NOT NULL PRIMARY KEY,
    NAME TEXT NOT NULL,
    DESCRIPTION TEXT NOT NULL,
    ROLE_STATUS BOOLEAN DEFAULT TRUE,
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UPDATED_AT TIMESTAMPTZ,
    UNIQUE (NAME)

);

CREATE TABLE IF NOT EXISTS USERS(

    ID TEXT NOT NULL PRIMARY KEY,
    COMPANY_ID TEXT NOT NULL REFERENCES COMPANIES(ID),
    ROLE_ID TEXT NOT NULL REFERENCES ROLES(ID),
    NAME TEXT,
    LASTNAME TEXT,
    EMAIL TEXT NOT NULL,
    PASSWORD TEXT NOT NULL,
    IS_ACTIVE BOOLEAN DEFAULT TRUE,
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UPDATED_AT TIMESTAMPTZ,
    UNIQUE (COMPANY_ID, EMAIL)

);

CREATE INDEX idx_users_email ON USERS(ID);

INSERT INTO companies (id, name, country, website) VALUES (1, 'first software company', 'USA', 'first.com');

INSERT INTO roles (id, name, description) VALUES 
(1, 'backend Dev', 'Make DB connection, data management and much more'),
(2, 'frontend Dev', 'Make responsive design, ux and much more');

INSERT INTO users (id, company_id, role_id, name, lastname, email, password) VALUES 
(1, 1, 1, 'Guillermo', 'Sanchez', 'guillermo@datacloudgui.com','1234'),
(2, 1, 2, 'Camilo', 'Pulido', 'camilo@datacloudgui.com','123456789');

COMMIT;