CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO Products (product_id, name, description, price, image_url)
VALUES
    (uuid_generate_v4(), 'Telefon X', 'Ajoyib smartfon', 1500.00, 'https://example.com/phone-x.jpg'),
    (uuid_generate_v4(), 'Kompyuter Y', 'Qulay va tezkor kompyuter', 2500.00, 'https://example.com/computer-y.jpg'),
    (uuid_generate_v4(), 'Televizor Z', 'Katta ekranli televizor', 1800.00, 'https://example.com/tv-z.jpg'),
    (uuid_generate_v4(), 'Konditsioner A', 'Kuchli isituvchi konditsioner', 1200.00, 'https://example.com/ac-a.jpg'),
    (uuid_generate_v4(), 'Qurilish Materiallari B', 'Yengil va ishonchli materiallar', 500.00, 'https://example.com/building-b.jpg'),
    (uuid_generate_v4(), 'Mebel C', 'Modern va qulay mebel', 800.00, 'https://example.com/furniture-c.jpg'),
    (uuid_generate_v4(), 'Oziq ovqat D', 'Tabiiy va salomat ovqatlar', 300.00, 'https://example.com/food-d.jpg'),
    (uuid_generate_v4(), 'Boshqa E', 'Boshqa mahsulot', 100.00, 'https://example.com/other-e.jpg'),
    (uuid_generate_v4(), 'Noutbuk F', 'Yengil va koproq rivojlanadigan noutbuk', 1800.00, 'https://example.com/laptop-f.jpg'),
    (uuid_generate_v4(), 'Monitor G', 'Katta ekranli monitor', 600.00, 'https://example.com/monitor-g.jpg');


INSERT INTO Carts (cart_id, user_id)
VALUES
    (uuid_generate_v4(), 'ca1ddd41-be35-4123-a8e0-be75ebbad39a'),
    (uuid_generate_v4(), 'faf7deb6-af44-4b09-8e8f-4582fed819c0'),
    (uuid_generate_v4(), '6dae588b-0339-4205-959e-9b62559660fd'),
    (uuid_generate_v4(), 'bbfe8ab6-2df2-4cae-ac28-d13289f90448'),
    (uuid_generate_v4(), 'df779633-deee-47c8-9b4e-7a1be57343c1'),
    (uuid_generate_v4(), '543cbe40-7630-498b-b5f0-e6b3cce2d6a0'),
    (uuid_generate_v4(), '88929e11-e201-4f27-8cd5-acba7162661e'),
    (uuid_generate_v4(), '462d860e-88d5-4800-8443-7b53531c794f'),
    (uuid_generate_v4(), '9a6b0590-c15e-4e53-8c67-fd0c0b7078ec'),
    (uuid_generate_v4(), '2d58ffc0-b064-4c9f-9aa6-081da88e07e7');


-- `cart_id` va `product_id` ustunlarini `Carts` va `Products` jadvallaridan olish
-- `options` ustuni JSONB formatida bo'lishi sababli, uni bosh qoldiramiz yoki xohlasangiz biror ma'lumot kiritishingiz mumkin.

INSERT INTO CartItems (cart_item_id, cart_id, product_id, quantity, options)
VALUES
    ('c1aafc77-32d7-4b89-b888-f04e3d72eb6d', 'd613380d-65cf-40be-b48e-b386a7fd9514', '135f9772-986c-436a-ad27-78d454e03176', 2, '{}'),
    ('f4b77dff-7a22-489d-9bff-3fdf9a6247b1', 'fdcf6b6f-e9db-49e5-be7f-cf429c53b462', 'd92a3181-a1fe-4c6b-a8d9-7c3d94513d35', 1, '{}'),
    ('93f2c646-4a5c-481d-8246-e858ff3a3bfa', 'ccb9f571-40e1-4036-a2a5-a4a071a87e46', '99a58a08-70cc-4c46-8933-3bf6f60a9e86', 3, '{}'),
    ('5dff3c1d-3d45-46b0-a10d-e7321c6f6df0', '977c57b1-302d-423c-a320-cb871b5ea303', '01f61592-b19e-4151-bc65-a35e58ebb3cf', 4, '{}'),
    ('b58f1c1d-34a2-45c8-8b97-5c2a3d1d6c6f', 'b0957566-39cc-4d26-8d21-43178e09cdac', '83779742-02ab-4a49-94d8-73ae43e41a36', 1, '{}'),
    ('7c8a2d66-8bde-4d5b-9214-4d975e8f6936', 'ba331326-5bf2-4f66-9136-c646c476b54f', '1f2b7d8d-0646-48f6-9b5b-d5c65607f22a', 5, '{}'),
    ('20bd43d6-6d83-4141-88a4-c4d5e2d272ff', '27a23f44-1b1c-4e39-b8a1-993df33d4337', '6e53d8cf-efdc-4487-9484-85e88ddd1e7c', 2, '{}'),
    ('19a6c348-fd6b-4c5f-89c3-3a92b9a3627f', '637209da-9aaa-4afa-9859-818fa06b4036', 'b112c051-1c88-4cce-bb67-388afde2228e', 3, '{}'),
    ('28bfb268-02d7-4419-a9d0-7c7a62765cb4', '09d59226-dcd3-4c0c-9b42-0a880af12dba', '176dab25-69a9-4195-9699-3fa52232c8d6', 1, '{}'),
    ('44b6c19f-7a8c-4748-b19f-20456947a3cb', '50657111-3030-4ecb-ba1c-f700d8d48d5d', '161c7497-f534-4130-99da-e4007303ebea', 2, '{}');


-- `P_Orders` jadvali uchun ma'lumotlar kiritish

INSERT INTO P_Orders (order_id, user_id, status, total_price)
VALUES 
    ('ca1ddd41-be35-4123-a8e0-be75ebbad39a', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'pending', 150.00),
    ('faf7deb6-af44-4b09-8e8f-4582fed819c0', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'completed', 250.50),
    ('6dae588b-0339-4205-959e-9b62559660fd', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'pending', 300.75),
    ('bbfe8ab6-2df2-4cae-ac28-d13289f90448', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'cancelled', 175.20),
    ('df779633-deee-47c8-9b4e-7a1be57343c1', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'completed', 220.30),
    ('543cbe40-7630-498b-b5f0-e6b3cce2d6a0', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'pending', 340.45),
    ('88929e11-e201-4f27-8cd5-acba7162661e', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'completed', 410.65),
    ('462d860e-88d5-4800-8443-7b53531c794f', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'cancelled', 290.90),
    ('9a6b0590-c15e-4e53-8c67-fd0c0b7078ec', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'completed', 305.80);



-- OrderItems jadvali ma'lumotlari
INSERT INTO OrderItems (order_item_id, order_id, product_id, quantity, options, created_at, updated_at)
VALUES
    ('1b3b9f19-53b3-467f-8d4a-6d7c22a87e99', 'd6d1aa0d-086c-4c2f-9c4e-c44c9b63dff1', '135f9772-986c-436a-ad27-78d454e03176', 1, '{"color": "black", "size": "XL"}'),
    ('fa0b2b05-29fa-42a8-a8c6-7b567db1d24a', '5fa7d26e-8db0-4b0f-8f56-2bfc23e5be56', 'd92a3181-a1fe-4c6b-a8d9-7c3d94513d35', 2, '{"memory": "16GB", "disk": "1TB SSD"}'),
    ('3a347b23-eb7b-46b3-bf61-3d1df8c6e6f8', '9da67d0d-33e3-4aef-915d-3f1bbceaa5b1', '99a58a08-70cc-4c46-8933-3bf6f60a9e86', 1, '{"resolution": "4K", "screen_size": "65 inches"}'),
    ('2f3c5f9f-6c72-4347-8650-6ac2667be7e7', '7881e0d1-3d1f-4e3e-8a07-8df5b9753409', '01f61592-b19e-4151-bc65-a35e58ebb3cf', 3, '{"type": "split", "capacity": "12000 BTU"}'),
    ('f3b99166-7a4c-4b16-96b4-4c7f4e00cfa1', '1d2b7328-f0e0-4d0d-bb84-8dfc6e1e1b15', '83779742-02ab-4a49-94d8-73ae43e41a36', 1, '{"quantity": "10 meters", "material": "steel"}'),
    ('eae8d76a-3fb8-4b1f-a8fe-7b849b9ea5cc', '4ab8e02d-efc1-40a5-9537-7022bfa87d70', '1f2b7d8d-0646-48f6-9b5b-d5c65607f22a', 2, '{"color": "brown", "style": "modern"}'),
    ('ab5b6a64-0303-4f4a-aac0-3058e361ed1b', 'ae0986b6-1c2e-4687-bf60-229b0fd9e4a7', '6e53d8cf-efdc-4487-9484-85e88ddd1e7c', 1, '{"type": "organic", "expiration_date": "2025-01-01"}'),
    ('80b34499-eed6-48b7-8d1e-902e7cc30a8b', '8a9e0d0b-1d6e-4b15-9f7d-3b9fd2b0a1a9', 'b112c051-1c88-4cce-bb67-388afde2228e', 4, '{"category": "miscellaneous", "description": "miscellaneous item"}'),
    ('a9e777f4-3eac-4ff4-b4ed-2a68e63c0d14', '3f8c8d2d-8e1a-4e20-8d4d-3f5fa3b0a6b2', '176dab25-69a9-4195-9699-3fa52232c8d6', 1, '{"ram": "8GB", "processor": "Intel Core i7"}'),
    ('1d3d8f49-898f-4f26-8826-13c43058f3b7', '2c1e8fbb-5f5f-4d5e-b39e-481d6328e8ef', '161c7497-f534-4130-99da-e4007303ebea', 2, '{"size": "small", "color": "blue"}');



INSERT INTO OrderRecommendations (recommendation_id, order_id, courier_id, recommended_at, status)
VALUES
    ('c5f5b30a-d1e8-46a7-9376-3ee1bb7a4f80', 'd6d1aa0d-086c-4c2f-9c4e-c44c9b63dff1', '0b48a25c-e01a-4e42-a3e7-b26208ff4670', '2024-07-05 18:01:16.287312+05', 'pending'),
    ('5a41a799-5243-4a90-9514-13a5c8a25a7c', '5fa7d26e-8db0-4b0f-8f56-2bfc23e5be56', 'e19c9b99-ca15-4aac-91f2-1545a8d1e20a', '2024-07-05 18:01:16.287312+05', 'completed'),
    ('e3b143c2-6b7f-4c26-87db-6060cc8be874', '9da67d0d-33e3-4aef-915d-3f1bbceaa5b1', 'fc67098c-d8c1-4bfd-be8b-716ddb3301e0', '2024-07-05 18:01:16.287312+05', 'pending'),
    ('a5589c1a-3a03-4e8a-bc71-ef5697f63a63', '7881e0d1-3d1f-4e3e-8a07-8df5b9753409', 'e23a658f-c311-4a63-9565-0f966243fdcf', '2024-07-05 18:01:16.287312+05', 'cancelled'),
    ('3ec555ea-5e2d-41ef-8a7b-39b6847a7db2', '1d2b7328-f0e0-4d0d-bb84-8dfc6e1e1b15', 'bd92467b-3299-4b03-8bc9-6aefcd2c15ec', '2024-07-05 18:01:16.287312+05', 'completed'),
    ('6a1a56b1-e4f4-4c63-aa3f-29e6275d2d05', 'ae0986b6-1c2e-4687-bf60-229b0fd9e4a7', 'efb24d0b-8950-49fd-91b8-832f5811fcb9', '2024-07-05 18:01:16.287312+05', 'completed'),
    ('d196d6f5-79a5-460f-90bc-36bbaaa6a142', '8a9e0d0b-1d6e-4b15-9f7d-3b9fd2b0a1a9', '726ea834-dcc3-4e6c-9eba-a0c8bcb25f68', '2024-07-05 18:01:16.287312+05', 'cancelled'),
    ('51fdaa03-6f5b-4e1d-b90a-af4641c4c1a3', '3f8c8d2d-8e1a-4e20-8d4d-3f5fa3b0a6b2', '87f8e776-e679-4770-bc95-cf7d31bfb6ff', '2024-07-05 18:01:16.287312+05', 'completed'),
    ('79fb5b2d-4d3f-4f2f-8bb7-11efbce2f151', '2c1e8fbb-5f5f-4d5e-b39e-481d6328e8ef', '36ec7595-490a-4f62-97c3-f0609494237a', '2024-07-05 18:01:16.287312+05', 'pending');