package jerasure

// #cgo CPPFLAGS: -I../gf-complete/internal/include
// #cgo CPPFLAGS: -Iinternal/include
// #cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
// #cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all
import "C"
