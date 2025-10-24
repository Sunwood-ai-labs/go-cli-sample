# 07. JSON パース・生成 📋

## 📖 概要

Go で JSON を扱う方法を学ぶよ〜！
設定ファイルの読み書きや、API のレスポンス処理に必須のテクニックだよ✨

## 🎯 学べること

- `json.Unmarshal()` で JSON を Go 構造体にデコード 🔄
- `json.Marshal()` で Go 構造体を JSON にエンコード 🔄
- `json.MarshalIndent()` できれいにフォーマットされた JSON を生成 ✨
- JSON タグ（`json:"name"`）の使い方 🏷️
- ネストした構造体の JSON 処理 🏗️
- JSON ファイルの読み書き 💾

## 💻 使い方

```bash
# このディレクトリに移動
cd examples/07-json-parsing

# 実行
go run main.go
```

実行すると、こんな流れで JSON 操作が行われるよ〜：

1. JSON 文字列をデコード 🔄
2. 構造体をコンパクトな JSON にエンコード 📦
3. きれいにフォーマットされた JSON を生成 ✨
4. ネストした構造体を処理 🏗️
5. JSON ファイルに保存 💾
6. JSON ファイルから読み込み 📖
7. クリーンアップ 🧹

## 🔥 ポイント

- **JSON タグ**：`json:"name"` でフィールド名を指定できるよ 🏷️
- **Unmarshal**：JSON → Go 構造体（ポインタを渡すこと！）
- **Marshal**：Go 構造体 → JSON（コンパクトな1行）
- **MarshalIndent**：きれいにフォーマットされた JSON を生成 ✨
- **大文字小文字**：構造体のフィールドは大文字で始めないとエクスポートされないよ！📌
- **省略可能なフィールド**：`omitempty` タグを使えばゼロ値を省略できる！

## 📚 JSON タグの例

```go
type Person struct {
    Name  string `json:"name"`                    // フィールド名を指定
    Age   int    `json:"age,omitempty"`           // ゼロ値なら省略
    Email string `json:"email_address"`           // 別名を使う
    Skip  string `json:"-"`                       // JSON に含めない
}
```

## 🎨 応用例

このテクニックを使えば、こんなことができちゃう：

- 設定ファイル（config.json）の読み書き ⚙️
- API レスポンスのパース 🌐
- データのエクスポート/インポート 💾
- ログの JSON フォーマット出力 📝

JSON は現代のプログラミングに欠かせないよ〜！マスターしよう💪✨
