CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO Users (user_id, email, password_hash, full_name)
VALUES
    (uuid_generate_v4(), 'aliyev@gmail.com', 'hashed_password1', 'Aliyev Alisher'),
    (uuid_generate_v4(), 'bekzod@gmail.com', 'hashed_password2', 'Bekzodov Bekzod'),
    (uuid_generate_v4(), 'jasur@gmail.com', 'hashed_password3', 'Jasurov Jasur'),
    (uuid_generate_v4(), 'nodir@gmail.com', 'hashed_password4', 'Nodirov Nodir'),
    (uuid_generate_v4(), 'ozod@gmail.com', 'hashed_password5', 'Ozodov Ozod'),
    (uuid_generate_v4(), 'gulnora@gmail.com', 'hashed_password6', 'Gulnoraova Gulnora'),
    (uuid_generate_v4(), 'dilshod@gmail.com', 'hashed_password7', 'Dilshodov Dilshod'),
    (uuid_generate_v4(), 'javlon@gmail.com', 'hashed_password8', 'Javlonov Javlon'),
    (uuid_generate_v4(), 'sojida@gmail.com', 'hashed_password9', 'Sojidova Sojida'),
    (uuid_generate_v4(), 'umar@gmail.com', 'hashed_password10', 'Umarov Umar');


INSERT INTO Couriers (courier_id, email, password_hash, full_name)
VALUES
    (uuid_generate_v4(), 'aziz@gmail.com', 'hashed_password1', 'Azizov Aziz'),
    (uuid_generate_v4(), 'mohira@gmail.com', 'hashed_password2', 'Mohiraeva Mohira'),
    (uuid_generate_v4(), 'isom@gmail.com', 'hashed_password3', 'Isomov Isom'),
    (uuid_generate_v4(), 'nozim@gmail.com', 'hashed_password4', 'Nozimov Nozim'),
    (uuid_generate_v4(), 'gulnaz@gmail.com', 'hashed_password5', 'Gulnazova Gulnaz'),
    (uuid_generate_v4(), 'toshmat@gmail.com', 'hashed_password6', 'Toshmatov Toshmat'),
    (uuid_generate_v4(), 'zumrad@gmail.com', 'hashed_password7', 'Zumradova Zumrad'),
    (uuid_generate_v4(), 'bexruz@gmail.com', 'hashed_password8', 'Bexruzov Bexruz'),
    (uuid_generate_v4(), 'yusuf@gmail.com', 'hashed_password9', 'Yusufov Yusuf'),
    (uuid_generate_v4(), 'nasiba@gmail.com', 'hashed_password10', 'Nasibaeva Nasiba');


INSERT INTO UserSessions (session_id, user_id, token, expires_at)
VALUES
    (uuid_generate_v4(), 'ca1ddd41-be35-4123-a8e0-be75ebbad39a', 'some_jwt_token_or_session_id1', NULL),
    (uuid_generate_v4(), 'faf7deb6-af44-4b09-8e8f-4582fed819c0', 'some_jwt_token_or_session_id2', NULL),
    (uuid_generate_v4(), '6dae588b-0339-4205-959e-9b62559660fd', 'some_jwt_token_or_session_id3', NULL),
    (uuid_generate_v4(), 'bbfe8ab6-2df2-4cae-ac28-d13289f90448', 'some_jwt_token_or_session_id4', NULL),
    (uuid_generate_v4(), 'df779633-deee-47c8-9b4e-7a1be57343c1', 'some_jwt_token_or_session_id5', NULL),
    (uuid_generate_v4(), '543cbe40-7630-498b-b5f0-e6b3cce2d6a0', 'some_jwt_token_or_session_id6', NULL),
    (uuid_generate_v4(), '88929e11-e201-4f27-8cd5-acba7162661e', 'some_jwt_token_or_session_id7', NULL),
    (uuid_generate_v4(), '462d860e-88d5-4800-8443-7b53531c794f', 'some_jwt_token_or_session_id8', NULL),
    (uuid_generate_v4(), '9a6b0590-c15e-4e53-8c67-fd0c0b7078ec', 'some_jwt_token_or_session_id9', NULL),
    (uuid_generate_v4(), '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'some_jwt_token_or_session_id10', NULL);
