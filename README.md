# Spokemon

A little project written in Golang.
It has a `/pokemon.csv` file with all the Pokemon data.

## Usage
```bash
# Clone the project
git clone https://github.com/matteac/spokemon.git

# Move to the project directory
cd spokemon

# Install dependencies
go get

# Build the project
go build

# And run it
./spokemon
```
or 
```bash
# See --help
./spokemon --name pikachu --type electric
# -----------------------------------
# 	Name: Pikachu
# 	Types: Electric,
#  	HP: 35
#	ATK: 55
#	DEF: 40
#	SP.ATK: 50
#	SP.DEF: 50
#	Speed: 90
#	Generation: 1
#	Legendary: FALSE
# -----------------------------------

./spokemon --name mewtwo --type psychic --legendary
# -----------------------------------
#	Name: Mewtwo
#	Types: Psychic,
#	HP: 106
#	ATK: 110
#	DEF: 90
#	SP.ATK: 154
#	SP.DEF: 90
#	Speed: 130
#	Generation: 1
#	Legendary: TRUE
# -----------------------------------

```
