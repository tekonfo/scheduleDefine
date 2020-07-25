所属しているアカペラサークルにおいて、文化祭における出演バンドのスケジュールを自動調整するシステムを作成する。

## 要件
- バンドは4~6人で構成されている
- 出演者は同時に出演することができない
- 出演場所は二つあり、カフェ、ストが存在する
- 移動時間は10分以上になるようにする必要がある
- カフェは同じ人が連続して出演することができない
- バンドの種類は、本バンド、企画バンド、OBバンドの三種類が存在する
- 希望時間をストは5,10、カフェは10,15で選ぶことができる
- 1回目はスト、カフェかを希望できる
- 2回目以降も出演することができる
- バンドごとに出演不可能時間帯を設けている
- コード巻きの時間が必要
  - とりあえず1時間あたり5分と仮定する

## エンティティの抽出
- バンド
- 出演者
- 出演場所
- スケジュール

バンドごとに行う登録アンケート
https://docs.google.com/forms/d/e/1FAIpQLSfnKAXA472-krfHU3hf8zalhOIFo6RB-y9YV0jYpPivp-4rRw/viewform

実際の進行表例
http://inspishinkou.web.fc2.com/


## このプロジェクトの概要

### 意味
- Goを利用して実際に稼働するプログラムを0から作成しきる
- 就活で見せるポートフォリオにする

### 達成の条件
各種CSVを読み込んで、結果をCSVに書き込む

### 期限
8月末

### 取り組み方
1日1時間やる。なので多分60hくらいある。
休日は2hやる

### タスクフロー
- [x] CSVImport 10h
- [ ] システムの実装 30h
- [ ] CSVOutput 10h
- [ ] システムの統合テスト 10h
- [ ] 簡単な入力でのテスト 10h

--- ここまではやりきる ---

- [ ] GoogleSpreadSheetによるフォームの作成
- [ ] フォームのバリデーションスクリプト記述
- [ ] データの入力・実行





## システム

### ディレクトリ構成


### アルゴリズム

この問題はNP困難である。
直近は深さ優先探索で一つずつ試す。
これで時間がかからないようであればこれでOK

### クラス図

<!-- ```mermaid
classDiagram
    Band : ID
    Band : Name
    Band : DesireLocationID
    Band : BandType
    Band : IsMultiPlay
    Member : ID
    Location: ID
    Schedule: ID
``` -->

### シーケンスフロー

基本的には、上から順に決めていく。
無理ならロールバックする

登録処理

1. 1日目のcafe
2. 1日目のstreet
3. 2日目のcafe
4. 2日目のstreet
5. ...

という順番でスケジュールを決定していく

`impossibleBandOrderは全日程、場所で共有のものを利用する`

スケジュール調整フロー
```mermaid
graph TD
AA{未登録のスケジュールが存在する <br/>  and  まだ歌っていないバンドが存在する} --> |Yes| AB{順序をすべて試した}
AA --> |No| END[終了]
AB --> |No| A
AB --> |Yes| BADEND[エラー発生 errorとして返す]
A[未登録のschedule取得] --> B{当てはまるバンド検索}
B --> |OK| C[scheduleにevent追加]
C --> E[対象bandのisMapped追加]
B --> |False:ロールバック実行| D[event一つ削除]
D --> Z[impossibleBandOrder追加]
Z --> AA
E --> F{コード巻き取りが必要}
F --> |必要| G[scheduleに追加] --> H[timeFromBeforeCodeRollUPを初期化] --> I[最初に戻る]
F --> |不必要| I
```

順序をすべて試したかどうかの判定は
`ImpossibleBandOrder`にbandsのスライスの末尾のbandのみが入っている順序が含まれているかどうかで判定
それが含まれていればすべて試している。

埋まりきらない場合、２回目出演バンドの決定方法

バンド検索の詳細

```mermaid
graph TD
Y{for で groupを一つ選択} --> |未検索のバンドがある| A[一回も歌っていないバンドである]
Y --> |未検索のバンドがない| Z[return false]
A --> B[impossibleBandOrderに含まれていない]
B --> C[不可能時間でない]
C --> D[同時刻にバンドメンバーが別の場所で歌っていない]
D --> E[移動時間の制約に引っかからない]
E --> EE[全ての条件をクリア]
EE --> |Yes| F[return band]
EE --> |No| Y
```


アドバイス
ヒューリスティック法っていうのもあるらしい
  ヒューリスティクス解法
著しくパフォーマンスが劣化する場合に、go routineで回すなどを考える
読み込んだ情報を
インメモリのデータベースにCSVをいれて、

イベントを登録できたら、長く使えるようになる

簡単に実装できる方法：
  UIを用意して、早い者勝ちにする
  衝突した場合は予約制にする？

確定したら通知を送る

空きが出てきて、それに対してわかるようにする
  slackに対して通知する

入力フォーム
slackはWebフォームの変わりもできる
- 埋まり具体
- バリデーションをかけられるのか？
  - CloudFunction → FireStoreでデータを流す
