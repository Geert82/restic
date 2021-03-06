package restic

import (
	"context"
	"restic/crypto"
)

// Repository stores data in a backend. It provides high-level functions and
// transparently encrypts/decrypts data.
type Repository interface {

	// Backend returns the backend used by the repository
	Backend() Backend

	Key() *crypto.Key

	SetIndex(Index)

	Index() Index
	SaveFullIndex(context.Context) error
	SaveIndex(context.Context) error
	LoadIndex(context.Context) error

	Config() Config

	LookupBlobSize(ID, BlobType) (uint, error)

	List(context.Context, FileType) <-chan ID
	ListPack(context.Context, ID) ([]Blob, int64, error)

	Flush() error

	SaveUnpacked(context.Context, FileType, []byte) (ID, error)
	SaveJSONUnpacked(context.Context, FileType, interface{}) (ID, error)

	LoadJSONUnpacked(context.Context, FileType, ID, interface{}) error
	LoadAndDecrypt(context.Context, FileType, ID) ([]byte, error)

	LoadBlob(context.Context, BlobType, ID, []byte) (int, error)
	SaveBlob(context.Context, BlobType, []byte, ID) (ID, error)

	LoadTree(context.Context, ID) (*Tree, error)
	SaveTree(context.Context, *Tree) (ID, error)
}

// Deleter removes all data stored in a backend/repo.
type Deleter interface {
	Delete(context.Context) error
}

// Lister allows listing files in a backend.
type Lister interface {
	List(context.Context, FileType) <-chan string
}

// Index keeps track of the blobs are stored within files.
type Index interface {
	Has(ID, BlobType) bool
	Lookup(ID, BlobType) ([]PackedBlob, error)
	Count(BlobType) uint
}
