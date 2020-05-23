package minicom_test

import (
	"testing"

	"github.com/hsmtkk/c024_minicom/pkg/minicom"
	"github.com/stretchr/testify/assert"
)

func TestCompute0(t *testing.T) {
	com := minicom.New()
	com.Set(1, 10)
	com.Set(2, 20)
	com.Add(40)
	want1 := 10
	want2 := 50
	got1, got2 := com.Output()
	assert.Equal(t, want1, got1, "should be equal")
	assert.Equal(t, want2, got2, "should be equal")
}

func TestCompute1(t *testing.T) {
	com := minicom.New()
	com.Set(1, -23)
	com.Sub(77)
	com.Set(1, 0)
	want1 := 0
	want2 := -100
	got1, got2 := com.Output()
	assert.Equal(t, want1, got1, "should be equal")
	assert.Equal(t, want2, got2, "should be equal")
}
