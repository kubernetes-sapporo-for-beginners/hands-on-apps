# hands-on-apps

[Kubernetes Sapporo for Beginners](https://sapporo-beginner-kubernetes.connpass.com/)で利用するハンズオン用のアプリケーションです。

ハンズオンに関しては、<br>
https://kubernetes-sapporo-for-beginners.github.io/hands-on/ <br>
を確認して下さい。

# Applications

## [greeting-api](./greeting-api)

APIアプリケーションです。

下記2つのエンドポイントを持ちます。

- `/hello` : 挨拶を返します。
- `/health` : ヘルスチェック用です。

### hello API

`こんにちわ` と返します。<br>
URLクエリーとして、 `id` が存在する場合、その情報を `${APP_LOG_DIR}/app.log` に追記書き込みします。

### health API

環境変数 `APP_LOG_DIR` が存在するディレクトリであれば、HTTP200,OKを返します。<br>
上記以外の場合は、HTTP500のエラーを返します。

