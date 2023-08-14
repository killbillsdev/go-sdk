package go_sdk

import (
	"fmt"
	"testing"
	"encoding/json"
)
// change to TestStore to run with the command go test
func testStore(t *testing.T) {
	stores, err := GetStores("prod", "e82376a1-2869-461b-9a6b-1f10bc87bedc")
    if err != nil {
        t.Errorf("Error: %v", err)
        return
    }

    fmt.Println("Stores:", stores)

    jsonData, _ := json.MarshalIndent(stores, "", "  ")
    fmt.Println(string(jsonData))

    // Add assertions here to validate the behavior of your code
    // For example:
    if len(stores) == 0 {
        t.Errorf("Expected stores to be non-empty, but got an empty slice.")
    }
}
