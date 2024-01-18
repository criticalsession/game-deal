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
documentation, **all prices shown are in USD**.

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
deal information for the specific game, or to add the game to your favorite list. (See "2.2 Adding Games to your Fav List" for 
more info.)

```
> search fallout 4 max=10
Searching for: "fallout 4"

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

### 2.2 Managing your Fav List

When you initially run **Game-Deal** it will create a `sqlite3` database in the root directory called `favs.db`. Do not
delete this file unless you want to lose all your favorited games.

**2.2.1 Add to Favs**

`> fav-add [ID]`

After using the `search` command you can add any of the returned games to your fav list by using the `fav-add` command followed
by the `ID` of one of the games returned.

```
> search lego batman max=5
Searching for: "lego batman"

┌─────┬───────────────────────────────────────────────────────────────┬───────────────┐
│ ID  │ TITLE                                                         │ CHEAPEST DEAL │
├─────┼───────────────────────────────────────────────────────────────┼───────────────┤
│ [1] │ LEGO Batman                                                   │ $3.89         │
│ [2] │ LEGO Batman 2                                                 │ $4.19         │
│ [3] │ LEGO Batman: The Videogame                                    │ $3.74         │
│ [4] │ LEGO Batman 3: Beyond Gotham                                  │ $3.89         │
│ [5] │ LEGO Batman 2: DC Super Heroes                                │ $3.89         │
│ [6] │ LEGO Batman 3: Beyond Gotham Season Pass                      │ $2.91         │
│ [7] │ LEGO DC Super-Villains Batman: The Animated Series Level Pack │ $1.04         │
└─────┴───────────────────────────────────────────────────────────────┴───────────────┘

Use "deals [ID] <max=?>" command to see deals for a game
> fav-add 1
"LEGO Batman" successfully added to favorites!

┌─────┬─────────────┐
│ ID  │ TITLE       │
├─────┼─────────────┤
│ [1] │ LEGO Batman │
└─────┴─────────────┘

> fav-add 4
"LEGO Batman 3: Beyond Gotham" successfully added to favorites!

┌─────┬──────────────────────────────┐
│ ID  │ TITLE                        │
├─────┼──────────────────────────────┤
│ [1] │ LEGO Batman                  │
│ [2] │ LEGO Batman 3: Beyond Gotham │
└─────┴──────────────────────────────┘
```

After searching for "lego batman", the command `fav-add 1` adds "LEGO Batman" to the favorites list and
`fav-add 4` adds "LEGO Batman 3: Beyond Gotham".

**2.2.2 List Favs**

`> fav-list`

Simply lists all games in fav list.

```
> fav-list
┌─────┬──────────────────────────────┐
│ ID  │ TITLE                        │
├─────┼──────────────────────────────┤
│ [1] │ LEGO Batman                  │
│ [2] │ LEGO Batman 3: Beyond Gotham │
└─────┴──────────────────────────────┘
```

**2.2.3 Remove Fav**

`> fav-remove [ID]`

To remove a game from your fav list use `fav-remove` followed by the `ID` of the game in your fav list (**not original
id!**) Use `fav-list` to get the list of ids available.

```
> fav-remove 2
"LEGO Batman 3: Beyond Gotham" successfully removed from favorites

┌─────┬─────────────┐
│ ID  │ TITLE       │
├─────┼─────────────┤
│ [1] │ LEGO Batman │
└─────┴─────────────┘
```

**2.2.4 Get Deals for all Favorited Games**

`> fav-deals <max=?>`

Once you have games in your fav list you can search for deals for all games using `fav-deals` followed by the optional
`max=` argument. This will return a table for each game in your fav list without having to search for them one by one. 
Same as in "2.1.3 Open Deal", the deal `ID` returned in the result can be used with the `open` command to open the
deal in your browser.

```
> fav-list
┌─────┬──────────────────────────────┐
│ ID  │ TITLE                        │
├─────┼──────────────────────────────┤
│ [1] │ LEGO Batman                  │
│ [2] │ LEGO Batman 3: Beyond Gotham │
└─────┴──────────────────────────────┘

> fav-deals max=5
Deals for: "LEGO Batman"
Historically cheapest price: $3.89 (2023-05-21 11:23:40)

┌─────┬───────────┬────────────────┬──────────────────┬─────────┐
│ ID  │ STORE     │ ORIGINAL PRICE │ DISCOUNTED PRICE │ SAVINGS │
├─────┼───────────┼────────────────┼──────────────────┼─────────┤
│ [1] │ IndieGala │ $19.99         │ $3.89            │ 80.54%  │
│ [2] │ Fanatical │ $19.99         │ $4.19            │ 79.04%  │
└─────┴───────────┴────────────────┴──────────────────┴─────────┘

Deals for: "LEGO Batman 3: Beyond Gotham"
Historically cheapest price: $2.39 (2023-11-17 14:59:43)

┌─────┬───────────┬────────────────┬──────────────────┬─────────┐
│ ID  │ STORE     │ ORIGINAL PRICE │ DISCOUNTED PRICE │ SAVINGS │
├─────┼───────────┼────────────────┼──────────────────┼─────────┤
│ [3] │ IndieGala │ $19.99         │ $3.89            │ 80.54%  │
│ [4] │ Fanatical │ $19.99         │ $4.19            │ 79.04%  │
│ [5] │ Noctre    │ $19.99         │ $4.80            │ 75.99%  │
│ [6] │ GOG       │ $19.99         │ $4.99            │ 75.04%  │
└─────┴───────────┴────────────────┴──────────────────┴─────────┘
```

Here we're getting deals for "LEGO Batman" and "LEGO Batman 3: Beyond Gotham" and limiting the price to a maximum
of $5.

### 2.3 Misc

**2.3.1 Help**

`> help`

If you're ever lost, use the `help` command to display a full list of commands and how to use them.

**2.3.2 Stores**

`> stores`

Currently unused, the `stores` command gives you a full list of stores in the CheapShark API. In the future, this will be
used to display deals for a given store, or hide deals from specific stores you don't want to use.

**2.3.3 Exit**

`> exit`

I mean...

## 3. Found a bug? Have suggestions?

Feel free to use the Issues tab above (or [click here](https://github.com/criticalsession/game-deal/issues)) if you've found 
bugs, have problems running **Game-Deal**, have suggestions for improvements or general tips on how I can make the Go code better.

## 4. To-Do

- [ ] Store search
- [ ] Hide deals from stores

## 5. Like Game-Deal?

If you're feeling generous, buy me a beer! - https://www.buymeacoffee.com/criticalsession 🍺❤️