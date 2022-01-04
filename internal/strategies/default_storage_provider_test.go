package strategies

import (
	"os"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultStorageProvider(t *testing.T) {

	baseDir := path.Join(os.TempDir(), strconv.FormatInt(time.Now().Unix(), 16))
	toMoveDir, _ := os.MkdirTemp("", "test")

	file, _ := os.Create(path.Join(toMoveDir, "content.txt"))
	file.WriteString("Ola mundo")
	file.Close()

	storage := NewDefaultStorageProvider()

	t.Run("O baseDir não existe", func(t *testing.T) {
		exists, err := storage.DirExists(baseDir)

		assert.Nil(t, err)
		assert.False(t, exists)
	})

	t.Run("O toMoveDir existe", func(t *testing.T) {
		exists, err := storage.DirExists(toMoveDir)

		assert.Nil(t, err)
		assert.True(t, exists)
	})

	t.Run("Cria a baseDir com directoria filha", func(t *testing.T) {
		err := storage.Mkdir(path.Join(baseDir, "filha"))
		assert.Nil(t, err)
	})

	t.Run("Move toMoveDir para a baseDir", func(t *testing.T) {
		err := storage.Move(toMoveDir, baseDir)
		assert.Nil(t, err)
	})
}
