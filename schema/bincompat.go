package schema

// The compatibility of the binaries, defaults to "auto" (automatically guessed), can be
// "full" (compatible with both Windows and Unix-based systems) and "proxy" (only bash-style
// proxy).
type BinCompat string
const (
	Auto BinCompat = "auto"
	Full BinCompat = "full"
	Proxy BinCompat = "proxy"
	Symlink BinCompat = "symlink"
)
