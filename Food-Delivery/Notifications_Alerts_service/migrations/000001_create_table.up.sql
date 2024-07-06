CREATE TYPE Order_type AS ENUM ('order_confirmation', 'status_update');
CREATE TYPE Courer_type AS ENUM ('new_recommendation', 'order_assignment', 'status_update');
CREATE TYPE Alert_type AS ENUM ('new_order', 'status_change', 'courier_activity');



--Foydalanuvchilarga buyurtma haqida ma'lumot berish uchun jadval.
CREATE TABLE UserNotifications (
    notification_id UUID PRIMARY KEY, -- Bildirishnoma ID'si
    user_id UUID NOT NULL, -- Bildirishnomani oladigan foydalanuvchi
    order_id UUID, -- Bildirishnomaga bog'langan buyurtma
    type Order_type NOT NULL, -- Bildirishnoma turi (masalan, 'buyurtmani_tasdiqlash", "holatni_yangilash')
    message TEXT NOT NULL, -- Bildirishnoma matni
    is_read BOOLEAN DEFAULT FALSE, -- Foydalanuvchi bildirishnomani o'qiganligini bildiradi
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);


--Kuryerlarga yangi buyurtma va holat yangilanishlari haqida ma'lumot berish uchun jadval.
CREATE TABLE CourierNotifications (
    notification_id UUID PRIMARY KEY, -- Bildirishnoma ID'si
    courier_id UUID NOT NULL, -- Bildirishnomani oladigan kuryer
    order_id UUID, -- Bildirishnomaga bog'langan buyurtma
    type Courer_type NOT NULL, -- Bildirishnoma turi (masalan, 'yangi_tavsiya', 'buyurtma_topshiriq', 'holatni_yangilash')
    message TEXT NOT NULL, -- Bildirishnoma matni
    is_read BOOLEAN DEFAULT FALSE, -- Kuryer bildirishnomani o'qiganligini bildiradi
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);


--Adminlarga yangi buyurtmalar va kuryer harakatlari haqida ma'lumot berish uchun jadval.
CREATE TABLE AdminAlerts (
    alert_id UUID PRIMARY KEY, -- Ogohlantirish ID'si
    admin_id UUID NOT NULL, -- Ogohlantirishni oladigan admin
    type Alert_type NOT NULL, -- Ogohlantirish turi (masalan, 'yangi_buyurtma', 'holatni_o'zgartirish', 'kuryer_faoliyati')
    message TEXT NOT NULL, -- Ogohlantirish matni
    is_read BOOLEAN DEFAULT FALSE, -- Admin ogohlantirishni o'qiganligini bildiradi
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);
