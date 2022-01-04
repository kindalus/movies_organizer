package strategies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOMDB(t *testing.T) {
	t.Run("Deve retornar o Matrix de 1999",
		func(t *testing.T) {
			omdb := NewOmdb("4be87309")
			spec, errFind := omdb.Find("The Matrix", 1999)

			assert.Nil(t, errFind)

			assert.Equal(t, spec.Title, "The Matrix")
			assert.Equal(t, spec.Genre, "Action")
			assert.Equal(t, spec.Year, "1999")

		})
}
