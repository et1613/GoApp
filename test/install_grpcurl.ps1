# grpcurl installer (Windows)
Write-Host "📥 Installing grpcurl..." -ForegroundColor Cyan

# Check if already installed
if (Get-Command grpcurl -ErrorAction SilentlyContinue) {
    Write-Host "✅ grpcurl already installed!" -ForegroundColor Green
    grpcurl --version
    exit 0
}

# Install via chocolatey
if (Get-Command choco -ErrorAction SilentlyContinue) {
    choco install grpcurl -y
} else {
    Write-Host "⚠️  Chocolatey not found. Installing manually..." -ForegroundColor Yellow
    
    # Download latest release
    $url = "https://github.com/fullstorydev/grpcurl/releases/download/v1.8.9/grpcurl_1.8.9_windows_x86_64.zip"
    $outFile = "$env:TEMP\grpcurl.zip"
    
    Invoke-WebRequest -Uri $url -OutFile $outFile
    
    # Extract
    Expand-Archive -Path $outFile -DestinationPath "$env:USERPROFILE\grpcurl" -Force
    
    # Add to PATH (session only)
    $env:Path += ";$env:USERPROFILE\grpcurl"
    
    Write-Host "✅ grpcurl installed to $env:USERPROFILE\grpcurl" -ForegroundColor Green
    Write-Host "   Add to PATH permanently or use full path" -ForegroundColor Yellow
}
