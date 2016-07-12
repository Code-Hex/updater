package updater

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdater(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(5)
	// Check safe
	go func() {
		result, err := CheckWithTag("Code-Hex", "pget", "0.0.1")
		if err != nil {
			t.Errorf("failed CheckWithTag method: %s", err.Error())
		}
		assert.Equal(t, result, "update available. version: 0.0.2", "expected 'update available. version: 0.0.2' got %s", result)
		wg.Done()
	}()

	go func() {
		result, err := CheckWithTitle("Code-Hex", "pget", "0.0.1")
		if err != nil {
			t.Errorf("failed CheckWithTitle method: %s", err.Error())
		}
		assert.Equal(t, result, "update available. version: v0.0.2", "expected 'update available. version: v0.0.2' got %s", result)
		wg.Done()
	}()

	// Check not update
	go func() {
		result, err := Check("Code-Hex", "pget", "0.0.3", "tag_name")
		if err != nil {
			t.Errorf("failed Check method: %s", err.Error())
		}
		assert.Equal(t, result, "update not available.", "expected 'update not available.' got %s", result)
		wg.Done()
	}()

	// Check error
	go func() {
		_, err := Check("&%&", "", "0.0.3", "tag_name")
		if err == nil {
			t.Errorf("failed check error")
		} else {
			fmt.Fprintf(os.Stdout, "gotcha error: %s\n", err.Error())
		}
		wg.Done()
	}()

	go func() {
		_, err := Check("Code-Hex", "pget", "0.0.3", "tme")
		if err == nil {
			t.Errorf("failed check error")
		} else {
			fmt.Fprintf(os.Stdout, "gotcha error: %s\n", err.Error())
		}
		wg.Done()
	}()
	wg.Wait()
}
