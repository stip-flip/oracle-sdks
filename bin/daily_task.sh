#!/bin/bash

# Define variables
SCRIPT_PATH="$1"
PLIST_PATH="/Library/LaunchDaemons/com.example.dailyscript.plist"
CRON_FILE="/etc/cron.d/wakeup"
UTC_TIME="00:30"

# Function to check if the script is running on macOS
is_macos() {
    [[ "$OSTYPE" == "darwin"* ]]
}

# Function to check if the script is running on Linux
is_linux() {
    [[ "$OSTYPE" == "linux-gnu"* ]]
}

# Function to convert UTC time to local time on macOS
convert_utc_to_local_macos() {
    date -j -f "%H:%M %Z" "$UTC_TIME UTC" +"%H:%M"
}

# Function to convert UTC time to local time on Linux
convert_utc_to_local_linux() {
    date -d "$UTC_TIME UTC" +"%H:%M"
}

# Function to set up macOS
setup_macos() {
    echo "Setting up macOS..."

    # Convert UTC time to local time
    local_time=$(convert_utc_to_local_macos)

    # Schedule wake-up using pmset
    sudo pmset repeat wakeorpoweron MTWRFSU "$local_time:00"

    # Parse hour and minute from local time
    local_hour=$(echo "$local_time" | cut -d':' -f1)
    local_minute=$(echo "$local_time" | cut -d':' -f2)

    # Create a launch daemon plist file
    sudo tee $PLIST_PATH > /dev/null <<EOL
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.example.dailyscript</string>
    <key>ProgramArguments</key>
    <array>
        <string>$SCRIPT_PATH</string>
    </array>
    <key>StartCalendarInterval</key>
    <dict>
        <key>Hour</key>
        <integer>$local_hour</integer>
        <key>Minute</key>
        <integer>$local_minute</integer>
    </dict>
    <key>RunAtLoad</key>
    <true/>
    <key>StandardOutPath</key>
    <string>/tmp/com.example.dailyscript.out</string>
    <key>StandardErrorPath</key>
    <string>/tmp/com.example.dailyscript.err</string>
</dict>
</plist>
EOL

    # Load the launch daemon
    sudo launchctl load $PLIST_PATH
}

# Function to set up Linux
setup_linux() {
    echo "Setting up Linux..."

    # Convert UTC time to local time
    local_time=$(convert_utc_to_local_linux)

    # Schedule the cron job to run the script at 00:30 UTC
    (crontab -l 2>/dev/null; echo "30 0 * * * $SCRIPT_PATH") | crontab -

    # Calculate rtcwake time (5 minutes before the script runs)
    wakeup_time=$(date -d "tomorrow $UTC_TIME UTC - 5 minutes" +%s)

    # Create a cron file for rtcwake
    sudo tee $CRON_FILE > /dev/null <<EOL
25 0 * * * root /usr/sbin/rtcwake -m no -t $wakeup_time
EOL
}

# Main execution
if is_macos; then
    setup_macos
elif is_linux; then
    setup_linux
else
    echo "Unsupported OS."
    exit 1
fi

echo "Setup complete."
