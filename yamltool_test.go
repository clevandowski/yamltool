package yamltool

import (
	"fmt"
  "testing"
)

func Equals(left []string, right []string) bool {
  if len(left) != len(right) {
    fmt.Printf("[DEBUG] len(left): %v != len(right): %v\n", len(left), len(right))
    return false
  }
  for n, v := range(left) {
    if v != right[n] {
      return false
    }
  }
  return true
}

func TestSplitDocumentsEmpty1(t *testing.T) {
  yaml := ``

  want := make([]string, 0)

  got := SplitDocuments(yaml)

  if ! Equals(want, got) {
    t.Fatalf("SplitDocuments():\nwant:\n'%v'\ngot:\n'%v'\n", want, got)
  }
}

func TestSplitDocumentsEmpty2(t *testing.T) {
  yaml := `---
---
---`

  want := make([]string, 0)

  got := SplitDocuments(yaml)

  if ! Equals(want, got) {
    t.Fatalf("SplitDocuments():\nwant:\n'%v'\ngot:\n'%v'\n", want, got)
  }
}

func TestSplitDocumentsOneDocument(t *testing.T) {
  yaml := `coincoin: 5`

  want := make([]string, 0)
  want = append(want, `coincoin: 5
`)

  got := SplitDocuments(yaml)

  if ! Equals(want, got) {
    t.Fatalf("SplitDocuments():\nwant:\n'%v'\ngot:\n'%v'\n", want, got)
  }
}


func TestSplitDocuments1(t *testing.T) {
  yaml := `coincoin: 5

---
cuicui: 10
`

  want := make([]string, 0)
  want = append(want, `coincoin: 5
`)
  want = append(want, `cuicui: 10
`)

  got := SplitDocuments(yaml)

  if ! Equals(want, got) {
    t.Fatalf("SplitDocuments(): want: '%v'\ngot: '%v'\n", want, got)
  }
}

func TestSplitDocuments2(t *testing.T) {
  yaml := `---
coincoin: 5

---
cuicui: 10
`

want := make([]string, 0)
want = append(want, `coincoin: 5
`)
want = append(want, `cuicui: 10
`)

  got := SplitDocuments(yaml)

  if ! Equals(want, got) {
    t.Fatalf("SplitDocuments(): want: '%v'\ngot: '%v'\n", want, got)
  }
}
