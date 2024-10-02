// Sample program to teach the mechanics of escape analysis.
package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// main is the entry point for the application.
func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
}

// createUserV1 creates a user value and passed
// a copy back to the caller.
//
//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)

	return u
}

// createUserV2 creates a user value and shares
// the value with the caller.
//
//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)

	return &u
}

/*
// See escape analysis and inlining decisions.

$ go build -gcflags -m=2
# github.com/weitecklee/ardanlabs_ultimate_go_advanced_concepts/chapter02/pointers/example4
./example4.go:21:6: cannot inline createUserV1: marked go:noinline
./example4.go:36:6: cannot inline createUserV2: marked go:noinline
./example4.go:10:6: cannot inline main: function too complex: cost 132 exceeds budget 80
./example4.go:37:2: u escapes to heap:
./example4.go:37:2:   flow: ~r0 = &u:
./example4.go:37:2:     from &u (address-of) at ./example4.go:44:9
./example4.go:37:2:     from return &u (return) at ./example4.go:44:2
./example4.go:37:2: moved to heap: u

// See the intermediate representation phase before
// generating the actual arch-specific assembly.

$ go build -gcflags -S

// See bounds checking decisions.

go build -gcflags="-d=ssa/check_bce/debug=1"

// See the actual machine representation by using
// the disasembler.

$ go tool objdump -s main.main example4

// See a list of the symbols in an artifact with
// annotations and size.

$ go tool nm example4
*/
