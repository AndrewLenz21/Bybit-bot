--USER PERMISSION

-- Otorgar permisos de uso sobre el esquema
GRANT ALL PRIVILEGES ON DATABASE EXAMPLE_DB TO pgGolangBot;
GRANT USAGE ON SCHEMA dbo TO "pgGolangBot";

GRANT USAGE, SELECT, UPDATE ON SEQUENCE dbo.test_table_todo_id_seq TO "pgGolangBot";


-- Otorgar permisos de select, insert, update, delete sobre todas las tablas del esquema
--NOT DELETE
GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA dbo TO "pgGolangBot";

-- Otorgar permisos para ejecutar todos los procedimientos almacenados en el esquema
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA dbo TO "pgGolangBot";
GRANT EXECUTE ON ALL PROCEDURES IN SCHEMA dbo TO "pgGolangBot";

-- Si en el futuro se crean nuevas tablas o funciones, otorgar automáticamente los mismos permisos
--NOT DELETE
ALTER DEFAULT PRIVILEGES IN SCHEMA dbo GRANT SELECT, INSERT, UPDATE ON TABLES TO "pgGolangBot";
ALTER DEFAULT PRIVILEGES IN SCHEMA dbo GRANT EXECUTE ON FUNCTIONS TO "pgGolangBot";
ALTER DEFAULT PRIVILEGES IN SCHEMA dbo GRANT EXECUTE ON PROCEDURES TO "pgGolangBot";

SELECT rolname FROM pg_roles;

/*SEQUENCES*/
SELECT nspname AS schema_name, relname AS sequence_name
FROM pg_class c
JOIN pg_namespace ns ON c.relnamespace = ns.oid
WHERE c.relkind = 'S' AND ns.nspname = 'dbo';

SELECT sequence_name
FROM information_schema.sequences
WHERE sequence_schema = 'dbo';


/*pgGolangBot permiss*/
SELECT nspname AS schema_name, relname AS sequence_name, relacl
FROM pg_class c
JOIN pg_namespace ns ON c.relnamespace = ns.oid
WHERE c.relkind = 'S' AND ns.nspname = 'dbo';


--SOLUTION
CREATE ROLE sequence_access_role NOLOGIN;

--new role sequence_access_role
SELECT 'GRANT USAGE, SELECT ON SEQUENCE ' || quote_ident(nspname) || '.' || quote_ident(relname) || ' TO sequence_access_role;'
FROM pg_class c
JOIN pg_namespace ns ON c.relnamespace = ns.oid
WHERE c.relkind = 'S' AND ns.nspname = 'dbo';


GRANT USAGE, SELECT ON SEQUENCE dbo.tip_order_status_order_status_id_seq TO sequence_access_role;
GRANT USAGE, SELECT ON SEQUENCE dbo.user_orders_id_order_seq TO sequence_access_role;
GRANT USAGE, SELECT ON SEQUENCE dbo.user_positions_id_position_seq TO sequence_access_role;
GRANT USAGE, SELECT ON SEQUENCE dbo.users_user_id_seq TO sequence_access_role;
GRANT USAGE, SELECT ON SEQUENCE dbo.user_trading_conf_id_configuration_seq TO sequence_access_role;
GRANT USAGE, SELECT ON SEQUENCE dbo.tip_trading_perp_pair_pair_id_seq TO sequence_access_role;
GRANT USAGE, SELECT ON SEQUENCE dbo.test_table_todo_id_seq TO sequence_access_role;
--revoke
REVOKE USAGE, SELECT, UPDATE ON SEQUENCE dbo.test_table_todo_id_seq FROM "pgGolangBot";

--give role to pg golang bot
GRANT sequence_access_role TO "pgGolangBot";

DROP TABLE dbo.TEST_TABLE_TODO_2
/**/
CREATE OR REPLACE FUNCTION grant_sequence_permissions_to_role(new_table_name text) RETURNS void AS $$
DECLARE
    sequence_name text;
BEGIN
    -- Look for the sequence name associated with the new table's serial or identity column.
    SELECT pg_get_serial_sequence('dbo.' || new_table_name, column_name) INTO sequence_name
    FROM information_schema.columns
    WHERE table_name = new_table_name AND (column_default LIKE 'nextval(%' OR column_default LIKE 'nextval(''%');

	-- Check if a sequence name was found.
    IF sequence_name IS NOT NULL THEN
        -- Give permission to the role.
        EXECUTE 'GRANT USAGE, SELECT ON SEQUENCE ' || sequence_name || ' TO sequence_access_role;';
    END IF;
END;
$$ LANGUAGE plpgsql;


SELECT pg_get_serial_sequence('"dbo"."test_table_todo"',column_name)
FROM information_schema.columns
WHERE table_name = 'test_table_todo' AND (column_default LIKE 'nextval(%' OR column_default LIKE 'nextval(''%');
--WHERE table_name LIKE 'test_table_todo' AND table_schema = 'dbo';


--VERIFY PERMISSIONS
SELECT
    n.nspname AS schema,
    c.relname AS sequence,
    c.relacl
FROM
    pg_class c
JOIN
    pg_namespace n ON n.oid = c.relnamespace
WHERE
    c.relkind = 'S' -- S = secuencia
    AND n.nspname = 'dbo' -- Esquema específico
    AND c.relacl IS NOT NULL;
	
--FUNCTION CREATED: assign permission
SELECT grant_sequence_permissions_to_role('user_orders');

--FUNCTION FOR REVOKE



