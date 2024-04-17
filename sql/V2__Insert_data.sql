INSERT INTO players (user_id, user_nick, jp_key) VALUES
    ('user01', 'NickOne', 'jp001'),
    ('user02', 'NickTwo', 'jp002'),
    ('user03', 'NickThree', 'jp003'),
    ('user04', 'NickFour', 'jp004'),
    ('user05', 'NickFive', 'jp005');


INSERT INTO balance (id, user_id, balance, denomination, currency, max_win, game_session_id) VALUES
    ('wallet01', 'user01', 10000, 2, 'USD', 50000, '16e2ea64-5725-40a4-99cc-c5c0deb8568d'),
    ('wallet02', 'user02', 15000, 2, 'EUR', 75000,'16e2ea64-5725-40a4-99cc-c5c0deb8568d'),
    ('wallet03', 'user03', 20000, 2, 'GBP', 100000,'16e2ea64-5725-40a4-99cc-c5c0deb8568d'),
    ('wallet04', 'user04', 25000, 2, 'BTC', 125000,'16e2ea64-5725-40a4-99cc-c5c0deb8568d'),
    ('wallet05', 'user05', 30000, 2, 'AUD', 150000,'16e2ea64-5725-40a4-99cc-c5c0deb8568d');