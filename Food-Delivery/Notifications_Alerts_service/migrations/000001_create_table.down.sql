-- Foydalanuvchilarga buyurtma haqida ma'lumot berish uchun jadvalni olib tashlash
DROP TABLE IF EXISTS UserNotifications CASCADE;

-- Kuryerlarga yangi buyurtma va holat yangilanishlari haqida ma'lumot berish uchun jadvalni olib tashlash
DROP TABLE IF EXISTS CourierNotifications CASCADE;

-- Adminlarga yangi buyurtmalar va kuryer harakatlari haqida ma'lumot berish uchun jadvalni olib tashlash
DROP TABLE IF EXISTS AdminAlerts CASCADE;
DROP TABLE IF EXISTS schema_migrations CASCADE;


-- Enum turlarini olib tashlash
DROP TYPE IF EXISTS Order_type CASCADE;
DROP TYPE IF EXISTS Courer_type CASCADE;
DROP TYPE IF EXISTS Alert_type CASCADE;
