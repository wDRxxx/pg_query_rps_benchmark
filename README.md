# How to run

1. (optionally) run docker and configure **create_tables.sql**

    `docker-compose up -d`
2. configure .env file for your purposes
    
    ```dotenv
    QUERY="SELECT * FROM test"
    DSN=postgresql://postgres:postgres@localhost:5432/benchmark
    DURATION=1000
   ```
3. Run benchmark
    
    `go run .`

To check data races use `-race` flag