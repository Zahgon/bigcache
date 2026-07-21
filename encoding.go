package bigcache

const (
	timestampSizeInBytes = 8
	hashSizeInBytes      = 8
	keySizeInBytes       = 2
	headersSizeInBytes   = timestampSizeInBytes + hashSizeInBytes + keySizeInBytes
)

func wrapEntry(timestamp uint64, hash uint64, key string, entry []byte, buffer *[]byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func appendToWrappedEntry(timestamp uint64, wrappedEntry []byte, entry []byte, buffer *[]byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func readEntry(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func readTimestampFromEntry(data []byte) uint64 { _ = "STUB: not implemented"; return 0 }

func readKeyFromEntry(data []byte) string { _ = "STUB: not implemented"; return "" }

func compareKeyFromEntry(data []byte, key string) bool { _ = "STUB: not implemented"; return false }

func readHashFromEntry(data []byte) uint64 { _ = "STUB: not implemented"; return 0 }

func resetHashFromEntry(data []byte) { _ = "STUB: not implemented"; return }
