CREATE OR REPLACE PROCEDURE update_player_balance(user_id_arg text, amount_arg bigint, transaction_type_arg text)
LANGUAGE plpgsql
AS $$
DECLARE
    current_balance bigint;
BEGIN
BEGIN
SELECT balance INTO current_balance FROM balance WHERE user_id = user_id_arg FOR UPDATE;

IF transaction_type_arg = 'CREDIT' OR (transaction_type_arg <> 'CREDIT' AND current_balance + amount_arg > 0) THEN
UPDATE balance SET balance = balance + amount_arg WHERE user_id = user_id_arg;
ELSE
            RAISE EXCEPTION 'Resulting balance would be non-positive for non-credit transactions.';
END IF;

COMMIT;
EXCEPTION
        WHEN OTHERS THEN
ROLLBACK;
RAISE;
END;
END;
$$;
