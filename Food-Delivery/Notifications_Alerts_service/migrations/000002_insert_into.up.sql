

-- Insert data into UserNotifications
INSERT INTO UserNotifications (notification_id, user_id, order_id, type, message)
VALUES 
    ('d6d1aa0d-086c-4c2f-9c4e-c44c9b63dff1', 'ca1ddd41-be35-4123-a8e0-be75ebbad39a', '9a6b0590-c15e-4e53-8c67-fd0c0b7078ec', 'order_confirmation', 'Buyurtmangiz tasdiqlandi va sizga yetkazib berish uchun tayyor!'),
    ('5fa7d26e-8db0-4b0f-8f56-2bfc23e5be56', 'faf7deb6-af44-4b09-8e8f-4582fed819c0', '462d860e-88d5-4800-8443-7b53531c794f', 'status_update', 'Sizning buyurtmangiz muvaffaqiyatli amalga oshirildi.'),
    ('9da67d0d-33e3-4aef-915d-3f1bbceaa5b1', '6dae588b-0339-4205-959e-9b62559660fd', '88929e11-e201-4f27-8cd5-acba7162661e', 'order_confirmation', 'Yangi buyurtmangizga maslahat berildi.'),
    ('7881e0d1-3d1f-4e3e-8a07-8df5b9753409', 'bbfe8ab6-2df2-4cae-ac28-d13289f90448', '543cbe40-7630-498b-b5f0-e6b3cce2d6a0', 'status_update', 'Buyurtmangiz bekor qilindi.'),
    ('1d2b7328-f0e0-4d0d-bb84-8dfc6e1e1b15', 'df779633-deee-47c8-9b4e-7a1be57343c1', '6dae588b-0339-4205-959e-9b62559660fd', 'status_update', 'Buyurtmangiz muvaffaqiyatli amalga oshirildi.'),
    ('4ab8e02d-efc1-40a5-9537-7022bfa87d70', '543cbe40-7630-498b-b5f0-e6b3cce2d6a0', 'faf7deb6-af44-4b09-8e8f-4582fed819c0', 'order_confirmation', 'Sizga yangi buyurtma taklif etildi.'),
    ('ae0986b6-1c2e-4687-bf60-229b0fd9e4a7', '88929e11-e201-4f27-8cd5-acba7162661e', 'ca1ddd41-be35-4123-a8e0-be75ebbad39a', 'status_update', 'Sizning buyurtmangiz muvaffaqiyatli amalga oshirildi.'),
    ('8a9e0d0b-1d6e-4b15-9f7d-3b9fd2b0a1a9', '462d860e-88d5-4800-8443-7b53531c794f', '3f8c8d2d-8e1a-4e20-8d4d-3f5fa3b0a6b2', 'status_update', 'Sizning buyurtmangiz bekor qilindi.'),
    ('3f8c8d2d-8e1a-4e20-8d4d-3f5fa3b0a6b2', '9a6b0590-c15e-4e53-8c67-fd0c0b7078ec', 'ae0986b6-1c2e-4687-bf60-229b0fd9e4a7', 'status_update', 'Sizning buyurtmangiz muvaffaqiyatli amalga oshirildi.'),
    ('2c1e8fbb-5f5f-4d5e-b39e-481d6328e8ef', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', '4ab8e02d-efc1-40a5-9537-7022bfa87d70', 'order_confirmation', 'Buyurtmangizga maslahat berildi.'),
    ('ca1ddd41-be35-4123-a8e0-be75ebbad39a', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', '1d2b7328-f0e0-4d0d-bb84-8dfc6e1e1b15', 'order_confirmation', 'Buyurtmangiz tasdiqlandi va sizga yetkazib berish uchun tayyor!'),
    ('faf7deb6-af44-4b09-8e8f-4582fed819c0', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', '7881e0d1-3d1f-4e3e-8a07-8df5b9753409', 'status_update', 'Sizning buyurtmangiz muvaffaqiyatli amalga oshirildi.');



INSERT INTO CourierNotifications (notification_id, courier_id, order_id, type, message)
VALUES 
    ('c0a2d512-13d4-4b7d-8d61-b50fbf0cfcd8', 'fea9df3c-3dba-46af-83c2-864943cc0861', 'd6d1aa0d-086c-4c2f-9c4e-c44c9b63dff1', 'order_assignment', 'Sizga yangi buyurtma tayinlandi.'),
    ('0c6b91b0-6827-4487-b801-18d92d455b50', '36ec7595-490a-4f62-97c3-f0609494237a', '5fa7d26e-8db0-4b0f-8f56-2bfc23e5be56', 'status_update', 'Buyurtma muvaffaqiyatli amalga oshirildi.'),
    ('fa54c5c1-3c30-4f43-8d9f-47c4e6a0e849', '87f8e776-e679-4770-bc95-cf7d31bfb6ff', '9da67d0d-33e3-4aef-915d-3f1bbceaa5b1', 'order_assignment', 'Buyurtma tayinlandi, yetkazish kerak.'),
    ('4f2c5ae8-0945-4c77-828b-6db7f61728e3', '726ea834-dcc3-4e6c-9eba-a0c8bcb25f68', '7881e0d1-3d1f-4e3e-8a07-8df5b9753409', 'new_recommendation', 'Buyurtma bekor qilindi.'),
    ('3b9c12af-44b1-4cc5-b9e8-5093e7cb2c3b', 'bd92467b-3299-4b03-8bc9-6aefcd2c15ec', '1d2b7328-f0e0-4d0d-bb84-8dfc6e1e1b15', 'status_update', 'Buyurtma muvaffaqiyatli amalga oshirildi.'),
    ('11ab3e17-92c3-4c76-b9f6-382d5db8b72e', 'efb24d0b-8950-49fd-91b8-832f5811fcb9', '4ab8e02d-efc1-40a5-9537-7022bfa87d70', 'order_assignment', 'Sizga yangi buyurtma tayinlandi.'),
    ('cdf4a93e-b920-40a2-a09c-dc1f98ffb89e', 'e23a658f-c311-4a63-9565-0f966243fdcf', 'ae0986b6-1c2e-4687-bf60-229b0fd9e4a7', 'status_update', 'Buyurtma muvaffaqiyatli amalga oshirildi.'),
    ('88b42e34-c353-4a1a-9172-dc19712bb7b0', 'e19c9b99-ca15-4aac-91f2-1545a8d1e20a', '8a9e0d0b-1d6e-4b15-9f7d-3b9fd2b0a1a9', 'new_recommendation', 'Buyurtma bekor qilindi.'),
    ('b4dbe71f-9b0e-48cf-95b3-477547dc0b27', 'fc67098c-d8c1-4bfd-be8b-716ddb3301e0', '3f8c8d2d-8e1a-4e20-8d4d-3f5fa3b0a6b2', 'status_update', 'Buyurtma muvaffaqiyatli amalga oshirildi.'),
    ('199f5fc2-e2ac-4963-97ae-4a5c7ebda102', '0b48a25c-e01a-4e42-a3e7-b26208ff4676', '2c1e8fbb-5f5f-4d5e-b39e-481d6328e8ef', 'new_recommendation', 'Sizga yangi buyurtma tayinlandi.');



INSERT INTO AdminAlerts (alert_id, admin_id, type, message)
VALUES 
    ('1a91f87c-1b9d-4c4e-9808-c7e1e569b0c7', 'ca1ddd41-be35-4123-a8e0-be75ebbad39a', 'new_order', 'Yangi buyurtma kelib tushdi.'),
    ('ed47b3e6-92c0-4701-8615-487b865d2673', 'faf7deb6-af44-4b09-8e8f-4582fed819c0', 'status_change', 'Buyurtma holati yangilandi.'),
    ('3f2a8b8e-6d2e-4b87-b50c-3a34f0d6475d', '6dae588b-0339-4205-959e-9b62559660fd', 'courier_activity', 'Kuryer faoliyati yangilandi.'),
    ('7f372196-f92c-4b8a-9ba0-7ebbd0a9d607', 'bbfe8ab6-2df2-4cae-ac28-d13289f90448', 'new_order', 'Yangi buyurtma kelib tushdi.'),
    ('af3c72be-00a2-4840-8c2f-13ea84ad1c5e', 'df779633-deee-47c8-9b4e-7a1be57343c1', 'status_change', 'Buyurtma holati yangilandi.'),
    ('95d4d635-b0d4-4312-a378-5e1b51c3028a', '543cbe40-7630-498b-b5f0-e6b3cce2d6a0', 'courier_activity', 'Kuryer faoliyati yangilandi.'),
    ('d9631f3c-5457-4c4d-b7c8-7272ad6b8e98', '88929e11-e201-4f27-8cd5-acba7162661e', 'new_order', 'Yangi buyurtma kelib tushdi.'),
    ('2f8e13a2-74c5-4608-a662-64836af3e037', '462d860e-88d5-4800-8443-7b53531c794f', 'status_change', 'Buyurtma holati yangilandi.'),
    ('6c6d4f12-3215-42e2-8ab3-7a7d086eefcb', '9a6b0590-c15e-4e53-8c67-fd0c0b7078ec', 'courier_activity', 'Kuryer faoliyati yangilandi.'),
    ('6a7290d1-02ec-4f2f-8d72-0ae89c2a25df', '2d58ffc0-b064-4c9f-9aa6-081da88e07e7', 'new_order', 'Yangi buyurtma kelib tushdi.');
