# Realtime Test Script
# Edit this file and fill in the values

# STEP 1: Get tokens from Postman and paste them here
$ALICE_TOKEN = "ALICE_ACCESS_TOKEN_HERE"
$BOB_TOKEN = "BOB_ACCESS_TOKEN_HERE"
$CONVERSATION_ID = "CONVERSATION_ID_HERE"

# STEP 2: Open two terminals and run these:

Write-Host "`n=== REALTIME TEST ===" -ForegroundColor Cyan
Write-Host "`n1. FIRST TERMINAL (Alice):" -ForegroundColor Yellow
Write-Host "   go run ./test/realtime_test.go Alice $ALICE_TOKEN $CONVERSATION_ID" -ForegroundColor Green

Write-Host "`n2. SECOND TERMINAL (Bob):" -ForegroundColor Yellow  
Write-Host "   go run ./test/realtime_test.go Bob $BOB_TOKEN $CONVERSATION_ID" -ForegroundColor Green

Write-Host "`n3. THIRD TERMINAL (Send Message - Alice):" -ForegroundColor Yellow
Write-Host @"
   Invoke-WebRequest -Method POST ``
     -Uri "http://localhost:50052/chat.ChatService/SendMessage" ``
     -Headers @{ "Authorization" = "Bearer $ALICE_TOKEN"; "Content-Type" = "application/json" } ``
     -Body '{"conversation_id":"$CONVERSATION_ID","content":"Hello Bob!","message_type":"text"}'
"@ -ForegroundColor Green

Write-Host "`n💡 When Alice sends a message, you'll see it in Bob's terminal!" -ForegroundColor Cyan
Write-Host "`n─────────────────────────────────────────`n"
