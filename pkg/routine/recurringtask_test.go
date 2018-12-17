// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package routine_test

import (
	"context"
	"testing"
	"time"

	"github.com/facebookgo/clock"
	"github.com/stretchr/testify/assert"

	"github.com/Vedaad-Shakib/IoTeX-Sim/pkg/routine"
)

type MockHandler struct {
	Count uint
}

func (h *MockHandler) Do() {
	h.Count++
}

func TestRecurringTask(t *testing.T) {
	h := &MockHandler{Count: 0}
	ctx := context.Background()
	ck := clock.NewMock()
	task := routine.NewRecurringTask(h.Do, 100*time.Millisecond, routine.WithClock(ck))
	task.Start(ctx)
	defer func() {
		task.Stop(ctx)
	}()

	ck.Add(600 * time.Millisecond)
	assert.True(t, h.Count >= 5)
}