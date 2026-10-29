package maybe

import (
	"strconv"
	"testing"
)

func TestJustAndNothing(t *testing.T) {
	j := Just(42)
	if !HasValue(j) {
		t.Error("Just should have a value")
	}
	if v, ok := j.Value(); !ok || v != 42 {
		t.Errorf("Just(42).Value() = (%v, %v), want (42, true)", v, ok)
	}

	n := Nothing[int]()
	if HasValue(n) {
		t.Error("Nothing should not have a value")
	}
	if v, ok := n.Value(); ok || v != 0 {
		t.Errorf("Nothing.Value() = (%v, %v), want (0, false)", v, ok)
	}
}

func TestJustZeroValueIsStillJust(t *testing.T) {
	if !HasValue(Just(0)) {
		t.Error("Just(0) must be distinguishable from Nothing")
	}
	if !HasValue(Just("")) {
		t.Error(`Just("") must be distinguishable from Nothing`)
	}
}

func TestZeroValueMaybeIsNothing(t *testing.T) {
	var m Maybe[int]
	if HasValue(m) {
		t.Error("zero value Maybe should be Nothing")
	}
}

func TestWithDefault(t *testing.T) {
	if got := WithDefault(0, Just(100)); got != 100 {
		t.Errorf("WithDefault(0, Just(100)) = %v, want 100", got)
	}
	if got := WithDefault(100, Nothing[int]()); got != 100 {
		t.Errorf("WithDefault(100, Nothing) = %v, want 100", got)
	}
}

func TestMap(t *testing.T) {
	double := func(x int) int { return x * 2 }

	if got := Map(double, Just(3)); got != Just(6) {
		t.Errorf("Map(double, Just(3)) = %v, want Just(6)", got)
	}
	if got := Map(double, Nothing[int]()); got != Nothing[int]() {
		t.Errorf("Map(double, Nothing) = %v, want Nothing", got)
	}
}

func TestMapChangesType(t *testing.T) {
	got := Map(strconv.Itoa, Just(7))
	if got != Just("7") {
		t.Errorf("Map(Itoa, Just(7)) = %v, want Just(\"7\")", got)
	}
}

func TestMap2(t *testing.T) {
	add := func(a, b int) int { return a + b }

	tests := []struct {
		name string
		a, b Maybe[int]
		want Maybe[int]
	}{
		{"both just", Just(1), Just(2), Just(3)},
		{"first nothing", Nothing[int](), Just(2), Nothing[int]()},
		{"second nothing", Just(1), Nothing[int](), Nothing[int]()},
		{"both nothing", Nothing[int](), Nothing[int](), Nothing[int]()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map2(add, tt.a, tt.b); got != tt.want {
				t.Errorf("Map2 = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap3(t *testing.T) {
	sum := func(a, b, c int) int { return a + b + c }

	if got := Map3(sum, Just(1), Just(2), Just(3)); got != Just(6) {
		t.Errorf("Map3 all just = %v, want Just(6)", got)
	}
	for i, args := range [][3]Maybe[int]{
		{Nothing[int](), Just(2), Just(3)},
		{Just(1), Nothing[int](), Just(3)},
		{Just(1), Just(2), Nothing[int]()},
	} {
		if got := Map3(sum, args[0], args[1], args[2]); HasValue(got) {
			t.Errorf("Map3 case %d with a Nothing should be Nothing, got %v", i, got)
		}
	}
}

func TestMap4(t *testing.T) {
	sum := func(a, b, c, d int) int { return a + b + c + d }

	if got := Map4(sum, Just(1), Just(2), Just(3), Just(4)); got != Just(10) {
		t.Errorf("Map4 all just = %v, want Just(10)", got)
	}
	for i, args := range [][4]Maybe[int]{
		{Nothing[int](), Just(2), Just(3), Just(4)},
		{Just(1), Nothing[int](), Just(3), Just(4)},
		{Just(1), Just(2), Nothing[int](), Just(4)},
		{Just(1), Just(2), Just(3), Nothing[int]()},
	} {
		if got := Map4(sum, args[0], args[1], args[2], args[3]); HasValue(got) {
			t.Errorf("Map4 case %d with a Nothing should be Nothing, got %v", i, got)
		}
	}
}

func TestMap5(t *testing.T) {
	sum := func(a, b, c, d, e int) int { return a + b + c + d + e }

	if got := Map5(sum, Just(1), Just(2), Just(3), Just(4), Just(5)); got != Just(15) {
		t.Errorf("Map5 all just = %v, want Just(15)", got)
	}
	for i, args := range [][5]Maybe[int]{
		{Nothing[int](), Just(2), Just(3), Just(4), Just(5)},
		{Just(1), Nothing[int](), Just(3), Just(4), Just(5)},
		{Just(1), Just(2), Nothing[int](), Just(4), Just(5)},
		{Just(1), Just(2), Just(3), Nothing[int](), Just(5)},
		{Just(1), Just(2), Just(3), Just(4), Nothing[int]()},
	} {
		if got := Map5(sum, args[0], args[1], args[2], args[3], args[4]); HasValue(got) {
			t.Errorf("Map5 case %d with a Nothing should be Nothing, got %v", i, got)
		}
	}
}

func TestAndThen(t *testing.T) {
	parse := func(s string) Maybe[int] {
		n, err := strconv.Atoi(s)
		if err != nil {
			return Nothing[int]()
		}
		return Just(n)
	}

	if got := AndThen(parse, Just("42")); got != Just(42) {
		t.Errorf("AndThen(parse, Just(\"42\")) = %v, want Just(42)", got)
	}
	if got := AndThen(parse, Just("abc")); HasValue(got) {
		t.Errorf("AndThen(parse, Just(\"abc\")) = %v, want Nothing", got)
	}
	if got := AndThen(parse, Nothing[string]()); HasValue(got) {
		t.Errorf("AndThen(parse, Nothing) = %v, want Nothing", got)
	}
}

func TestAndThenNotCalledOnNothing(t *testing.T) {
	called := false
	AndThen(func(x int) Maybe[int] {
		called = true
		return Just(x)
	}, Nothing[int]())
	if called {
		t.Error("AndThen must not call f when the input is Nothing")
	}
}

func TestMapNotCalledOnNothing(t *testing.T) {
	called := false
	Map(func(x int) int {
		called = true
		return x
	}, Nothing[int]())
	if called {
		t.Error("Map must not call f when the input is Nothing")
	}
}

func TestMapNotCalledWhenAnyArgIsNothing(t *testing.T) {
	called := false
	Map2(func(a, b int) int {
		called = true
		return a + b
	}, Just(1), Nothing[int]())
	if called {
		t.Error("Map2 must not call f when any argument is Nothing")
	}
}
