-- -- `Couriers` jadvali uchun ma'lumotlar kiritish
-- INSERT INTO Couriers (courier_id, full_name, phone_number, email, status)
-- VALUES
--     ('b1a3e980-1d1b-4d0f-9e4f-2b0b4d6b6f1a', 'Ali Valiyev', '+998901234567', 'ali.valiyev@example.com', 'available'),
--     ('7b8c2b40-39b1-4c3a-8a64-2b7f4e5f1a2b', 'Salim Karimov', '+998901234568', 'salim.karimov@example.com', 'unavailable'),
--     ('e2d3a880-5e6f-4d0a-9a7b-1d0f6b7c8e9d', 'Madina Akbarova', '+998901234569', 'madina.akbarova@example.com', 'on_delivery'),
--     ('c4d5a690-7b8c-4e3a-8c7b-3e5f1a9b2c3d', 'Rustam Qodirov', '+998901234570', 'rustam.qodirov@example.com', 'available'),
--     ('f6e7b980-8c7b-4d2a-9c8d-4e7f2b0f3a4e', 'Zebo Rahimova', '+998901234571', 'zebo.rahimova@example.com', 'unavailable'),
--     ('a8b9c1d0-9e8f-4c3a-8f9d-5b9f3a0b4c5e', 'Aziz Eshonov', '+998901234572', 'aziz.eshonov@example.com', 'on_delivery'),
--     ('9f0a1b2c-0e9f-4d3a-8f9e-6a1f4b0c5d7e', 'Sevinch Karimova', '+998901234573', 'sevinch.karimova@example.com', 'available'),
--     ('b2c3d4e0-1e0f-4a3a-8f9f-7b2c5d1f8e9e', 'Dilshod Turgunov', '+998901234574', 'dilshod.turgunov@example.com', 'unavailable'),
--     ('d5e6f7a0-2e1f-4b3a-8f9f-8b3d6e2f9e1f', 'Ravshan Xasanov', '+998901234575', 'ravshan.xasanov@example.com', 'on_delivery'),
--     ('e0f1a2b0-3f2f-4c3a-8f9f-9c4e7f2f0a1b', 'Laylo Bozorova', '+998901234576', 'laylo.bozorova@example.com', 'available');


-- DO $$ 
-- BEGIN
--     -- Agar qiymatlar mavjud bo'lmasa, 'available', 'unavailable' va 'on_delivery' qiymatlarini qo'shamiz
--     IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'courier_status') THEN
--         CREATE TYPE courier_status AS ENUM ('available', 'unavailable', 'on_delivery');
--     END IF;
--     -- Agar 'courier_status' mavjud bo'lsa, faqat mavjud bo'lmagan qiymatlarni qo'shamiz
--     IF NOT EXISTS (SELECT 1 FROM pg_enum WHERE enumlabel = 'available' AND enumtypid = 'courier_status'::regtype) THEN
--         ALTER TYPE courier_status ADD VALUE 'available';
--     END IF;
--     IF NOT EXISTS (SELECT 1 FROM pg_enum WHERE enumlabel = 'unavailable' AND enumtypid = 'courier_status'::regtype) THEN
--         ALTER TYPE courier_status ADD VALUE 'unavailable';
--     END IF;
--     IF NOT EXISTS (SELECT 1 FROM pg_enum WHERE enumlabel = 'on_delivery' AND enumtypid = 'courier_status'::regtype) THEN
--         ALTER TYPE courier_status ADD VALUE 'on_delivery';
--     END IF;
-- END $$;


-- INSERT INTO Orders (order_id, customer_id, order_details, delivery_address, payment_status)
-- VALUES 
--     ('d3b4a4b3-7841-4a6e-b0c7-54e450dc8a8d', 'ca1ddd41-be35-4123-a8e0-be75ebbad39a', '5 ta olma va 2 ta anor', 'Toshkent shahri, Mirzo Ulug`bek tumani, Alisher Navoi kochasi 15-uy', 'pending'),
--     ('fbf24974-3cf1-44f7-93e7-938b6d8dc5b0', 'faf7deb6-af44-4b09-8e8f-4582fed819c0', '10 ta non va 3 ta go`sht', 'Samarqand shahri, Navoiy tumani, Bekzod kocha 25-uy', 'pending'),
--     ('e7284e62-1f92-4b7b-97cf-6bcd8e0e4aeb', '6dae588b-0339-4205-959e-9b62559660fd', '3 ta piyoz va 5 ta sabzi', 'Buxoro shahri, Mirzo Ulug`bek tumani, Jasur kochasi 10-uy', 'pending'),
--     ('a23f14e3-867c-498b-a596-c0c2584f2f0d', 'bbfe8ab6-2df2-4cae-ac28-d13289f90448', '7 ta kartoshka va 4 ta morkov', 'Andijon viloyati, Asaka tumani, Nodir kochasi 30-uy', 'pending'),
--     ('b7d32c5d-05a5-4857-9508-3d16593b3b4f', 'df779633-deee-47c8-9b4e-7a1be57343c1', '2 ta go`sht va 5 ta non', 'Namangan shahri, Chust tumani, Ozod kochasi 5-uy', 'pending'),
--     ('c2d6f3b7-5e72-4f8d-8c89-51e01e3f7c08', '543cbe40-7630-498b-b5f0-e6b3cce2d6a0', '8 ta sabzi va 3 ta non', 'Farg`ona shahri, Yozyovon tumani, Gulnora kochasi 20-uy', 'pending'),
--     ('d1f7a2e8-0e3f-4d1a-87f2-601ecad5b58b', '88929e11-e201-4f27-8cd5-acba7162661e', '6 ta go`sht va 2 ta morkov', 'Qarshi shahri, Kattaqorgon tumani, Dilshod kochasi 12-uy', 'pending'),
--     ('e8b4d0d9-3f07-44e1-9b73-4e1f9b8c432d', '462d860e-88d5-4800-8443-7b53531c794f', '10 ta non va 4 ta piyoz', 'Jizzax shahri, Dostlik tumani, Javlon kochasi 8-uy', 'pending'),
--     ('f6d2a5ea-78b8-44a1-b7c5-3a9e647c9b9d', '9a6b0590-c15e-4e53-8c67-fd0c0b7078ec', '3 ta qovun va 6 ta olcha', 'Xorazm viloyati, Urganch shahri, Sojida kochasi 25-uy', 'pending');



-- INSERT INTO CourierOrders (courier_order_id, courier_id, order_id, status)
-- VALUES 
--     ('d3b4a4b3-7841-4a6e-b0c7-54e450dc8a8d', 'fea9df3c-3dba-46af-83c2-864943cc0861', 'd3b4a4b3-7841-4a6e-b0c7-54e450dc8a8d', 'assigned'),
--     ('fbf24974-3cf1-44f7-93e7-938b6d8dc5b0', '36ec7595-490a-4f62-97c3-f0609494237a', 'fbf24974-3cf1-44f7-93e7-938b6d8dc5b0', 'assigned'),
--     ('e7284e62-1f92-4b7b-97cf-6bcd8e0e4aeb', '87f8e776-e679-4770-bc95-cf7d31bfb6ff', 'e7284e62-1f92-4b7b-97cf-6bcd8e0e4aeb', 'assigned'),
--     ('a23f14e3-867c-498b-a596-c0c2584f2f0d', '726ea834-dcc3-4e6c-9eba-a0c8bcb25f68', 'a23f14e3-867c-498b-a596-c0c2584f2f0d', 'assigned'),
--     ('b7d32c5d-05a5-4857-9508-3d16593b3b4f', 'bd92467b-3299-4b03-8bc9-6aefcd2c15ec', 'b7d32c5d-05a5-4857-9508-3d16593b3b4f', 'assigned'),
--     ('c2d6f3b7-5e72-4f8d-8c89-51e01e3f7c08', 'efb24d0b-8950-49fd-91b8-832f5811fcb9', 'c2d6f3b7-5e72-4f8d-8c89-51e01e3f7c08', 'assigned'),
--     ('d1f7a2e8-0e3f-4d1a-87f2-601ecad5b58b', 'e23a658f-c311-4a63-9565-0f966243fdcf', 'd1f7a2e8-0e3f-4d1a-87f2-601ecad5b58b', 'assigned'),
--     ('e8b4d0d9-3f07-44e1-9b73-4e1f9b8c432d', 'e19c9b99-ca15-4aac-91f2-1545a8d1e20a', 'e8b4d0d9-3f07-44e1-9b73-4e1f9b8c432d', 'assigned'),
--     ('f6d2a5ea-78b8-44a1-b7c5-3a9e647c9b9d', 'fc67098c-d8c1-4bfd-be8b-716ddb3301e0', 'f6d2a5ea-78b8-44a1-b7c5-3a9e647c9b9d', 'assigned');
