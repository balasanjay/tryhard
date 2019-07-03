// +build ignore

package foo_test

func bar() error {
	return nil
}

func foo() error {
	err := bar()
	if err != nil {
		// Zero statements
	}

	err = bar()
	if err != nil {
		// Multiple statements.
		println("foo")
		println("bar")
	}

	err = bar()
	if err != nil {
		// Single statement return.
		return nil
	}

	err = bar()
	if err != nil {
		// Single statement not-return.
		println("foo")
	}
}
