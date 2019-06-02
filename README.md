# Pay Later Service

## Installation 
* Copy config.json to $HOME/configs/config.json and change properties if needed
* `go build ` will create a payLaterService binary
* Run `go test ./...` to run all the testcases
* Run `./payLaterService` to get the cli
* This app uses postgres

## Usage
Use help (in cli) to get all the available commands
| Action |Command |
| --- | ---|
| New User | new-user name email credit-limit |
| New Merchant | new-merchant name interest |
| Transaction | transaction user-name merchant-name amount |
| Update merchant interest rate | update-merchant-interest merchant-name interest |
| Merchant details | report discount merchant-name |
| User dues for a particular user | report dues user-name |
| Find users at credit limit | report users-at-credit-limit |
| All user dues | report total-dues |
| Ctrl + c / exit | Safe exit |