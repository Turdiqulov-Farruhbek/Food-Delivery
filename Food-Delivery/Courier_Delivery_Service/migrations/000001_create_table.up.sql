-- 1. ENUM turini yaratish
CREATE TYPE courier_status AS ENUM ('available', 'unavailable', 'on_delivery');
CREATE TYPE pay_status AS ENUM ('pending', 'paid');
CREATE TYPE CourOrd_status AS ENUM ('assigned', 'picked_up', 'en_route', 'delivered', 'payment_collected');


-- 2. Couriers jadvalini yaratish
--Kuryerlar haqida ma'lumot saqlaydi.
CREATE TABLE Couriers (
    courier_id UUID PRIMARY KEY,              -- Kuryerning noyob ID'si
    name VARCHAR(100) NOT NULL,               -- Kuryerning ismi
    phone_number VARCHAR(20) UNIQUE,          -- Kuryerning telefon raqami
    email VARCHAR(100) UNIQUE,                -- Kuryerning elektron pochtasi
    status courier_status DEFAULT 'available' -- Kuryerning holati ("mavjud", "mavjud emas", "etkazib berishda")
);



--Buyurtmalar haqida ma'lumot saqlaydi
CREATE TABLE Orders (
    order_id UUID PRIMARY KEY,                      -- Buyurtmaning noyob ID'si
    customer_id UUID NOT NULL,                      -- Mijozning ID'si
    order_details TEXT,                             -- Buyurtma tafsilotlari
    delivery_address VARCHAR(255),                  -- Yetkazib berish manzili
    payment_status pay_status DEFAULT 'pending',    -- To'lov holati ("kutishda", "to'langan")
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Buyurtma sanasi
);


-- Kuryer va buyurtma o'rtasidagi munosabat va yetkazib berish holatini kuzatadi.
CREATE TABLE CourierOrders (
    courier_order_id UUID PRIMARY KEY,                                          -- Kuryer-buyurtma munosabatining noyob ID'si
    courier_id UUID NOT NULL REFERENCES Couriers(courier_id) ON DELETE CASCADE, -- Kuryer ID'si
    order_id UUID NOT NULL REFERENCES Orders(order_id) ON DELETE CASCADE,       -- Buyurtma ID'si
    status CourOrd_status DEFAULT 'assigned',                                   -- Yetkazib berish holati ('tayinlangan', 'olib_oldi', 'marshrutda', 'etkazib berildi', 'to'lov_yig'ildi')
    assigned_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                          -- Tayinlangan vaqt
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP                            -- Oxirgi yangilangan vaqt
);


