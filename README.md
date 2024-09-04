# What is this?

A go package for retrieving predetermined, extended statistics about an (currently only stable) osu! lobby from the osu! api.

# Usage Example

First, retrieve the `clientId` and `clientSecret` from [here](https://osu.ppy.sh/home/account/edit#new-oauth-application)

```go
package main

import (
	"encoding/json"
	"fmt"
	osu_lobby_inspector "github.com/SnakeTwix/osu-lobby-inspector"
)

const clientId = 123456
const clientSecret = "secret"

func main() {
	fetcher, _ := osu_lobby_inspector.NewStatisticsFetcher(clientId, clientSecret)
	lobbyStats, _ := fetcher.FetchLobbyStatistics(115172885)
	jsonContent, _ := json.Marshal(lobbyStats)

	fmt.Println(string(jsonContent))
}

```

