-- Buyurtma holatlari uchun enum turi
CREATE TYPE Card_type AS ENUM ('pending', 'completed', 'canceled');

-- Mahsulotlarning asosiy ma'lumotlarini saqlaydi.
CREATE TABLE Products (
    product_id UUID PRIMARY KEY, -- Mahsulotning noyob ID'si
    name VARCHAR(100) NOT NULL, -- Mahsulot nomi
    description TEXT, -- Mahsulot tavsifi
    price DECIMAL(10, 2) NOT NULL, -- Mahsulot narxi
    image_url VARCHAR(255), -- Mahsulotning rasmini saqlovchi URL
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- Foydalanuvchi savatlari haqida ma'lumot saqlaydi.
CREATE TABLE Carts (
    cart_id UUID PRIMARY KEY, -- Savatning noyob ID'si
    user_id UUID NOT NULL, -- Savatga ega foydalanuvchi ID'si
    UNIQUE (user_id), -- Har bir foydalanuvchida faqat bitta faol savat bo'lishi mumkin
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Savat yaratilgan vaqti
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- Savatdagi mahsulotlar haqida ma'lumot saqlaydi
CREATE TABLE CartItems (
    cart_item_id UUID PRIMARY KEY, -- Savat mahsuloti noyob ID'si
    cart_id UUID NOT NULL REFERENCES Carts(cart_id) ON DELETE CASCADE, -- Savat ID'si
    product_id UUID NOT NULL REFERENCES Products(product_id) ON DELETE CASCADE, -- Mahsulot ID'si
    quantity INT NOT NULL CHECK (quantity > 0), -- Mahsulot miqdori
    options JSONB, -- Variantlar, masalan, o'lcham, qo'shimchalar va hokazo
    UNIQUE (cart_id, product_id), -- Har bir mahsulot savatda faqat bir marta bo'lishi mumkin
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

/*Farqi va bog'lanishi
Carts jadvali savatlarning asosiy ma'lumotlarini saqlaydi, masalan, foydalanuvchi bilan bog'lanishi va savatning qachon yaratilganligi.
CartItems jadvali esa savatdagi individual mahsulotlar haqida batafsil ma'lumot beradi, masalan, mahsulot ID'si, miqdori va variantlari.
Carts va CartItems jadvallari o'rtasida bir-to-ko'p (one-to-many) bog'lanish mavjud: har bir savatda (Carts jadvali) ko'p mahsulotlar bo'lishi mumkin 
(CartItems jadvali), ammo har bir mahsulot faqat bitta savatga tegishli. Bu bog'lanish cart_id ustuni orqali amalga oshiriladi.*/


-- Buyurtmalar haqida ma'lumot saqlaydi.
CREATE TABLE P_Orders (
    order_id UUID PRIMARY KEY, -- Buyurtma noyob ID'si
    user_id UUID NOT NULL, -- Buyurtma bergan foydalanuvchi ID'si
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Buyurtma yaratilgan vaqti
    status VARCHAR(50) DEFAULT 'pending', -- Buyurtma holati (masalan, "kutishda", "tugallangan", "bekor qilingan")
    total_price DECIMAL(10, 2) NOT NULL, -- Buyurtma umumiy summasi
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- Buyurtmadagi mahsulotlar haqida ma'lumot saqlaydi.
CREATE TABLE OrderItems (
    order_item_id UUID PRIMARY KEY, -- Buyurtma mahsuloti noyob ID'si
    order_id UUID NOT NULL REFERENCES P_Orders(order_id) ON DELETE CASCADE, -- Buyurtma ID'si
    product_id UUID NOT NULL REFERENCES Products(product_id) ON DELETE CASCADE, -- Mahsulot ID'si
    quantity INT NOT NULL CHECK (quantity > 0), -- Mahsulot miqdori
    options JSONB, -- Variantlar, masalan, o'lcham, qo'shimchalar va hokazo
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- Buyurtma tavsiyalari haqida ma'lumot saqlaydi.
CREATE TABLE OrderRecommendations (
    recommendation_id UUID PRIMARY KEY, -- Tavsiya noyob ID'siOrders
    order_id UUID NOT NULL REFERENCES P_Orders(order_id) ON DELETE CASCADE, -- Buyurtma ID'si
    courier_id UUID NOT NULL REFERENCES Couriers(courier_id) ON DELETE CASCADE, -- Kuryer ID'si
    recommended_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Tavsiya qilingan vaqt
    status VARCHAR  DEFAULT 'pending', -- Tavsiya holati (masalan, 'pending', 'accepted', 'rejected')
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);


