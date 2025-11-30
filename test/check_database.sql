-- Test Sonuçlarını Kontrol Et
-- psql'de çalıştır: \i test/check_database.sql

-- 1. Toplam İstatistikler
SELECT 
    'Users' as table_name, COUNT(*) as count FROM users
UNION ALL
    SELECT 'Conversations', COUNT(*) FROM conversations
UNION ALL
    SELECT 'Messages', COUNT(*) FROM messages
UNION ALL
    SELECT 'Participants', COUNT(*) FROM conversation_participants;

-- 2. En Son Eklenen Kullanıcılar (Son 5)
SELECT 
    id,
    phone_number,
    display_name,
    created_at
FROM users 
ORDER BY created_at DESC 
LIMIT 5;

-- 3. Tüm Conversation'lar ve Detayları
SELECT 
    c.id,
    c.is_group,
    c.group_name,
    c.created_at,
    COUNT(DISTINCT cp.user_id) as participants,
    COUNT(m.id) as messages
FROM conversations c
LEFT JOIN conversation_participants cp ON c.id = cp.conversation_id
LEFT JOIN messages m ON c.id = m.conversation_id
GROUP BY c.id
ORDER BY c.created_at DESC;

-- 4. Son Mesajlar (Kimden Kime)
SELECT 
    m.id,
    sender.phone_number as sender,
    m.content,
    m.created_at
FROM messages m
JOIN users sender ON m.sender_id = sender.id
ORDER BY m.created_at DESC
LIMIT 10;

-- 5. Kullanıcı Başına Device Sayısı
SELECT 
    u.phone_number,
    COUNT(d.device_name) as device_count,
    STRING_AGG(d.device_name, ', ') as devices
FROM users u
LEFT JOIN user_devices d ON u.id = d.user_id
GROUP BY u.id, u.phone_number
ORDER BY device_count DESC;

-- 6. Bugün Oluşturulan Kayıtlar
SELECT 
    'Users Today' as type, COUNT(*) as count
FROM users 
WHERE created_at::date = CURRENT_DATE
UNION ALL
SELECT 'Conversations Today', COUNT(*)
FROM conversations
WHERE created_at::date = CURRENT_DATE
UNION ALL
SELECT 'Messages Today', COUNT(*)
FROM messages
WHERE created_at::date = CURRENT_DATE;
