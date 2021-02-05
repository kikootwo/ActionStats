package ActionStats

import (
	"encoding/json"
	"errors"
	"sync"
)

// Hashed Map to hold Actions and their respective times. Keep all times to calculate average.
// Could possibly keep the current sum and an integer representing the number of actions added
var values map[string][]float32

// Read/Write Mutex to protect the map in concurrent workloads.
var lock = sync.RWMutex{}

// Struct representing the input JSON data for addAction
type inputAction struct{
	Action string	`json:"action"`
	Time float32	`json:"time"`
}

// Struct representing the output JSON data for each Action's Statistic
type outputStats struct {
	Action string	`json:"action"`
	Avg float32		`json:"avg"`
}

// Reset the current stored action statistics
func Reset() {
	values = nil
}

// Adds an action and its time to the 'values' map, under the relevant action key
func AddAction(jsonString string) error{
	// On the first call, the values map will be nil. Ensure that it is instantiated.
	if (values == nil){
		values = make(map[string][]float32)
	}

	// An empty string will never be a valid json dataset
	if(len(jsonString) == 0){
		return errors.New("invalid json input; input was empty")
	}

	// Attempt to deserialize the JSON string. Return error if there is one.
	inputAction := inputAction{}
	error := json.Unmarshal([]byte(jsonString), &inputAction)
	if error != nil {
		return error
	}

	// If the key is present, add the time to the Times Slice. If not, add a new slice with the Action's time
	if timesSlice, keyPresent := readValues(inputAction.Action); keyPresent{
		writeValues(inputAction.Action, append(timesSlice, inputAction.Time))
	} else {
		writeValues(inputAction.Action, []float32{inputAction.Time})
	}

	// No errors, succeeded.
	return nil
}

func GetStats() string{
	if values == nil {
		return ""
	}
	if len(values) == 0 {
		return ""
	}

	var stats []outputStats

	// Activate Read Lock on Map for 'range values'
	lock.RLock()
	defer lock.RUnlock()

	for action, timesSlice := range values{
		// Create a new stats object. The Avg is the sum of all of the elements of the slice divided by the number of elements
		stat := outputStats{
			Action: action,
			Avg:    sumSlice(timesSlice) / float32(len(timesSlice)),
		}
		stats = append(stats, stat)
	}

	// Serialize the object to JSON. This should not contain any bad data, so we can ignore the error with _
	jsonString, _ := json.Marshal(stats)

	return string(jsonString)
}


// Wrap Reading from the map with RLock/RUnlock for concurrent workloads
func readValues(key string) ([]float32, bool){
	lock.RLock()
	defer lock.RUnlock()
	timesSlice, keyPresent := values[key]
	return timesSlice, keyPresent
}

// Wrap Writing to the map with RLock/RUnlock for concurrent workloads
func writeValues(key string, timesSlice []float32) {
	lock.Lock()
	defer lock.Unlock()
	values[key] = timesSlice
}

// Function to sum the entirety of the passed slice. Return that sum.
func sumSlice(slice []float32) float32{
	sum := float32(0)

	for _, number := range slice{
		sum += number
	}

	return sum
}
