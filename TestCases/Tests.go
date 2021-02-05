package main

import (
	"../ActionStats"
	"fmt"
	"math/rand"
	"time"
)


func main() {
	actions := []string{"run", "jump", "slide", "attack"}
	rand.Seed(time.Now().UnixNano())
	GivenTestCase()
	RandomTestCases(10, actions)
	RandomTestCases(100, actions)
	RandomTestCases(1000, actions)
	RandomTestCases(10000, actions)
	IntentionalBadData()
}

func GivenTestCase()  {
	action1 := "{\"action\":\"jump\", \"time\":100}"
	action2 := "{\"action\":\"run\", \"time\":75}"
	action3 := "{\"action\":\"jump\", \"time\":200}"
	ActionStats.AddAction(action1)
	ActionStats.AddAction(action2)
	ActionStats.AddAction(action3)

	fmt.Println(ActionStats.GetStats())
	ActionStats.Reset()
}

func RandomTestCases(numberOfRuns int, actions []string)  {
	for i := 0; i < numberOfRuns; i++ {
		action := actions[rand.Intn(len(actions))]
		time := rand.Intn(2000)
		ActionStats.AddAction(fmt.Sprint("{\"action\":\"", action, "\", \"time\":", time, "}"))
	}
	fmt.Println(ActionStats.GetStats())
	ActionStats.Reset()
}

func IntentionalBadData(){
	action1 := ""
	action2 := "J̷̞̟̜̱̺͗̀͜u̶̡̧̻̥͔͛̔͆̒̂͑̓͝m̴̧̧̧͔̖̩̬̋͛̊̎̑̌́̾͘p̵̨̞̦̪͇̎͊̒̿c̵̞͓̥̳̞̪̜̍̚͘l̴̢̢̜̗̫͚̫̭̠̓̔͐̏ŏ̶̧̼̜͉͍͎͉̜̟́̀͊u̸͉͎̫̣̠̱̩̙̔̓̈́̊̽̓̓̔̅̌d̴̢͓͕͍̎̅͛̃"
	action3 := "{\"action\":\"jump\", \"time\",200}"
	ActionStats.AddAction(action1)
	ActionStats.AddAction(action2)
	ActionStats.AddAction(action3)

	fmt.Println(ActionStats.GetStats())
	ActionStats.Reset()
}
