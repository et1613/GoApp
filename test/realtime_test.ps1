# Realtime Test Script
# Bu dosyayı düzenle ve değerleri doldur

# ADIM 1: Postman'den token'ları al ve buraya yapıştır
$ALICE_TOKEN = "ALICE_ACCESS_TOKEN_BURAYA"
$BOB_TOKEN = "BOB_ACCESS_TOKEN_BURAYA"
$CONVERSATION_ID = "CONVERSATION_ID_BURAYA"

# ADIM 2: İki terminal aç ve şunları çalıştır:

Write-Host "`n=== REALTIME TEST ===" -ForegroundColor Cyan
Write-Host "`n1. İLK TERMINAL (Alice):" -ForegroundColor Yellow
Write-Host "   go run ./test/realtime_test.go Alice $ALICE_TOKEN $CONVERSATION_ID" -ForegroundColor Green

Write-Host "`n2. İKİNCİ TERMINAL (Bob):" -ForegroundColor Yellow  
Write-Host "   go run ./test/realtime_test.go Bob $BOB_TOKEN $CONVERSATION_ID" -ForegroundColor Green

Write-Host "`n3. ÜÇÜNCÜ TERMINAL (Mesaj Gönder - Alice):" -ForegroundColor Yellow
Write-Host @"
   Invoke-WebRequest -Method POST ``
     -Uri "http://localhost:50052/chat.ChatService/SendMessage" ``
     -Headers @{ "Authorization" = "Bearer $ALICE_TOKEN"; "Content-Type" = "application/json" } ``
     -Body '{"conversation_id":"$CONVERSATION_ID","content":"Merhaba Bob!","message_type":"text"}'
"@ -ForegroundColor Green

Write-Host "`n💡 Alice mesaj gönderdiğinde, Bob'un terminalinde göreceksin!" -ForegroundColor Cyan
Write-Host "`n─────────────────────────────────────────`n"
