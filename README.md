# Starchips Tournament
Starchips battle tournament system.
Include API Servers and Client/Admin Interactive UI
Admin can create user and generate QR, show to player, then player can see his stats.
Player should use his stats page with identical physical name card to identify ownership.
Player can select bet amount from UI and generate battle QR for admin/referree to initiate battles.
Battle histories are to be shown on each player's stats page.
Leaderboard is to show ranking of players with highest starchips at a time.

Planned
- Player Tiers
    - players battling should be on same tier with different minimum bets of each tier.
- Battle Phases
    - Each phase of tournament have minimum bets
    - The Tournament have N phases
    - Current Phase and time until next phase should show on player stats and leaderboard page
    - Ongoing battles initiated before end of phases should continue as normal with already placed bets.
    - Server does not need to know ongoing battles. It needs only results submitted.

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
