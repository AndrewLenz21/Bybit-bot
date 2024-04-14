CREATE TABLE dbo.TEST_TABLE_TODO (
	id SERIAL PRIMARY KEY
	, number_test INT NOT NULL
	, desc_test VARCHAR(255)
	, ins_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
DROP TABLE dbo.TEST_TABLE_TODO
--FUNCTION
CREATE OR REPLACE FUNCTION dbo.insert_test_table_todo(number_test INT, desc_test VARCHAR(255))
RETURNS void AS $$
BEGIN
    INSERT INTO dbo.TEST_TABLE_TODO (number_test, desc_test)
    VALUES (number_test, desc_test);
END;
$$ LANGUAGE plpgsql;
--PROCEDURE
CREATE OR REPLACE PROCEDURE dbo.insert_test_table_todo_proc(number_test INT, desc_test VARCHAR(255))
LANGUAGE plpgsql 
AS $$
BEGIN
    INSERT INTO dbo.TEST_TABLE_TODO (number_test, desc_test)
    VALUES (number_test, desc_test);
END;
$$;

DROP PROCEDURE IF EXISTS dbo.insert_test_table_todo_proc;
--TRANSFER TO FUNCTION
CREATE OR REPLACE FUNCTION dbo.insert_test_table_todo_func(number_test INT, desc_test VARCHAR(255))
RETURNS TEXT
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO dbo.TEST_TABLE_TODO (number_test, desc_test)
    VALUES (number_test, desc_test);
    
    -- Succeded
    RETURN 'Operation succeded';
EXCEPTION
    WHEN OTHERS THEN
    -- Error
    RETURN 'Failed: ' || SQLERRM;
END;
$$;

--TRANSFER TO FUNCTION
CREATE OR REPLACE FUNCTION dbo.update_test_table_todo_func(id INT, new_desc_test VARCHAR(255))
RETURNS TEXT
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE dbo.TEST_TABLE_TODO
    SET desc_test = new_desc_test  -- Corrección aquí para usar el parámetro correctamente
    WHERE id = id;
    
    -- Si la operación es exitosa
    RETURN 'Operation succeeded';
EXCEPTION
    WHEN OTHERS THEN
    -- En caso de error
    RETURN 'Failed: ' || SQLERRM;
END;
$$;

SELECT dbo.insert_test_table_todo_func(3, 'Prova 4');

--
SELECT * FROM dbo.TEST_TABLE_TODO

SELECT dbo.insert_test_table_todo('aa', 'Hello world')
CALL dbo.insert_test_table_todo_proc(2, 'Hello world')

CALL dbo.insert_test_table_todo_proc(1, 'Test Directo');


--NEXT STORED PROCEDURE
CREATE OR REPLACE FUNCTION select_from_test_table_todo(number_test_param INT DEFAULT NULL)
RETURNS TABLE (
	id INT,
	number_test INT,
	number_plus INT,
	ins_date TIMESTAMP WITH TIME ZONE
)
LANGUAGE plpgsql AS $$
BEGIN
	RETURN QUERY
	SELECT 
		t.id,
		t.number_test,
		t.number_test + 4 as number_plus,
		t.ins_date
	FROM dbo.TEST_TABLE_TODO t
	WHERE (number_test_param IS NULL OR t.number_test = number_test_param);
END;
$$;

--
CREATE OR REPLACE PROCEDURE select_from_test_table_todo_2(number_test_param INT DEFAULT NULL)
RETURNS TABLE (
	id INT,
	number_test INT,
	number_plus INT,
	ins_date TIMESTAMP WITH TIME ZONE
)
LANGUAGE plpgsql AS $$
BEGIN
    -- Simplemente ejecutar la consulta; en un procedimiento almacenado no se puede retornar directamente
    -- Esta consulta se ejecutaría pero no retornaría resultados al cliente
    SELECT 
        id,
        number_test,
        number_test + 4 as number_plus,
        ins_date
    FROM dbo.TEST_TABLE_TODO
    WHERE (number_test_param IS NULL OR number_test = number_test_param);
END;
$$;

CALL select_from_test_table_todo_2(2)

SELECT * FROM select_from_test_table_todo(2); -- Para filtrar por number_test = 1

