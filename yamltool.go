package yamltool

import (
  "strings"
)

func SplitDocuments(rawDocuments string) []string {
  documents := make([]string, 0)
  document := ""
  for _, line := range strings.Split(rawDocuments, "\n") {
    if line == "---" {
      documentTrimed := strings.TrimSpace(document)
      if documentTrimed != "" && documentTrimed != "---" {
        documents = append(documents, document)
        document = ""
      }
    }
    trimedLine := strings.TrimSpace(line)
    if trimedLine != "" && trimedLine != "---" {
      document += line + "\n"
    }
  }
  documentTrimed := strings.TrimSpace(document)
  if documentTrimed != "" && documentTrimed != "---" {
    documents = append(documents, document)
  }

  return documents
}
