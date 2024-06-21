# Define variables
$binaryPath = Join-Path (Split-Path -Parent $MyInvocation.MyCommand.Path) $args[0]  # Prepend the absolute path of the current directory
$envPath = Join-Path (Split-Path -Parent $MyInvocation.MyCommand.Path) .env  # Prepend the absolute path of the current directory
$taskName = "DailyBinaryExecution"
$timeUTC = "00:30"

# Convert UTC time to local time
$localTime = [System.TimeZoneInfo]::ConvertTimeBySystemTimeZoneId(
    [DateTime]::ParseExact($timeUTC, "HH:mm", $null), 
    "UTC", 
    [System.TimeZoneInfo]::Local.Id
)

# Create a new task action
$action = New-ScheduledTaskAction -Execute $binaryPath -Argument $envPath

# Create a daily trigger at the local time
$trigger = New-ScheduledTaskTrigger -Daily -At $localTime.TimeOfDay

# Create task settings with WakeToRun enabled
$settings = New-ScheduledTaskSettingsSet -AllowStartIfOnBatteries -DontStopIfGoingOnBatteries -StartWhenAvailable -WakeToRun

# Register the task
Register-ScheduledTask -Action $action -Trigger $trigger -TaskName $taskName -Settings $settings -Description "Daily binary execution at 00:30 UTC"

Write-Output "Task scheduled successfully. It will run daily at $localTime (local time)."
