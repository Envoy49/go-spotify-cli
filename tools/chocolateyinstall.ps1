$ErrorActionPreference = 'Stop';
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"
$fileLocation = Join-Path $toolsDir 'go-spotify-cli.exe'

$binPath = Join-Path (Get-ToolsLocation) 'go-spotify-cli.exe'
Copy-Item $fileLocation $binPath
