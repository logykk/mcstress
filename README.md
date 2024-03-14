# MCstress
Maximum intensity stress testing utility for Minecraft servers

This tool works on all versions of Minecraft after 1.7.2 excluding CraftBukkit and derivatives (Spigot, Paper, etc.).

MCstress works by creating thousands of simultaneous bot users (default 2048) and repeatedly logging them into a target server.

With thousands of bots, this quickly overwhelms the network thread on the server.
This causes currently online players to be disconnected, or in extreme situations, crashes the entire server.

With less bots and/or by running MCstress in short bursts, you can cause intense lag spikes resulting in huge rubber-banding for server players.

In some cases a server will attempt to validate the usernames of the bots joining.
Due to how many there are and how frequent they're joining, the Mojang session server will rate limit the server's address, preventing players from reconnecting for some time after stress testing.

## Installation
- Install go, instructions can be found [here](https://www.google.com/search?q=install+go)
- Then clone this repo `git clone https://github.com/logykk/mcstress`
- Build it `cd mcstress; go build -o mcstress .`

## Usage
You need to provide at least the server's IP, port and protocol version

`./mcstress -address <IP>:<port> -protocol <protocol>`

### Options
| Name      | Description                                     |
| --------- | ----------------------------------------------- |
| -address  | Server IP:port                                  |
| -username | Base username for the bots                      |
| -uuid     | UUID for bots (1.19.2) only                     |
| -protocol | Protocol version of server                      |
| -number   | Numer of bots running in parallel               |
| -wait     | Milliseconds each bot waits before reconnecting |

To get the protocol version from a release version check [wiki.vg](https://wiki.vg/Protocol_version_numbers)
