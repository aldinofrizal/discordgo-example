please run to start the app
```bash
docker build --tag [image_name] .

docker create --name [container_name] -e BOT_TOKEN=[DISCORD_BOT_TOKEN] -e CHANNEL_GENERAL_ID=[DISCORD_CHANNEL_ID] [image_name]

docker run [container_name]
```