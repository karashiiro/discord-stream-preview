# Discord Stream Preview API
Small API for grabbing the Go Live preview image from a user in an accessible server.

**Requires a user bot token; use at your own risk. You will probably get banned if you use this too frequently.**

Requires the following environment variables:

`DISCORD_STREAM_PREVIEW_API_TOKEN` - A user bot token.

`DISCORD_STREAM_PREVIEW_API_PORT` - The port this API should run on.

### API routes
`/streams/:guildID/:channelID/:userID` - Returns an object like `{"url":"https://..."}`
containing a URL that points to the preview image.
