# Run Migrations
make migrate-up branch=test-mysql

# Starts the hot reload with air
air -c .\.air.toml

# Postman collection
Check eulabs.postman_collection.json file

# Set up your .env based on .env.example