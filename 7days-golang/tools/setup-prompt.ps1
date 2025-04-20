$profileContent = @'
function prompt {
    $projectRoot = "C:\Users\Administrator\Documents"
    $currentPath = (Get-Location).Path
    
    if ($currentPath.StartsWith($projectRoot)) {
        $relativePath = $currentPath.Substring($projectRoot.Length)
        if ($relativePath) {
            "doc$relativePath> "
        } else {
            "doc> "
        }
    } else {
        "$currentPath> "
    }
}
'@

# Create profile directory if it doesn't exist
$profileDir = Split-Path $PROFILE -Parent
if (!(Test-Path $profileDir)) {
    New-Item -ItemType Directory -Path $profileDir -Force
}

# Write the profile content
Set-Content -Path $PROFILE -Value $profileContent

Write-Host "PowerShell profile has been updated. Please restart your PowerShell terminal or run:"
Write-Host ". `$PROFILE"