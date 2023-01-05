mock:
	mockgen --build_flags=--mod=mod -package mockdb -destination storage/mockdb/storage.go github.com/ahror0204/mocking/storage StorageI