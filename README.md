# Docs

```shell
# Step 1: DB initialize
make db_up

# Step 2: Nats initialize

# add nats in docker (if not exists)
docker pull nats
# run nats server
docker run --name nats --rm -p 4222:4222 -p 6222:6222 -p 8222:8222 nats

# Step 3: Run app
make dev # or make prod

# Step 4: Run sub
make run_sub
```