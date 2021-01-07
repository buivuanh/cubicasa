**Start**
<br />
`go run main.go`
<br />
<br />

**Config**
<br />
`config.json`
<br />
<br />

**Migration Databases**
<br />
`migrations/*`
<br />
<br />

**APIs**
<br />
Get a hub: `GET: http://localhost:5000/api/hubs/:hubID`
<br />
Create a hub: `POST: http://localhost:5000/api/hubs`
<br />
<br />
Get a team: `GET: http://localhost:5000/api/prodteams/:teamID`
<br />
Create a team: `POST: http://localhost:5000/api/prodteams`
<br />
Team join hub: `POST: http://localhost:5000/api/prodteams/:teamID/join-hub`
<br />
`payload: {"id": <hubID>}`
<br />
<br />
Get a user: `GET: http://localhost:5000/api/users/:userID`
<br />
Create a user: `POST: http://localhost:5000/api/prodteams/:teamID/users`