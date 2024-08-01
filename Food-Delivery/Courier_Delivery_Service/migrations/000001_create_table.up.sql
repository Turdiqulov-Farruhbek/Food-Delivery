BEGIN;

CREATE TYPE Courier_status AS ENUM ('available', 'unavailable', 'on_delivery');
CREATE TYPE User_role AS ENUM ('user', 'admin');
CREATE TYPE OrderStatus AS ENUM ('PENDING_Status', 'ACCEPTED', 'IN_PROGRESS', 'DELIVERED', 'CANCELED');
CREATE TYPE pay_status AS ENUM ('pending', 'paid');
CREATE TYPE CourOrd_status AS ENUM ('assigned', 'picked_up', 'en_route', 'delivered', 'payment_collected');

/*BARCHA TYPELARNI KO`RISH QUERY 

SELECT n.nspname AS schema_name,
       t.typname AS enum_name,
       e.enumlabel AS enum_value
FROM pg_catalog.pg_type t
     JOIN pg_catalog.pg_namespace n ON n.oid = t.typnamespace
     JOIN pg_catalog.pg_enum e ON t.oid = e.enumtypid
WHERE t.typtype = 'e'
ORDER BY schema_name, enum_name, e.enumsortorder;

*/


CREATE TABLE Users (
    user_id UUID PRIMARY KEY, -- Foydalanuvchi noyob ID'si
    email VARCHAR(100) UNIQUE NOT NULL, -- Foydalanuvchi email manzili
    password_hash VARCHAR(255) NOT NULL, -- Parolni hash qilingan qiymati
    full_name VARCHAR(100) NOT NULL, -- Foydalanuvchi to'liq ismi
    role User_role DEFAULT 'user', -- Foydalanuvchi rol (masalan, 'user', 'admin')
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Foydalanuvchi yaratilgan vaqti
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

COMMIT;



CREATE TABLE UserSessions (
    session_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES Users(user_id) ON DELETE CASCADE,
    token VARCHAR(500) UNIQUE NOT NULL, -- JWT token or session ID
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP, -- Token expiration time
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);





-- 2. Couriers jadvalini yaratish
--Kuryerlar haqida ma'lumot saqlaydi.
CREATE TABLE Couriers (
    courier_id UUID PRIMARY KEY,              -- Kuryerning noyob ID'si
    email VARCHAR(100) UNIQUE,                -- Kuryerning elektron pochtasi
    password_hash VARCHAR(255) NOT NULL, -- Parolni hash qilingan qiymati
    full_name VARCHAR(100) NOT NULL,          -- Kuryerning ismi
    phone_number VARCHAR(20) UNIQUE,          -- Kuryerning telefon raqami
    status Courier_status DEFAULT 'available', -- Kuryer holati ("mavjud", "mavjud emas", "etkazib berishda")
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE courier(
    customer_id UUID NOT NULL,                      -- Mijozning ID'si
    order_details TEXT,                             -- Buyurtma tafsilotlari
    delivery_address VARCHAR(255),                  -- Yetkazib berish manzili
    payment_status pay_status DEFAULT 'pending',    -- To'lov holati ("kutishda", "to'langan")
    order_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Buyurtma sanasi
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);



-- Buyurtmalar jadvalini yaratish
CREATE TABLE orders (
    order_id UUID PRIMARY KEY,                          -- Buyurtmaning noyob ID'si
    customer_id UUID NOT NULL,                          -- Mijozning ID'si
    order_details TEXT NOT NULL,                        -- Buyurtma tafsilotlari
    delivery_address VARCHAR(255) NOT NULL,             -- Yetkazib berish manzili
    payment_status pay_status NOT NULL,                 -- To'lov holati
    order_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,  -- Buyurtma sanasi
    status OrderStatus DEFAULT 'PENDING_Status'         -- Buyurtma holati
);


-- Kuryer va buyurtma o'rtasidagi munosabat va yetkazib berish holatini kuzatadi.
CREATE TABLE CourierOrders (
    courier_order_id UUID PRIMARY KEY,                                          -- Kuryer-buyurtma munosabatining noyob ID'si
    courier_id UUID NOT NULL REFERENCES Couriers(courier_id) ON DELETE CASCADE, -- Kuryer ID'si
    order_id UUID NOT NULL REFERENCES orders(order_id),                                -- Buyurtma ID'si
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
    assigned_to UUID REFERENCES users(user_id),          -- "assigned_to" ustuni vazifa qaysi foydalanuvchiga tayinlanganligini ko'rsatadi va "users" jadvalidagi "id" ustuniga chet el kaliti sifatida murojaat qiladi.
    due_date DATE,                                     -- "due_date" ustuni vazifaning oxirgi muddatini saqlaydi.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- "created_at" ustuni vazifa qachon yaratilganligini ko'rsatadi va sukut bo'yicha joriy vaqtni saqlaydi.
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP     -- "updated_at" ustuni vazifa qachon yangilanganligini ko'rsatadi va sukut bo'yicha joriy vaqtni saqlaydi.
);


CREATE TABLE courier_locations (
    id UUID PRIMARY KEY,                               -- "id" ustuni jadvalning birlamchi kaliti bo'lib, u avtomatik ravishda o'suvchi raqamdir.
    courier_id UUID REFERENCES users(user_id),           -- "courier_id" ustuni kurerning foydalanuvchi identifikatori bo'lib, "users" jadvalidagi "id" ustuniga chet el kaliti sifatida murojaat qiladi.
    latitude DECIMAL(10, 8) NOT NULL,                  -- "latitude" ustuni kurerning kenglik koordinatasini saqlaydi va bu ustun bo'sh bo'lishi mumkin emas.
    longitude DECIMAL(11, 8) NOT NULL,                 -- "longitude" ustuni kurerning uzunlik koordinatasini saqlaydi va bu ustun bo'sh bo'lishi mumkin emas.
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP     -- "updated_at" ustuni kurer joylashuvi qachon yangilanganligini ko'rsatadi va sukut bo'yicha joriy vaqtni saqlaydi.
);


