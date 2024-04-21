CREATE TABLE dbo.USERS (
    user_id SERIAL PRIMARY KEY
    , username VARCHAR(255) NOT NULL
    , email VARCHAR(255) UNIQUE NOT NULL
    , salt VARCHAR(255) NOT NULL
    , hash VARCHAR(255) NOT NULL
	, flg_active BOOLEAN DEFAULT TRUE
    , ins_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , ins_user INT
    , upt_date TIMESTAMP WITH TIME ZONE
    , upt_user INT
);

SELECT * FROM dbo.USERS

/*
INSERT INTO dbo.USERS 
(
	username
	, email
	, hash
	, salt
)
VALUES 
(
	'Andrew Lenz'
	, 'bybit_admin@pro.com'
	, 'butowski21'
	, 'prova_salt'
)
*/
