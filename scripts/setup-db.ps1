# Database Setup Script
# Run this once after starting Postgres to enable UUID extension and apply migrations

Write-Host "Enabling UUID extension..." -ForegroundColor Cyan
Get-Content -Raw migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -f -

Write-Host "`nApplying migration 0001..." -ForegroundColor Cyan
Get-Content -Raw migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

Write-Host "`nApplying migration 0002..." -ForegroundColor Cyan
Get-Content -Raw migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

Write-Host "`nApplying migration 0003..." -ForegroundColor Cyan
Get-Content -Raw migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

Write-Host "`nMigrations complete!" -ForegroundColor Green
