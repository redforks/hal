package timeth

import (
	"reflect"
	"spork/testing/reset"
	"time"

	bdd "github.com/onsi/ginkgo"
	"github.com/redforks/hal"
	"github.com/stretchr/testify/assert"
)

var _ = bdd.Describe("timeth", func() {
	var (
		assertTimeNotMocked = func() {
			assert.Equal(t(), reflect.ValueOf(time.Now), reflect.ValueOf(hal.Now))
		}

		assertTimeMocked = func() {
			assert.Equal(t(), reflect.ValueOf(now), reflect.ValueOf(hal.Now))
		}
	)

	bdd.BeforeEach(func() {
		reset.Enable()
	})

	bdd.AfterEach(func() {
		reset.Disable()
	})

	bdd.It("Default Time", func() {
		Install()
		assert.Equal(t(), time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local), hal.Now())
	})

	bdd.It("Install", func() {
		assertTimeNotMocked()
		Install()
		assertTimeMocked()
	})

	bdd.It("Auto install", func() {
		assertTimeNotMocked()
		Set(time.Unix(11000000, 0))
		assertTimeMocked()
	})

	bdd.It("Set", func() {
		ti := time.Unix(100000, 0)
		Set(ti)
		assert.Equal(t(), ti, hal.Now())
	})

	bdd.It("Auto Install on Tick", func() {
		assertTimeNotMocked()
		Tick(time.Second)
		assertTimeMocked()
	})

	bdd.It("Tick", func() {
		Tick(time.Second)
		assert.Equal(t(), time.Date(2016, 1, 1, 0, 0, 1, 0, time.Local), hal.Now())

		Tick(-time.Second)
		assert.Equal(t(), time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local), hal.Now())
	})

})
