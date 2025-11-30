# QUICK REALTIME TEST
# Usage: .\test_realtime.ps1 <ACCESS_TOKEN> <CONVERSATION_ID>

param(
    [string]$Token,
    [string]$ConversationId
)

if (-not $Token -or -not $ConversationId) {
    Write-Host "❌ Usage: .\test_realtime.ps1 <ACCESS_TOKEN> <CONVERSATION_ID>" -ForegroundColor Red
    Write-Host ""
    Write-Host "📝 Steps:" -ForegroundColor Cyan
    Write-Host "   1. Postman: VerifyOTP → Copy access_token"
    Write-Host "   2. Postman: CreateConversation → Copy id"
    Write-Host "   3. Run: .\test_realtime.ps1 'TOKEN' 'CONV_ID'"
    exit 1
}

Write-Host "🚀 Starting Realtime Test..." -ForegroundColor Green
Write-Host "Token: $($Token.Substring(0, 20))..." -ForegroundColor Gray
Write-Host "Conversation: $ConversationId" -ForegroundColor Gray
Write-Host ""

# Run test in background
$process = Start-Process -FilePath "go" -ArgumentList "run", ".\test\quick_test.go" `
    -WorkingDirectory "C:\Users\pgadmin\Desktop\GoApp" `
    -NoNewWindow -PassThru

Write-Host "✅ Test client running (PID: $($process.Id))" -ForegroundColor Green
Write-Host ""
Write-Host "📤 Now send a message from Postman:" -ForegroundColor Cyan
Write-Host "   POST localhost:50050" -ForegroundColor White
Write-Host "   chat.ChatService/SendMessage" -ForegroundColor White
Write-Host "   {" -ForegroundColor White
Write-Host '     "conversation_id": "' -NoNewline -ForegroundColor White
Write-Host $ConversationId -NoNewline -ForegroundColor Yellow
Write-Host '",' -ForegroundColor White
Write-Host '     "content": "Test message!"' -ForegroundColor White
Write-Host "   }" -ForegroundColor White
Write-Host ""
Write-Host "⏳ Waiting 60 seconds..." -ForegroundColor Yellow

# Wait for process
Wait-Process -Id $process.Id -Timeout 65 -ErrorAction SilentlyContinue

Write-Host ""
Write-Host "✅ Test completed!" -ForegroundColor Green
