package main

import (
	"fmt"
	_ "github.com/shogo82148/go-shuffle"
	"math/rand"
	"time"
)

var left = []string{
	"ambitious",
	"anxious",
	"awesome",
	"beautiful",
	"boring",
	"busy",
	"cool",
	"cute",
	"cynical",
	"ecstatic",
	"evil",
	"ex",
	"fair",
	"funny",
	"furious",
	"good",
	"goofy",
	"grumpy",
	"genious",
	"hungry",
	"helpful",
	"horrific",
	"incredible",
	"intrepid",
	"inoffensive",
	"jealous",
	"jolly",
	"joyful",
	"keen",
	"kind",
	"lazy",
	"lucky",
	"lost",
	"mad",
	"misty",
	"moody",
	"nasty",
	"nice",
	"nonviolent",
	"peaceful",
	"peaky",
	"pretty",
	"ridiculous",
	"romantic",
	"rude",
	"sad",
	"shy",
	"sweet",
	"tender",
	"terrific",
	"tiny",
	"ugly",
	"upset",
	"uptight",
	"violent",
}
var right = []string{
	"albatross",
	"alligator",
	"anaconda",
	"bandicoot",
	"bear",
	"butterfly",
	"camel",
	"cat",
	"cow",
	"dingo",
	"dolphin",
	"dragonfly",
	"eagle",
	"earwig",
	"elephant",
	"falcon",
	"flamingo",
	"fox",
	"gibbon",
	"goose",
	"gorilla",
	"hamster",
	"heron",
	"hyena",
	"ibis",
	"impala",
	"iguana",
	"jackrussel",
	"jaguar",
	"jellyfish",
	"kangaroo",
	"kiwi",
	"koala",
	"leopard",
	"lion",
	"lobster",
	"meerkat",
	"mole",
	"monkey",
	"octopus",
	"otter",
	"oyster",
	"panda",
	"penguin",
	"puma",
	"quetzal",
	"rabbit",
	"raccoon",
	"rhinoceros",
	"seal",
	"snake",
	"squid",
	"tiger",
	"toucan",
	"turkey",
	"wallaby",
	"weasel",
	"wolf",
	"yak",
	"zebra",
	"zebu",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	leftIndex := rand.Intn(len(left))
	rand.Seed(time.Now().UTC().UnixNano())
	rightIndex := rand.Intn(len(right))
	fmt.Printf("%s-%s\n", left[leftIndex], right[rightIndex])
}
