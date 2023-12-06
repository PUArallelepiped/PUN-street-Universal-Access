CREATE OR REPLACE FUNCTION Valid_phone(phone_number VARCHAR(16))
RETURNS BOOLEAN AS
$Valid_phone$
DECLARE
    e164Pattern VARCHAR(50) := '^(\+8869[0-9]{8}|09[0-9]{8})$';
BEGIN
    RETURN phone_number ~ e164Pattern;
END;
$Valid_phone$
LANGUAGE plpgsql;

-- SELECT Valid_phone('098765432'); 

    