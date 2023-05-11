package testutil

import (
	"os"

	. "github.com/onsi/gomega"
)

func WriteFile(filepath string, content string) {
	Expect(os.WriteFile(filepath, []byte(content), 0644)).To(Succeed())
}
