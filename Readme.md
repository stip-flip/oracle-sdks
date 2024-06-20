# SF Oracles

This repository contains the SF Oracle SDKs, which is a collection of software development kits (SDKs) for submitting price automatically to SF contract oracles on the Ether.

## Installation

To use the SF Oracles SDKs in your Go project, you need to have Go installed on your machine. You can download and install Go from the official Go website.

To install the SDKs, you can either:

### Build it from source

1. Clone this repository to your local machine.
2. Navigate to the root directory of the repository.
3. Build the SDKs by running the following command:

```shell
go build
```

### Install the Go SDK by running the following command:

```shell
go get -u github.com/stip-flip/oracle-sdks
```

### Download the appropriate binary from the releases page

1. Navigate to the releases page of this repository.
2. Download the appropriate binary for your operating system.
3. Extract the contents of the archive to a directory on your machine.

### Use the docker image available on the [docker hub](https://hub.docker.com/r/sotachi/sf-oracle)

```shell
docker pull sotachi/sf-oracle
```

## Usage

You have 2 possibility to run the sdk, either as a server or as a cron job.

1. As a server

   Once the sdk is installed and ready to run, you need to set the following environment variables:

   ```shell

   # The private key of the account that will be used to sign transactions
   PRIVATE_KEY=your-private-key
   # Cryptocompare API key
   CC_API_KEY=your-api-key
   # The contract oracle address that will be used to store crypto compare oracle prices
   CC_ORACLE_ADDRESS=cc-oracle-address

   POLL_INTERVAL=300s

   CHAIN_ID=61

   RPC_URL=https://etc.etcdesktop.com
   ```

   This .env file hold important informations, ie the private key, make sure it is kept safe and not shared with anyone.

   Note that in order to submit prices to an oracle, you need to stake a minimum amount in the contract first. Visit https://sf.exchange/oracle/deposit to stake the minimum amount.

   You can then run this program as a standalone binary or as a docker container.

2. As a CRON job

   Download the appropriate binary ([cronv1.0.1](https://github.com/stip-flip/oracle-sdks/releases)) from the releases page or build it from source.

   Create a .env file next to the cron binary with the following content:

   ```shell
   # The private key of the account that will be used to sign transactions
   PRIVATE_KEY=your-private-key

   CHAIN_ID=61

   RPC_URL=https://etc.etcdesktop.com
   ```

   This .env file hold important informations, ie the private key, make sure it is kept safe and not shared with anyone.

   You can then automate the execution of the binary by using the daily_task bash script available in the repository.

   [Download](https://github.com/stip-flip/oracle-sdks/releases/download/v1.0.1/daily_task.sh) and place the daily_task.sh script in the same directory as the cron binary and the .env files.

   Windows:

   Open Powersell and run the following command:

   ```shell
   chmod +x .\daily_task.ps1
   chmod +x "cron-binary-name"
   .\daily_task.ps1 "cron-binary-name"
   ```

   Linux and MacOS:

   Open a terminal and run the following command:

   ```shell
    chmod +x ./daily_task.sh
    chmod +x "cron-binary-name"
   ./daily_task.sh "cron-binary-name"
   ```

   This will wake your computer up from sleep and run the cron binary at 00:30 UTC every day.

   MacOS will by default block the execution of the binary, you need to allow it in the security settings.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
