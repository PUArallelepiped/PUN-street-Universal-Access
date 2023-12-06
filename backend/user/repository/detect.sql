DROP FUNCTION IF EXISTS Valid_phone(VARCHAR);

CREATE OR REPLACE FUNCTION Valid_phone(phone_number VARCHAR(16))
RETURNS BOOLEAN AS
$Valid_phone$
DECLARE
    e164Pattern VARCHAR(50) := '^(\+8869[0-9]{8}|09[0-9]{8})$';
BEGIN
    IF phone_number ~ e164Pattern THEN
        RAISE NOTICE 'OK';
        RETURN TRUE;
       
    ELSE
        RAISE NOTICE 'Non-standard';
        RETURN FALSE;
        
    END IF;
    
END;
$Valid_phone$
LANGUAGE plpgsql;
----------------------------------------------------------------------------------------------------------------------
DROP FUNCTION IF EXISTS Valid_password(VARCHAR, INT, INT);

CREATE OR REPLACE FUNCTION Valid_password(passwd VARCHAR, min_length INT, required_patterns INT) 
RETURNS BOOLEAN AS 
$$
DECLARE
    count INT := 0;
BEGIN
    IF LENGTH(passwd) < min_length THEN
        RAISE NOTICE 'short password';
        RETURN FALSE;
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
        RAISE NOTICE 'OK';
        RETURN TRUE;
        
    ELSE
        
        RAISE NOTICE 'Password is too simple';
        RETURN FALSE;
    END IF;
END;
$$ LANGUAGE plpgsql;
----------------------------------------------------------------------------------------------------------------------
DROP FUNCTION IF EXISTS Valid_email(VARCHAR);

CREATE OR REPLACE FUNCTION Valid_email(email VARCHAR(20))
RETURNS BOOLEAN AS
$valid_email$
BEGIN
    IF email LIKE '%_@_%' THEN
        RAISE NOTICE 'OK';
        RETURN TRUE;
    ELSE
        RAISE NOTICE 'Invalid email format';
        RETURN FALSE;
    END IF;
END;
$valid_email$ LANGUAGE plpgsql;
----------------------------------------------------------------------------------------------------------------------
DROP FUNCTION IF EXISTS Valid_data(VARCHAR, VARCHAR);

CREATE OR REPLACE FUNCTION Valid_data(phone_number VARCHAR(20), email VARCHAR(20), password VARCHAR)
RETURNS VARCHAR AS
$valid_data$
BEGIN

    IF Valid_email(email) AND Valid_phone(phone_number) AND Valid_password(password,6,3) THEN
        IF EXISTS (SELECT 1 FROM user_data WHERE user_data.phone_number = search_data.phone_number) THEN
        RETURN 'phoneNumber has been registered';
        END IF;

        IF EXISTS (SELECT 1 FROM user_data WHERE user_data.email = search_data.email) THEN
            RETURN 'email has been registered';
        END IF;

        RETURN 'OK';
    END IF;
    RETURN 'InValid Data';
END;
$valid_data$ LANGUAGE plpgsql;

-- SELECT Search_data('091234567','a@gmail.com','orangeOAO123');
