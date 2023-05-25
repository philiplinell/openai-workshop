# github.com/golang/go

## commit message

```sh
commit 21e1ffe7bc7793a8f1c1fca4dfa32f1f3f70681c (HEAD)
Author: Carlo Alberto Ferraris <cafxx@strayorange.com>
Date:   Fri Jan 21 14:21:38 2022 +0900

    bufio: implement large write forwarding in Writer.WriteString

    Currently bufio.Writer implements forwarding to the underlying Writer
    for large writes via Write, but it does not do the same for large
    writes via WriteString.

    If the underlying Writer is also a StringWriter, use the same "large
    writes" logic also in WriteString while taking care to only check
    once per call to WriteString whether the underlying Writer implements
    StringWriter.

    Change-Id: Id81901c07b035936816b9e41b1f5688e699ee8e9
    Reviewed-on: https://go-review.googlesource.com/c/go/+/380074
    Reviewed-by: Ian Lance Taylor <iant@google.com>
    Run-TryBot: Ian Lance Taylor <iant@google.com>
    Auto-Submit: Ian Lance Taylor <iant@google.com>
    TryBot-Result: Gopher Robot <gobot@golang.org>
    Reviewed-by: Dmitri Shuralyov <dmitshur@google.com>


```

## diff

```sh
diff --git a/src/bufio/bufio.go b/src/bufio/bufio.go
index 7483946fc0..bcc273c78b 100644
--- a/src/bufio/bufio.go
+++ b/src/bufio/bufio.go
@@ -731,13 +731,28 @@ func (b *Writer) WriteRune(r rune) (size int, err error) {
 // If the count is less than len(s), it also returns an error explaining
 // why the write is short.
 func (b *Writer) WriteString(s string) (int, error) {
+	var sw io.StringWriter
+	tryStringWriter := true
+
 	nn := 0
 	for len(s) > b.Available() && b.err == nil {
-		n := copy(b.buf[b.n:], s)
-		b.n += n
+		var n int
+		if b.Buffered() == 0 && sw == nil && tryStringWriter {
+			// Check at most once whether b.wr is a StringWriter.
+			sw, tryStringWriter = b.wr.(io.StringWriter)
+		}
+		if b.Buffered() == 0 && tryStringWriter {
+			// Large write, empty buffer, and the underlying writer supports
+			// WriteString: forward the write to the underlying StringWriter.
+			// This avoids an extra copy.
+			n, b.err = sw.WriteString(s)
+		} else {
+			n = copy(b.buf[b.n:], s)
+			b.n += n
+			b.Flush()
+		}
 		nn += n
 		s = s[n:]
-		b.Flush()
 	}
 	if b.err != nil {
 		return nn, b.err
diff --git a/src/bufio/bufio_test.go b/src/bufio/bufio_test.go
index ff3396e946..b3456d2341 100644
--- a/src/bufio/bufio_test.go
+++ b/src/bufio/bufio_test.go
@@ -762,6 +762,67 @@ func TestWriteString(t *testing.T) {
 	}
 }
 
+func TestWriteStringStringWriter(t *testing.T) {
+	const BufSize = 8
+	{
+		tw := &teststringwriter{}
+		b := NewWriterSize(tw, BufSize)
+		b.WriteString("1234")
+		tw.check(t, "", "")
+		b.WriteString("56789012")   // longer than BufSize
+		tw.check(t, "12345678", "") // but not enough (after filling the partially-filled buffer)
+		b.Flush()
+		tw.check(t, "123456789012", "")
+	}
+	{
+		tw := &teststringwriter{}
+		b := NewWriterSize(tw, BufSize)
+		b.WriteString("123456789")   // long string, empty buffer:
+		tw.check(t, "", "123456789") // use WriteString
+	}
+	{
+		tw := &teststringwriter{}
+		b := NewWriterSize(tw, BufSize)
+		b.WriteString("abc")
+		tw.check(t, "", "")
+		b.WriteString("123456789012345")      // long string, non-empty buffer
+		tw.check(t, "abc12345", "6789012345") // use Write and then WriteString since the remaining part is still longer than BufSize
+	}
+	{
+		tw := &teststringwriter{}
+		b := NewWriterSize(tw, BufSize)
+		b.Write([]byte("abc")) // same as above, but use Write instead of WriteString
+		tw.check(t, "", "")
+		b.WriteString("123456789012345")
+		tw.check(t, "abc12345", "6789012345") // same as above
+	}
+}
+
+type teststringwriter struct {
+	write       string
+	writeString string
+}
+
+func (w *teststringwriter) Write(b []byte) (int, error) {
+	w.write += string(b)
+	return len(b), nil
+}
+
+func (w *teststringwriter) WriteString(s string) (int, error) {
+	w.writeString += s
+	return len(s), nil
+}
+
+func (w *teststringwriter) check(t *testing.T, write, writeString string) {
+	t.Helper()
+	if w.write != write {
+		t.Errorf("write: expected %q, got %q", write, w.write)
+	}
+	if w.writeString != writeString {
+		t.Errorf("writeString: expected %q, got %q", writeString, w.writeString)
+	}
+}
+
 func TestBufferFull(t *testing.T) {
 	const longString = "And now, hello, world! It is the time for all good men to come to the aid of their party"
 	buf := NewReaderSize(strings.NewReader(longString), minReadBufferSize)

```
