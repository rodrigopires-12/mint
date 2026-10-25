# Verification of deliverable E1.1 (base project and gateway doctor).
# Prints PASS/FAIL/SKIP per step and RESULTADO: PASS|FAIL at the end. Exit code 1 on FAIL.
$ErrorActionPreference = 'Continue'
# gateway.exe writes UTF-8; make PowerShell decode native output as UTF-8.
$savedEncoding = [Console]::OutputEncoding
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
$root = Resolve-Path (Join-Path $PSScriptRoot '..\..')
Set-Location $root
if (Test-Path 'D:\') {
    $env:GOMODCACHE = 'D:\mint\go\mod'
    $env:GOCACHE = 'D:\mint\go\cache'
}
$script:fails = 0
$script:skips = 0

function Step([string]$name, [scriptblock]$body) {
    try {
        $r = @(& $body)
        $last = if ($r.Count -gt 0) { $r[$r.Count - 1] } else { $null }
        if (($last -is [string]) -and ($last -eq 'SKIP')) { Write-Host ("SKIP  {0}" -f $name) -ForegroundColor Yellow; $script:skips++ }
        elseif (($last -is [bool]) -and $last) { Write-Host ("PASS  {0}" -f $name) -ForegroundColor Green }
        else { Write-Host ("FAIL  {0}: {1}" -f $name, ($last | Out-String).Trim()) -ForegroundColor Red; $script:fails++ }
    } catch {
        Write-Host ("FAIL  {0}: {1}" -f $name, $_.Exception.Message) -ForegroundColor Red
        $script:fails++
    }
}

function Run([string]$command, [string[]]$argList) {
    $out = (& $command @argList 2>&1 | ForEach-Object { "$_" }) -join "`n"
    return @{ Code = $LASTEXITCODE; Out = $out }
}

Step 'go build ./...' {
    $r = Run 'go' @('build', './...')
    if ($r.Code -ne 0) { return $r.Out }
    $true
}

Step 'go vet + go test ./...' {
    $r = Run 'go' @('vet', './...')
    if ($r.Code -ne 0) { return $r.Out }
    $r = Run 'go' @('test', '-count=1', './...')
    if ($r.Code -ne 0) { return $r.Out }
    $true
}

$exe = Join-Path $root 'dist\gateway.exe'
Step 'build gateway.exe (windows/amd64)' {
    New-Item -ItemType Directory -Force (Join-Path $root 'dist') | Out-Null
    $env:GOOS = 'windows'; $env:GOARCH = 'amd64'; $env:CGO_ENABLED = '0'
    $r = Run 'go' @('build', '-trimpath', '-ldflags', '-s -w -X main.version=verify', '-o', $exe, './gateway/cmd/gateway')
    Remove-Item Env:GOOS, Env:GOARCH, Env:CGO_ENABLED -ErrorAction SilentlyContinue
    if ($r.Code -ne 0) { return $r.Out }
    if (-not (Test-Path $exe)) { return 'gateway.exe not produced' }
    $true
}

Step 'read-only guard refuses forbidden.sql' {
    $r = Run $exe @('sqlcheck', 'scripts\verify\testdata\forbidden.sql')
    if ($r.Code -ne 1) { return "exit code $($r.Code), expected 1`n$($r.Out)" }
    if ($r.Out -match '\[OK\]') { return "some statement was accepted`n$($r.Out)" }
    $true
}

Step 'read-only guard accepts allowed.sql' {
    $r = Run $exe @('sqlcheck', 'scripts\verify\testdata\allowed.sql')
    if ($r.Code -ne 0) { return "exit code $($r.Code)`n$($r.Out)" }
    $true
}

$wavesCount = (Select-String -Path 'sources\winthor\_waves.yaml' -Pattern '^\s+- PC[A-Z0-9_]+\s*$').Count
Step "_waves.yaml lists 46 tables (found $wavesCount)" { $wavesCount -eq 46 }

$hasOracle = (Test-Path 'config.local.yaml') -or ($env:ORACLE_DSN)
Step 'doctor against the test Oracle (no FAIL, all tables checked)' {
    if (-not $hasOracle) { return 'SKIP' }
    $r = Run $exe @('doctor', '--json')
    $json = $null
    try { $json = $r.Out | ConvertFrom-Json } catch { return "invalid JSON`n$($r.Out)" }
    if ($r.Code -ne 0) { return "exit code $($r.Code); FAIL checks: " + (($json.results | Where-Object status -eq 'FAIL' | ForEach-Object { $_.check + ': ' + $_.detail }) -join ' | ') }
    if ($json.tables_expected -ne $wavesCount) { return "tables_expected=$($json.tables_expected), waves=$wavesCount" }
    if ($json.tables_checked -ne $json.tables_expected) { return "tables_checked=$($json.tables_checked) of $($json.tables_expected)" }
    $true
}

Step 'doctor with wrong password: clear message, no secret, no stack trace' {
    if (-not $hasOracle) { return 'SKIP' }
    $bad = 'wrong-password-E11-' + (Get-Random)
    $saved = $env:ORACLE_PASSWORD
    $env:ORACLE_PASSWORD = $bad
    try { $r = Run $exe @('doctor') } finally {
        if ($null -eq $saved) { Remove-Item Env:ORACLE_PASSWORD -ErrorAction SilentlyContinue } else { $env:ORACLE_PASSWORD = $saved }
    }
    if ($r.Code -ne 1) { return "exit code $($r.Code), expected 1`n$($r.Out)" }
    if ($r.Out -match [regex]::Escape($bad)) { return 'the wrong password appears in the output' }
    if ($r.Out -match 'goroutine |panic:') { return "stack trace in output`n$($r.Out)" }
    if ($r.Out -notmatch 'ORA-01017|usu.rio ou senha') { return "no clear authentication message`n$($r.Out)" }
    $true
}

[Console]::OutputEncoding = $savedEncoding
Write-Host ''
if ($script:fails -eq 0) {
    Write-Host ("RESULTADO: PASS ({0} passos pulados)" -f $script:skips) -ForegroundColor Green
    exit 0
}
Write-Host ("RESULTADO: FAIL ({0} falhas, {1} pulados)" -f $script:fails, $script:skips) -ForegroundColor Red
exit 1
