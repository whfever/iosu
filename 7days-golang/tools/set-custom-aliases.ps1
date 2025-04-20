# Create PowerShell profile content
$profileContent = @'
# Custom aliases for directory navigation
function Set-AdminLocation { Set-Location "C:\Users\Administrator" }
function Set-DocsLocation { Set-Location "C:\Users\Administrator\Documents" }

Set-Alias -Name cda -Value Set-AdminLocation
Set-Alias -Name cdd -Value Set-DocsLocation
'@

# Create profile directory if needed
$profileDir = Split-Path $PROFILE -Parent
if (!(Test-Path $profileDir)) {
    New-Item -ItemType Directory -Path $profileDir -Force
}

# Write to profile
Set-Content -Path $PROFILE -Value $profileContent

Write-Host "Aliases have been added. To use them, restart PowerShell or run:"
Write-Host ". `$PROFILE"