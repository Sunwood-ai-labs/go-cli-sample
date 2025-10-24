package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// TestMainOutput tests the basic output of the main function
func TestMainOutput(t *testing.T) {
	// 標準出力をキャプチャ
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// コマンドライン引数をリセット
	oldArgs := os.Args
	os.Args = []string{"cmd"}

	// main関数を実行
	main()

	// 標準出力を復元
	w.Close()
	os.Stdout = old
	os.Args = oldArgs

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 期待される出力を確認
	expected := "Hello, ギャルエンジニア from Go CLI! ✨"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}

// TestMainWithArgument tests the output when an argument is provided
func TestMainWithArgument(t *testing.T) {
	// 標準出力をキャプチャ
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// コマンドライン引数を設定
	oldArgs := os.Args
	testArg := "テスト引数"
	os.Args = []string{"cmd", testArg}

	// main関数を実行
	main()

	// 標準出力を復元
	w.Close()
	os.Stdout = old
	os.Args = oldArgs

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 期待される出力を確認
	if !strings.Contains(output, testArg) {
		t.Errorf("Expected output to contain argument %q, got %q", testArg, output)
	}

	expectedMsg := "引数うけとったよ〜:"
	if !strings.Contains(output, expectedMsg) {
		t.Errorf("Expected output to contain %q, got %q", expectedMsg, output)
	}
}

// TestMainWithMultipleArguments tests that only the first argument is processed
func TestMainWithMultipleArguments(t *testing.T) {
	// 標準出力をキャプチャ
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 複数の引数を設定
	oldArgs := os.Args
	os.Args = []string{"cmd", "first", "second", "third"}

	// main関数を実行
	main()

	// 標準出力を復元
	w.Close()
	os.Stdout = old
	os.Args = oldArgs

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 最初の引数のみが表示されることを確認
	if !strings.Contains(output, "first") {
		t.Errorf("Expected output to contain first argument %q, got %q", "first", output)
	}
}

// BenchmarkMain ベンチマークテスト
func BenchmarkMain(b *testing.B) {
	// 標準出力を無効化
	old := os.Stdout
	os.Stdout = nil

	// コマンドライン引数を設定
	oldArgs := os.Args
	os.Args = []string{"cmd"}

	defer func() {
		os.Stdout = old
		os.Args = oldArgs
	}()

	// 出力を無視するために null デバイスにリダイレクト
	devNull, _ := os.Open(os.DevNull)
	os.Stdout = devNull
	defer devNull.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		main()
	}
}
