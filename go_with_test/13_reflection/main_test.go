package main

import (
	"reflect"
	"testing"
)

// func TestWalk(t *testing.T) {
// 	expected := "SachinMaurya"

// 	var got []string

// 	x := struct {
// 		Name string
// 	}{expected}

// 	walk(x, func(input string) {
// 		got = append(got, input)
// 	})
// 	if got[0] != expected {
// 		t.Errorf("got %q, want %q", got[0], expected)
// 	}

// 	if len(got) != 1 {
// 		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
// 	}
// }

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Sachin"},
			[]string{"Sachin"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Sachin", "Mumbai"},
			[]string{"Sachin", "Mumbai"},
		},
		{
			"struct with int and  string field",
			struct {
				Name string
				Age  int
			}{"Sachin", 22},
			[]string{"Sachin"},
		},
		{
			"nested struct",
			Person{
				"Sachin",
				Profile{22, "Mumbai"},
			},
			[]string{"Sachin", "Mumbai"},
		},
		{
			"slices",
			[]Profile{
				{22, "Mumbai"},
				{20, "Goa"},
			},
			[]string{"Mumbai", "Goa"},
		},
		{
			"Array",
			[2]Profile{
				{33, "Mumbai"},
				{12, "Raju"},
			},
			[]string{"Mumbai", "Raju"},
		},
		{
			"map",
			map[string]string{
				"Name":   "Sachin",
				"Gender": "M",
			},
			[]string{"Sachin", "M"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Baz",
			"Baz": "Boz",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Baz")
		assertContains(t, got, "Boz")
	})

	t.Run("with channel", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{22, "Mumbai"}
			aChannel <- Profile{31, "Goa"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Mumbai", "Goa"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)

		}
	})

	t.Run("with Function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{22, "Mumbai"}, Profile{12, "Goa"}
		}

		var got []string
		want := []string{"Mumbai", "Goa"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)

		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)

	}
}
