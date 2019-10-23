package telemetry_test

import (
	"github.com/FactomProject/factomd/testHelper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimulation(t *testing.T) {
	// Just load simulator
	assert.NotPanics(t, func(){
		testHelper.SetupSim("LFF", map[string]string{}, 10, 0, 0, t)
	})
}
