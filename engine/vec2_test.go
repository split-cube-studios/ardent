package engine

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := Vec2{X: 1, Y: 2}
	actual := Vec2{X: 0, Y: 1}.Add(Vec2{X: 1, Y: 1})
	if actual != expected {
		t.Fatalf("Expected %v got %v", expected, actual)
	}
}

func TestSub(t *testing.T) {
	expected := Vec2{X: -1, Y: 0}
	actual := Vec2{X: 0, Y: 1}.Sub(Vec2{X: 1, Y: 1})
	if actual != expected {
		t.Fatalf("Expected %v got %v", expected, actual)
	}
}

func TestAngle(t *testing.T) {
	for _, test := range []struct {
		input  Vec2
		output float64
	}{
		{Vec2{X: 1, Y: 0}, 0},
		{Vec2{X: 1, Y: 1}, math.Pi / 4},
		{Vec2{X: 0, Y: 1}, math.Pi / 2},
		{Vec2{X: -1, Y: 1}, math.Pi * 3 / 4},
		{Vec2{X: -1, Y: -1}, -math.Pi * 3 / 4},
		{Vec2{X: 0, Y: -1}, -math.Pi / 2},
		{Vec2{X: 1, Y: -1}, -math.Pi / 4},
	} {
		angle := test.input.Angle()
		if angle != test.output {
			t.Fatalf(
				"Expected %f for %v got %f",
				test.output,
				test.input,
				angle,
			)
		}
	}
}

func TestAngleTo(t *testing.T) {}

func TestLerp(t *testing.T) {}

func TestDistance(t *testing.T) {}

func TestTranslate(t *testing.T) {}

func TestScale(t *testing.T) {}

func TestDot(t *testing.T) {}

func TestLength(t *testing.T) {}

func TestNormalize(t *testing.T) {}
