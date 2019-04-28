package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search map", func(t *testing.T) {
		dictionary := map[string]string{
			"foo": "bar",
		}

		got := Search(dictionary, "foo")
		want := "bar"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("search dictionary", func(t *testing.T) {
		dictionary := Dictionary{
			"foo": "bar",
		}

		got, _ := dictionary.Search("foo")
		want := "bar"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("search dictionary not found", func(t *testing.T) {
		dictionary := Dictionary{
			"foo": "bar",
		}

		_, err := dictionary.Search("foofoo")

		if err == nil {
			t.Fatal("expected to get an error")
		}
		if err != ErrNotFound {
			t.Errorf("unexpected error: %s", err.Error())
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("add dictionary OK", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Add("foo", "bar")
		if err != nil {
			t.Fatal("not expecting error:", err)
		}

		want := "bar"
		got, err := dictionary.Search("foo")

		if err != nil {
			t.Fatal("not expecting error:", err)
		}
		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	})

	t.Run("add dictionary rejected", func(t *testing.T) {
		dictionary := Dictionary{
			"foo": "bar",
		}

		err := dictionary.Add("foo", "baz")
		if err == nil {
			t.Fatal("was expecting error")
		}
		if err != ErrExists {
			t.Errorf("want '%s' got '%s'", ErrExists, err)
		}

		want := "bar"
		got, err := dictionary.Search("foo")

		if err != nil {
			t.Fatal("not expecting error:", err)
		}
		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing value", func(t *testing.T) {
		dictionary := Dictionary{
			"foo": "bar",
		}

		dictionary.Delete("foo")
		_, err := dictionary.Search("foo")

		if err == nil {
			t.Fatal("expecting error")
		}
		if err != ErrNotFound {
			t.Errorf("want '%s' got '%s'", ErrNotFound, err)
		}
	})
}
