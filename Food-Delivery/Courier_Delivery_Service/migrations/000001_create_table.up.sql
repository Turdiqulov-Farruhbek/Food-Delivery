-- 1. ENUM turini yaratish
CREATE TYPE courier_status AS ENUM ('available', 'unavailable', 'on_delivery');
CREATE TYPE pay_status AS ENUM ('pending', 'paid');
CREATE TYPE CourOrd_status AS ENUM ('assigned', 'picked_up', 'en_route', 'delivered', 'payment_collected');


-- 2. Couriers jadvalini yaratish
--Kuryerlar haqida ma'lumot saqlaydi.
CREATE TABLE Couriers (
    courier_id UUID PRIMARY KEY,              -- Kuryerning noyob ID'si
    full_name VARCHAR(100) NOT NULL,          -- Kuryerning ismi
    phone_number VARCHAR(20) UNIQUE,          -- Kuryerning telefon raqami
    email VARCHAR(100) UNIQUE,                -- Kuryerning elektron pochtasi
    status courier_status DEFAULT 'available', -- Kuryerning holati ("mavjud", "mavjud emas", "etkazib berishda")
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);



--Buyurtmalar haqida ma'lumot saqlaydi
CREATE TABLE Orders (
    order_id UUID PRIMARY KEY,                      -- Buyurtmaning noyob ID'si
    customer_id UUID NOT NULL,                      -- Mijozning ID'si
    order_details TEXT,                             -- Buyurtma tafsilotlari
    delivery_address VARCHAR(255),                  -- Yetkazib berish manzili
    payment_status pay_status DEFAULT 'pending',    -- To'lov holati ("kutishda", "to'langan")
    order_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Buyurtma sanasi
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);


-- Kuryer va buyurtma o'rtasidagi munosabat va yetkazib berish holatini kuzatadi.
CREATE TABLE CourierOrders (
    courier_order_id UUID PRIMARY KEY,                                          -- Kuryer-buyurtma munosabatining noyob ID'si
    courier_id UUID NOT NULL REFERENCES Couriers(courier_id) ON DELETE CASCADE, -- Kuryer ID'si
    order_id UUID NOT NULL REFERENCES Orders(order_id) ON DELETE CASCADE,       -- Buyurtma ID'si
    status CourOrd_status DEFAULT 'assigned',                                   -- Yetkazib berish holati ('tayinlangan', 'olib_oldi', 'marshrutda', 'etkazib berildi', 'to'lov_yig'ildi')
    assigned_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,           -- Tayinlangan vaqt
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,            -- Oxirgi yangilangan vaqt
    deleted_at BIGINT DEFAULT 0                           
);


CREATE TABLE tasks (
    id UUID PRIMARY KEY,                               -- "id" ustuni jadvalning birlamchi kaliti bo'lib, u avtomatik ravishda o'suvchi raqamdir.
    title VARCHAR(255) NOT NULL,                       -- "title" ustuni vazifaning nomini saqlaydi va maksimal uzunligi 255 ta belgidan iborat bo'lishi mumkin. Bu ustun bo'sh bo'lishi mumkin emas.
    description TEXT,                                  -- "description" ustuni vazifaning batafsil tavsifini saqlaydi va uzunligi cheksiz bo'lishi mumkin.
    status VARCHAR(50) NOT NULL,                       -- "status" ustuni vazifaning holatini (masalan, "to do", "in progress", "done") saqlaydi va maksimal uzunligi 50 ta belgidan iborat bo'lishi mumkin. Bu ustun bo'sh bo'lishi mumkin emas.
    assigned_to INTEGER REFERENCES users(id),          -- "assigned_to" ustuni vazifa qaysi foydalanuvchiga tayinlanganligini ko'rsatadi va "users" jadvalidagi "id" ustuniga chet el kaliti sifatida murojaat qiladi.
    due_date DATE,                                     -- "due_date" ustuni vazifaning oxirgi muddatini saqlaydi.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- "created_at" ustuni vazifa qachon yaratilganligini ko'rsatadi va sukut bo'yicha joriy vaqtni saqlaydi.
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP     -- "updated_at" ustuni vazifa qachon yangilanganligini ko'rsatadi va sukut bo'yicha joriy vaqtni saqlaydi.
);


CREATE TABLE courier_locations (
    id UUID PRIMARY KEY,                               -- "id" ustuni jadvalning birlamchi kaliti bo'lib, u avtomatik ravishda o'suvchi raqamdir.
    courier_id INTEGER REFERENCES users(id),           -- "courier_id" ustuni kurerning foydalanuvchi identifikatori bo'lib, "users" jadvalidagi "id" ustuniga chet el kaliti sifatida murojaat qiladi.
    latitude DECIMAL(10, 8) NOT NULL,                  -- "latitude" ustuni kurerning kenglik koordinatasini saqlaydi va bu ustun bo'sh bo'lishi mumkin emas.
    longitude DECIMAL(11, 8) NOT NULL,                 -- "longitude" ustuni kurerning uzunlik koordinatasini saqlaydi va bu ustun bo'sh bo'lishi mumkin emas.
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP     -- "updated_at" ustuni kurer joylashuvi qachon yangilanganligini ko'rsatadi va sukut bo'yicha joriy vaqtni saqlaydi.
);


