package module02

import (
	"math/rand"
	"time"
)

// SecretSanta shuffles a list of players and assigns each player a secret santa
// based on their position in the list. If two players have the same name, they
// will be paired with each other. The function returns a map of players to their
// secret santas.
func SecretSanta(s []string) {
	secretSanta(s)
}

func secretSanta(players []string) map[string]string {
	// generate a random seed using the current time
	rand.Seed(time.Now().UnixNano())

	// shuffle the list of players
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	// create a slice to hold the secret santas
	secretSantas := make([]string, len(players))
	// copy the players into the secret santa slice
	copy(secretSantas, players)

	// loop through each player and check if they have the same name as another
	// player. If they do, swap their names with the name of the next player in the
	// list.
	for i := range players {
		if players[i] == secretSantas[i] {
			secretSantas[i], secretSantas[(i+1)%len(players)] = secretSantas[(i+1)%len(players)], secretSantas[i]
		}
	}

	// create a map to hold the santas
	santas := make(map[string]string)
	// loop through each player and add them to the map with their secret santa
	for i := range players {
		santas[players[i]] = secretSantas[i]
	}

	// return the map of santas
	return santas
}

// func main() {
// 	players := []string{"Fred", "Wilma", "Barney", "Pebbles", "Bam Bam"}
// 	santas := SecretSanta(players)
// 	fmt.Println(santas)
// }
