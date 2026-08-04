## カラム名について
| 元の名前 | 正式名称 | 説明 |
| --- | --- | --- |
| country_code | iso_country_code | ISO 3166-1 alpha-2 国コード |
| name_jp | country_name_ja_formal | 日本語の正式国名 |
| name_jps | country_name_ja_common | 日本語の通称（短縮名） |
| capital_jp | capital_name_ja | 日本語の首都名 |
| name_en | country_name_en_formal | 英語の正式国名 |
| name_ens | country_name_en_common | 英語の通称（短縮名） |
| capital_en | capital_name_en | 英語の首都名 |
| lat | capital_latitude | 首都の緯度 |
| lon | capital_longitude | 首都の経度 |

## できそうなこと
- 距離計算（ハバーサイン）
- 最北・最南・最東・最西の国を探す
- 日本から近い国ランキング
- 地図にプロット
- 方角（ベアリング）計算
- 2点間の中間点を求める
- 距離ヒートマップを作る
- 国ごとの行政区分みたいな

## 追加した方がいいカラムのまとめ（一覧）
| カラム名 | 用途 |
| ------ | --------- |
| region_code | 地域分類（アジア・欧州など） |
| sub_region_code | サブ地域分類（東アジアなど） |
| iso_country_code_3 | ISO 3桁コード |
| phone_country_code | 国際電話コード |
| currency_code | 通貨コード |
| is_active | 論理削除 |
| population | 人口 |
| area_sq_km | 国土面積 |
| timezone | 標準タイムゾーン |
| official_language | 公用語 |
| created_by / updated_by | 監査ログ |

# 🌍 地理用語の英語化：日本語訳＋説明（体系別）
## 1. 方位（Cardinal Directions）
| 英語 | 日本語 | 説明 | 
|------|--------|------| 
| north | 北 | 地図の上方向。北極側。 | 
| south | 南 | 地図の下方向。南極側。 | 
| east | 東 | 太陽が昇る方向。 | 
| west | 西 | 太陽が沈む方向。 | 
| northeast | 北東 | 北と東の中間方向。 | 
| northwest | 北西 | 北と西の中間方向。 | 
| southeast | 南東 | 南と東の中間方向。 | 
| southwest | 南西 | 南と西の中間方向。 | 
| northern | 北の | 北側に位置することを示す形容詞。 | 
| southern | 南の | 南側に位置することを示す形容詞。 | 

## 2. 半球（Hemispheres） 
| 英語 | 日本語 | 説明 | 
|------|--------|------| 
| Northern Hemisphere | 北半球 | 赤道より北側の地域。日本・欧州・北米など。 | 
| Southern Hemisphere | 南半球 | 赤道より南側の地域。オーストラリア・南米など。 | 
| Eastern Hemisphere | 東半球 | 本初子午線より東側。アジア・オセアニアなど。 | 
| Western Hemisphere | 西半球 | 本初子午線より西側。アメリカ大陸など。 | 

## 3. 極（Poles / Polar Regions）
| 英語 | 日本語 | 説明 | 
|------|--------|------| 
| North Pole | 北極 | 地球の最北端。 | 
| South Pole | 南極 | 地球の最南端。 | 
| Arctic | 北極圏 | 北緯66.5度以北の地域。 | 
| Antarctic | 南極圏 | 南緯66.5度以南の地域。 | 
| Arctic Circle | 北極圏境界線 | 北緯66.5度の緯線。 | 
| Antarctic Circle | 南極圏境界線 | 南緯66.5度の緯線。 | 

## 4. 緯度・経度（Latitude / Longitude）
| 英語 | 日本語 | 説明 | 
|------|--------|------| 
| latitude | 緯度 | 南北方向の位置。赤道を0度とする。 | 
| longitude | 経度 | 東西方向の位置。本初子午線を0度とする。 | 
| north latitude | 北緯 | 赤道より北の緯度。 | 
| south latitude | 南緯 | 赤道より南の緯度。 | 
| east longitude | 東経 | 本初子午線より東の経度。 | 
| west longitude | 西経 | 本初子午線より西の経度。 | 
| coordinates | 座標 | 緯度・経度の組。 | 
| Prime Meridian | 本初子午線 | 経度0度の基準線。 | 
| Equator | 赤道 | 緯度0度の基準線。 | 

## 5. 地形（Landforms）
| 英語 | 日本語 | 説明 |
|------|--------|------|
| mountain | 山 | 高く盛り上がった地形。 | 
| hill | 丘 | 山より低い高まり。 | 
| valley | 谷 | 山に挟まれた低地。 | 
| plain | 平野 | 平坦な広い土地。 | 
| plateau | 高原 | 高い位置にある平坦地。 | 
| basin | 盆地 | 周囲を山に囲まれた低地。 | 
| canyon | 峡谷 | 深く狭い谷。 | 
| ridge | 尾根 | 山の連なりの高い部分。 | 
| peak | 山頂 | 山の最も高い点。 | 
| volcano | 火山 | マグマが噴出する山。 | 
| crater | 火口 | 火山の噴出口。 | 
| desert | 砂漠 | 降水量が極端に少ない地域。 | 
| oasis | オアシス | 砂漠の中の水がある場所。 | 
| glacier | 氷河 | 長期間の氷の流れ。 | 

## 6. 水域（Bodies of Water）
| 英語 | 日本語 | 説明 |
|------|--------|------|
| ocean | 大洋 | 地球の大きな海域。 |
| sea | 海 | 大洋より小さい海域。 |
| bay | 湾 | 海が陸地に入り込んだ部分。 |
| gulf | 大湾 | bay より大きい湾。 |
| strait | 海峡 | 海と海をつなぐ狭い水路。 |
| lake | 湖 | 陸地に囲まれた水域。 |
| river | 川 | 流れる水の通り道。 |
| delta | 三角州 | 川の河口にできる堆積地形。 |
| waterfall | 滝 | 水が落下する場所。 |

## 7. 気候（Climate）
| 英語 | 日本語 | 説明 |
|------|--------|------|
| climate | 気候 | 長期間の平均的な天気。 |
| weather | 天気 | 短期的な大気の状態。 |
| humidity | 湿度 | 空気中の水分量。 |
| precipitation | 降水 | 雨・雪などの総称。 |
| monsoon | モンスーン | 季節風。 |
| drought | 干ばつ | 長期間の降雨不足。 |
| typhoon | 台風 | 北西太平洋の熱帯低気圧。 |
| hurricane | ハリケーン | 大西洋・東太平洋の熱帯低気圧。 | 

## 8. 地域区分（Regions）
| 英語 | 日本語 | 説明 |
|------|--------|------|
| continent | 大陸 | 広大な陸地。 |
| country | 国 | 主権を持つ地域。 |
| state | 州 | 国の下位行政区。 |
| province | 省 | 国の行政区画。 |
| prefecture | 県 | 日本の行政区画。 |
| municipality | 自治体 | 市町村など。 |
| capital city | 首都 | 国の中心都市。 |

## 9. 地図・測量（Maps）
| 英語 | 日本語 | 説明 |
|------|--------|------|
| map | 地図 | 地理情報の図示。 |
| atlas | 地図帳 | 多数の地図をまとめた書籍。 |
| topographic map | 地形図 | 高低差を示す地図。 |
| contour line | 等高線 | 同じ高さを結ぶ線。 |
| compass | コンパス | 方位を測る道具。 |

## 必要なら、
- 200語すべてに日本語訳＋説明を付けた完全版 
- あなたの TSV カラムに合わせた地理用語辞書（Go 用） 
- 地理用語の英語ラベル命名規則（API 風） 
