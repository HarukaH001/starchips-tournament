# Starchips Tournament
Starchips battle tournament system.

## Prerequisite
- Go
- PostgreSQL

## API

### Common

#### Health check

`GET /health`

```
curl 127.0.0.1:5000/health
```

```
curl 127.0.0.1:5001/health
```

Response

`200 OK`

### Admin
Admin API included all Client API 's requests

#### Init Tournament

`POST /init`

```
curl -X POST 127.0.0.1:5001/init \
     -H "Content-Type: application/json" \
     -d '{"players": 128}'
```

Body

```
{
    players: 128
}
```

Response

`200 OK`
All new player slots generated.
Allowed if new player limits are greater than current limits.
Admin can then register users and generate printings
of player's QR with player name or empty line for manually filling walk-in players.

`410 Gone`
Player slots are not empty and request's player size now greater than current size,
indicated that the tournament already initiated.

#### Register Player

`POST /player/register`

```
curl -X POST 127.0.0.1:5001/player/register \
     -H "Content-Type: application/json" \
     -d '{"name": "player1", "contact_number": "0912345678"}'
```

Body

```
{
    name: "player1",
    contact_number: "0912345678"
}
```

Response

`200 OK`
Return player ID and Serial

`409 Conflict`
Player limits reached.

#### Submit battle result

`POST /battle/submit`

### Client

#### Get player stats

`GET /player?id={uuid}`

#### Get player battle histories

`GET /battle/history?player_id={uuid}`

#### Get leaderboard

`GET /leaderboard`
