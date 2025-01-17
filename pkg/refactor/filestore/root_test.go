package filestore_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/direktiv/direktiv/pkg/refactor/filestore"
	"github.com/direktiv/direktiv/pkg/refactor/filestore/filestoresql"
	"github.com/direktiv/direktiv/pkg/refactor/utils"
	"github.com/google/uuid"
)

func TestRoot_CreateFileWithoutRootDirectory(t *testing.T) {
	db, err := utils.NewMockGorm()
	if err != nil {
		t.Fatalf("unepxected NewMockGorm() error = %v", err)
	}
	fs := filestoresql.NewSQLFileStore(db)

	root, err := fs.CreateRoot(context.Background(), uuid.UUID{})
	if err != nil {
		t.Fatalf("unepxected CreateRoot() error = %v", err)
	}

	assertRootErrorFileCreation(t, fs, root.ID, "/file1.text", filestore.ErrNoParentDirectory)
	assertRootErrorFileCreation(t, fs, root.ID, "/dir1", filestore.ErrNoParentDirectory)

	assertRootCorrectFileCreation(t, fs, root.ID, "/")
	assertRootCorrectFileCreation(t, fs, root.ID, "/file1.text")
	assertRootCorrectFileCreation(t, fs, root.ID, "/dir1")
}

func TestRoot_CreateFile(t *testing.T) {
	db, err := utils.NewMockGorm()
	if err != nil {
		t.Fatalf("unepxected NewMockGorm() error = %v", err)
	}
	fs := filestoresql.NewSQLFileStore(db)

	root, err := fs.CreateRoot(context.Background(), uuid.UUID{})
	if err != nil {
		t.Fatalf("unepxected CreateRoot() error = %v", err)
	}

	tests := []struct {
		path    string
		payload string
	}{
		{"/", ""},
		{"/example.text", "abcd"},
		{"/example1.text", "abcd"},
		{"/example2.text", "abcd"},
	}
	for _, tt := range tests {
		t.Run("valid", func(t *testing.T) {
			assertRootCorrectFileCreation(t, fs, root.ID, tt.path)
		})
	}
}

func TestRootQuery_IsEmptyDirectory(t *testing.T) {
	db, err := utils.NewMockGorm()
	if err != nil {
		t.Fatalf("unepxected NewMockGorm() error = %v", err)
	}
	fs := filestoresql.NewSQLFileStore(db)

	root, err := fs.CreateRoot(context.Background(), uuid.New())
	if err != nil {
		t.Fatalf("unepxected CreateRoot() error = %v", err)
	}

	assertEmptyDirectory(t, fs, root.ID, "/", false, filestore.ErrNotFound)
	assertEmptyDirectory(t, fs, root.ID, "/dir1", false, filestore.ErrNotFound)

	assertRootCorrectFileCreation(t, fs, root.ID, "/")
	assertEmptyDirectory(t, fs, root.ID, "/", true, nil)
	assertEmptyDirectory(t, fs, root.ID, "/dir1", false, filestore.ErrNotFound)

	assertRootCorrectFileCreation(t, fs, root.ID, "/file1.text")
	assertRootCorrectFileCreation(t, fs, root.ID, "/file2.text")

	assertRootCorrectFileCreation(t, fs, root.ID, "/dir1")
	assertRootCorrectFileCreation(t, fs, root.ID, "/dir1/file3.text")
	assertRootCorrectFileCreation(t, fs, root.ID, "/dir1/file4.text")

	assertRootCorrectFileCreation(t, fs, root.ID, "/dir2")

	assertEmptyDirectory(t, fs, root.ID, "/", false, nil)
	assertEmptyDirectory(t, fs, root.ID, "/dir1", false, nil)

	assertEmptyDirectory(t, fs, root.ID, "/dir2", true, nil)
}

func assertEmptyDirectory(t *testing.T, fs filestore.FileStore, rootID uuid.UUID, path string, wantEmpty bool, wantErr error) {
	t.Helper()

	gotEmpty, gotErr := fs.ForRootID(rootID).IsEmptyDirectory(context.Background(), path)
	if !errors.Is(gotErr, wantErr) {
		t.Errorf("unexpected IsEmptyDirectory() error, got: %v, want: %v", gotErr, wantErr)

		return
	}
	if gotEmpty != wantEmpty {
		t.Errorf("unexpected IsEmptyDirectory(), got: %v, want %v", gotEmpty, wantEmpty)
	}
}

func assertRootCorrectFileCreation(t *testing.T, fs filestore.FileStore, rootID uuid.UUID, path string) {
	t.Helper()

	var data []byte
	typ := filestore.FileTypeDirectory
	if strings.Contains(path, ".text") {
		data = []byte("some data")
		typ = filestore.FileTypeFile
	}

	assertRootCorrectFileCreationWithContent(t, fs, rootID, path, string(typ), data)
}

func assertRootCorrectFileCreationWithContent(t *testing.T, fs filestore.FileStore, rootID uuid.UUID, path string, typ string, data []byte) {
	t.Helper()

	file, _, err := fs.ForRootID(rootID).CreateFile(context.Background(), path, filestore.FileType(typ), bytes.NewReader(data))
	if err != nil {
		t.Errorf("unexpected CreateFile() error: %v", err)

		return
	}
	if file == nil {
		t.Errorf("unexpected nil file CreateFile()")

		return
	}
	if file.Path != path {
		t.Errorf("unexpected file.Path, got: >%s<, want: >%s<", file.Path, path)

		return
	}

	if typ != "directory" {
		reader, _ := fs.ForFile(file).GetData(context.Background())
		createdData, _ := io.ReadAll(reader)
		if string(createdData) != string(data) {
			t.Errorf("unexpected GetPath(), got: >%s<, want: >%s<", createdData, data)

			return
		}
	}

	file, err = fs.ForRootID(rootID).GetFile(context.Background(), path)
	if err != nil {
		t.Errorf("unexpected GetFile() error: %v", err)

		return
	}
	if file == nil {
		t.Errorf("unexpected nil file GetFile()")

		return
	}
	if file.Path != path {
		t.Errorf("unexpected file.Path, got: >%s<, want: >%s<", file.Path, path)

		return
	}
}

func assertRootErrorFileCreation(t *testing.T, fs filestore.FileStore, rootID uuid.UUID, path string, wantErr error) {
	t.Helper()

	typ := filestore.FileTypeDirectory
	if strings.Contains(path, ".text") {
		typ = filestore.FileTypeFile
	}

	file, rev, gotErr := fs.ForRootID(rootID).CreateFile(context.Background(), path, typ, strings.NewReader(""))
	if file != nil {
		t.Errorf("unexpected none nil CreateFile().file")

		return
	}
	if rev != nil {
		t.Errorf("unexpected none nil CreateFile().revsion")

		return
	}
	if !errors.Is(gotErr, wantErr) {
		t.Errorf("unexpected CreateFile() error, got: %v, want: %v", gotErr, wantErr)

		return
	}
}

func TestRoot_CorrectReadDirectory(t *testing.T) {
	db, err := utils.NewMockGorm()
	if err != nil {
		t.Fatalf("unepxected NewMockGorm() error = %v", err)
	}
	fs := filestoresql.NewSQLFileStore(db)

	root, err := fs.CreateRoot(context.Background(), uuid.New())
	if err != nil {
		t.Fatalf("unepxected CreateRoot() error = %v", err)
	}

	// Test root directory:
	{
		assertRootCorrectFileCreation(t, fs, root.ID, "/")
		assertRootCorrectFileCreation(t, fs, root.ID, "/file1.text")
		assertRootCorrectFileCreation(t, fs, root.ID, "/file2.text")

		assertRootFilesInPath(t, fs, root.ID, "/",
			"/file1.text",
			"/file2.text",
		)
	}

	// Add /dir1 directory:
	{
		assertRootCorrectFileCreation(t, fs, root.ID, "/dir1")
		assertRootCorrectFileCreation(t, fs, root.ID, "/dir1/file3.text")
		assertRootCorrectFileCreation(t, fs, root.ID, "/dir1/file4.text")

		assertRootFilesInPath(t, fs, root.ID, "/dir1",
			"/dir1/file3.text",
			"/dir1/file4.text",
		)
		assertRootFilesInPath(t, fs, root.ID, "/",
			"/file1.text",
			"/file2.text",
			"/dir1",
		)
	}

	// Add /dir1/dir2 directory:
	{
		assertRootCorrectFileCreation(t, fs, root.ID, "/dir1/dir2")
		assertRootCorrectFileCreation(t, fs, root.ID, "/dir1/dir2/file5.text")
		assertRootCorrectFileCreation(t, fs, root.ID, "/dir1/dir2/file6.text")

		assertRootFilesInPath(t, fs, root.ID, "/dir1/dir2",
			"/dir1/dir2/file5.text",
			"/dir1/dir2/file6.text",
		)
		assertRootFilesInPath(t, fs, root.ID, "/dir1",
			"/dir1/file3.text",
			"/dir1/file4.text",
			"/dir1/dir2",
		)
		assertRootFilesInPath(t, fs, root.ID, "/",
			"/file1.text",
			"/file2.text",
			"/dir1",
		)
	}
}

func TestRoot_RenamePath(t *testing.T) {
	db, err := utils.NewMockGorm()
	if err != nil {
		t.Fatalf("unepxected NewMockGorm() error = %v", err)
	}
	fs := filestoresql.NewSQLFileStore(db)

	root, err := fs.CreateRoot(context.Background(), uuid.New())
	if err != nil {
		t.Fatalf("unepxected CreateRoot() error = %v", err)
	}

	// Test root directory:
	{
		assertRootCorrectFileCreation(t, fs, root.ID, "/")
		assertRootCorrectFileCreation(t, fs, root.ID, "/dir1")
		assertRootCorrectFileCreation(t, fs, root.ID, "/dir1/dir2")
		assertRootCorrectFileCreation(t, fs, root.ID, "/dir1/file.text")
	}

	f, err := fs.ForRootID(root.ID).GetFile(context.Background(), "/dir1/file.text")
	if err != nil {
		t.Fatalf("unepxected GetFile() error = %v", err)
	}
	err = fs.ForFile(f).SetPath(context.Background(), "/file.text")
	if err != nil {
		t.Fatalf("unepxected SetPath() error = %v", err)
	}

	assertRootFilesInPath(t, fs, root.ID, "/",
		"/dir1",
		"/file.text",
	)

	f, err = fs.ForRootID(root.ID).GetFile(context.Background(), "/dir1/dir2")
	if err != nil {
		t.Fatalf("unepxected GetFile() error = %v", err)
	}
	err = fs.ForFile(f).SetPath(context.Background(), "/dir2")
	if err != nil {
		t.Fatalf("unepxected SetPath() error = %v", err)
	}

	assertRootFilesInPath(t, fs, root.ID, "/",
		"/dir1",
		"/dir2",
		"/file.text",
	)
}

func TestRoot_CalculateChecksumDirectory(t *testing.T) {
	db, err := utils.NewMockGorm()
	if err != nil {
		t.Fatalf("unepxected NewMockGorm() error = %v", err)
	}
	fs := filestoresql.NewSQLFileStore(db)

	root, err := fs.CreateRoot(context.Background(), uuid.New())
	if err != nil {
		t.Fatalf("unepxected CreateRoot() error = %v", err)
	}

	filestore.DefaultCalculateChecksum = func(data []byte) []byte {
		return []byte(fmt.Sprintf("---%s---", data))
	}

	// Test root directory:
	{
		assertRootCorrectFileCreationWithContent(t, fs, root.ID, "/", "directory", nil)
		assertRootCorrectFileCreationWithContent(t, fs, root.ID, "/file1.text", "text", []byte("content1"))
		assertRootCorrectFileCreationWithContent(t, fs, root.ID, "/file2.text", "text", []byte("content2"))
		assertRootCorrectFileCreationWithContent(t, fs, root.ID, "/dir1", "directory", nil)
		assertRootCorrectFileCreationWithContent(t, fs, root.ID, "/empty_dir", "directory", nil)
		assertRootCorrectFileCreationWithContent(t, fs, root.ID, "/dir1/file3.text", "text", []byte("content3"))
		assertRootCorrectFileCreationWithContent(t, fs, root.ID, "/dir1/file4.text", "text", []byte("content4"))

		assertChecksums(t, fs, root.ID,
			"/", "",
			"/file1.text", "---content1---",
			"/file2.text", "---content2---",
			"/dir1", "",
			"/empty_dir", "",
			"/dir1/file3.text", "---content3---",
			"/dir1/file4.text", "---content4---",
		)
	}
}

func assertRootFilesInPath(t *testing.T, fs filestore.FileStore, rootID uuid.UUID, searchPath string, paths ...string) {
	t.Helper()

	files, err := fs.ForRootID(rootID).ReadDirectory(context.Background(), searchPath)
	if err != nil {
		t.Errorf("unepxected ReadDirectory() error = %v", err)

		return
	}
	if len(files) != len(paths) {
		t.Errorf("unexpected ReadDirectory() length, got: %d, want: %d", len(files), len(paths))

		return
	}

	for i := range paths {
		if files[i].Path != paths[i] {
			t.Errorf("unexpected files[%d].Path , got: >%s<, want: >%s<", i, files[i].Path, paths[i])

			return
		}
	}
}

func assertChecksums(t *testing.T, fs filestore.FileStore, rootID uuid.UUID, paths ...string) {
	t.Helper()

	checksumsMap, err := fs.ForRootID(rootID).CalculateChecksumsMap(context.Background())
	if err != nil {
		t.Errorf("unepxected CalculateChecksumsMap() error = %v", err)
	}
	if len(checksumsMap)*2 != len(paths) {
		t.Errorf("unexpected CalculateChecksumsMap() length, got: %d, want: %d", len(checksumsMap), len(paths)/2)
	}

	wantChecksumsMap := make(map[string]string)

	for i := 0; i < len(paths)-1; i += 2 {
		wantChecksumsMap[paths[i]] = paths[i+1]
	}

	if !reflect.DeepEqual(checksumsMap, wantChecksumsMap) {
		t.Errorf("unexpected CalculateChecksumsMap() result, got: %v, want: %v", checksumsMap, wantChecksumsMap)
	}
}
