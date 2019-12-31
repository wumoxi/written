# Javascript 正则表达式应用


比如有这么一个情况，服务器端通过第三方编辑器如Editor生成的文章会有一些样式不是终端(小程序、IOS、安卓)想要的数据，这个时候，就需要使用强大的正则表达式，来解决问题了！
具体情况，如将返回数据中的所有img标签样式去除，然后在终端自行样式处理。


_**比如服务器返回的内容如下**_

```html
<p style="text-align:center;">
    <strong><span style="color:#ffffff;background-color:#c00000;font-family:隶书, simli;font-size:20px;">广德梅花清水湾</span></strong>
</p>
<p style="text-align:center;">
    <br />
</p>
<p style="text-align:center;">
    <br />
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">1#、2#、3#、4#、5#、6#主体封顶</span></strong></span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">   </span><img src="/ueditor/php/upload/20191226/15773633336354.png" title="1.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" /><span style="font-family:隶书, simli;font-size:14px;">   </span><img src="/ueditor/php/upload/20191226/15773634167470.png" title="2.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" /></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">           </span></strong>
</p>
<p style="text-align:center;line-height:normal;margin-bottom:5px;margin-top:5px;">
    <strong><span style="font-size:14px;font-family:宋体;"><span style="font-family:隶书, simli;font-size:14px;"> </span><img src="/ueditor/php/upload/20191226/15773636855039.png" title="3.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" /><span style="font-family:隶书, simli;font-size:14px;"> </span><img src="/ueditor/php/upload/20191226/15773637304356.png" title="4.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" /></span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="text-align:center;font-family:隶书, simli;font-size:14px;">   </span><span style="font-size:14px;font-family:宋体;"><br />
        </span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="text-align:center;font-size:14px;font-family:宋体;"><span style="text-align:center;font-family:隶书, simli;font-size:14px;">         </span><img src="/ueditor/php/upload/20191226/15773638408650.png" title="5.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" /><span style="text-align:center;font-family:隶书, simli;font-size:14px;"> </span><img src="/ueditor/php/upload/20191226/15773639731391.png" title="6.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" /><span style="text-align:center;font-family:隶书, simli;font-size:14px;">           </span></span></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">   7#楼7层施工完成 </span></strong></span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">     </span><img src="/ueditor/php/upload/20191226/15773643381997.png" title="7.png" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" /></strong>
</p>
<p style="text-align:center;">
    <strong><span style="text-align:center;font-family:隶书, simli;font-size:14px;">            </span></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">            8#楼主体封顶     </span></strong></span></strong><strong><span style="font-family:隶书, simli;font-size:14px;">        </span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">  </span><img src="/ueditor/php/upload/20191226/15773643385494.png" title="8.png" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" /></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">    9#楼17层结构完成</span></strong></span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">  </span></strong>
</p>
<p style="text-align:center;">
    <strong><img src="/ueditor/php/upload/20191226/15773644809572.png" title="9.png" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" /></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">10#楼1层顶板施工中</span></strong></span></strong>
</p>
<p style="text-align:center;">
    <strong><img src="/ueditor/php/upload/20191226/15773645155559.png" title="10.png" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" /></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">11#楼1层结构完成</span></strong></span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">  </span><img src="/ueditor/php/upload/20191226/15773645915796.png" title="11" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" /><span style="font-family:隶书, simli;font-size:14px;"> </span></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">12#楼地下室墙柱钢筋施工中</span></strong></span></strong>
</p>
<p style="text-align:center;">
    <strong><img src="/ueditor/php/upload/20191226/15773645912754.png" title="12" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" /><span style="font-family:隶书, simli;font-size:14px;"></span></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">        13#楼主体封顶       </span></strong></span></strong><strong><span style="text-align:center;font-family:隶书, simli;font-size:14px;">     </span><span style="text-align:center;font-family:宋体;font-size:14px;"><span style="font-family:隶书, simli;font-size:14px;"> </span></span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">  </span><img src="/ueditor/php/upload/20191226/15773646908736.png" style="float:none;width:500px;height:300px;" title="13" width="500" height="300" border="2" hspace="20" vspace="20" /></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">14#楼17层墙柱钢筋施工中</span></strong></span></strong>
</p>
<p style="text-align:center;">
    <strong><img src="/ueditor/php/upload/20191226/15773646902951.png" title="14" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" /><br />
    </strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;"><span style="text-align:center;font-family:隶书, simli;font-size:14px;"> </span></span></strong><strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">      15#楼16层顶板施工中     </span></strong></span></strong><strong><span style="font-family:隶书, simli;font-size:14px;"><span style="font-family:隶书, simli;font-size:14px;">       </span></span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">  </span><img src="/ueditor/php/upload/20191226/1577364796222.png" style="float:none;width:500px;height:500px;" title="15.png" width="500" height="500" border="2" hspace="20" vspace="20" /></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">16#楼16层顶板施工中</span></strong></span></strong>
</p>
<p style="text-align:center;">
    <strong><img src="/ueditor/php/upload/20191226/15773647979544.png" title="16.png" width="500" height="500" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:500px;" /></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">            </span><span style="font-family:宋体;font-size:14px"><span style="font-family:隶书, simli;font-size:14px;">                                                       </span></span></strong>
</p>
<p style="text-align:center;">
    <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">17#、18#、19#楼主体封顶</span></strong></span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">        </span><img src="/ueditor/php/upload/20191226/15773648826610.png" style="float:none;width:200px;height:350px;" title="17" width="200" height="350" border="2" hspace="25" vspace="25" /><span style="font-family:隶书, simli;font-size:14px;"> </span><img src="/ueditor/php/upload/20191226/15773648827731.png" title="18" width="200" height="350" border="2" hspace="25" vspace="25" style="float:none;width:200px;height:350px;" /><span style="font-family:隶书, simli;font-size:14px;">  </span><img src="/ueditor/php/upload/20191226/15773648822557.png" title="19" width="200" height="350" border="2" hspace="25" vspace="25" style="float:none;width:200px;height:350px;" /></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:宋体;font-size:14px;"><br />
        </span></strong>
</p>
<p style="text-align:center;">
    <strong><span style="font-family:隶书, simli;font-size:14px;">             </span><span style="font-family:隶书, simli;font-size:14px;">                                     </span></strong>
</p>
<section data-role="outer" label="Powered by 365editor" style="font-family:微软雅黑;text-align:center;">
    <section style="margin:0 auto;display:flex;justify-content:center;align-items:center;">
        <section style="display:flex;flex-direction:column;justify-content:center;align-items:center;width:100%;box-sizing:border-box;">
            <section style="display:flex;flex-direction:column;justify-content:center;align-items:center;background:#fff;padding:0 4px;box-sizing:border-box;margin-bottom:-9px;z-index:1;">
                <section style="display:flex;flex-direction:row;justify-content:center;align-items:center;">
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;margin:0px 2px;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;margin:0px 2px;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;margin:0px 2px;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;margin-left:2px;">
                    </section>
                </section>
                <section style="display:flex;flex-direction:row;justify-content:center;align-items:center;margin-top:2px;">
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;margin:0px 2px;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;margin:0px 2px;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;margin:0px 2px;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;">
                    </section>
                    <section style="width:4px;height:4px;background:#651f06;border-radius:50%;flex-shrink:0;box-sizing:border-box;margin-left:2px;">
                    </section>
                </section>
            </section>
            <section style="width:100%;border:solid 1px #651f06;padding:6px;box-sizing:border-box;">
                <section style="width:100%;border:solid 1px #651f06;padding:6px;box-sizing:border-box;padding:18px 8px;">
                    <p style="text-align:justify;background-color:#ffffff;font-size:13px;color:#333333;letter-spacing:1.5px;line-height:1.75;">
                        <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">1、本工程现状信息采集最晚日期为2019年12月25日；</span></strong></span></strong>
                    </p>
                    <p style="text-align:justify;background-color:#ffffff;font-size:13px;color:#333333;letter-spacing:1.5px;line-height:1.75;">
                        <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">2、本系统展示之项目宣传介绍，如与买卖合同约定及政府批准文件不符的，不符之处均以合同约定及政府批文为准；                                        </span></strong></span></strong>
                    </p>
                    <p style="text-align:justify;background-color:#ffffff;font-size:13px;color:#333333;letter-spacing:1.5px;line-height:1.75;">
                        <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">3、文图展示内容如有修订的恕不另行通知；                                        </span></strong></span></strong>
                    </p>
                    <p style="text-align:justify;background-color:#ffffff;font-size:13px;color:#333333;letter-spacing:1.5px;line-height:1.75;">
                        <strong style="text-align:center;"><span style="font-family:宋体;font-size:14px;"><strong style="text-indent:48px;"><span style="font-size:14px;color:#0000ff;font-family:隶书, simli;">4、图文未经允许不得转载。 </span></strong></span></strong><strong><span style="font-family:隶书, simli;font-size:14px;">  </span></strong>
                    </p>
                </section>
            </section>
        </section>
    </section>
</section>
```

_**书写正则表达式**_

```regexp
<img\s+(src="[/\w.]+")[\s\w=".]+(style="[/\w:;]+")[\s\w=".]+\/>
```

_**获取返回匹配字符串**_

```js
// Regexp
var regexp = /<img\s+(src="[/\w.]+")[\s\w=".]+(style="[/\w:;]+")[\s\w=".]+\/>/g;

// 匹配所有符合条件的字符串返回一个二维数组
function getMatchesAll(string, regex) {
    var matches = [];
    var match;
    while (match = regex.exec(string)) {
        matches.push(match);
    }
    return matches;
}

// 获取所有匹配
var matches = getMatchesAll(str, regexp);

console.log(JSON.stringify(matches))
```

控制台将输出如下

```json
[
    [
        "<img src="/ueditor/php/upload/20191226/15773633336354.png" title="1.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" />",
        "src="/ueditor/php/upload/20191226/15773633336354.png"",
        "style="float:none;width:350px;height:250px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773634167470.png" title="2.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" />",
        "src="/ueditor/php/upload/20191226/15773634167470.png"",
        "style="float:none;width:350px;height:250px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773636855039.png" title="3.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" />",
        "src="/ueditor/php/upload/20191226/15773636855039.png"",
        "style="float:none;width:350px;height:250px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773637304356.png" title="4.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" />",
        "src="/ueditor/php/upload/20191226/15773637304356.png"",
        "style="float:none;width:350px;height:250px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773638408650.png" title="5.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" />",
        "src="/ueditor/php/upload/20191226/15773638408650.png"",
        "style="float:none;width:350px;height:250px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773639731391.png" title="6.png" width="350" height="250" border="2" hspace="20" vspace="20" style="float:none;width:350px;height:250px;" />",
        "src="/ueditor/php/upload/20191226/15773639731391.png"",
        "style="float:none;width:350px;height:250px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773643381997.png" title="7.png" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" />",
        "src="/ueditor/php/upload/20191226/15773643381997.png"",
        "style="float:none;width:500px;height:300px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773643385494.png" title="8.png" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" />",
        "src="/ueditor/php/upload/20191226/15773643385494.png"",
        "style="float:none;width:500px;height:300px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773644809572.png" title="9.png" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" />",
        "src="/ueditor/php/upload/20191226/15773644809572.png"",
        "style="float:none;width:500px;height:300px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773645155559.png" title="10.png" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" />",
        "src="/ueditor/php/upload/20191226/15773645155559.png"",
        "style="float:none;width:500px;height:300px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773645915796.png" title="11" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" />",
        "src="/ueditor/php/upload/20191226/15773645915796.png"",
        "style="float:none;width:500px;height:300px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773645912754.png" title="12" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" />",
        "src="/ueditor/php/upload/20191226/15773645912754.png"",
        "style="float:none;width:500px;height:300px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773646908736.png" style="float:none;width:500px;height:300px;" title="13" width="500" height="300" border="2" hspace="20" vspace="20" />",
        "src="/ueditor/php/upload/20191226/15773646908736.png"",
        "style="float:none;width:500px;height:300px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773646902951.png" title="14" width="500" height="300" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:300px;" />",
        "src="/ueditor/php/upload/20191226/15773646902951.png"",
        "style="float:none;width:500px;height:300px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/1577364796222.png" style="float:none;width:500px;height:500px;" title="15.png" width="500" height="500" border="2" hspace="20" vspace="20" />",
        "src="/ueditor/php/upload/20191226/1577364796222.png"",
        "style="float:none;width:500px;height:500px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773647979544.png" title="16.png" width="500" height="500" border="2" hspace="20" vspace="20" style="float:none;width:500px;height:500px;" />",
        "src="/ueditor/php/upload/20191226/15773647979544.png"",
        "style="float:none;width:500px;height:500px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773648826610.png" style="float:none;width:200px;height:350px;" title="17" width="200" height="350" border="2" hspace="25" vspace="25" />",
        "src="/ueditor/php/upload/20191226/15773648826610.png"",
        "style="float:none;width:200px;height:350px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773648827731.png" title="18" width="200" height="350" border="2" hspace="25" vspace="25" style="float:none;width:200px;height:350px;" />",
        "src="/ueditor/php/upload/20191226/15773648827731.png"",
        "style="float:none;width:200px;height:350px;""
    ],
    [
        "<img src="/ueditor/php/upload/20191226/15773648822557.png" title="19" width="200" height="350" border="2" hspace="25" vspace="25" style="float:none;width:200px;height:350px;" />",
        "src="/ueditor/php/upload/20191226/15773648822557.png"",
        "style="float:none;width:200px;height:350px;""
    ]
]
```

现在有了这些数组数据，那处理这个数组后，将原有字符串中这些img标签替换掉，那这件事情就算完成了，实现就是字符串，正则，数组结合处理的过程！

