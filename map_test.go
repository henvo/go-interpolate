package interpolate

import (
	"testing"
)

func TestFromMapSingleValue(t *testing.T) {
	m := make(map[string]interface{})

	m["name"] = "Bob"

	result := FromMap("Hello %{name}, how are you?", m)
	expected := "Hello Bob, how are you?"

	if result != expected {
		t.Fatalf("Expected: '" + expected + "'. But got: '" + result + "'")
	}
}

func TestFromMapDoubleValue(t *testing.T) {
	m := make(map[string]interface{})

	m["animal"] = "Turtles"

	result := FromMap("I like %{animal}! %{animal} are awesome!", m)
	expected := "I like Turtles! Turtles are awesome!"

	if result != expected {
		t.Fatalf("Expected: '" + expected + "'. But got: '" + result + "'")
	}
}

func TestFromMapMultiValues(t *testing.T) {
	m := make(map[string]interface{})

	m["type"] = "Info"
	m["pre"] = "You've got"
	m["post"] = "mails!"
	m["count"] = 3

	result := FromMap("%{type}: %{pre} %{count} %{post}", m)
	expected := "Info: You've got 3 mails!"

	if result != expected {
		t.Fatalf("Expected: '" + expected + "'. But got: '" + result + "'")
	}
}

func TestFromMapMissingValue(t *testing.T) {
	m := make(map[string]interface{})

	result := FromMap("Hello %{audience}!", m)
	expected := "Hello !"

	if result != expected {
		t.Fatalf("Expected: '" + expected + "'. But got: '" + result + "'")
	}
}
