DROP FUNCTION IF EXISTS is_valid_password(character varying, integer, integer);

CREATE OR REPLACE FUNCTION is_valid_password(passwd VARCHAR, min_length INT, required_patterns INT) 
RETURNS VARCHAR AS 
$$
DECLARE
    count INT := 0;
BEGIN
    IF LENGTH(passwd) < min_length THEN
        RETURN 'short password';
    END IF;

    IF passwd ~ '[a-z]' THEN
        count := count + 1;
    END IF;

    IF passwd ~ '[A-Z]' THEN
        count := count + 1;
    END IF;

    IF passwd ~ '[0-9]' THEN
        count := count + 1;
    END IF;

    IF passwd ~ '[~!@#$%^&*]' THEN
        count := count + 1;
    END IF;

    IF count >= required_patterns THEN
        RETURN 'OK';
    ELSE
        RETURN 'Password is too simple';
    END IF;
END;
$$ LANGUAGE plpgsql;


SELECT is_valid_password('orangeOAO123', 6, 3) AS result;
