# bytewriter
Writer for fixed size byte slices.

## description
This writer allows you to set a byte slice as the target of your writes without the dangers of it getting reallocated when the original slice if full. In some circumstances this is more appropriate than the use of `bytes.Buffer`, for example when trying to insert data into a bigger slice.

## docs
https://pkg.go.dev/github.com/noxer/bytewriter
