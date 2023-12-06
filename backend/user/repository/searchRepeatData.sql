CREATE OR REPLACE FUNCTION search_data(phone_number VARCHAR(20), email VARCHAR(20))
RETURNS BOOLEAN AS
$search_data$
BEGIN
    IF email LIKE '%_@_%' THEN
        RETURN TRUE;
    ELSE
        RETURN FALSE;
    END IF;
END;
$search_data$ LANGUAGE plpgsql;

SELECT search_data('0912345628','a@gmail.com');
