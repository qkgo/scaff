package buffers

import (
	"context"
	"io"
)

// https://ixday.github.io/post/golang-cancel-copy/

// here is some syntaxic sugar inspired by the Tomas Senart's video,
// it allows me to inline the Reader interface
type readerFunc func(p []byte) (n int, err error)

func (rf readerFunc) Read(p []byte) (n int, err error) { return rf(p) }

// slightly modified function signature:
// - context has been added in order to propagate cancelation
// - I do not return the number of bytes written, has it is not useful in my use case
func CancelableCopy(ctx context.Context, dst io.Writer, src io.Reader, closeFunc func()) (int64, error) {

	// Copy will call the Reader and Writer interface multiple time, in order
	// to copy by chunk (avoiding loading the whole file in memory).
	// I insert the ability to cancel before read time as it is the earliest
	// possible in the call process.
	size, err := io.Copy(dst, readerFunc(func(p []byte) (int, error) {
		// golang non-blocking channel: https://gobyexample.com/non-blocking-channel-operations
		//log.Println("io copy loop")
		select {
		// if context has been canceled

		case <-ctx.Done():
			// stop process and propagate "context canceled" error
			closeFunc()
			return 0, ctx.Err()
		default:
			// otherwise just run default io.Reader implementation
			//log.Println("reading bytes ")
			return src.Read(p)
		}
	}))
	return size, err
}
