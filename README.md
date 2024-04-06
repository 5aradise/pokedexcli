# Pokédex CLI

### Description

A simple console pokedex application written in Go programming language based on [pokeapi.co](https://pokeapi.co/api/v2)

## Build and start using

For this application to work, you'll need to install Go environment. See [go.dev](https://go.dev/) for installation instructions.

When you have installed Go, clone this repository

```
git clone https://github.com/5aradise/pokedexcli.git
```

Then you should create an .exe file by running _go build_

```
go build
```

To start play with Pokédex CLI, simply open this .exe file

```
.\pokedexcli.exe
```

## Command

To see the list of available commands, just write _help_

```
>help
help - Displays a list of all commands
config - Displays current player configuration
pokedex - Displays your pokemons
inspect {pokemon_name} - Displays information about a {pokemon_name} pokemon
map [offset] [limit] - Displays the names of [limit] locations, starting from the [offset], in the Pokemon world
mapf [step] - Displays the next [step] locations
mapb [step] - Displays the previous [step] locations
areas {location_name} - Displays areas in a {location_name} location
explore {area_name} - Displays pokemons in a {area_name} area
catch {pokemon_name} - Trys to catches a {pokemon_name} pokemon
exit - Exit the Pokedex
```
