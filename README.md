# Game-Deal

## Game-Deal is a CLI application written in Go that lets you search for game deals across multiple stores using the CheapShark API.

[![GitHub release (with filter)](https://img.shields.io/github/v/release/criticalsession/game-deal)](https://github.com/criticalsession/game-deal/releases)
[![GitHub issues](https://img.shields.io/github/issues/criticalsession/game-deal)](https://github.com/criticalsession/game-deal/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/criticalsession/game-deal)](https://goreportcard.com/report/github.com/criticalsession/game-deal)
[![GitHub License](https://img.shields.io/github/license/criticalsession/game-deal)](https://github.com/criticalsession/criticalsession/blob/main/LICENSE)
[![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/criticalsession)](https://twitter.com/criticalsession)

<p align="center">
  <img src="https://github.com/criticalsession/game-deal/blob/main/docs/scrn.png?raw=true" width="608" />
</p>

**Game-Deal** is an open source command-line tool written in Go. With **Game-Deal** you can search games, find and compare 
deals, then open the store page to buy. You can also add games to a favorite list and find deals for all games in your
favorite list with one command.

## 1. DISCLAIMER

**Game-Deal** uses the CheapShark API to find and track game deals but is not affiliated with CheapShark in any way. 
Deal links have to go through CheapShark's redirect URL but this should not affect your price. As per CheapShark's API
document, **all prices shown are in USD**.

**Game-Deal** has no control over the prices, it simply plugs into the CheapShark API and displays the information as is.

## 2. How to Use

After running **Game-Deal** you'll be prompted with a `>`. This is where you'll enter commands. If you're ever unsure of
what commands you have available you can always use `> help` to get the full list and a short description of each command.

### 2.1 Searching and Finding Deals

**2.1.1 Search**

`> search [keywords] <max=?>`

Use `search` followed by any number of keywords to look up games. You can use the optional `max=` to filter the results by 
maximum price.

This command returns a list of games. Each entry will have an `ID`, `Title` and `Cheapest Deal`. The `Cheapest Deal` field
reflects the cheapest price available at any of the stores in the system. The `ID` value is used to either display 
deal information for the specific game, or to add the game to your favorite list. (See "Adding Games to your Fav List" for 
more info.)

```
> search fallout 4 max=10

┌─────┬────────────────────────────────────┬───────────────┐
│ ID  │ TITLE                              │ CHEAPEST DEAL │
├─────┼────────────────────────────────────┼───────────────┤
│ [1] │ Fallout 4                          │ $5.79         │
│ [2] │ Fallout 4 - Nuka-World             │ $8.39         │
│ [3] │ Fallout 4 - Automatron             │ $8.39         │
│ [4] │ Fallout 4 - Vault-Tec Workshop     │ $4.19         │
│ [5] │ Fallout 4 - Wasteland Workshop     │ $4.19         │
│ [6] │ Fallout 4 Game of the Year Edition │ $8.79         │
│ [7] │ Fallout 4  Contraptions Workshop   │ $4.19         │
└─────┴────────────────────────────────────┴───────────────┘
```

This will search for all games with "fallout 4" in the title, and filter out any games that are more expensive than $10.

**2.1.2 List Deals**

`> deals [ID] <max=?>`

Use `deals` after searching for games to display a list of deals for the given game. The `ID` argument must match one of
the ids of the games returned by the `search` command. The `max=` argument is optional and will filter out any deals
with price higher than the given value.

This command returns the historically cheapest price seen (and the earliest it was seen), and a table with the following
fields `ID`, `Store`, `Original Price`, `Discounted Price` and `Savings`. The `ID` argument can be used to open the deal
in browser. The rest should be self-explanatory.

```
> deals 1 max=10

Deals for: "Fallout 4"
Historically cheapest price: $4.25 (2021-10-15 17:09:59)

┌─────┬───────────┬────────────────┬──────────────────┬─────────┐
│ ID  │ STORE     │ ORIGINAL PRICE │ DISCOUNTED PRICE │ SAVINGS │
├─────┼───────────┼────────────────┼──────────────────┼─────────┤
│ [1] │ Fanatical │ $19.99         │ $5.79            │ 71.04%  │
└─────┴───────────┴────────────────┴──────────────────┴─────────┘
```

**2.1.3 Open Deal**

`> open [ID]`

If any of the deals interest you, you can now open the deal in your browser by using the `open` command. Again, the `ID`
argument here should match one of the returned deals.

```
> open 1

Opening deal in browser 🚀
```

### Adding Games to your Fav List