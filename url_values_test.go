package interpolate

import (
	"net/url"
	"testing"
)

func TestFromURLValuesSingleValue(t *testing.T) {
	values := make(url.Values)
	values.Add("name", "Alice")

	result := FromURLValues("Hello %{name}, how are you?", values)
	expected := "Hello Alice, how are you?"

	if result != expected {
		t.Fatalf("Expected: '" + expected + "'. But got: '" + result + "'")
	}
}

func TestFromURLValuesDoubleValue(t *testing.T) {
	values := make(url.Values)
	values.Add("animal", "Turtles")

	result := FromURLValues("I like %{animal}. %{animal} are awesome!", values)
	expected := "I like Turtles. Turtles are awesome!"

	if result != expected {
		t.Fatalf("Expected: '" + expected + "'. But got: '" + result + "'")
	}
}

func TestFromURLValuesIdenticalValue(t *testing.T) {
	values := make(url.Values)
	values.Add("animal", "Turtles")
	values.Add("animal", "Wolves")

	result := FromURLValues("I like %{animal}. %{animal} are awesome!", values)
	expected := "I like Turtles. Turtles are awesome!"

	if result != expected {
		t.Fatalf("Expected: '" + expected + "'. But got: '" + result + "'")
	}
}

func TestFromURLValuesMultipleValues(t *testing.T) {
	values := make(url.Values)
	values.Add("subject", "Info")
	values.Add("message", "New email received!")

	result := FromURLValues("%{subject}: %{message}", values)
	expected := "Info: New email received!"

	if result != expected {
		t.Fatalf("Expected: '" + expected + "'. But got: '" + result + "'")
	}
}

func TestFromURLValuesMissingValue(t *testing.T) {
	values := make(url.Values)

	result := FromURLValues("My secret: %{secret}", values)
	expected := "My secret: "

	if result != expected {
		t.Fatalf("Expected: '" + expected + "'. But got: '" + result + "'")
	}
}
