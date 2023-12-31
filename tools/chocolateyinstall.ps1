$ErrorActionPreference = 'Stop';
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"
$fileLocation = Join-Path $toolsDir '.'

$binPath = Join-Path (Get-ToolsLocation) '.'
Copy-Item $fileLocation $binPath
