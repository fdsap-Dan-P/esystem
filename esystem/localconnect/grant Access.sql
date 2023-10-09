-- 1. Grant CONNECT to the database:
GRANT CONNECT ON DATABASE eSystemCentral TO esystem_user;

-- 2. Grant USAGE on schema:
GRANT USAGE ON SCHEMA schema_name TO esystem_user;

-- 3. Grant on all tables for DML statements: SELECT, INSERT, UPDATE, DELETE:
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA schema_name TO esystem_user;

-- 4. Grant all privileges on all tables in the schema:
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA schema_name TO esystem_user;

-- 5. Grant all privileges on all sequences in the schema:
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA schema_name TO esystem_user;

-- 6. Grant all privileges on the database:
GRANT ALL PRIVILEGES ON DATABASE eSystemCentral TO esystem_user;

-- 7. Grant permission to create database:
ALTER USER esystem_user CREATEDB;

-- 8. Make a user superuser:
ALTER USER esystem_user WITH SUPERUSER;

-- 9. Remove superuser status:
ALTER USER esystem_user WITH NOSUPERUSER;

-- Those statements above only affect the current existing tables. To apply to newly created tables, 
-- you need to use alter default. For example:
ALTER DEFAULT PRIVILEGES
FOR USER esystem_user
IN SCHEMA schema_name
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO esystem_user;