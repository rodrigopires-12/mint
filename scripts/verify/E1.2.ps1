# Verification of deliverable E1.2 (table inventory).
$ErrorActionPreference = 'Continue'
$savedEncoding = [Console]::OutputEncoding
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
$root = Resolve-Path (Join-Path $PSScriptRoot '..\..')
Set-Location $root
if (Test-Path 'D:\') { $env:GOMODCACHE = 'D:\mint\go\mod'; $env:GOCACHE = 'D:\mint\go\cache' }
$script:fails = 0; $script:skips = 0

function Step([string]$name, [scriptblock]$body) {
    try {
        $r = @(& $body)
        $last = if ($r.Count -gt 0) { $r[$r.Count - 1] } else { $null }
        if (($last -is [string]) -and ($last -eq 'SKIP')) { Write-Host ("SKIP  {0}" -f $name) -ForegroundColor Yellow; $script:skips++ }
        elseif (($last -is [bool]) -and $last) { Write-Host ("PASS  {0}" -f $name) -ForegroundColor Green }
        else { Write-Host ("FAIL  {0}: {1}" -f $name, ($last | Out-String).Trim()) -ForegroundColor Red; $script:fails++ }
    } catch { Write-Host ("FAIL  {0}: {1}" -f $name, $_.Exception.Message) -ForegroundColor Red; $script:fails++ }
}
function Run([string]$command, [string[]]$argList) {
    $out = (& $command @argList 2>&1 | ForEach-Object { "$_" }) -join "`n"
    return @{ Code = $LASTEXITCODE; Out = $out }
}

$exe = Join-Path $root 'dist\gateway.exe'
Step 'build gateway.exe' {
    New-Item -ItemType Directory -Force (Join-Path $root 'dist') | Out-Null
    $r = Run 'go' @('build', '-o', $exe, './gateway/cmd/gateway')
    if ($r.Code -ne 0) { return $r.Out }
    $true
}
Step 'unit tests (inventory, catalog)' {
    $r = Run 'go' @('test', '-count=1', './gateway/internal/inventory/...', './internal/catalog/...', './internal/oradict/...')
    if ($r.Code -ne 0) { return $r.Out }
    $true
}

$tables = Select-String -Path 'sources\winthor\_waves.yaml' -Pattern '^\s+- (PC[A-Z0-9_]+)\s*$' | ForEach-Object { $_.Matches[0].Groups[1].Value }
Step "one <TABLE>.yaml per catalog table ($($tables.Count))" {
    $missing = $tables | Where-Object { -not (Test-Path ("sources\winthor\{0}.yaml" -f $_)) }
    if ($missing) { return "missing: " + ($missing -join ', ') }
    $true
}
Step 'every <TABLE>.yaml has a key, columns and a status' {
    $bad = @()
    foreach ($t in $tables) {
        $p = "sources\winthor\$t.yaml"
        if (-not (Test-Path $p)) { continue }
        $y = Get-Content $p -Raw
        if ($y -notmatch '(?m)^status: (proposed|approved)\s*$') { $bad += "$t (status)" }
        if ($y -notmatch '(?m)^columns:') { $bad += "$t (columns)" }
        if ($y -notmatch '(?m)^primary_key:') { $bad += "$t (primary_key)" }
    }
    if ($bad) { return ($bad -join ', ') }
    $true
}
$approved = ($tables | Where-Object { (Test-Path "sources\winthor\$_.yaml") -and ((Get-Content "sources\winthor\$_.yaml" -Raw) -match '(?m)^status: approved\s*$') }).Count
Write-Host ("INFO  approved specs: {0}/{1}" -f $approved, $tables.Count)

Step 'docs/tables-inventory.md answers points 1, 2, 11, 12, 13 and 14' {
    if (-not (Test-Path 'docs\tables-inventory.md')) { return 'report missing' }
    $md = Get-Content 'docs\tables-inventory.md' -Raw
    $missing = @('1', '2', '11', '12', '13', '14') | Where-Object { $md -notmatch ("(?m)^### #{0} " -f $_) }
    if ($missing) { return "missing answers: #" + ($missing -join ', #') }
    $true
}
Step 'doctor reports table files (OK when all approved, WARN otherwise)' {
    if (-not ((Test-Path 'config.local.yaml') -or $env:ORACLE_DSN)) { return 'SKIP' }
    $r = Run $exe @('doctor', '--json')
    try { $json = $r.Out | ConvertFrom-Json } catch { return "invalid JSON`n$($r.Out)" }
    $res = $json.results | Where-Object check -eq 'arquivos de tabela'
    if (-not $res) { return 'check not found' }
    if ($approved -eq $tables.Count -and $res.status -ne 'OK') { return "expected OK, got $($res.status): $($res.detail)" }
    if ($approved -lt $tables.Count -and $res.status -ne 'WARN') { return "expected WARN, got $($res.status): $($res.detail)" }
    $true
}

[Console]::OutputEncoding = $savedEncoding
Write-Host ''
if ($script:fails -eq 0) { Write-Host ("RESULTADO: PASS ({0} passos pulados)" -f $script:skips) -ForegroundColor Green; exit 0 }
Write-Host ("RESULTADO: FAIL ({0} falhas, {1} pulados)" -f $script:fails, $script:skips) -ForegroundColor Red; exit 1
