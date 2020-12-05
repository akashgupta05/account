printf "Running Migrations: "

eval $(cat development.env) migrate -source file://migrations -database $DATABASE_URL up