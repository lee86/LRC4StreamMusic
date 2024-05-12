# qq音乐歌词接口

## 描述

此程序包含两个接口，分别用于返回`歌曲列表`和`歌词`

本意是针对群晖audio歌词插件进行精简后修改，主要是原php不会，所以分离成server + client形式，方便直接调用

流程相对于原插件多了一次中间查询返回的操作，个人自选

## 调用方式

http GET

http://127.0.0.1:8080/music?musicname=[查询歌曲、歌手]
<p>返回消息：</p>

```json
{
  "code": 0,
  "data": {
    "keyword": "邓丽君 Love Story",
    "priority": 0,
    "qc": [],
    "semantic": {
      "curnum": 0,
      "curpage": 1,
      "list": [],
      "totalnum": 0
    },
    "song": {
      "curnum": 3,
      "curpage": 1,
      "list": [
        {
          "albumid": 48800,
          "albummid": "001txRLw1UpVdS",
          "albumname": "往日情怀",
          "albumname_hilight": "往日情怀",
          "alertid": 24,
          "belongCD": 0,
          "cdIdx": 14,
          "chinesesinger": 0,
          "docid": "15949678681133688759",
          "grp": [
            {
              "albumid": 88889,
              "albummid": "002bCWzL1dFPGO",
              "albumname": "With Love from… 爱之世界",
              "albumname_hilight": "With Love from… 爱之世界",
              "alertid": 24,
              "belongCD": 0,
              "cdIdx": 6,
              "chinesesinger": 0,
              "docid": "4902863076468054777",
              "interval": 215,
              "isonly": 0,
              "lyric": "",
              "lyric_hilight": "",
              "media_mid": "001VxSlC3Ya29z",
              "msgid": 15,
              "newStatus": 2,
              "nt": 458904454,
              "pay": {
                "payalbum": 0,
                "payalbumprice": 0,
                "paydownload": 1,
                "payinfo": 1,
                "payplay": 1,
                "paytrackmouth": 1,
                "paytrackprice": 200
              },
              "preview": {
                "trybegin": 81763,
                "tryend": 127489,
                "trysize": 960887
              },
              "pubtime": 189273600,
              "pure": 5,
              "singer": [
                {
                  "id": 4713,
                  "mid": "000kfe022gdaAn",
                  "name": "邓丽君",
                  "name_hilight": "<em>邓丽君</em>"
                }
              ],
              "size128": 3447569,
              "size320": 8618534,
              "sizeape": 0,
              "sizeflac": 23185857,
              "sizeogg": 5414446,
              "songid": 4904655,
              "songmid": "002oXogB0pggDy",
              "songname": "Love Story",
              "songname_hilight": "<em>Love Story</em>",
              "strMediaMid": "001VxSlC3Ya29z",
              "stream": 1,
              "switch": 16880385,
              "t": 1,
              "tag": 10,
              "type": 0,
              "ver": 0,
              "vid": ""
            },
            {
              "albumid": 11631,
              "albummid": "0004Xx211ImYFN",
              "albumname": "天国的情人 - 邓丽君逝世十周年纪念声影存集",
              "albumname_hilight": "天国的情人 - <em>邓丽君</em>逝世十周年纪念声影存集",
              "alertid": 24,
              "belongCD": 0,
              "cdIdx": 4,
              "chinesesinger": 0,
              "docid": "5104710424092121090",
              "interval": 216,
              "isonly": 0,
              "lyric": "",
              "lyric_hilight": "",
              "media_mid": "001Nr3Aj3TDtJ7",
              "msgid": 15,
              "newStatus": 2,
              "nt": 458904454,
              "pay": {
                "payalbum": 0,
                "payalbumprice": 0,
                "paydownload": 1,
                "payinfo": 1,
                "payplay": 1,
                "paytrackmouth": 1,
                "paytrackprice": 200
              },
              "preview": {
                "trybegin": 88841,
                "tryend": 128948,
                "trysize": 960887
              },
              "pubtime": 1262275200,
              "pure": 5,
              "singer": [
                {
                  "id": 4713,
                  "mid": "000kfe022gdaAn",
                  "name": "邓丽君",
                  "name_hilight": "<em>邓丽君</em>"
                }
              ],
              "size128": 3458406,
              "size320": 0,
              "sizeape": 0,
              "sizeflac": 0,
              "sizeogg": 0,
              "songid": 132581,
              "songmid": "003PP1dl0aFtlK",
              "songname": "Love Story",
              "songname_hilight": "<em>Love Story</em>",
              "strMediaMid": "001Nr3Aj3TDtJ7",
              "stream": 1,
              "switch": 16880385,
              "t": 1,
              "tag": 10,
              "type": 0,
              "ver": 0,
              "vid": ""
            }
          ],
          "interval": 214,
          "isonly": 0,
          "lyric": "",
          "lyric_hilight": "",
          "media_mid": "004PEwSh4OfTcU",
          "msgid": 15,
          "newStatus": 2,
          "nt": 458904454,
          "pay": {
            "payalbum": 0,
            "payalbumprice": 0,
            "paydownload": 1,
            "payinfo": 1,
            "payplay": 1,
            "paytrackmouth": 1,
            "paytrackprice": 200
          },
          "preview": {
            "trybegin": 165726,
            "tryend": 197676,
            "trysize": 960887
          },
          "pubtime": 867686400,
          "pure": 5,
          "singer": [
            {
              "id": 4713,
              "mid": "000kfe022gdaAn",
              "name": "邓丽君",
              "name_hilight": "<em>邓丽君</em>"
            }
          ],
          "size128": 3462583,
          "size320": 8656144,
          "sizeape": 0,
          "sizeflac": 22213765,
          "sizeogg": 5314692,
          "songid": 576122,
          "songmid": "002UDcFa1eUAju",
          "songname": "Love Story",
          "songname_hilight": "<em>Love Story</em>",
          "strMediaMid": "004PEwSh4OfTcU",
          "stream": 1,
          "switch": 16880385,
          "t": 1,
          "tag": 12,
          "type": 0,
          "ver": 0,
          "vid": ""
        }
      ],
      "totalnum": 600
    },
    "tab": 0,
    "taglist": [],
    "totaltime": 0,
    "zhida": {
      "type": null
    }
  },
  "message": "",
  "notice": "",
  "subcode": 0,
  "time": 1646295296,
  "tips": ""
}
```

http://127.0.0.1:8080/lrc?musicmid=【歌曲mid编号】
<p>返回消息：</p>

```text
[ti:Love story ]
[ar:邓丽君]
[al:216241]
[by:]
[offset:0]
[00:24.32]Where do i begin
[00:28.10]To tell a story of how great a love can be
[00:34.82]The sweet love story that is older than the sea
[00:41.07]The simple truth about the love she brings to me
[00:49.29]Where do i start
[00:57.10]With her first hello
[01:00.85]She gave a meaning to this empty world of mine
[01:07.40]There&apos;d never be another love
[01:11.01]Another time
[01:13.86]She came into my life and made the living fine
[01:22.07]She fills my heart 
[01:28.66]She fills my heart with very special things
[01:35.19]With angel songs 
[01:38.46]With wild imaginings
[01:41.74]She fills my soul with so much love
[01:47.49]That anywhere i go
[01:51.69]I&apos;m never lonely 
[01:54.87]With her along who could be lonely
[02:01.57]I reach for her hand
[02:04.83]It&apos;s always there
[02:12.64]How long does it last
[02:16.79]Can love be measured by the hours in a day
[02:23.19]I have no answers now but this much i can say:
[02:29.85]I know i&apos;ll need her till the stars all burn away
[02:38.08]And she&apos;ll be there
[02:46.81]How long does it last
[02:51.79]Can love be measured by the hours in a day
[02:57.99]I have no answers now but this much i can say:
[03:04.13]I know i&apos;ll need her till the stars all burn away
[03:10.14]And she&apos;ll be there
```

## 免责声明及法律参考：

使用目的：该仓库及其子目录下的所有资源仅供学习和研究使用。其旨在为学术和研究人员提供参考和资料，任何其他目的均不适用。

非商业与非法用途：严禁将此项目及其内容用于任何商业或非法用途。对于因违反此规定而产生的任何法律后果，用户需自行承担全部责任。

来源与版权：该仓库目录下的所有资源信息均来源于网络。如有关于版权的争议或问题，请联系原作者或权利人。本声明者与版权问题无关且不承担任何相关责任。

下载后的处理：请注意，您在下载任何资源后，务必在24小时内从您的电脑或存储设备中彻底删除上述资源，无论这些资源是软件、文档还是其他形式的数据。

支持正版：如果您发现某个程序或资源对您有帮助或您喜欢它，请积极支持正版。购买和注册正版软件不仅可以获取官方的支持和更新，而且可以享受更多的功能和服务，如正版的人工智能服务。