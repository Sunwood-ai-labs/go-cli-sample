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

## 📺 実行結果

実行すると、こんな流れで JSON 操作が行われるよ〜：

```bash
$ go run main.go
```

```
📋 JSON パース・生成サンプル ✨
========================================

1️⃣ JSON をデコードするよ〜
✅ デコード成功！
名前: ギャルエンジニア
年齢: 25
メール: gal@example.com
趣味: [コーディング カフェ巡り ショッピング]
アクティブ: true

2️⃣ 構造体を JSON にエンコードするよ〜
📄 コンパクトな JSON:
{"name":"イケてるエンジニア","age":30,"email":"cool@example.com","hobbies":["Go言語","OSS開発","技術ブログ"],"is_active":true}

3️⃣ きれいにフォーマットされた JSON を生成するよ〜
📄 フォーマット済み JSON:
{
  "name": "イケてるエンジニア",
  "age": 30,
  "email": "cool@example.com",
  "hobbies": [
    "Go言語",
    "OSS開発",
    "技術ブログ"
  ],
  "is_active": true
}

4️⃣ ネストした構造体を処理するよ〜
... (省略)

5️⃣ JSON ファイルに保存するよ〜
✅ config.json に保存したよ！

6️⃣ JSON ファイルから読み込むよ〜
✅ ファイルから読み込み成功！
アプリ名: Go CLI Sample
バージョン: 1.0.0
管理者: 管理者ちゃん (admin@example.com)

7️⃣ クリーンアップするよ〜
✅ config.json を削除したよ！

========================================
すべての JSON 操作が完了したよ〜！🎉
```

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
