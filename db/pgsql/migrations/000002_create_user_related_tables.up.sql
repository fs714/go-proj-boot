CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    name varchar(64) NOT NULL,
    passhash varchar(64) NOT NULL,
    alias text,
    email varchar(64),
    description text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE TABLE roles (
    id uuid PRIMARY KEY NOT NULL,
    name varchar(64) NOT NULL,
    alias text,
    description text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE TABLE organization (
    id uuid PRIMARY KEY NOT NULL,
    name varchar(64) NOT NULL,
    alias text,
    description text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE TABLE user_role (
    user_id uuid NOT NULL UNIQUE,
    role_id uuid NOT NULL UNIQUE,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE TABLE user_organization (
    user_id uuid NOT NULL UNIQUE,
    org_id uuid NOT NULL UNIQUE,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);