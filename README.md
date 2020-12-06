# account
Account service provides implementation of credit and debit for a user account. 


### Steps to up and running

1. Install Go on your local machine. You will also need `Postgres` running locally.

2. `git clone git@github.com:akashgupta05/account.git`

3. Create a dev environment file

   `cp env.sample development.env`

   Edit values in `development.env` to match your needs

4. Run Migrations

   - Install [golang-migrate](https://github.com/golang-migrate/migrate) tool

     `brew install golang-migrate`

   -  Create database

     `createdb account_development`

   - Use helper bash script to run migrations

     `./scripts/migrations.sh`

5. Install `gin`. This is needed for live-reload during development

   `GOBIN=/usr/local/bin/ go install github.com/codegangsta/gin`

6. Use helper bash script to run locally

   `./scripts/local_run.sh`
