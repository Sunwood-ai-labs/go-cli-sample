# Test Reports

このディレクトリには、Go CLI Sampleプロジェクトのテスト結果とカバレッジレポートが含まれています。

## 生成日時

2025-10-23

## テスト結果サマリー

### ユニットテスト

- **テスト数**: 3件
- **成功**: 3件
- **失敗**: 0件
- **カバレッジ**: 100.0%

#### テストケース詳細

1. **TestMainOutput**: 基本的な出力をテスト
2. **TestMainWithArgument**: 引数付き実行のテスト
3. **TestMainWithMultipleArguments**: 複数引数のテスト

### ベンチマーク結果

- **BenchmarkMain**:
  - 実行回数: 227,283回
  - 平均実行時間: 5,159 ns/op
  - メモリ割り当て: 48 B/op
  - アロケーション数: 1 allocs/op

## ファイル一覧

- `test-results.txt` - ユニットテスト実行結果の詳細
- `coverage.out` - カバレッジデータ（go tool cover用）
- `coverage.html` - HTMLカバレッジレポート（ブラウザで閲覧可能）
- `benchmark-results.txt` - ベンチマークテスト実行結果

## レポートの閲覧方法

### HTMLカバレッジレポート

```bash
# ブラウザでHTMLレポートを開く
open test-reports/coverage.html  # macOS
xdg-open test-reports/coverage.html  # Linux
start test-reports/coverage.html  # Windows
```

### カバレッジデータの確認

```bash
# ターミナルでカバレッジを確認
go tool cover -func=test-reports/coverage.out
```

## テストの再実行方法

```bash
# すべてのテストとレポートを再生成
go test -v -cover 2>&1 | tee test-reports/test-results.txt
go test -coverprofile=test-reports/coverage.out
go tool cover -html=test-reports/coverage.out -o test-reports/coverage.html
go test -bench=. -benchmem 2>&1 | tee test-reports/benchmark-results.txt
```
