/*
 * Food tasting.  Yum.
 */

package food

import (
	"math"
	"testing"

	"github.com/DiscoViking/goBrains/locationmanager"
)

// Basic food verification.
func TestFood(t *testing.T) {
	var food *Food
	cm := locationmanager.NewLocationManager(800, 800)

	testContents := []float64{
		0, // Empty food object.
		16,
		30,
		64,
	}

	for _, val := range testContents {
		t.Log("Test food with contents:", val)
		food = New(cm, val)

		// Content should be as entered.
		checkContent(t, food, val)

		// Food should not immediately disappear unless it has zero content.
		// If it does disappear, then abort this test.
		if (val != 0) && (!checkCheck(t, food, false)) {
			return
		}

		// Content should have decayed.
		val = val - decay_rate
		if val < 0 {
			val = 0
		}
		checkContent(t, food, val)

		// Radius is the square root of the content.
		checkRadius(t, food, math.Sqrt(val))

		// Test simple food consumption.
		checkConsumption(t, food, val)

		// Get the food colour.
		checkColour(t, food)

		// Reset.
		food = New(cm, val)

		// Ensure that we cannot eat more food than there is in the instance.
		checkEmptying(t, food, val)

		// Food should disappear once finished.
		checkCheck(t, food, true)
	}
}

// Content checking.  Verify that the appropriate function used by external tests also works.
func checkContent(t *testing.T, food *Food, content float64) {
	if food.content != content {
		t.Errorf("Expected content of %v, found %v.", content, food.content)
	}
	if food.content != food.GetContent() {
		t.Errorf("Different results from direct and indirect content queries: %v/%v",
			food.content,
			food.GetContent(),
		)
	}
}

// Checking the check function.  Poor naming conventions bite badly.
// Return whether the food still exists after the check.
func checkCheck(t *testing.T, food *Food, result bool) bool {
	if !(result == food.Check()) {
		t.Errorf("Expected check result of %v, got %v.", result, !result)
		return false
	}
	return true
}

// Radius checking.
func checkRadius(t *testing.T, food *Food, radius float64) {
	if food.Radius() != radius {
		t.Errorf("Expected radius of %v, found %v.", food.Radius(), radius)
	}
}

// Food consumption.
func checkConsumption(t *testing.T, food *Food, content float64) {
	resp := food.Consume()

	switch {
	case (content > 0) && (resp == 0):
		t.Errorf("Expected consumption of food, but none reported.")
	case (content == 0) && (resp != 0):
		t.Errorf("Expected no comsumption of food, but %v reported.", resp)
	}
}

// Check food has a colour. Belt-and-braces test only.
func checkColour(t *testing.T, f *Food) {
	f.Color()
}

// Check being greedy.  If we try and eat all the food, does it stop?
func checkEmptying(t *testing.T, food *Food, content float64) {
	var ii float64
	currCont := content

	for ii = 0; ii < (content + 2); ii++ {
		checkConsumption(t, food, currCont)

		// Food consumed.
		currCont--
		if currCont < 0 {
			currCont = 0
		}

		checkContent(t, food, currCont)
	}

	// Radius should zero now.
	checkRadius(t, food, 0)
}
