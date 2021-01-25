package pa

import (
	"fmt"
	"regexp"
)

func GetInfo() {
	testVarrr := `
<!DOCTYPE html>
<html>
<!--  当前的类别归属ID-->
<!--  当前的类别表现ID-->
<!--  当前的类别listname (eg:banjia)-->
<!--  当前的类别catename (eg:搬家)-->
<!--  一级类别归属ID-->
<!--  一级类别表现ID-->
<!--  一级类别listname eg:jiazheng-->
<!--  一级类别catename eg:家政-->
<!--  二级类别归属ID [Ljava.lang.Integer;@17b4c919 数字长度：2-->
<!--  二级类别表现ID-->
<!--  二级类别listname eg:jiazhuang-->
<!--  二级类别catename eg:家装-->
<!-- Pc端帖子信息体改版 -->
<!-- 列表页的第一条帖子为精准/置顶帖时，相应的标识更改为“广告”标识 -->
<!-- 表现类别的全路径   eg:8512,96-->

			<head>
        	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<title>【58同城】北京二手手机交易网_二手手机价格_二手手机市场 </title>
	<meta name="description" content='北京二手手机交易网为您提供大量北京二手手机报价、交易等信息，在这里您可以免费查看和发布北京二手手机转让、求购、价格、图片等信息，买卖北京二手手机请到58同城北京二手手机交易市场。'/>
	<meta name="keywords" content='北京二手手机转让,北京二手手机交易市场,二手手机价格'/>
	<meta http-equiv="X-UA-Compatible" content="edge" />
		

<meta http-equiv="mobile-agent" content="format=xhtml; url=//m.58.com/bj/shouji/?from=ganji" />
<meta http-equiv="mobile-agent" content="format=html5; url=//m.58.com/bj/shouji/?from=ganji" />
<meta http-equiv="mobile-agent" content="format=wml; url=//m.58.com/bj/shouji/?from=ganji" />
            <script type="text/javascript">
    try{
        var ____json4fe = {"catentry":[{"dispid":"5","name":"二手市场","listname":"sale"},{"dispid":"36","name":"二手手机","listname":"shouji"}],"locallist":[{"dispid":"1","name":"北京","listname":"bj"}],"start":(new Date()).getTime(),"modules":"list"};
    ____json4fe.version = 'A';
    ____json4fe.ABVersion = 'ZUCHE_B';
    ____json4fe.modules = 'listpage';
    ____json4fe.sid = '136460514211221353733011037';
    ____json4fe.sessionid = '68cb5299-033d-4f4a-8ced-3ad991cad5cc';
    ____json4fe.smsc = '0ab2e8dd936f9980fc199aa7beb1490e107e29ef6b6a06bc39bf3fcbd6166caea326d982d647ee11378d6d304e4a498465299c700b3d6a1dd71642185b6f7fb19f854bdc6707c73990e49192b4c247948f36d13e954a84203ccf40c8a14ff33aa929df43feda637ba648d23e0d6f722307fa2290d777d0787026b94425a118e4';
    ____json4fe.newcateid='';
    ____json4fe.keyword='';
    ____json4fe.page='';
    ____json4fe.param='';
    ____json4fe.req_version= "1.0.0";
        }catch(e){};
    (function(){
        var e = "header,nav,section,aside,article,footer".split(','),
                i=0,
                length=e.length;
        while(i<length){
            document.createElement(e[i++])
        }
    })();
    var ajax_param = '{"infoMethod":["wltAge","renzheng","gfrz","wxicon"],"dataParam":"42643242627073_22999328008455_0_adinfo_rb,42255837070882_68006518347521_0_adinfo_rb,42291153297829_68006518347521_0_adinfo_rb,39860104407835_67036155779094_0_adinfo_rb,39860104407835_67036155779094_0_adsumplayinfo,42255837070882_68006518347521_0_adsumplayinfo,42291153297829_68006518347521_0_adsumplayinfo,42643242627073_22999328008455_0_adsumplayinfo,42643242627073_22999328008455_0_adinfo,42255837070882_68006518347521_0_adinfo,42291153297829_68006518347521_0_adinfo,39860104407835_67036155779094_0_adinfo,44776197052933_31090533585931_0_promationinfo,32180495876939_22999328008455_0_promationinfo,39860104407835_67036155779094_0_promationinfo,39857474720807_34782619515151_0_promationinfo,43892140061590_73183543809563_0_commoninfo,44990544254365_34782619515151_0_commoninfo,44989562911777_62059603549955_0_commoninfo,44989251661577_56548663934991_0_commoninfo,44988863522196_22999328008455_0_commoninfo,44988863582866_68006518347521_0_commoninfo,44988806973845_40711099314454_0_commoninfo,44987795256866_54619981_0_commoninfo,44987404463765_46284061248275_0_commoninfo,44987023698077_4571913731846_0_commoninfo,44987022991394_25544712154374_0_commoninfo,44987019222946_77017467144464_0_commoninfo,44986720269349_59649892143117_0_commoninfo,44986676109191_34424738104076_0_commoninfo,44986663067780_31090533585931_0_commoninfo,44986609917329_34955377567495_0_commoninfo,44985806888605_57105124693266_0_commoninfo,44985776128775_29893618950153_0_commoninfo,44982906536999_77306537192211_0_commoninfo,44981448941060_74360061016578_0_commoninfo,44980932428289_74347704561948_0_commoninfo,44980530438548_73800594833447_0_commoninfo,44980490112912_76839291462698_0_commoninfo,44980114581924_61970159569932_0_commoninfo,44978344875930_74489129007374_0_commoninfo,44976958686220_23332494403847_0_commoninfo,44976883227788_71081438984715_0_commoninfo,44975710510235_77291979006723_0_commoninfo,44970093341570_5302699079175_0_commoninfo,44969051370765_75193648407059_0_commoninfo,44968566010148_75903659898644_0_commoninfo,44967359846426_69799800295437_0_commoninfo,44967274528665_57897448065043_0_commoninfo,44965805624714_62382003827204_0_commoninfo,44965778130181_55995344896271_0_commoninfo,44454973371414_31090533585931_0_commonflowlayer,44987943091204_68006518347521_0_commonflowlayer,44986118421537_22999328008455_0_commonflowlayer,44990544254365_34782619515151_0_commonflowlayer,44956039288720_77101043764236_0_commonflowlayer,37833978957841_61495355181071_0_commonflowlayer,44654071468451_13115728365830_0_commonflowlayer,44989251661577_56548663934991_0_commonflowlayer,44643123610529_70792921227537_0_commonflowlayer,44791601055650_23660948632070_0_commonflowlayer,44965764135072_55995344896271_0_commonflowlayer,44965805624714_62382003827204_0_commonflowlayer,44646388028425_76139802851849_0_commonflowlayer,44978226983690_74489129007374_0_commonflowlayer,44969051370765_75193648407059_0_commonflowlayer,44422693433239_28058740286727_0_commonflowlayer,44966870406283_69799800295437_0_commonflowlayer,41626248209174_61130970780940_0_commonflowlayer,43473734796548_73774298680835_0_commonflowlayer,44945924884500_75024332721171_0_commonflowlayer","dispCateId":36,"dispCateName":"shouji","platform":"pc","pageIndex":1,"paramMap":{}}';
    var s_ajax_param = 's_contact_shouji_136460514211221353733011037_';
    // 类别和城市全局变量
    var currentListName = 'shouji',
        citylistname = 'bj',
        firstListName = 'sale',
        secListName = 'shouji';
    
    // 是否接入担保交易
    var onlinedealicon = '';
</script>
<script type='text/javascript'>
!function(t, e, w) {
    w._ty_key = 'cGL4cfmjZpA';
    var a = t.createElement(e);
    a.async = !0,a.src = ('https:' == t.location.protocol ? 'https://' : 'http://') + 'j2.58cdn.com.cn/pc/hy/tingyun-rum.min.js';
    var c = t.getElementsByTagName(e)[0];
    c.parentNode.insertBefore(a, c)
}(document, 'script', window);
</script>
<script type="text/javascript">
    window.WMDA_SDK_CONFIG = ({
    api_v: 1,
    sdk_v: 0.1,
    mode: 'report',
    appid: 1411632341505,
    key: 'p2ikeuny',
    project_id: '1409632296065'
});
      var _vds = _vds || [];
      window._vds = _vds;
      (function(){
        _vds.push(['setAccountId', '98e5a48d736e5e14']);
        _vds.push(['setImp', false]);
        _vds.push(['setPageGroup', 'list']);
        (function() {
          var vds = document.createElement('script');
          vds.type='text/javascript';
          vds.async = true;
          vds.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'dn-growing.qbox.me/vds.js';
          var s = document.getElementsByTagName('script')[0];
          s.parentNode.insertBefore(vds, s);
        })();
      })();
</script>
<script type="text/javascript" src='//j1.58cdn.com.cn/wmda/js/statistic_v20201103104710.js'></script>
<script type="text/javascript">
    var ____loadCfg ="";
    if(currentListName=="changtuzhuanche"|| currentListName=="pinche")
        ____loadCfg = ['huangye', 'changtuzhuanche', 'listpage'];
    else
        ____loadCfg = ['huangye', 'shenghuo', 'listpage'];
</script>
<script src='//j2.58cdn.com.cn/js/require_jquery_load.js'></script>
<script type="text/javascript" src='//j1.58cdn.com.cn/js/v8/boot/boot_huangye_v20201022172444.js'></script>
<script type="text/javascript">
    require(['_pkg/huangye/huangye_shenghuo_list_dom'], function () {});
</script>

<script type="text/javascript" src="//cbjs.baidu.com/js/m.js"></script>

<script type="text/javascript" src="//j2.58cdn.com.cn/js/saveKeyword.js"></script>
<script type="text/javascript">
    var xxfwConfig = {namespace: 'huangyelistpc'};
</script>
<script type="text/javascript" src='//j1.58cdn.com.cn/resource/xxzl/xxfw/xxfw.min_v20200907150609.js'></script>
<script type="application/ld+json">
{
    "@context": "https://ziyuan.baidu.com/contexts/cambrian.jsonld",
    "@id":https://bj.58.com/shouji/?from=ganji,
    "upDate": 2021-01-24T20:24:01
}
</script>

<script>

(function(){

    var bp = document.createElement('script');

    var curProtocol = window.location.protocol.split(':')[0];

    if (curProtocol === 'https') {

        bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';

    }

    else {

        bp.src = 'http://push.zhanzhang.baidu.com/push.js';

    }

    var s = document.getElementsByTagName("script")[0];

    s.parentNode.insertBefore(bp, s);

})();

</script>
<script>
var fromMess = "hylbsidebar";
</script>
<script src='//j1.58cdn.com.cn/resource/xxzl/public/index_v0.js'></script>
<script type="text/javascript" src='//j1.58cdn.com.cn/webim/js/entry_v20201125103402.js'></script>
            
<link rel="stylesheet" href="//c.58cdn.com.cn/lbg/58.com/ui8/modules/header/huangye/retable-ca53c955b6.css">
<link href='//c.58cdn.com.cn/componentsLoader/dist/CompontsLoader_v20201229143446.css' rel="stylesheet" type="text/css">
<link rel="stylesheet" href="//c.58cdn.com.cn/lbg/58.com/ui8/modules/header/huangye/hy_freetel-63d7c224dc.css">


<link rel="stylesheet" href="//c.58cdn.com.cn/lbg/58.com/ui8/modules/header/huangye/hy_acp_item-04cfa4a9c5.css">
    </head>
<body>

<div id="commonTopbar" class="commonTopbar"></div>

<!-- =S banner  div -->
<div class="topbannerbar"><div id="brand_list_top_banner" class="brandad1000" onclick="clickLog('from=pc_list_shouji_top_banner');" ></div></div>
<!-- =E banner  div--><header id="header">
    <div class="header-inner">
    <a class="logo float_l" href="/huangye/" target="_blank" onclick="clickLog('from=pc_list_logo_shouji');">58同城</a>
        <a href='//post.58.com/hy/1/36/s5?ertry_source=info_list&page_from=info_list' target="_blank" onclick="clickLog('from=pc_list_fabu_shouji');" class="postbtn float_r">免费发布</a>
                <div class="header-search float_r">
         <em class="ico-searchkey"></em>
         <form action="" method="get" onsubmit="b_query('final=1&amp;searchtype=3');return false;">
                                                    <input lang="zh-CN" x-webkit-grammar="builtin:translate" x-webkit-speech="" class="search-input c_000" id="keyword1" name="b_q" para="key" autocomplete="off" value="" onkeyup="win1.GetContentData();"   onclick="">
            <span onMouseOut="this.className='sbtn'" onMouseOver="this.className='sbtn hover'" class="sbtn">
            <input type="submit" value="搜服务" onclick="clickLog('from=pc_list_search_shouji');" class="but-bl" name="" id="searchbtn1" ></span>
                                               		<div class="clear"></div>
        				
	       		        	  </form>
       	</div>
		        	<div class="nav float_l" >
	<a href="/">北京58同城</a>
    	   &gt; <a href="/sale.shtml">北京二手市场</a>
					 &gt; <a href="/shouji/" itemprop="url">北京二手手机</a>
	

</div>
	</div>
</header>




<section id="selection">
	
	<dl class="secitem" >
		<dt class="secitem_brand">类别 ：</dt>
		<dd id="ObjectType" >
																
									
			<a href="/shouji/" class='select' onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">全部</a>
								
									
			<a href="/iphonesj/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">苹果</a>
								
									
			<a href="/sanxing/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">三星</a>
								
									
			<a href="/xiaomi/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">小米</a>
								
									
			<a href="/huaweisj/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">华为</a>
								
									
			<a href="/oppo/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">OPPO</a>
								
									
			<a href="/vivo/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">vivo</a>
								
									
			<a href="/meizu/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">魅族</a>
								
									
			<a href="/shoujipeijian/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">手机配件</a>
								
									
			<a href="/nuojiya/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">诺基亚</a>
								
									
			<a href="/htc/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">HTC</a>
								
									
			<a href="/ailixin/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">索尼爱立信</a>
								
									
			<a href="/nubiya/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">努比亚</a>
								
									
			<a href="/xiapushouji/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">夏普</a>
								
									
			<a href="/motuoluola/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">摩托罗拉</a>
								
									
			<a href="/lianxiang/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">联想</a>
								
									
			<a href="/lgsj/"  onclick="clickLog('from=pc_list_shaixuan_shouji&abVersion=$!{virtualCatePcAB}&filterparams=$!{filterParams})">LG</a>
															    									   <a href="javascript:void(0);" class="itemmore" id="more_ObjectType">更多<i></i></a>
		   <div style="display: none;" class="moreBrandct" id="moreBrandct">
				<dl class="hotarea">
								<dt class=""><span>C</span></dt>
					<dd style="display: none;">
																			<a href="/chuizisj/" >锤子</a>
									</dd>
								<dt class=""><span>H</span></dt>
					<dd style="display: none;">
																			<a href="/huaweisj/" >华为</a>
															<a href="/htc/" >HTC</a>
									</dd>
								<dt class=""><span>J</span></dt>
					<dd style="display: none;">
																			<a href="/jinlisj/" >金立</a>
									</dd>
								<dt class=""><span>K</span></dt>
					<dd style="display: none;">
																			<a href="/kupaisj/" >酷派</a>
									</dd>
								<dt class=""><span>L</span></dt>
					<dd style="display: none;">
																			<a href="/lianxiang/" >联想</a>
															<a href="/lgsj/" >LG</a>
															<a href="/leshisj/" >乐视</a>
									</dd>
								<dt class=""><span>M</span></dt>
					<dd style="display: none;">
																			<a href="/meizu/" >魅族</a>
															<a href="/motuoluola/" >摩托罗拉</a>
									</dd>
								<dt class=""><span>N</span></dt>
					<dd style="display: none;">
																			<a href="/nuojiya/" >诺基亚</a>
															<a href="/nubiya/" >努比亚</a>
									</dd>
								<dt class=""><span>O</span></dt>
					<dd style="display: none;">
																			<a href="/oppo/" >OPPO</a>
									</dd>
								<dt class=""><span>P</span></dt>
					<dd style="display: none;">
																			<a href="/iphonesj/" >苹果</a>
									</dd>
								<dt class=""><span>Q</span></dt>
					<dd style="display: none;">
																			<a href="/qitasj/" >其他</a>
									</dd>
								<dt class=""><span>S</span></dt>
					<dd style="display: none;">
																			<a href="/sanxing/" >三星</a>
															<a href="/shoujipeijian/" >手机配件</a>
															<a href="/ailixin/" >索尼爱立信</a>
															<a href="/shoujihuishou/" >手机回收</a>
									</dd>
								<dt class=""><span>V</span></dt>
					<dd style="display: none;">
																			<a href="/vivo/" >vivo</a>
									</dd>
								<dt class=""><span>X</span></dt>
					<dd style="display: none;">
																			<a href="/xiaomi/" >小米</a>
															<a href="/xiapushouji/" >夏普</a>
									</dd>
								<dt class=""><span>Y</span></dt>
					<dd style="display: none;">
																			<a href="/yijiasj/" >一加</a>
									</dd>
								<dt class=""><span>Z</span></dt>
					<dd style="display: none;">
																			<a href="/zhongxingshouji/" >中兴</a>
									</dd>
							</dl>
			</div>
			    	</dd>
    </dl>
	
	<dl class="secitem" >
		<dt class="secitem_brand">价格 ：</dt>
		<dd id="MinPriceqj" >
																
									
			<a href="/shouji/" class='select' >全部</a>
								
									
			<a href="/shouji/pve_5537_0_100/"  >100元以下</a>
								
									
			<a href="/shouji/pve_5537_101_200/"  >100-200元</a>
								
									
			<a href="/shouji/pve_5537_201_500/"  >200-500元</a>
								
									
			<a href="/shouji/pve_5537_501_1000/"  >500-1000元</a>
								
									
			<a href="/shouji/pve_5537_1001_2000/"  >1000-2000元</a>
								
									
			<a href="/shouji/pve_5537_2001_3500/"  >2000-3500元</a>
								
									
			<a href="/shouji/pve_5537_3501_999999/"  >3500元以上</a>
											<span class="prifilter">
				<span class="text text2">
					<input type="text"  para="minprice"  size="3" muti="1" onkeyup="cknum(this)" min="0" name="b_q" /> -
					<input type="text"  para="minprice"  size="3" muti="1" onkeyup="cknum(this)" min="0" name="b_q" /> 元
				</span>
				<span class="btn btn2"><input type="button" value="确定" onclick="b_query_minprice()" /></span>
			</span>
									    	</dd>
    </dl>
	
	<dl class="secitem" >
		<dt class="secitem_brand">新旧 ：</dt>
		<dd id="oldlevel" >
																
									
			<a href="/shouji/" class='select' >全部</a>
								
									
			<a href="/shouji/pve_1230_7/"  >全新</a>
								
									
			<a href="/shouji/pve_1230_2/"  >95成新</a>
								
									
			<a href="/shouji/pve_1230_3/"  >9成新</a>
								
									
			<a href="/shouji/pve_1230_4/"  >8成新</a>
								
									
			<a href="/shouji/pve_1230_5/"  >7成新及以下</a>
															    	</dd>
    </dl>
	
	<dl class="secitem" >
		<dt class="secitem_brand">物品状况 ：</dt>
		<dd id="shoujizhuangkuang" >
																
									
			<a href="/shouji/" class='select' >全部</a>
								
									
			<a href="/shouji/pve_6681_511789/"  >正品发票</a>
								
									
			<a href="/shouji/pve_6681_511790/"  >保修期内</a>
								
									
			<a href="/shouji/pve_6681_511791/"  >主配齐全</a>
								
									
			<a href="/shouji/pve_6681_511792/"  >无拆无修</a>
															    	</dd>
    </dl>
				<dl class="secitem secitem_area clearfix">
			<dt>地址 ：</dt>
			<dd zwnameid="quyu" zwname="地址" >
						        		        	<a  class='select'
					   					   para='local' ck='bj'   href="/shouji/" data-finalurl="/shouji/" onClick="clickLog('from= pc_list_quyu_bj');">全北京</a>
							        		        	<a   href="/chaoyang/shouji/" data-finalurl="/chaoyang/shouji/" onClick="clickLog('from= pc_list_quyu_chaoyang');">朝阳</a>
							        		        	<a   href="/haidian/shouji/" data-finalurl="/haidian/shouji/" onClick="clickLog('from= pc_list_quyu_haidian');">海淀</a>
							        		        	<a   href="/changping/shouji/" data-finalurl="/changping/shouji/" onClick="clickLog('from= pc_list_quyu_changping');">昌平</a>
							        		        	<a   href="/fengtai/shouji/" data-finalurl="/fengtai/shouji/" onClick="clickLog('from= pc_list_quyu_fengtai');">丰台</a>
							        		        	<a   href="/daxing/shouji/" data-finalurl="/daxing/shouji/" onClick="clickLog('from= pc_list_quyu_daxing');">大兴</a>
							        		        	<a   href="/tongzhouqu/shouji/" data-finalurl="/tongzhouqu/shouji/" onClick="clickLog('from= pc_list_quyu_tongzhouqu');">通州</a>
							        		        	<a   href="/fangshan/shouji/" data-finalurl="/fangshan/shouji/" onClick="clickLog('from= pc_list_quyu_fangshan');">房山</a>
							        		        	<a   href="/shunyi/shouji/" data-finalurl="/shunyi/shouji/" onClick="clickLog('from= pc_list_quyu_shunyi');">顺义</a>
							        		        	<a   href="/xicheng/shouji/" data-finalurl="/xicheng/shouji/" onClick="clickLog('from= pc_list_quyu_xicheng');">西城</a>
							        		        	<a   href="/dongcheng/shouji/" data-finalurl="/dongcheng/shouji/" onClick="clickLog('from= pc_list_quyu_dongcheng');">东城</a>
							        		        	<a   href="/miyun/shouji/" data-finalurl="/miyun/shouji/" onClick="clickLog('from= pc_list_quyu_miyun');">密云</a>
							        		        	<a   href="/shijingshan/shouji/" data-finalurl="/shijingshan/shouji/" onClick="clickLog('from= pc_list_quyu_shijingshan');">石景山</a>
							        		        	<a   href="/huairou/shouji/" data-finalurl="/huairou/shouji/" onClick="clickLog('from= pc_list_quyu_huairou');">怀柔</a>
							        		        	<a   href="/mentougou/shouji/" data-finalurl="/mentougou/shouji/" onClick="clickLog('from= pc_list_quyu_mentougou');">门头沟</a>
							        		        	<a   href="/yanqing/shouji/" data-finalurl="/yanqing/shouji/" onClick="clickLog('from= pc_list_quyu_yanqing');">延庆</a>
							        		        	<a   href="/pinggu/shouji/" data-finalurl="/pinggu/shouji/" onClick="clickLog('from= pc_list_quyu_pinggu');">平谷</a>
							        		        	<a   href="/beijingzhoubian/shouji/" data-finalurl="/beijingzhoubian/shouji/" onClick="clickLog('from= pc_list_quyu_beijingzhoubian');">北京周边</a>
							        			    			    <div class="fl clearfix recent_div" id="shangquan" style=" display: none;"></div>

														    </dd>
		</dl>
		<div id="selected" class="bar"></div>
	<i class="shadow"></i>
</section>
<div id="adtop_1" class="hc">
    <div class="mouter defa">
        <div class="inner">
            <div id="md3" class="mdul"></div>
            <ul id="md1" class="mdul"></ul>
            <ul id="md2" class="mdul"></ul>
            <ul id="md4" class="mdul"></ul>
            <ul id="md5" class="mdul"></ul>
            <div class="clear"></div>
        </div>
    </div>
</div>


<div class="containnerWrap clearfix">
        <section id="mainlist" class="main clearfix pr">
                            <div class="pr tabsbar_pop_wrap">
	<div class="pop lamp fang_01 tabsbar_pop" id="tabsbar"><span></span><font class="close"></font><font class="arr_btm"></font>
	</div>
</div>
<div class="tabsbar" id="xgtjtab">
	<div class="list-tabs list-tabs-es">
			<a class='sel' name='b_link' para='custom' cl='https://bj.58.com/shouji/?from=ganji&rb=1'		   data-finalurl="https://bj.58.com/shouji/?from=ganji" onClick="clickLog('from=qbshouji');" href="https://bj.58.com/shouji/?from=ganji"><h1>全部</h1>
		</a>
		<a 		   data-finalurl="https://bj.58.com/shouji/0/?from=ganji" onClick="clickLog('from=grshouji');" href="https://bj.58.com/shouji/0/?from=ganji"><h1>个人</h1>
		</a>
		<a 		   data-finalurl="https://bj.58.com/shouji/1/?from=ganji" onClick="clickLog('from=sjshouji');" href="https://bj.58.com/shouji/1/?from=ganji"><h1>商家</h1>
		</a>
		        		</div>
			<span class="zhaobiao">
									<a id="es-tg" rel="nofollow" class="pubtopright" target="_blank" href="//e.58.com/jiaoyi.html">
					<!-- ES-2332 二手设备，办公设备,艺术收藏 ,成人用品,通讯业务,二手求购-->

											<img src="//img.58cdn.com.cn/ui7/ershou/list/tg.png" />
									</a>
						</span>
</div>
                
                                                        <div class="filterbar">
	<div class="filterbar-l">
		<div class="filterup" id="filterup">
            			<h1>北京二手手机</h1>
		</div>
		<span class="checkbox"><input name="" type="checkbox" value="" id="ispic" /><label for="ispic">只看有图</label></span>
	</div>

	<div class="filterbar-md">
        <div class="sort_wrap sort_wrap-es">
                        <span class="default  active" >
                <a href="https://bj.58.com/shouji/?from=ganji" rel="nofollow">默认排序</a>
            </span>
            <span  class="time" >
                <a href="https://bj.58.com/shouji/?from=ganji&sort=sortid_desc" rel="nofollow">按时间</a>
            </span>
            <span  class="sort_price" >
                <a href="https://bj.58.com/shouji/?from=ganji&sort=minprice_asc" rel="nofollow">按价格</a>
                            </span>
        </div>
	</div>
</div>
                                
        <div id="infolist" class="cleft" >
                        <table id="jingzhun"  class="small-tbimg ac_container list-new-table"   cellspacing="0" cellpadding="0">
				
									<tr logr='q_2_22999328008455_42643242627073_4_sortid:646494332@ses:@npos:1@pos:1' infotag="42643242627073_adinfo" data-sign=""  class="new-list" >
		<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_adinfo_tupian&entityId=42643242627073&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
			<div class="ac_linkurl">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EzPWE1nWEzPWcLnj01XaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDznWbOrHnzrjTkrjEdPEDKPjcvPjnzPjcvnW0kP1nKrHbknTDdP1ckTHnvTHDKnHmQnHEOnHTYnHTdPkDQTyQG0Lw_uyuYTHDKnEDKsHDKTHD1n1mLnWmzrHnLnjNvPWcYPWEKnTDKnEDduW7bnAc3uidhnH-WsHEvmvcVrH9QuiYdm1Tdmyu-mvNzmH0KnHn1PW0zPWcOPjTdP1b3PjTknEDQn1nvP1cvnWb1rHnznHNYrj9kTED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrED3P1b1nA7bPidWnvm3sHEdrAcVmWELmiYznAPbrHD3nhcLm1NKTHTKTEDkTHDKnT78IyQ_THTOmvuBmWbzmHm1PyP6nAn&adact=4&psid=136460514211221353733011037&entinfo=42643242627073_q&slot=1000019' tongji_label="listclick" target="_blank" rel="nofollow" >
				<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://t1.58cdn.com.cn/bidding/small/n_v24a047d973f274c3f9c6fb1f6d55db435_130_130.jpg'/>
								</a>
							</div>

					</td>
				<td class="t" colspan="2">
					<div class="tdiv">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EzPWE1nWEzPWcLnj01XaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDznWbOrHnzrjTkrjEdPEDKPjcvPjnzPjcvnW0kP1nKrHbknTDdP1ckTHnvTHDKnHmQnHEOnHTYnHTdPkDQTyQG0Lw_uyuYTHDKnEDKsHDKTHD1n1mLnWmzrHnLnjNvPWcYPWEKnTDKnEDduW7bnAc3uidhnH-WsHEvmvcVrH9QuiYdm1Tdmyu-mvNzmH0KnHn1PW0zPWcOPjTdP1b3PjTknEDQn1nvP1cvnWb1rHnznHNYrj9kTED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrED3P1b1nA7bPidWnvm3sHEdrAcVmWELmiYznAPbrHD3nhcLm1NKTHTKTEDkTHDKnT78IyQ_THTOmvuBmWbzmHm1PyP6nAn&adact=3&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42643242627073_q&slot=1000019' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_adinfo_bj');" target="_blank" class="t ac_linkurl" rel="nofollow" name='42643242627073'>
				<div class="new-long-tit new-long-tit2">
					高价回收苹果12pro-Max-11三星、华为等
				</div>
														<span class='tu'></span>
								<span class="jingpin">精</span>
				            					</a>

														<div class="item-desc">
						全北京
											</div>
									
									<p class="seller">
								
				<a class="c_666" href="/shouji/">北京</a>
					&nbsp;/ 今天
							</p>
				
				<p class="item-tags">
			<span class="async-tags"></span>
			</p>
    		</div>
    	</td>
    	
<td class="vertop-es">
	<b class='pri'>6000</b> <span>元</span>
</td>
	</tr>
		
		
				
									<tr logr='q_2_68006518347521_42255837070882_4_sortid:643467726@ses:@npos:2@pos:2' infotag="42255837070882_adinfo" data-sign=""  class="new-list" >
		<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_adinfo_tupian&entityId=42255837070882&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
			<div class="ac_linkurl">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWNdrjnLnj0krj9zXaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDvrjTkPWNQrjnYP1NznEDKPjczPHN3n10kP1T3rjcKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTdPkDzTyQG0Lw_uyuYTHDKnEDKsHDKTHD1n1m3nWEzrj93PWmYP1mkn1cKnTDKnED1rjn1n1FhPaYLrAEYsHEzuWcVrjNQnzYznhnvPvNQPjDzuWnKnHn1PW9zPjc3rHTdrHEQPjTQPkDQn1nvrjcYnW93rj9LPjNQPWE3TED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrE7hPH9vnhFhriYOmyP-sHEOP10Vrjmznzd6rHwbujmkrjnLPWmKTHTKTEDkTHD_nHDYPakdnHEQTHTKUMR_UTDknHTvuAEOnhP6mvR6PvDk&adact=4&psid=136460514211221353733011037&entinfo=42255837070882_q&slot=1000019' tongji_label="listclick" target="_blank" rel="nofollow" >
				<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://t1.58cdn.com.cn/bidding/small/n_v2398b44955f284c5f8c0133745fe90d8b_130_130.jpg'/>
								</a>
								<div class="video-num"><span class="bg"></span><em>1视频</em></div>
							</div>

						<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWNdrjnLnj0krj9zXaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDvrjTkPWNQrjnYP1NznEDKPjczPHN3n10kP1T3rjcKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTdPkDzTyQG0Lw_uyuYTHDKnEDKsHDKTHD1n1m3nWEzrj93PWmYP1mkn1cKnTDKnED1rjn1n1FhPaYLrAEYsHEzuWcVrjNQnzYznhnvPvNQPjDzuWnKnHn1PW9zPjc3rHTdrHEQPjTQPkDQn1nvrjcYnW93rj9LPjNQPWE3TED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrE7hPH9vnhFhriYOmyP-sHEOP10Vrjmznzd6rHwbujmkrjnLPWmKTHTKTEDkTHD_nHDYPakdnHEQTHTKUMR_UTDknHTvuAEOnhP6mvR6PvDk&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42255837070882_q&slot=1000019' target="_blank" rel="nofollow" class="video_icon"></a>
					</td>
				<td class="t" colspan="2">
					<div class="tdiv">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWNdrjnLnj0krj9zXaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDvrjTkPWNQrjnYP1NznEDKPjczPHN3n10kP1T3rjcKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTdPkDzTyQG0Lw_uyuYTHDKnEDKsHDKTHD1n1m3nWEzrj93PWmYP1mkn1cKnTDKnED1rjn1n1FhPaYLrAEYsHEzuWcVrjNQnzYznhnvPvNQPjDzuWnKnHn1PW9zPjc3rHTdrHEQPjTQPkDQn1nvrjcYnW93rj9LPjNQPWE3TED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrE7hPH9vnhFhriYOmyP-sHEOP10Vrjmznzd6rHwbujmkrjnLPWmKTHTKTEDkTHD_nHDYPakdnHEQTHTKUMR_UTDknHTvuAEOnhP6mvR6PvDk&adact=3&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42255837070882_q&slot=1000019' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_adinfo_bj');" target="_blank" class="t ac_linkurl" rel="nofollow" name='42255837070882'>
				<div class="new-long-tit new-long-tit2">
					高价上门回收手机
				</div>
														<span class='tu'></span>
								<span class="jingpin">精</span>
				            					</a>

														<div class="item-desc">
						北京时代通讯、高价上门回收各种品牌手机、全新手机、二手手机、
											</div>
									
									<p class="seller">
								
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/fengtaizhoubian/shouji/">丰台周边</a>
					&nbsp;/ 今天
							</p>
				
				<p class="item-tags">
			<span class="async-tags"></span>
			</p>
    		</div>
    	</td>
    	
<td class="vertop-es">
	<b class='pri'>6666</b> <span>元</span>
</td>
	</tr>
		
		
				
									<tr logr='q_2_68006518347521_42291153297829_4_sortid:643743634@ses:@npos:3@pos:3' infotag="42291153297829_adinfo" data-sign=""  class="new-list" >
		<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_adinfo_tupian&entityId=42291153297829&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
			<div class="ac_linkurl">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWbQnHN1nWbLrjcOXaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDvrjTkPWNQrjnYP1NznEDKPjczrHDQPHnzrH03nWbKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTdPkD1TyQG0Lw_uyuYTHDKnEDKsHDKTHD1n1mLnWDkPWDOPH01Pjb1P1mKnTDKnE7BnhELuHP6midWmW9dsHwhPWnVrAEOPBdbPywWnHNdnHP-P1EKnHn1PW0znHTvnWDdnjc3P1nvnTDQn1nvP1cQnjmznjDvnjmOPWn1TED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrED1mHN1njPBmidbPjEOsHw-PvEVmWIBPzYQnhEdPvELnjDzPADKTHTKTEDkTHD_nHDYnBkdrHbYTHTKUMR_UTD3PHPhnvnLPyNQuWmvnvc1&adact=4&psid=136460514211221353733011037&entinfo=42291153297829_q&slot=1000019' tongji_label="listclick" target="_blank" rel="nofollow" >
				<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://t1.58cdn.com.cn/bidding/small/n_v2c88532c2cdbc4849bc21479817f510c5_130_130.jpg'/>
								</a>
								<div class="video-num"><span class="bg"></span><em>2视频</em></div>
							</div>

						<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWbQnHN1nWbLrjcOXaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDvrjTkPWNQrjnYP1NznEDKPjczrHDQPHnzrH03nWbKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTdPkD1TyQG0Lw_uyuYTHDKnEDKsHDKTHD1n1mLnWDkPWDOPH01Pjb1P1mKnTDKnE7BnhELuHP6midWmW9dsHwhPWnVrAEOPBdbPywWnHNdnHP-P1EKnHn1PW0znHTvnWDdnjc3P1nvnTDQn1nvP1cQnjmznjDvnjmOPWn1TED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrED1mHN1njPBmidbPjEOsHw-PvEVmWIBPzYQnhEdPvELnjDzPADKTHTKTEDkTHD_nHDYnBkdrHbYTHTKUMR_UTD3PHPhnvnLPyNQuWmvnvc1&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42291153297829_q&slot=1000019' target="_blank" rel="nofollow" class="video_icon"></a>
					</td>
				<td class="t" colspan="2">
					<div class="tdiv">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWbQnHN1nWbLrjcOXaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDvrjTkPWNQrjnYP1NznEDKPjczrHDQPHnzrH03nWbKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTdPkD1TyQG0Lw_uyuYTHDKnEDKsHDKTHD1n1mLnWDkPWDOPH01Pjb1P1mKnTDKnE7BnhELuHP6midWmW9dsHwhPWnVrAEOPBdbPywWnHNdnHP-P1EKnHn1PW0znHTvnWDdnjc3P1nvnTDQn1nvP1cQnjmznjDvnjmOPWn1TED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrED1mHN1njPBmidbPjEOsHw-PvEVmWIBPzYQnhEdPvELnjDzPADKTHTKTEDkTHD_nHDYnBkdrHbYTHTKUMR_UTD3PHPhnvnLPyNQuWmvnvc1&adact=3&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42291153297829_q&slot=1000019' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_adinfo_bj');" target="_blank" class="t ac_linkurl" rel="nofollow" name='42291153297829'>
				<div class="new-long-tit new-long-tit2">
					全北京 高价上门回收手机、置换、出售全新手机二手手
				</div>
														<span class='tu'></span>
								<span class="jingpin">精</span>
				            					</a>

														<div class="item-desc">
						北京时代通讯收售、全新手机、二手手机、平板电脑、笔记本、等，
											</div>
									
									<p class="seller">
								
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/aolinpikegognyuan/shouji/">奥林匹克公园</a>
					&nbsp;/ 今天
							</p>
				
				<p class="item-tags">
			<span class="async-tags"></span>
			</p>
    		</div>
    	</td>
    	
<td class="vertop-es">
	<b class='pri'>9000</b> <span>元</span>
</td>
	</tr>
		
		
				
									<tr logr='q_2_67036155779094_39860104407835_4_sortid:625214086@ses:@npos:4@pos:4' infotag="39860104407835_adinfo" data-sign=""  class="new-list" >
		<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_adinfo_tupian&entityId=39860104407835&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
			<div class="ac_linkurl">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1nOrjmknHTYPjTLrjndXaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDvP1T1PWDdPH0LrHTOPTDKn1b3PWTQnjEYnj03n1NKPHTknTDdnjTkTHnvTHDKnHmQnHEOnHTYnHTdPkDYTyQG0Lw_uyuYTHDKnEDKsHDKTHD1n10QP1m3rHmvPjmOPjm3nH0KnTDKnE76ujNLmvcLniYYnj-hsHEvPjbVmH7buBYzuWRBnjmYrHmOujEKnHn1P1DLPW9OPW91rH93Pj9knTDQn1nLnH0vrjbvP1TOrjmQn10vTED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrEDvrjnzn17WPBd6n1bvsHwWrAnVmHuWmBYznvcYuWRWrH7bn1EKTHTKTEDkTHD_nHDYnBkvrjnYTHTKUMR_UT7-nj-hPWPhmWb1mWuhm1NQ&adact=4&psid=136460514211221353733011037&entinfo=39860104407835_q&slot=1000019' tongji_label="listclick" target="_blank" rel="nofollow" >
				<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://t1.58cdn.com.cn/bidding/small/n_v20a89f5cb498344a08d46705ebeb8454f_130_130.jpg'/>
								</a>
								<div class="video-num"><span class="bg"></span><em>2视频</em></div>
							</div>

						<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1nOrjmknHTYPjTLrjndXaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDvP1T1PWDdPH0LrHTOPTDKn1b3PWTQnjEYnj03n1NKPHTknTDdnjTkTHnvTHDKnHmQnHEOnHTYnHTdPkDYTyQG0Lw_uyuYTHDKnEDKsHDKTHD1n10QP1m3rHmvPjmOPjm3nH0KnTDKnE76ujNLmvcLniYYnj-hsHEvPjbVmH7buBYzuWRBnjmYrHmOujEKnHn1P1DLPW9OPW91rH93Pj9knTDQn1nLnH0vrjbvP1TOrjmQn10vTED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrEDvrjnzn17WPBd6n1bvsHwWrAnVmHuWmBYznvcYuWRWrH7bn1EKTHTKTEDkTHD_nHDYnBkvrjnYTHTKUMR_UT7-nj-hPWPhmWb1mWuhm1NQ&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=39860104407835_q&slot=1000019' target="_blank" rel="nofollow" class="video_icon"></a>
					</td>
				<td class="t" colspan="2">
					<div class="tdiv">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1nOrjmknHTYPjTLrjndXaO1pZwVUT7bsHPWnWNLmvmYsHw-nWTVPjKbPiYOuHIBsH-hPjFbPvPhmhnvPkDvP1T1PWDdPH0LrHTOPTDKn1b3PWTQnjEYnj03n1NKPHTknTDdnjTkTHnvTHDKnHmQnHEOnHTYnHTdPkDYTyQG0Lw_uyuYTHDKnEDKsHDKTHD1n10QP1m3rHmvPjmOPjm3nH0KnTDKnE76ujNLmvcLniYYnj-hsHEvPjbVmH7buBYzuWRBnjmYrHmOujEKnHn1P1DLPW9OPW91rH93Pj9knTDQn1nLnH0vrjbvP1TOrjmQn10vTED1P9DQTHD8nTDKsEDKTy6YIZK1rBtfmhC8PH98mvqVsLPCULRJpit4uMFfUHdMmyOJpEDKnHT1sW03sWDzPi3QPjNKnTDkTEDdsjnvTgVqTHDKnHTknjTQrEDvrjnzn17WPBd6n1bvsHwWrAnVmHuWmBYznvcYuWRWrH7bn1EKTHTKTEDkTHD_nHDYnBkvrjnYTHTKUMR_UT7-nj-hPWPhmWb1mWuhm1NQ&adact=3&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=39860104407835_q&slot=1000019' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_adinfo_bj');" target="_blank" class="t ac_linkurl" rel="nofollow" name='39860104407835'>
				<div class="new-long-tit new-long-tit2">
					诚信经营苹果 华为 三星各大品牌手机（支持分期回收
				</div>
														<span class='tu'></span>
								<span class="jingpin">精</span>
				            					</a>

														<div class="item-desc">
						苹果7-32g 国行二手1000 苹果7-128g 国行二手
											</div>
									
									<p class="seller">
								
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/cbd/shouji/">CBD </a>
					&nbsp;/ 今天
							</p>
				
				<p class="item-tags">
			<span class="async-tags"></span>
			</p>
    		</div>
    	</td>
    	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
	</tr>
		
		
				
						<tr logr='j_2_31090533585931_44776197052933_3_sortid:663158039@ses:3!rankershoulbsclickranknew|2!bendifuwudebusirule@npos:1@pos:5' infotag="44776197052933_promationinfo"  data-sign=""  class="ac_item new-list" >
	<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_promationinfo_tupian&entityId=44776197052933&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
		<div  class="ac_linkurl">
			<a href='https://bj.58.com/shouji/44776197052933x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44776197052933_j&slot=-1' tongji_label="listclick" target="_blank" rel="nofollow">
			<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://pic2.58cdn.com.cn/p1/tiny/n_v26d648d11a8544abdb0bdf5f95cd291f8_130_130.jpg' alt='北京创新科技' />
			</a>
					</div>

			</td>
		<td class="t" colspan="2">
				<a href='https://bj.58.com/shouji/44776197052933x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44776197052933_j&slot=-1' tongji_label="listclick" target="_blank" class="t ac_linkurl" >
				<div class="new-long-tit new-long-tit2">
										北京回收苹果手机，回收华为手机，回收小米，OPPO
				</div>
										<span class='tu'></span>
		    					</a>

				
						<div class="item-desc">
														高价上门回收电脑、服务器、办公设备、网络设备的正规企业。本公...
					<span> (今天)</span>
												</div>
					
							<p class="seller">
						
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/guomao/shouji/">国贸</a>
					&nbsp;/ 今天
								</p>
					<p class="item-tags">
			<span class="async-tags"></span>
			</p>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>

				
						<tr logr='j_2_22999328008455_32180495876939_3_sortid:629923842@ses:3!rankershoulbsclickranknew|2!bendifuwudebusirule@npos:2@pos:6' infotag="32180495876939_promationinfo"  data-sign=""  class="ac_item new-list" >
	<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_promationinfo_tupian&entityId=32180495876939&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
		<div  class="ac_linkurl">
			<a href='https://bj.58.com/shouji/32180495876939x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=32180495876939_j&slot=-1' tongji_label="listclick" target="_blank" rel="nofollow">
			<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://pic2.58cdn.com.cn/p1/tiny/n_v2457ccdeaa11d4d7f86e0a17230a26951_130_130.jpg' alt='北京德馨数码' />
			</a>
					</div>

			</td>
		<td class="t" colspan="2">
				<a href='https://bj.58.com/shouji/32180495876939x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=32180495876939_j&slot=-1' tongji_label="listclick" target="_blank" class="t ac_linkurl" >
				<div class="new-long-tit new-long-tit2">
										高价回收苹果11pro-Max-8三星、华为、小米
				</div>
										<span class='tu'></span>
		    					</a>

				
						<div class="item-desc">
														德馨数码回收中心！全国连锁，专业放心！快速检测付款！支持：现...
					<span> (今天)</span>
												</div>
					
							<p class="seller">
						
				<a class="c_666" href="/shouji/">北京</a>
					&nbsp;/ 今天
								</p>
					<p class="item-tags">
			<span class="async-tags"></span>
			</p>
	</td>
	
<td class="vertop-es">
	<b class='pri'>6000</b> <span>元</span>
</td>
</tr>

				
						<tr logr='j_2_67036155779094_39860104407835_3_sortid:625214086@ses:3!rankershoulbsclickranknew|2!bendifuwudebusirule@npos:3@pos:7' infotag="39860104407835_promationinfo"  data-sign=""  class="ac_item new-list" >
	<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_promationinfo_tupian&entityId=39860104407835&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
		<div  class="ac_linkurl">
			<a href='https://bj.58.com/shouji/39860104407835x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=39860104407835_j&slot=-1' tongji_label="listclick" target="_blank" rel="nofollow">
			<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://pic3.58cdn.com.cn/p1/tiny/n_v25b5065a7b22449d79b83ea946d6a4175_130_130.jpg' alt='小孟手机专卖' />
			</a>
						<div class="video-num"><span class="bg"></span><em>2视频</em></div>
					</div>

				<a href='https://bj.58.com/shouji/39860104407835x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=39860104407835_j&slot=-1' target="_blank" rel="nofollow" class="video_icon"></a>
			</td>
		<td class="t" colspan="2">
				<a href='https://bj.58.com/shouji/39860104407835x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=39860104407835_j&slot=-1' tongji_label="listclick" target="_blank" class="t ac_linkurl" >
				<div class="new-long-tit new-long-tit2">
										诚信经营苹果 华为 三星各大品牌手机（回收置换）
				</div>
										<span class='tu'></span>
		    					</a>

				
						<div class="item-desc">
														正规专卖店，正规分期，原封手机，正品国行支持以旧换新！！！正...
					<span> (今天)</span>
												</div>
					
							<p class="seller">
						
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/cbd/shouji/">CBD </a>
					&nbsp;/ 今天
								</p>
					<p class="item-tags">
			<span class="async-tags"></span>
			</p>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>

				
						<tr logr='j_2_34782619515151_39857474720807_3_sortid:624730521@ses:3!rankershoulbsclickranknew|2!bendifuwudebusirule@npos:4@pos:8' infotag="39857474720807_promationinfo"  data-sign=""  class="ac_item new-list" >
	<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_promationinfo_tupian&entityId=39857474720807&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
		<div  class="ac_linkurl">
			<a href='https://bj.58.com/shouji/39857474720807x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=39857474720807_j&slot=-1' tongji_label="listclick" target="_blank" rel="nofollow">
			<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://pic5.58cdn.com.cn/p1/tiny/n_v27e97605a44f24ab58d6ad471baad172a_130_130.jpg' alt='北京江汗科技' />
			</a>
					</div>

			</td>
		<td class="t" colspan="2">
				<a href='https://bj.58.com/shouji/39857474720807x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=39857474720807_j&slot=-1' tongji_label="listclick" target="_blank" class="t ac_linkurl" >
				<div class="new-long-tit new-long-tit2">
										高价上门回收苹果三星手机华为小米手机ipad平板
				</div>
										<span class='tu'></span>
		    					</a>

				
						<div class="item-desc">
														大量回收苹果三星手机OPPO手机vivo手机回收北京江汗科技...
					<span> (今天)</span>
												</div>
					
							<p class="seller">
						
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/chaoyangzhoubian/shouji/">朝阳周边</a>
					&nbsp;/ 今天
								</p>
					<p class="item-tags">
			<span class="async-tags"></span>
			</p>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>

				
									<tr logr='p_2_31090533585931_44454973371414_0_sortid:660648479@ses:@npos:1@pos:9@nodeType:bflowlayer' infotag="44454973371414_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44454973371414x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				上门回收苹果11，回收华为P40Pro，回收三星手
			</div>
														</a>
				<div class="item-desc">
							
				(12-07)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/shouji/">北京</a>
					&nbsp;/ 12-07
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>
		
				
									<tr logr='p_2_68006518347521_44987943091204_0_sortid:664812305@ses:@npos:2@pos:10@nodeType:bflowlayer' infotag="44987943091204_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44987943091204x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				北京高价回收手机 快速上门 优质服务
			</div>
														</a>
				<div class="item-desc">
							
				(6小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/anzhenqiao/shouji/">安贞</a>
					&nbsp;/ 6小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>
		
				
									<tr logr='p_2_22999328008455_44986118421537_0_sortid:664798050@ses:@npos:3@pos:11@nodeType:bflowlayer' infotag="44986118421537_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44986118421537x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				高价回收苹果11pro-Max-8三星、华为、小米
			</div>
														</a>
				<div class="item-desc">
							
				(10小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/shouji/">北京</a>
					&nbsp;/ 10小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>6000</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_2_34782619515151_44990544254365_0_sortid:664832626@ses:@npos:4@pos:12@nodeType:bflowlayer' infotag="44990544254365_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44990544254365x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				高价回收苹果11Promax、华为mate40Pr
			</div>
														</a>
				<div class="item-desc">
							
				(40分钟)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/chaoyangzhoubian/shouji/">朝阳周边</a>
					&nbsp;/ 40分钟
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>
		
				
									<tr logr='p_1_77101043764236_44956039288720_0_sortid:664563056@ses:@npos:5@pos:13@nodeType:bflowlayer' infotag="44956039288720_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44956039288720x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				闲置手机
			</div>
														</a>
				<div class="item-desc">
							
				(01-21)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/tongzhouqu/shouji/">通州</a>
		<span> - </span>
		<a class="c_666" href="/bgbeiguan/shouji/">北关</a>
					&nbsp;/ 01-21
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>1500</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_61495355181071_37833978957841_0_sortid:662487409@ses:@npos:6@pos:14@nodeType:bflowlayer' infotag="37833978957841_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/37833978957841x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为mate10pro 蓝 6+128G 国行全网
			</div>
														</a>
				<div class="item-desc">
							
				(01-02)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fangshan/shouji/">房山</a>
		<span> - </span>
		<a class="c_666" href="/liulihe/shouji/">琉璃河</a>
					&nbsp;/ 01-02
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>400</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_13115728365830_44654071468451_0_sortid:662203933@ses:@npos:7@pos:15@nodeType:bflowlayer' infotag="44654071468451_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44654071468451x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				小米8 128GB
			</div>
														</a>
				<div class="item-desc">
							
				(12-25)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/chaowaidajie/shouji/">朝外大街</a>
					&nbsp;/ 12-25
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>600</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_56548663934991_44989251661577_0_sortid:664822528@ses:@npos:8@pos:16@nodeType:bflowlayer' infotag="44989251661577_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44989251661577x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为畅享10plus全网通手机
			</div>
														</a>
				<div class="item-desc">
							
				(3小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/liuliqiao/shouji/">六里桥</a>
					&nbsp;/ 3小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>699</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_70792921227537_44643123610529_0_sortid:662118402@ses:@npos:9@pos:17@nodeType:bflowlayer' infotag="44643123610529_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44643123610529x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				苹果11 黑色 256g
			</div>
														</a>
				<div class="item-desc">
							
				(12-24)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/guomao/shouji/">国贸</a>
					&nbsp;/ 12-24
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>3500</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_23660948632070_44791601055650_0_sortid:663278383@ses:@npos:10@pos:18@nodeType:bflowlayer' infotag="44791601055650_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44791601055650x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				mate20 9成新 6+128 黑色
			</div>
														</a>
				<div class="item-desc">
							
				(01-06)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/haidian/shouji/">海淀</a>
		<span> - </span>
		<a class="c_666" href="/sijiqinghd/shouji/">四季青</a>
					&nbsp;/ 01-06
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>1550</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_55995344896271_44965764135072_0_sortid:664639032@ses:@npos:11@pos:19@nodeType:bflowlayer' infotag="44965764135072_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44965764135072x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为mate10pro 蓝 6+128G便宜出
			</div>
														</a>
				<div class="item-desc">
							
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fangshan/shouji/">房山</a>
		<span> - </span>
		<a class="c_666" href="/zhoukd/shouji/">周口店</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>400</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_62382003827204_44965805624714_0_sortid:664639356@ses:@npos:12@pos:20@nodeType:bflowlayer' infotag="44965805624714_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44965805624714x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为mate10pro 蓝 6+128G
			</div>
														</a>
				<div class="item-desc">
							
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fangshan/shouji/">房山</a>
		<span> - </span>
		<a class="c_666" href="/zhoukd/shouji/">周口店</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>400</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_76139802851849_44646388028425_0_sortid:662143906@ses:@npos:13@pos:21@nodeType:bflowlayer' infotag="44646388028425_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44646388028425x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				自用苹果x完好无损目前不需要了，便宜转让，中介勿扰
			</div>
														</a>
				<div class="item-desc">
							
				(12-25)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/changping/shouji/">昌平</a>
		<span> - </span>
		<a class="c_666" href="/shahe/shouji/">沙河</a>
					&nbsp;/ 12-25
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>2200</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_74489129007374_44978226983690_0_sortid:664736398@ses:@npos:14@pos:22@nodeType:bflowlayer' infotag="44978226983690_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44978226983690x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为Honor 7x 4+32GB
			</div>
														</a>
				<div class="item-desc">
							
				(01-23)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/caoqiao/shouji/">草桥</a>
					&nbsp;/ 01-23
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>520</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_75193648407059_44969051370765_0_sortid:664664713@ses:@npos:15@pos:23@nodeType:bflowlayer' infotag="44969051370765_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44969051370765x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				苹果11，iPhone11
			</div>
														</a>
				<div class="item-desc">
							
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/bjshuangqiao/shouji/">双桥</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>2550</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_28058740286727_44422693433239_0_sortid:660396292@ses:@npos:16@pos:24@nodeType:bflowlayer' infotag="44422693433239_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44422693433239x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				高价回收苹果华为手机提供苹果服务
			</div>
														</a>
				<div class="item-desc">
							
				(12-04)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/caoqiao/shouji/">草桥</a>
					&nbsp;/ 12-04
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>
		
				
									<tr logr='p_1_69799800295437_44966870406283_0_sortid:664647674@ses:@npos:17@pos:25@nodeType:bflowlayer' infotag="44966870406283_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44966870406283x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				OPPOr11s黑色4+64g全网通
			</div>
														</a>
				<div class="item-desc">
							
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/haidian/shouji/">海淀</a>
		<span> - </span>
		<a class="c_666" href="/xierqi/shouji/">西二旗</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>500</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_61130970780940_41626248209174_0_sortid:638549063@ses:@npos:18@pos:26@nodeType:bflowlayer' infotag="41626248209174_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/youxiting/41626248209174x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				王者荣耀帐号 铂金账号
			</div>
														</a>
				<div class="item-desc">
							
				(03-26)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/beijingzhoubian/shouji/">北京周边</a>
		<span> - </span>
		<a class="c_666" href="/beijingzhoubianqita/shouji/">其他</a>
					&nbsp;/ 03-26
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>
		
				
									<tr logr='p_1_73774298680835_43473734796548_0_sortid:652982552@ses:@npos:19@pos:27@nodeType:bflowlayer' infotag="43473734796548_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/43473734796548x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				二手苹果11国行128G99新无拆无修
			</div>
														</a>
				<div class="item-desc">
							
				(09-09)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
					&nbsp;/ 09-09
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>
		
				
									<tr logr='p_1_75024332721171_44945924884500_0_sortid:664484037@ses:@npos:20@pos:28@nodeType:bflowlayer' infotag="44945924884500_commonflowlayer"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44945924884500x.shtml?nodeType=bflowlayer&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commonflowlayer_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				小米8探索版8加128
			</div>
														</a>
				<div class="item-desc">
							
				(01-20)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/guomao/shouji/">国贸</a>
					&nbsp;/ 01-20
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>600</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_73183543809563_43892140061590_0_sortid:664834084@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:1@pos:29' infotag="43892140061590_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/43892140061590x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=43892140061590_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				苹果手机，苹果7P/8P/X原装二手
			</div>
														</a>
				<div class="item-desc">
							出售苹果7P/8P/X 原装二手，无暗病，无进水，无拆无修全网通，移动...
				(15分钟)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/shouji/">北京</a>
					&nbsp;/ 15分钟
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>
		
				
									<tr logr='p_2_34782619515151_44990544254365_0_sortid:664832626@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:2@pos:30' infotag="44990544254365_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44990544254365x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44990544254365_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				高价回收苹果11Promax、华为mate40Pr
			</div>
														</a>
				<div class="item-desc">
							北京上门回收苹果三星手机OPPO手机vivo手机回收北京江汗科技电脑回...
				(40分钟)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/chaoyangzhoubian/shouji/">朝阳周边</a>
					&nbsp;/ 40分钟
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>
		
				
									<tr logr='p_1_62059603549955_44989562911777_0_sortid:664824960@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:3@pos:31' infotag="44989562911777_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44989562911777x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44989562911777_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为Mate40 pro秘银色8+256G全新未开
			</div>
														</a>
				<div class="item-desc">
							华为Mate40 pro秘银色8+256G全新未开封
				(2小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/gaobeidianbj/shouji/">高碑店</a>
					&nbsp;/ 2小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>7250</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_56548663934991_44989251661577_0_sortid:664822528@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:4@pos:32' infotag="44989251661577_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44989251661577x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44989251661577_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为畅享10plus全网通手机
			</div>
														</a>
				<div class="item-desc">
							纯原美女一手自用 无拆修，客服保修还有大包年，秒杀不议价谢谢。
				(3小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/liuliqiao/shouji/">六里桥</a>
					&nbsp;/ 3小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>699</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_2_22999328008455_44988863522196_0_sortid:664819496@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:5@pos:33' infotag="44988863522196_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44988863522196x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44988863522196_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				高价回收苹果11pro-Max-8三星、华为、小米
			</div>
														</a>
				<div class="item-desc">
							德馨数码回收中心！全国连锁，专业放心！快速检测付款！支持：现金支付，网...
				(4小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/shouji/">北京</a>
					&nbsp;/ 4小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>6000</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_2_68006518347521_44988863582866_0_sortid:664819496@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:6@pos:34' infotag="44988863582866_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44988863582866x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44988863582866_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				高价上门回收手机
			</div>
														</a>
				<div class="item-desc">
							北京时代通讯、高价上门回收各种品牌手机、全新手机、二手手机、笔记本、平...
				(4小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/fengtaizhoubian/shouji/">丰台周边</a>
					&nbsp;/ 4小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>6666</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_40711099314454_44988806973845_0_sortid:664819054@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:7@pos:35' infotag="44988806973845_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44988806973845x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44988806973845_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				九成新手机
			</div>
														</a>
				<div class="item-desc">
							voviX20！！！！6＋128G超大内存，不卡顿，流畅清新
				(4小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/dongdaqiao/shouji/">东大桥</a>
					&nbsp;/ 4小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>613</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_54619981_44987795256866_0_sortid:664811150@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:8@pos:36' infotag="44987795256866_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44987795256866x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44987795256866_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为b5 商务金 
			</div>
														</a>
				<div class="item-desc">
							此品为公司库存 二手 有磨损 成色很好 支持电话接听 华为b5商务金 ...
				(6小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/muxiyuan/shouji/">木樨园</a>
					&nbsp;/ 6小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>299</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_46284061248275_44987404463765_0_sortid:664808097@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:9@pos:37' infotag="44987404463765_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44987404463765x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44987404463765_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				全新2021年新款vivo手机5G网络全网通
			</div>
														</a>
				<div class="item-desc">
							手机1月24日购买，给老人用，结果不会用，特价出售，保证正品，绝无套路...
				(7小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/yayuncun/shouji/">亚运村</a>
					&nbsp;/ 7小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>1200</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_4571913731846_44987023698077_0_sortid:664805122@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:10@pos:38' infotag="44987023698077_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44987023698077x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44987023698077_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				个人闲置vivo手机 低价转让
			</div>
														</a>
				<div class="item-desc">
							个人因换电话了，这台手机闲置。电话型号：vivo Y81s，内存64G...
				(8小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/jiuxianqiao/shouji/">酒仙桥</a>
					&nbsp;/ 8小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>550</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_25544712154374_44987022991394_0_sortid:664805116@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:11@pos:39' infotag="44987022991394_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44987022991394x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44987022991394_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				自用国行苹果x手机，三网通，无划痕
			</div>
														</a>
				<div class="item-desc">
							苹果x手机，自用的苹果x，256g内存！刚出的时候买的，电池77容量，...
				(8小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/haidian/shouji/">海淀</a>
		<span> - </span>
		<a class="c_666" href="/xibeiwang/shouji/">西北旺</a>
					&nbsp;/ 8小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>2200</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_77017467144464_44987019222946_0_sortid:664805087@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:12@pos:40' infotag="44987019222946_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44987019222946x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44987019222946_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				三星s20 手机
			</div>
														</a>
				<div class="item-desc">
							国行全网通 5G 8+128 无拆机无维修 99新
				(8小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/beijingzhoubian/shouji/">北京周边</a>
		<span> - </span>
		<a class="c_666" href="/beijingzhoubianqita/shouji/">其他</a>
					&nbsp;/ 8小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>1200</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_59649892143117_44986720269349_0_sortid:664802751@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:13@pos:41' infotag="44986720269349_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44986720269349x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44986720269349_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				二手苹果无线耳机.换华为了用不上
			</div>
														</a>
				<div class="item-desc">
							原厂苹果无线耳机.用了一年
				(7小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/shibalidian/shouji/">十八里店</a>
					&nbsp;/ 7小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>200</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_34424738104076_44986676109191_0_sortid:664802406@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:14@pos:42' infotag="44986676109191_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44986676109191x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44986676109191_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				自用华为mate9便宜出
			</div>
														</a>
				<div class="item-desc">
							自用华为手机mate9，保养良好，运行顺畅，便宜出
				(9小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/haidian/shouji/">海淀</a>
		<span> - </span>
		<a class="c_666" href="/gongzhufen/shouji/">公主坟</a>
					&nbsp;/ 9小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>400</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_2_31090533585931_44986663067780_0_sortid:664802305@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:15@pos:43' infotag="44986663067780_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44986663067780x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44986663067780_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				北京回收vivo手机OPPO手机，回收华为苹果手机
			</div>
														</a>
				<div class="item-desc">
							高价上门回收电脑、服务器、办公设备、网络设备的正规企业。本公司成立于2...
				(9小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/changping/shouji/">昌平</a>
		<span> - </span>
		<a class="c_666" href="/tiantongyuan/shouji/">天通苑</a>
					&nbsp;/ 9小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
</tr>
		
				
									<tr logr='p_1_34955377567495_44986609917329_0_sortid:664801889@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:16@pos:44' infotag="44986609917329_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44986609917329x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44986609917329_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				全新未开封华为matexs
			</div>
														</a>
				<div class="item-desc">
							国行全新未开封华为折叠机皇8+512星际蓝
				(9小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/chaoqingbankuai/shouji/">朝青板块</a>
					&nbsp;/ 9小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>19500</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_57105124693266_44985806888605_0_sortid:664795616@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:17@pos:45' infotag="44985806888605_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44985806888605x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44985806888605_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				换手机了用不上了
			</div>
														</a>
				<div class="item-desc">
							换手机了用不上了便宜买了
				(10小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/haidian/shouji/">海淀</a>
		<span> - </span>
		<a class="c_666" href="/zhongguancun/shouji/">中关村</a>
					&nbsp;/ 10小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>760</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_29893618950153_44985776128775_0_sortid:664795375@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:18@pos:46' infotag="44985776128775_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44985776128775x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44985776128775_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				年度机皇
			</div>
														</a>
				<div class="item-desc">
							先到先得！较低不低于7200！
				(11小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/fengtaizhoubian/shouji/">丰台周边</a>
					&nbsp;/ 11小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>7299</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_77306537192211_44982906536999_0_sortid:664772957@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:19@pos:47' infotag="44982906536999_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44982906536999x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44982906536999_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				想买的可以私信聊聊
			</div>
														</a>
				<div class="item-desc">
							手机特别好 需要点钱急卖
				(17小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/beijingzhoubian/shouji/">北京周边</a>
		<span> - </span>
		<a class="c_666" href="/beijingzhoubianqita/shouji/">其他</a>
					&nbsp;/ 17小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>2000</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_74360061016578_44981448941060_0_sortid:664761569@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:20@pos:48' infotag="44981448941060_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44981448941060x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44981448941060_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				vivoX20plusA全网通4+64GB手机
			</div>
														</a>
				<div class="item-desc">
							vivoX20plusA全网通4+64GB手机，个人手机，非诚勿扰！
				(20小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/haidian/shouji/">海淀</a>
		<span> - </span>
		<a class="c_666" href="/jimenqiao/shouji/">蓟门桥</a>
					&nbsp;/ 20小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>520</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_74347704561948_44980932428289_0_sortid:664757534@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:21@pos:49' infotag="44980932428289_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44980932428289x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44980932428289_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				黑色全网通大屏viviX20plusA 4+64G
			</div>
														</a>
				<div class="item-desc">
							黑色全网通大屏viviX20plusA 4+64GB手机，超大屏幕，影...
				(21小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/haidian/shouji/">海淀</a>
		<span> - </span>
		<a class="c_666" href="/dazhongsi/shouji/">大钟寺</a>
					&nbsp;/ 21小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>666</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_73800594833447_44980530438548_0_sortid:664754393@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:22@pos:50' infotag="44980530438548_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44980530438548x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44980530438548_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				苹果x银白色
			</div>
														</a>
				<div class="item-desc">
							本人自用的，需要的可以联系我，非诚勿扰
				(22小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/haidian/shouji/">海淀</a>
		<span> - </span>
		<a class="c_666" href="/sijiqinghd/shouji/">四季青</a>
					&nbsp;/ 22小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>1800</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_76839291462698_44980490112912_0_sortid:664754078@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:23@pos:51' infotag="44980490112912_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44980490112912x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44980490112912_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为荣耀8+256G高配5G版手机
			</div>
														</a>
				<div class="item-desc">
							个人闲置，华为荣耀正品，八核CPU麒麟，8+256高配，蓝牙导航wif...
				(22小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/bjxiaotunlu/shouji/">小屯路</a>
					&nbsp;/ 22小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>500</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_61970159569932_44980114581924_0_sortid:664751145@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:24@pos:52' infotag="44980114581924_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44980114581924x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44980114581924_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				poopreno5pro 5G
			</div>
														</a>
				<div class="item-desc">
							用了一天。孩子不要安卓的。手机是曲屏的 四个摄像头。是Pro哦。5G哦
				(23小时)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/beijingzhoubian/shouji/">北京周边</a>
		<span> - </span>
		<a class="c_666" href="/yantaibj/shouji/">烟台</a>
					&nbsp;/ 23小时
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>3200</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_74489129007374_44978344875930_0_sortid:664737319@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:25@pos:53' infotag="44978344875930_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44978344875930x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44978344875930_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				小米note LET全网通16G经典黑
			</div>
														</a>
				<div class="item-desc">
							小米note LET全网通16G，后盖经典黑带闪星，适合老年人接打电话...
				(01-23)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/lizeqiao/shouji/">丽泽桥</a>
					&nbsp;/ 01-23
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>220</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_23332494403847_44976958686220_0_sortid:664726489@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:26@pos:54' infotag="44976958686220_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44976958686220x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44976958686220_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				16g苹果
			</div>
														</a>
				<div class="item-desc">
							完美备用机便宜出有要的联系
				(01-23)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/shunyi/shouji/">顺义</a>
		<span> - </span>
		<a class="c_666" href="/nancai/shouji/">南彩</a>
					&nbsp;/ 01-23
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>300</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_71081438984715_44976883227788_0_sortid:664725900@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:27@pos:55' infotag="44976883227788_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44976883227788x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44976883227788_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				喜欢的朋友赶快下手，处理了放假
			</div>
														</a>
				<div class="item-desc">
							128g内存8核哦，8.6成新，喜欢的朋友可以聊聊，价格优惠
				(01-23)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/chaoyanggongyuan/shouji/">朝阳公园</a>
					&nbsp;/ 01-23
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>380</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_77291979006723_44975710510235_0_sortid:664716738@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:28@pos:56' infotag="44975710510235_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44975710510235x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44975710510235_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为20i 4+128
			</div>
														</a>
				<div class="item-desc">
							华为手机 二手 因为换了手机所以卖掉，只是后壳有裂痕，就像手机钢化膜一...
				(01-23)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/huairou/shouji/">怀柔</a>
		<span> - </span>
		<a class="c_666" href="/huairoubj/shouji/">怀柔城区</a>
					&nbsp;/ 01-23
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>450</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_5302699079175_44970093341570_0_sortid:664672854@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:29@pos:57' infotag="44970093341570_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44970093341570x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44970093341570_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				转让全新华为5G手机
			</div>
														</a>
				<div class="item-desc">
							华为mate30pro 5g手机（内存8g+256g）全新未拆封亲戚送...
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/dongcheng/shouji/">东城</a>
		<span> - </span>
		<a class="c_666" href="/dongzhimen/shouji/">东直门</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>6000</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_75193648407059_44969051370765_0_sortid:664664713@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:30@pos:58' infotag="44969051370765_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44969051370765x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44969051370765_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				苹果11，iPhone11
			</div>
														</a>
				<div class="item-desc">
							苹果11，自己用的，没有磕碰划痕，耳机充电器都在，内存是128的，也没...
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/bjshuangqiao/shouji/">双桥</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>2550</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_75903659898644_44968566010148_0_sortid:664660921@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:31@pos:59' infotag="44968566010148_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44968566010148x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44968566010148_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				oppoA3全网通128G手机
			</div>
														</a>
				<div class="item-desc">
							oppoA3全网通手机，4+128GB内存组合，我在通州买的，用了不到...
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/lizeqiao/shouji/">丽泽桥</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>650</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_69799800295437_44967359846426_0_sortid:664651498@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:32@pos:60' infotag="44967359846426_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44967359846426x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44967359846426_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				三星SM-G9250 3+32G白色全网通
			</div>
														</a>
				<div class="item-desc">
							三星SM-G9250 3+32G白色全网通
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/caoqiao/shouji/">草桥</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>299</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_57897448065043_44967274528665_0_sortid:664650832@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:33@pos:61' infotag="44967274528665_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44967274528665x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44967274528665_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				苹果xs
			</div>
														</a>
				<div class="item-desc">
							这台苹果xs功能强大全原装无拆无修，送18w快充
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/datun/shouji/">大屯</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>2200</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_62382003827204_44965805624714_0_sortid:664639356@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:34@pos:62' infotag="44965805624714_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44965805624714x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44965805624714_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为mate10pro 蓝 6+128G
			</div>
														</a>
				<div class="item-desc">
							华为mate10pro 蓝 6+128G 国行全网通双卡双待，成色很好...
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fangshan/shouji/">房山</a>
		<span> - </span>
		<a class="c_666" href="/zhoukd/shouji/">周口店</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>400</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='p_1_55995344896271_44965778130181_0_sortid:664639141@ses:3!rankershoudefaultranknew|2!bendifuwupublicclickrule@npos:35@pos:63' infotag="44965778130181_commoninfo"  data-sign=""  class="ac_item new-list" >
		<td class="t" colspan="3">
			<a href='https://bj.58.com/shouji/44965778130181x.shtml?link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=44965778130181_p&slot=-1' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_commoninfo_bj');" target="_blank" class="t ac_linkurl" >
			<div class="new-long-tit new-long-tit2">
				华为mate10pro 蓝 6+128G便宜出
			</div>
														</a>
				<div class="item-desc">
							华为mate10pro 蓝 6+128G 国行全网通双卡双待，成色很好...
				(01-22)
					</div>
				<p class="seller">
				
				<a class="c_666" href="/fangshan/shouji/">房山</a>
		<span> - </span>
		<a class="c_666" href="/zhoukd/shouji/">周口店</a>
					&nbsp;/ 01-22
							</p>

		<p class="item-tags">
			<span class="async-tags"></span>
			</p>
		</div>
	</td>
	
<td class="vertop-es">
	<b class='pri'>400</b> <span>元</span>
</td>
</tr>
		
				
									<tr logr='q_2_67036155779094_39860104407835_4_sortid:625214086@ses:@npos:1@pos:64' infotag="39860104407835_adsumplayinfo" data-sign=""  class="new-list" >
		<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_adsumplayinfo_tupian&entityId=39860104407835&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
			<div class="ac_linkurl">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1nOrjmknHTYPjTLrjndXaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDvP1T1PWDdPH0LrHTOPTDKn1b3PWTQnjEYnj03n1NKPHTknTDdnjTkTHnvTHDKnHmQnHEOnHTYnHTvnkDQTRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nLnH0vrjbvPWEvrHEvrjDLTHTKTHDKmWmkmyEvP1cVrHN1PzYYmHFhsyckmWcVuju-mHNdnHnvP1K6THD1n10QP1m3rHm3n1b3rjE3njTKnHn1P1DLPW9OPW0krH9vnHnLP9DKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKn1b1PyNznWnVPW0duiYYnAmzsHbznymVujmkPH01ujnYnhP-TEDkTEDKnTDQsjDQPjc_PW91PTDkTyOdUAkKrAP6P1cQmyD1mhuWnhP6P9&adact=4&psid=136460514211221353733011037&entinfo=39860104407835_q&slot=1002464' tongji_label="listclick" target="_blank" rel="nofollow" >
				<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://t1.58cdn.com.cn/bidding/small/n_v20a89f5cb498344a08d46705ebeb8454f_130_130.jpg'/>
								</a>
								<div class="video-num"><span class="bg"></span><em>2视频</em></div>
							</div>

						<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1nOrjmknHTYPjTLrjndXaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDvP1T1PWDdPH0LrHTOPTDKn1b3PWTQnjEYnj03n1NKPHTknTDdnjTkTHnvTHDKnHmQnHEOnHTYnHTvnkDQTRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nLnH0vrjbvPWEvrHEvrjDLTHTKTHDKmWmkmyEvP1cVrHN1PzYYmHFhsyckmWcVuju-mHNdnHnvP1K6THD1n10QP1m3rHm3n1b3rjE3njTKnHn1P1DLPW9OPW0krH9vnHnLP9DKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKn1b1PyNznWnVPW0duiYYnAmzsHbznymVujmkPH01ujnYnhP-TEDkTEDKnTDQsjDQPjc_PW91PTDkTyOdUAkKrAP6P1cQmyD1mhuWnhP6P9&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=39860104407835_q&slot=1002464' target="_blank" rel="nofollow" class="video_icon"></a>
					</td>
				<td class="t" colspan="2">
					<div class="tdiv">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1nOrjmknHTYPjTLrjndXaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDvP1T1PWDdPH0LrHTOPTDKn1b3PWTQnjEYnj03n1NKPHTknTDdnjTkTHnvTHDKnHmQnHEOnHTYnHTvnkDQTRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nLnH0vrjbvPWEvrHEvrjDLTHTKTHDKmWmkmyEvP1cVrHN1PzYYmHFhsyckmWcVuju-mHNdnHnvP1K6THD1n10QP1m3rHm3n1b3rjE3njTKnHn1P1DLPW9OPW0krH9vnHnLP9DKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKn1b1PyNznWnVPW0duiYYnAmzsHbznymVujmkPH01ujnYnhP-TEDkTEDKnTDQsjDQPjc_PW91PTDkTyOdUAkKrAP6P1cQmyD1mhuWnhP6P9&adact=3&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=39860104407835_q&slot=1002464' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_adsumplayinfo_bj');" target="_blank" class="t ac_linkurl" rel="nofollow" name='39860104407835'>
				<div class="new-long-tit new-long-tit2">
					诚信经营苹果 华为 三星各大品牌手机（支持分期回收
				</div>
														<span class='tu'></span>
								<span class="jingpin">精</span>
				            					</a>

														<div class="item-desc">
						苹果7-32g 国行二手1000 苹果7-128g 国行二手
											</div>
									
									<p class="seller">
								
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/cbd/shouji/">CBD </a>
					&nbsp;/ 今天
							</p>
				
				<p class="item-tags">
			<span class="async-tags"></span>
			</p>
    		</div>
    	</td>
    	
<td class="vertop-es">
	<b class='pri'>面议</b>
</td>
	</tr>
		
		
				
									<tr logr='q_2_68006518347521_42255837070882_4_sortid:643467726@ses:@npos:2@pos:65' infotag="42255837070882_adsumplayinfo" data-sign=""  class="new-list" >
		<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_adsumplayinfo_tupian&entityId=42255837070882&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
			<div class="ac_linkurl">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWNdrjnLnj0krj9zXaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDvrjTkPWNQrjnYP1NznEDKPjczPHN3n10kP1T3rjcKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTvnkDzTRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nvrjcYnW93rjmvPj0vnjnzTHTKTHDKuWTvmW0YP1nVnAc1PBYYPvm1sycknWbVmhnzujI-PWcQuWTOTHD1n1m3nWEzrjbkPHbYnHEknH0KnHn1PW9zPjc3rj93P1EdnHmYrTDKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKuWnduyw6mW9VnWD3raYYmWbYsyDvPWcVmhN3uANYm1nYmvw6TEDkTEDKnTDQsjDQPjE_PHDYnEDkTyOdUAkKn1NOPWEdPvNYrj-WnhRWnk&adact=4&psid=136460514211221353733011037&entinfo=42255837070882_q&slot=1002464' tongji_label="listclick" target="_blank" rel="nofollow" >
				<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://t1.58cdn.com.cn/bidding/small/n_v2398b44955f284c5f8c0133745fe90d8b_130_130.jpg'/>
								</a>
								<div class="video-num"><span class="bg"></span><em>1视频</em></div>
							</div>

						<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWNdrjnLnj0krj9zXaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDvrjTkPWNQrjnYP1NznEDKPjczPHN3n10kP1T3rjcKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTvnkDzTRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nvrjcYnW93rjmvPj0vnjnzTHTKTHDKuWTvmW0YP1nVnAc1PBYYPvm1sycknWbVmhnzujI-PWcQuWTOTHD1n1m3nWEzrjbkPHbYnHEknH0KnHn1PW9zPjc3rj93P1EdnHmYrTDKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKuWnduyw6mW9VnWD3raYYmWbYsyDvPWcVmhN3uANYm1nYmvw6TEDkTEDKnTDQsjDQPjE_PHDYnEDkTyOdUAkKn1NOPWEdPvNYrj-WnhRWnk&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42255837070882_q&slot=1002464' target="_blank" rel="nofollow" class="video_icon"></a>
					</td>
				<td class="t" colspan="2">
					<div class="tdiv">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWNdrjnLnj0krj9zXaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDvrjTkPWNQrjnYP1NznEDKPjczPHN3n10kP1T3rjcKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTvnkDzTRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nvrjcYnW93rjmvPj0vnjnzTHTKTHDKuWTvmW0YP1nVnAc1PBYYPvm1sycknWbVmhnzujI-PWcQuWTOTHD1n1m3nWEzrjbkPHbYnHEknH0KnHn1PW9zPjc3rj93P1EdnHmYrTDKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKuWnduyw6mW9VnWD3raYYmWbYsyDvPWcVmhN3uANYm1nYmvw6TEDkTEDKnTDQsjDQPjE_PHDYnEDkTyOdUAkKn1NOPWEdPvNYrj-WnhRWnk&adact=3&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42255837070882_q&slot=1002464' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_adsumplayinfo_bj');" target="_blank" class="t ac_linkurl" rel="nofollow" name='42255837070882'>
				<div class="new-long-tit new-long-tit2">
					高价上门回收手机
				</div>
														<span class='tu'></span>
								<span class="jingpin">精</span>
				            					</a>

														<div class="item-desc">
						北京时代通讯、高价上门回收各种品牌手机、全新手机、二手手机、
											</div>
									
									<p class="seller">
								
				<a class="c_666" href="/fengtai/shouji/">丰台</a>
		<span> - </span>
		<a class="c_666" href="/fengtaizhoubian/shouji/">丰台周边</a>
					&nbsp;/ 今天
							</p>
				
				<p class="item-tags">
			<span class="async-tags"></span>
			</p>
    		</div>
    	</td>
    	
<td class="vertop-es">
	<b class='pri'>6666</b> <span>元</span>
</td>
	</tr>
		
		
				
									<tr logr='q_2_68006518347521_42291153297829_4_sortid:643743634@ses:@npos:3@pos:66' infotag="42291153297829_adsumplayinfo" data-sign=""  class="new-list" >
		<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_adsumplayinfo_tupian&entityId=42291153297829&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
			<div class="ac_linkurl">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWbQnHN1nWbLrjcOXaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDvrjTkPWNQrjnYP1NznEDKPjczrHDQPHnzrH03nWbKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTvnkD1TRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nvP1cQnjmQrHNLn1EOn10vTHTKTHDKnhuWujKbPyEVnj9zPiYYuHu-sy7BPymVnyEQuAn1PhNYnvD1THD1n1mLnWDkPWcQPHTzrj01PWTKnHn1PW0znHTvnWTQPWTvrHm1nkDKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKPWFhmWuhnH9VuH6BuBYYPj-bsyD3mhEVnv7huWEYuHKWnj6bTEDkTEDKnTDQsjDQPjc_PHbOPTDkTyOdUAkKnWDkuWw-mW-6myuWujK-n9&adact=4&psid=136460514211221353733011037&entinfo=42291153297829_q&slot=1002464' tongji_label="listclick" target="_blank" rel="nofollow" >
				<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://t1.58cdn.com.cn/bidding/small/n_v2c88532c2cdbc4849bc21479817f510c5_130_130.jpg'/>
								</a>
								<div class="video-num"><span class="bg"></span><em>2视频</em></div>
							</div>

						<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWbQnHN1nWbLrjcOXaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDvrjTkPWNQrjnYP1NznEDKPjczrHDQPHnzrH03nWbKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTvnkD1TRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nvP1cQnjmQrHNLn1EOn10vTHTKTHDKnhuWujKbPyEVnj9zPiYYuHu-sy7BPymVnyEQuAn1PhNYnvD1THD1n1mLnWDkPWcQPHTzrj01PWTKnHn1PW0znHTvnWTQPWTvrHm1nkDKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKPWFhmWuhnH9VuH6BuBYYPj-bsyD3mhEVnv7huWEYuHKWnj6bTEDkTEDKnTDQsjDQPjc_PHbOPTDkTyOdUAkKnWDkuWw-mW-6myuWujK-n9&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42291153297829_q&slot=1002464' target="_blank" rel="nofollow" class="video_icon"></a>
					</td>
				<td class="t" colspan="2">
					<div class="tdiv">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EznWbQnHN1nWbLrjcOXaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDvrjTkPWNQrjnYP1NznEDKPjczrHDQPHnzrH03nWbKPHDknTDdnHTkTHnvTHDKnHmQnHEOnHTYnHTvnkD1TRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nvP1cQnjmQrHNLn1EOn10vTHTKTHDKnhuWujKbPyEVnj9zPiYYuHu-sy7BPymVnyEQuAn1PhNYnvD1THD1n1mLnWDkPWcQPHTzrj01PWTKnHn1PW0znHTvnWTQPWTvrHm1nkDKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKPWFhmWuhnH9VuH6BuBYYPj-bsyD3mhEVnv7huWEYuHKWnj6bTEDkTEDKnTDQsjDQPjc_PHbOPTDkTyOdUAkKnWDkuWw-mW-6myuWujK-n9&adact=3&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42291153297829_q&slot=1002464' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_adsumplayinfo_bj');" target="_blank" class="t ac_linkurl" rel="nofollow" name='42291153297829'>
				<div class="new-long-tit new-long-tit2">
					全北京 高价上门回收手机、置换、出售全新手机二手手
				</div>
														<span class='tu'></span>
								<span class="jingpin">精</span>
				            					</a>

														<div class="item-desc">
						北京时代通讯收售、全新手机、二手手机、平板电脑、笔记本、等，
											</div>
									
									<p class="seller">
								
				<a class="c_666" href="/chaoyang/shouji/">朝阳</a>
		<span> - </span>
		<a class="c_666" href="/aolinpikegognyuan/shouji/">奥林匹克公园</a>
					&nbsp;/ 今天
							</p>
				
				<p class="item-tags">
			<span class="async-tags"></span>
			</p>
    		</div>
    	</td>
    	
<td class="vertop-es">
	<b class='pri'>9000</b> <span>元</span>
</td>
	</tr>
		
		
				
									<tr logr='q_2_22999328008455_42643242627073_4_sortid:646494332@ses:@npos:4@pos:67' infotag="42643242627073_adsumplayinfo" data-sign=""  class="new-list" >
		<td class="img" onclick="clickLog('from=pc_list_shouji_xinxi_adsumplayinfo_tupian&entityId=42643242627073&entityType=0&psid=136460514211221353733011037&abVersion=&filterparams=');">
			<div class="ac_linkurl">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EzPWE1nWEzPWcLnj01XaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDznWbOrHnzrjTkrjEdPEDKPjcvPjnzPjcvnW0kP1nKrHbknTDdP1ckTHnvTHDKnHmQnHEOnHTYnHTvnkDYTRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nvP1cvnWb1P1TdPWmzPjmYTHTKTHDKPjbOnAnOrANVPjT1nBYYnWb3sHbYPWTVuHwbPAn3ryNknAuhTHD1n1mLnWmzrHEkPH0OrjEknjDKnHn1PW0zPWcOn1b1nWDdPj93nTDKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKPHNdmWubrHbVrHn1PiYYPH03syF-rHbVryEOnWPBPAEzPh76TEDkTEDKnTDQTHTKUMR_UT7hP1-WujbLPAwBmvubuWcv&adact=4&psid=136460514211221353733011037&entinfo=42643242627073_q&slot=1002464' tongji_label="listclick" target="_blank" rel="nofollow" >
				<img src="//img.58cdn.com.cn/n/images/list/noimg.gif" lazy_src='https://t1.58cdn.com.cn/bidding/small/n_v24a047d973f274c3f9c6fb1f6d55db435_130_130.jpg'/>
								</a>
							</div>

					</td>
				<td class="t" colspan="2">
					<div class="tdiv">
				<a href='https://legoclick.58.com/jump?target=szqBpB3draOWUvYf0v6fIyGGs1EzPWE1nWEzPWcLnj01XaO1pZwVUT7bsHnOPAR6PA7hsHNkrH9VPj93uaYOPjmLsyDvuHT1PAEvP1nYuTDznWbOrHnzrjTkrjEdPEDKPjcvPjnzPjcvnW0kP1nKrHbknTDdP1ckTHnvTHDKnHmQnHEOnHTYnHTvnkDYTRKjsNQFNdwnwNuNsNF5R7w5HEDQTHDKTiYQTEDQn1nvP1cvnWb1P1TdPWmzPjmYTHTKTHDKPjbOnAnOrANVPjT1nBYYnWb3sHbYPWTVuHwbPAn3ryNknAuhTHD1n1mLnWmzrHEkPH0OrjEknjDKnHn1PW0zPWcOn1b1nWDdPj93nTDKn1mKnEDQsWTKTiYKTE7CIZwk01CfsvFJsWN3shPfUiq1pAqdphbf5vuzUvYquv78phbKTHDknz3Lra3QnWN8nHEdTHTKnTDKPik1P97exEDQTHDknjcYPWEKPHNdmWubrHbVrHn1PiYYPH03syF-rHbVryEOnWPBPAEzPh76TEDkTEDKnTDQTHTKUMR_UT7hP1-WujbLPAwBmvubuWcv&adact=3&link_abtest=PCFlowSupport_B&psid=136460514211221353733011037&entinfo=42643242627073_q&slot=1002464' tongji_label="listclick" onclick="clickLog('from=pc_list_shouji_xinxi_adsumplayinfo_bj');" target="_blank" class="t ac_linkurl" rel="nofollow" name='42643242627073'>
				<div class="new-long-tit new-long-tit2">
					高价回收苹果12pro-Max-11三星、华为等
				</div>
														<span class='tu'></span>
								<span class="jingpin">精</span>
				            					</a>

														<div class="item-desc">
						全北京
											</div>
									
									<p class="seller">
								
				<a class="c_666" href="/shouji/">北京</a>
					&nbsp;/ 今天
							</p>
				
				<p class="item-tags">
			<span class="async-tags"></span>
			</p>
    		</div>
    	</td>
    	
<td class="vertop-es">
	<b class='pri'>6000</b> <span>元</span>
</td>
	</tr>
		
		
		</table>
                         	<div id="gjwb_bottom_jz_ad"></div>
	<div id="bottom_google_ad"></div>
<div class="pager" >
	 <strong><span>1</span></strong><a href="/shouji/pn2/?from=ganji" onClick="clickLog('from=huangye_list_fanye_2');"><span>2</span></a><a href="/shouji/pn3/?from=ganji" onClick="clickLog('from=huangye_list_fanye_3');"><span>3</span></a><a href="/shouji/pn4/?from=ganji" onClick="clickLog('from=huangye_list_fanye_4');"><span>4</span></a><a href="/shouji/pn5/?from=ganji" onClick="clickLog('from=huangye_list_fanye_5');"><span>5</span></a><a href="/shouji/pn6/?from=ganji" onClick="clickLog('from=huangye_list_fanye_6');"><span>6</span></a><a href="/shouji/pn7/?from=ganji" onClick="clickLog('from=huangye_list_fanye_7');"><span>7</span></a><a href="/shouji/pn8/?from=ganji" onClick="clickLog('from=huangye_list_fanye_8');"><span>8</span></a><a href="/shouji/pn9/?from=ganji" onClick="clickLog('from=huangye_list_fanye_9');"><span>9</span></a><a href="/shouji/pn10/?from=ganji" onClick="clickLog('from=huangye_list_fanye_10');"><span>10</span></a><a href="/shouji/pn11/?from=ganji" onClick="clickLog('from=huangye_list_fanye_11');"><span>11</span></a><a href="/shouji/pn12/?from=ganji" onClick="clickLog('from=huangye_list_fanye_12');"><span>12</span></a><a class="next" href="/shouji/pn2/?from=ganji" onClick="clickLog('from=huangye_list_fanye_2');"><span>下一页</span></a>
</div>

<div id="infocont">
						<a target="_blank" rel="nofollow" href="//post.58.com/hy/1/36/s5" onclick="clickLog('pc_list_shouji_bottom_fabu');" >马上发布一条二手手机信息&raquo;</a>
			</div>

	<div id="gjwb_bottom_jz_ad"></div>
	<div id="bottom_google_ad"></div>



	<div id="direct_ad_bottom" onclick="clickLog('from=pc_list_shouji_bottom_zhineng1');" style="" ></div>
	<div id="bottom_zhaoshangjiameng_ad" onclick="clickLog('from=pc_list_shouji_bottom_tuiguang2');" style="" ></div>

		<div id="wangmeng_list_bottom_ditong" onclick="clickLog('from=pc_list_shouji_bottom_tuiguang3');"   ></div>        </div>    

    </section>
            <div id="rightframe" class="cright">
	
<div style="margin-bottom:10px">
	<a onclick="pc_list_xianxingpeifu_shouji" href='//weiquan.58.com/complaint/search/?businessline=1&pageid=2' target="_blank">
		<img src="//img.58cdn.com.cn/ds/other/bannerz_03.png">
	</a>
</div>
<div style="margin-bottom:10px">
	
		<img width="100%" src="//img.58cdn.com.cn/ds/other/wxts_list.png">
	
</div>



		




	
	<div class="tuijian" id="ylad" style="display:none"></div>

				<div id="brand_tcrm"></div>
		<script>
		$("#brand_md").bind("bindBrandEvent",function(){ 
            if($('.brand_md200').length > 0){
				$('.brand_md_tit').show();
			}
        }); 
	</script>
					<div id="developmentcontent"></div>
	</div>    </div>


	<section id="links">
<i class="linksTopline"></i>

		<dl class="linksItem relatelink" id="relatelink">
	<dt class="b-left"><h2>北京二手手机品牌导航</h2></dt>
	<dd class="b-right">
	<dl id="relateSelect" class="relateSelect">
 	 	<dt><span>HTC</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手htc evo 4g/" target="_blank">武汉二手htc evo 4g</a>
 		 		<a href="//wh.58.com/shouji/jh_二手g11 htc/" target="_blank">武汉二手g11 htc</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc evo 3d/" target="_blank">武汉二手htc evo 3d</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc sync/" target="_blank">武汉二手htc sync</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc g16/" target="_blank">武汉二手htc g16</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc s710e/" target="_blank">武汉二手htc s710e</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc evo design 4g/" target="_blank">武汉二手htc evo design 4g</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc g2/" target="_blank">武汉二手htc g2</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc desire s/" target="_blank">武汉二手htc desire s</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc chacha/" target="_blank">武汉二手htc chacha</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc sensation/" target="_blank">武汉二手htc sensation</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc s510e/" target="_blank">武汉二手htc s510e</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc sense/" target="_blank">武汉二手htc sense</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc desire hd/" target="_blank">武汉二手htc desire hd</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc desire/" target="_blank">武汉二手htc desire</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc rezound/" target="_blank">武汉二手htc rezound</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc g1/" target="_blank">武汉二手htc g1</a>
 		 		<a href="//wh.58.com/shouji/jh_二手htc t9199/" target="_blank">武汉二手htc t9199</a>
 			</dd>
 	 	<dt><span>LG</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手LG P970/" target="_blank">武汉二手LG P970</a>
 		 		<a href="//wh.58.com/shouji/jh_二手LG LU6200/" target="_blank">武汉二手LG LU6200</a>
 		 		<a href="//wh.58.com/shouji/jh_二手LG S310/" target="_blank">武汉二手LG S310</a>
 		 		<a href="//wh.58.com/shouji/jh_二手LG P990/" target="_blank">武汉二手LG P990</a>
 		 		<a href="//wh.58.com/shouji/jh_二手lg g2/" target="_blank">武汉二手lg g2</a>
 		 		<a href="//wh.58.com/shouji/jh_二手lg ls970/" target="_blank">武汉二手lg ls970</a>
 		 		<a href="//wh.58.com/shouji/jh_二手lg vs920/" target="_blank">武汉二手lg vs920</a>
 		 		<a href="//wh.58.com/shouji/jh_二手lg e970/" target="_blank">武汉二手lg e970</a>
 		 		<a href="//wh.58.com/shouji/jh_二手lg f160/" target="_blank">武汉二手lg f160</a>
 		 		<a href="//wh.58.com/shouji/jh_二手lg手机/" target="_blank">武汉二手lg手机</a>
 			</dd>
 	 	<dt><span>OPPO</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手OPPO A520 /" target="_blank">武汉二手OPPO A520 </a>
 		 		<a href="//wh.58.com/shouji/jh_二手OPPO A100/" target="_blank">武汉二手OPPO A100</a>
 		 		<a href="//wh.58.com/shouji/jh_二手OPPO U525/" target="_blank">武汉二手OPPO U525</a>
 		 		<a href="//wh.58.com/shouji/jh_二手OPPO U529 /" target="_blank">武汉二手OPPO U529 </a>
 		 		<a href="//wh.58.com/shouji/jh_二手OPPO A201/" target="_blank">武汉二手OPPO A201</a>
 		 		<a href="//wh.58.com/shouji/jh_二手OPPO A115/" target="_blank">武汉二手OPPO A115</a>
 			</dd>
 	 	<dt><span>三星</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手三星5320xm/" target="_blank">武汉二手三星5320xm</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星i509/" target="_blank">武汉二手三星i509</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星i9023/" target="_blank">武汉二手三星i9023</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星9100/" target="_blank">武汉二手三星9100</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星s5570/" target="_blank">武汉二手三星s5570</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星p1000/" target="_blank">武汉二手三星p1000</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星w899/" target="_blank">武汉二手三星w899</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星s5830/" target="_blank">武汉二手三星s5830</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星i9008l/" target="_blank">武汉二手三星i9008l</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星s5660/" target="_blank">武汉二手三星s5660</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星5570/" target="_blank">武汉二手三星5570</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星i579/" target="_blank">武汉二手三星i579</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星i9000/" target="_blank">武汉二手三星i9000</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星i809/" target="_blank">武汉二手三星i809</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星i9100g/" target="_blank">武汉二手三星i9100g</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星5660/" target="_blank">武汉二手三星5660</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星i9250/" target="_blank">武汉二手三星i9250</a>
 		 		<a href="//wh.58.com/shouji/jh_二手三星s5230c/" target="_blank">武汉二手三星s5230c</a>
 			</dd>
 	 	<dt><span>中兴</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手中兴V880/" target="_blank">武汉二手中兴V880</a>
 		 		<a href="//wh.58.com/shouji/jh_二手中兴U960S/" target="_blank">武汉二手中兴U960S</a>
 		 		<a href="//wh.58.com/shouji/jh_二手中兴S302/" target="_blank">武汉二手中兴S302</a>
 		 		<a href="//wh.58.com/shouji/jh_二手中兴S202/" target="_blank">武汉二手中兴S202</a>
 		 		<a href="//wh.58.com/shouji/jh_二手中兴N600/" target="_blank">武汉二手中兴N600</a>
 		 		<a href="//wh.58.com/shouji/jh_二手中兴U830/" target="_blank">武汉二手中兴U830</a>
 		 		<a href="//wh.58.com/shouji/jh_二手中兴N880S/" target="_blank">武汉二手中兴N880S</a>
 			</dd>
 	 	<dt><span>夏普</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手夏普SH7218U/" target="_blank">武汉二手夏普SH7218U</a>
 		 		<a href="//wh.58.com/shouji/jh_二手夏普SH9020C/" target="_blank">武汉二手夏普SH9020C</a>
 		 		<a href="//wh.58.com/shouji/jh_二手夏普SH5010C/" target="_blank">武汉二手夏普SH5010C</a>
 		 		<a href="//wh.58.com/shouji/jh_二手夏普V903SH/" target="_blank">武汉二手夏普V903SH</a>
 		 		<a href="//wh.58.com/shouji/jh_二手夏普SH9110C/" target="_blank">武汉二手夏普SH9110C</a>
 		 		<a href="//wh.58.com/shouji/jh_二手夏普SH7118C/" target="_blank">武汉二手夏普SH7118C</a>
 		 		<a href="//wh.58.com/shouji/jh_二手夏普SH9010C/" target="_blank">武汉二手夏普SH9010C</a>
 			</dd>
 	 	<dt><span>天语</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手天语S810/" target="_blank">武汉二手天语S810</a>
 		 		<a href="//wh.58.com/shouji/jh_二手天语Q10/" target="_blank">武汉二手天语Q10</a>
 		 		<a href="//wh.58.com/shouji/jh_二手天语X90/" target="_blank">武汉二手天语X90</a>
 		 		<a href="//wh.58.com/shouji/jh_二手天语X9/" target="_blank">武汉二手天语X9</a>
 		 		<a href="//wh.58.com/shouji/jh_二手天语M610/" target="_blank">武汉二手天语M610</a>
 		 		<a href="//wh.58.com/shouji/jh_二手天语W700/" target="_blank">武汉二手天语W700</a>
 		 		<a href="//wh.58.com/shouji/jh_二手天语W366/" target="_blank">武汉二手天语W366</a>
 			</dd>
 	 	<dt><span>摩托罗拉</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉xt531/" target="_blank">武汉二手摩托罗拉xt531</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉me865/" target="_blank">武汉二手摩托罗拉me865</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉mb525/" target="_blank">武汉二手摩托罗拉mb525</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉xt319/" target="_blank">武汉二手摩托罗拉xt319</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉me525 defy/" target="_blank">武汉二手摩托罗拉me525 defy</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉defy/" target="_blank">武汉二手摩托罗拉defy</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉mb200/" target="_blank">武汉二手摩托罗拉mb200</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉mt810/" target="_blank">武汉二手摩托罗拉mt810</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉ex128/" target="_blank">武汉二手摩托罗拉ex128</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉atrix 4g/" target="_blank">武汉二手摩托罗拉atrix 4g</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉xt711/" target="_blank">武汉二手摩托罗拉xt711</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉a1200/" target="_blank">武汉二手摩托罗拉a1200</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉mt720/" target="_blank">武汉二手摩托罗拉mt720</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉xt502/" target="_blank">武汉二手摩托罗拉xt502</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉ex212/" target="_blank">武汉二手摩托罗拉ex212</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉e6/" target="_blank">武汉二手摩托罗拉e6</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉ex211/" target="_blank">武汉二手摩托罗拉ex211</a>
 		 		<a href="//wh.58.com/shouji/jh_二手摩托罗拉v3/" target="_blank">武汉二手摩托罗拉v3</a>
 			</dd>
 	 	<dt><span>步步高</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手步步高I518 /" target="_blank">武汉二手步步高I518 </a>
 		 		<a href="//wh.58.com/shouji/jh_二手步步高i531 /" target="_blank">武汉二手步步高i531 </a>
 		 		<a href="//wh.58.com/shouji/jh_二手步步高i508/" target="_blank">武汉二手步步高i508</a>
 		 		<a href="//wh.58.com/shouji/jh_二手步步高V205/" target="_blank">武汉二手步步高V205</a>
 		 		<a href="//wh.58.com/shouji/jh_二手步步高i606 /" target="_blank">武汉二手步步高i606 </a>
 		 		<a href="//wh.58.com/shouji/jh_二手步步高i6/" target="_blank">武汉二手步步高i6</a>
 		 		<a href="//wh.58.com/shouji/jh_二手步步高K203M/" target="_blank">武汉二手步步高K203M</a>
 			</dd>
 	 	<dt><span>索爱</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手索爱mt15i/" target="_blank">武汉二手索爱mt15i</a>
 		 		<a href="//wh.58.com/shouji/jh_二手索爱x8/" target="_blank">武汉二手索爱x8</a>
 		 		<a href="//wh.58.com/shouji/jh_二手索爱st18i/" target="_blank">武汉二手索爱st18i</a>
 		 		<a href="//wh.58.com/shouji/jh_二手索爱wt18i/" target="_blank">武汉二手索爱wt18i</a>
 		 		<a href="//wh.58.com/shouji/jh_二手索爱u1/" target="_blank">武汉二手索爱u1</a>
 		 		<a href="//wh.58.com/shouji/jh_二手索爱u5i/" target="_blank">武汉二手索爱u5i</a>
 		 		<a href="//wh.58.com/shouji/jh_二手索爱e15i/" target="_blank">武汉二手索爱e15i</a>
 			</dd>
 	 	<dt><span>联想</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手联想A500/" target="_blank">武汉二手联想A500</a>
 		 		<a href="//wh.58.com/shouji/jh_二手联想S630/" target="_blank">武汉二手联想S630</a>
 		 		<a href="//wh.58.com/shouji/jh_二手联想37AH0/" target="_blank">武汉二手联想37AH0</a>
 		 		<a href="//wh.58.com/shouji/jh_二手联想S500/" target="_blank">武汉二手联想S500</a>
 		 		<a href="//wh.58.com/shouji/jh_二手联想S800/" target="_blank">武汉二手联想S800</a>
 		 		<a href="//wh.58.com/shouji/jh_二手联想A336/" target="_blank">武汉二手联想A336</a>
 			</dd>
 	 	<dt><span>苹果</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/iphone4/" target="_blank">武汉二手iphone4 </a>
 		 		<a href="//wh.58.com/iphone4s/" target="_blank">武汉二手iphone4s </a>
 		 		<a href="//wh.58.com/iphone3gs/" target="_blank">武汉二手iphone 3gs</a>
 		 		<a href="//wh.58.com/iphone3g/" target="_blank">武汉二手iphone 3g</a>
 		 		<a href="//wh.58.com/iphone4s/" target="_blank">武汉二手苹果4s </a>
 		 		<a href="//wh.58.com/iphonesj/" target="_blank">武汉二手苹果手机</a>
 		 		<a href="//wh.58.com/shouji/jh_二手苹果3/" target="_blank">武汉二手苹果3</a>
 			</dd>
 	 	<dt><span>诺基亚</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚1/" target="_blank">武汉二手诺基亚1</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚c1-02/" target="_blank">武汉二手诺基亚c1-02</a>
 		 		<a href="//wh.58.com/shouji/jh_诺基亚/" target="_blank">武汉诺基亚</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚1100/" target="_blank">武汉二手诺基亚1100</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚5000/" target="_blank">武汉二手诺基亚5000</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚5238/" target="_blank">武汉二手诺基亚5238</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚c5-04/" target="_blank">武汉二手诺基亚c5-04</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚700/" target="_blank">武汉二手诺基亚700</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚702t/" target="_blank">武汉二手诺基亚702t</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚5235/" target="_blank">武汉二手诺基亚5235</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚5800w/" target="_blank">武汉二手诺基亚5800w</a>
 		 		<a href="//wh.58.com/nuojiyac503/" target="_blank">武汉二手诺基亚</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚6120/" target="_blank">武汉二手诺基亚6120</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚3208c/" target="_blank">武汉二手诺基亚3208c</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚5700/" target="_blank">武汉二手诺基亚5700</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚500/" target="_blank">武汉二手诺基亚500</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚5233/" target="_blank">武汉二手诺基亚5233</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚5236/" target="_blank">武汉二手诺基亚5236</a>
 		 		<a href="//wh.58.com/shouji/jh_二手诺基亚c1-01/" target="_blank">武汉二手诺基亚c1-01</a>
 			</dd>
 	 	<dt><span>金立</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手金立V109/" target="_blank">武汉二手金立V109</a>
 		 		<a href="//wh.58.com/shouji/jh_二手金立A350/" target="_blank">武汉二手金立A350</a>
 		 		<a href="//wh.58.com/shouji/jh_二手金立A320/" target="_blank">武汉二手金立A320</a>
 		 		<a href="//wh.58.com/shouji/jh_二手金立GN205/" target="_blank">武汉二手金立GN205</a>
 		 		<a href="//wh.58.com/shouji/jh_二手金立V303/" target="_blank">武汉二手金立V303</a>
 		 		<a href="//wh.58.com/shouji/jh_二手金立M508/" target="_blank">武汉二手金立M508</a>
 		 		<a href="//wh.58.com/shouji/jh_二手金立W100/" target="_blank">武汉二手金立W100</a>
 			</dd>
 	 	<dt><span>飞利浦</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手飞利浦X100 /" target="_blank">武汉二手飞利浦X100 </a>
 		 		<a href="//wh.58.com/shouji/jh_二手飞利浦X513 /" target="_blank">武汉二手飞利浦X513 </a>
 		 		<a href="//wh.58.com/shouji/jh_二手飞利浦X216/" target="_blank">武汉二手飞利浦X216</a>
 		 		<a href="//wh.58.com/shouji/jh_二手飞利浦D613/" target="_blank">武汉二手飞利浦D613</a>
 		 		<a href="//wh.58.com/shouji/jh_二手飞利浦W727 /" target="_blank">武汉二手飞利浦W727 </a>
 		 		<a href="//wh.58.com/shouji/jh_二手飞利浦X518/" target="_blank">武汉二手飞利浦X518</a>
 		 		<a href="//wh.58.com/shouji/jh_二手飞利浦X622/" target="_blank">武汉二手飞利浦X622</a>
 			</dd>
 	 	<dt><span>魅族</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手魅族m8 se/" target="_blank">武汉二手魅族m8 se</a>
 		 		<a href="//wh.58.com/shouji/jh_二手魅族i509/" target="_blank">武汉二手魅族i509</a>
 		 		<a href="//wh.58.com/shouji/jh_二手魅族m8 wifi/" target="_blank">武汉二手魅族m8 wifi</a>
 		 		<a href="//wh.58.com/shouji/jh_二手魅族m9白色/" target="_blank">武汉二手魅族m9白色</a>
 			</dd>
 	 	<dt><span>黑莓</span></dt>
 	<dd>
 		 		<a href="//wh.58.com/shouji/jh_二手黑莓9900/" target="_blank">武汉二手黑莓9900</a>
 		 		<a href="//wh.58.com/shouji/jh_二手黑莓9000/" target="_blank">武汉二手黑莓9000</a>
 		 		<a href="//wh.58.com/shouji/jh_二手黑莓9530/" target="_blank">武汉二手黑莓9530</a>
 		 		<a href="//wh.58.com/shouji/jh_二手黑莓9780/" target="_blank">武汉二手黑莓9780</a>
 		 		<a href="//wh.58.com/shouji/jh_二手黑莓8310/" target="_blank">武汉二手黑莓8310</a>
 		 		<a href="//wh.58.com/shouji/jh_二手黑莓8830/" target="_blank">武汉二手黑莓8830</a>
 		 		<a href="//wh.58.com/shouji/jh_二手黑莓9800/" target="_blank">武汉二手黑莓9800</a>
 			</dd>
 		</dl>
	</dd>
	</dl>
	
<dl class="linksItem friendtips" id="friendtips">
<dt class="b-left"><h2>北京二手手机转让频道介绍</h2></dt>
<dd class="b-right"><p>北京二手手机频道为您提供北京二手手机转让/出售信息，在此有大量北京二手手机报价/图片/价格信息供您选择，您可以免费查看北京<a href=//www.58.com/shoujijiagelist/ target=_blank>手机价格大全</a>并发布北京二手手机信息。 <a href=//m.58.com/bj/shouji/ target=_blank>北京二手手机</a>触屏版,<a href=//wap.58.com/bj/shouji/ target=_blank>北京二手手机</a>极速版。</p></dd>
</dl>

<dl class="linksItem hotcity" id="hotcity">
<dt class="b-left"><h2>北京二手手机转让相关城市</h2></dt>
<dd class="b-right">
<ul class="b-ul">
<li><a href='//sh.58.com/shouji/'>上海二手手机</a></li>
<li><a href='//bj.58.com/shouji/'>北京二手手机</a></li>
<li><a href='//sz.58.com/shouji/'>深圳二手手机</a></li>
<li><a href='//gz.58.com/shouji/'>广州二手手机</a></li>
<li><a href='//tj.58.com/shouji/'>天津二手手机</a></li>
<li><a href='//nj.58.com/shouji/'>南京二手手机</a></li>
<li><a href='//dl.58.com/shouji/'>大连二手手机</a></li>
<li><a href='//hz.58.com/shouji/'>杭州二手手机</a></li>
<li><a href='//sy.58.com/shouji/'>沈阳二手手机</a></li>
<li><a href='//hrb.58.com/shouji/'>哈尔滨二手手机</a></li>
<li><a href='//cd.58.com/shouji/'>成都二手手机</a></li>
<li><a href='//dg.58.com/shouji/'>东莞二手手机</a></li>
<li><a href='//jn.58.com/shouji/'>济南二手手机</a></li>
<li><a href='//fs.58.com/shouji/'>佛山二手手机</a></li>
<li><a href='//wx.58.com/shouji/'>无锡二手手机</a></li>
<li><a href='//cs.58.com/shouji/'>长沙二手手机</a></li>
<li><a href='//wh.58.com/shouji/'>武汉二手手机</a></li>
<li><a href='//nb.58.com/shouji/'>宁波二手手机</a></li>
<li><a href='//cc.58.com/shouji/'>长春二手手机</a></li>
<li><a href='//su.58.com/shouji/'>苏州二手手机</a></li>
<li><a href='//qd.58.com/shouji/'>青岛二手手机</a></li>
<li><a href='//cq.58.com/shouji/'>重庆二手手机</a></li>
<li><a href='//fz.58.com/shouji/'>福州二手手机</a></li>
</ul>
</dd>
</dl>
<dl class="linksItem hotcity" id="hotcity">
<dt class="b-left"><h2>北京二手手机相关分类</h2></dt>
<dd class="b-right">
<ul class="b-ul">
<li>
<a href="//bj.58.com/shouji/">北京二手手机</a>
</li>
<li>
<a href="//bj.58.com/diannao/">北京二手电脑</a>
</li>
<li>
<a href="//bj.58.com/shuma/">北京二手数码产品</a>
</li>
<li>
<a href="//bj.58.com/jiadian/">北京二手家电</a>
</li>
<li>
<a href="//bj.58.com/ershoujiaju/">北京二手家具</a>
</li>
<li>
<a href="//bj.58.com/shoujihao/">
北京手机号码转让</a>
</li>
<li>
<a href="//bj.58.com/tongxunyw/">北京通讯业务</a>
</li>
<li>
<a href="//bj.58.com/yingyou/">北京母婴玩具转让</a>
</li>
<li>
<a href="//bj.58.com/fushi/">北京服装箱包转让</a>
</li>
<li>
<a href="//bj.58.com/meirong/">北京美容化妆品转让</a>
</li>
<li>
<a href="//bj.58.com/tushu/">
北京图书音像转让</a>
</li>
<li>
<a href="//bj.58.com/bangong/">北京办公设备转让</a>
</li>
<li>
<a href="//bj.58.com/ershoushebei/">北京二手设备</a>
</li>
<li>
<a href="//bj.58.com/yishu/">北京艺术收藏品转让</a>
</li>
<li>
<a href="//bj.58.com/wenti/">北京户外体育用品转让</a>
</li>
<li>
<a href="//bj.58.com/danche/">北京二手摩托车</a>
</li>
<li>
<a href="//bj.58.com/diandongche/">北京二手电动车</a>
</li>
</ul>
</dd>
</dl>

<dl class="linksItem hotcity" id="hotcity">
<dt class="b-left">
<h2>
北京二手手机周边城市
</h2>
</dt>
<dd class="b-right">
<ul class="b-ul">
<li>
<a target="_blank" href="//tj.58.com/shouji/">天津二手手机</a>
</li>
<li>
<a target="_blank" href="//sjz.58.com/shouji/">石家庄二手手机</a>
</li>
<li>
<a target="_blank" href="//zjk.58.com/shouji/">张家口二手手机</a>
</li>
<li>
<a target="_blank" href="//bd.58.com/shouji/">保定二手手机</a>
</li>
<li>
<a target="_blank" href="//lf.58.com/shouji/">廊坊二手手机</a>
</li>
<li>
<a target="_blank" href="//ts.58.com/shouji/">唐山二手手机</a>
</li>
<li>
<a target="_blank" href="//chengde.58.com/shouji/">承德二手手机</a>
</li>
<li>
<a target="_blank" href="//qhd.58.com/shouji/">秦皇岛二手手机</a>
</li>
<li>
<a target="_blank" href="//cangzhou.58.com/shouji/">沧州二手手机</a>
</li>
<li>
<a target="_blank" href="//hs.58.com/shouji/">衡水二手手机</a>
</li>
</ul>
</dd>
</dl>

<dl class="linksItem friendlink" id="friendlink">
<dt class="b-left">
<h2>友情链接</h2></dt>
<dd class="b-right"><a href='//bj.58.com/shouji/' target='_blank'>北京二手手机</a>&nbsp;<a href='//t.58.com/bj' target='_blank'>北京团购</a>&nbsp;<a href='http://www.158300.com.cn/' target='_blank'>密山</a>&nbsp;<a href='http://3c.chengw.com/' target='_blank'>重庆家电网</a>&nbsp;<a href='//www.58.com/changhongsjzmy/' target='_blank'>长虹手机怎么样</a>&nbsp;<a href='//www.58.com/shouji/' target='_blank'>二手</a>&nbsp;<a href='http://www.sou300.com/' target='_blank'>搜三百</a>&nbsp;<a href='//www.58.com/1000shoujituijian/' target='_blank'>1000左右的手机推荐</a>&nbsp;<a href='//ku.m.58.com/sjdq/' target='_blank'>手机品牌大全</a>&nbsp;<a href='//pinpai.58.com/shouji/' target='_blank'>手机品牌大全</a>&nbsp;<a href='http://beitang.xbaixing.com/' target='_blank'>北塘吧</a>&nbsp;<a href='http://xc.58fenlei.com/shouji/' target='_blank'>许昌二手手机转让</a>&nbsp;<a href='http://nj.fenlei168.com/' target='_blank'>南京生活信息</a>&nbsp;<a href='//www.58.com/shouji/' target='_blank'>二手手机交易网</a>&nbsp;<a href='//www.58.com/shoujidaquan/' target='_blank'>手机报价大全</a>&nbsp;<a href='http://m.zhuanzhuan.com/' target='_blank'>转转</a>&nbsp;<a href='http://product.cnmo.com/' target='_blank'>手机大全</a>&nbsp;<a href='http://bj.ganji.com/shouji/' target='_blank'>北京二手手机</a>&nbsp;<a href='//www.58.com/pplbgypmap/' target='_blank'>办公用品</a>&nbsp;<a href='//www.58.com/pplbgypz172/' target='_blank'>办公用品</a>&nbsp;<a href='http://www.zhuanzhuan.com/contact.html' target='_blank'>转转客服</a>&nbsp;<a href='//tj.58.com/yuanzhuangshengang/' target='_blank'>天津二手原装神钢工程机械</a>&nbsp;<a href='http://www.aihuishou.com' target='_blank'>手机回收</a>&nbsp;<a href='//youpin.m.58.com/' target='_blank'>转转优品</a>&nbsp;<a href='//zhuanzhuan.58.com/product/pc/' target='_blank'>台式机</a>&nbsp;<a href='/link.html' target='_blank'>更多...</a></dd>
</dl>

</section>


<footer id="footer" class="footer"><p class="copyright">&copy; 58.com <a id="askicon" target="_blank" title="有问题请与58客服进行对话" rel="nofollow" href="//about.58.com/v5/">帮助中心</a></p></footer>
	<footer id="footer" class="footer">
	    <div id="upWrap">
			<a target="_blank" href="//about.58.com/help.html" rel="nofollow">常见问题</a><span>|</span>
			<a target="_blank" href="//about.58.com/" rel="nofollow">帮助中心</a><span>|</span>
			<a target="_blank" href="//about.58.com/feedback.html" rel="nofollow">意见反馈</a><span>|</span>
			<a target="_blank" href="//about.58.com/home/" rel="nofollow">了解58同城</a><span>|</span>
			<a target="_blank" href="//about.58.com/hr/" rel="nofollow">加入58同城</a><span>|</span>
			<a target="_blank" href="//fanqizha.58.com/" rel="nofollow">反欺诈联盟</a><span>|</span>
			<a target="_blank" href="//110.58.com" rel="nofollow">报案平台</a><span>|</span>
			<a target="_blank" href="//e.58.com" rel="nofollow">推广服务</a><span>|</span>
			<a target="_blank" href="//biz.58.com" rel="nofollow">渠道招商</a><span>|</span>
			<a target="_blank" href="//baozhang.58.com" rel="nofollow">先行赔付</a><span>|</span>
			<a target="_blank" href="//ir.58.com?PGTID=0d100000-0000-13da-8041-51f9a3a739fa&amp;ClickID=2" rel="nofollow">Investor Relations</a>
		</div>
	    <div id="downWrap">
			<em>2005-2035 58.com版权所有</em><span>|</span>
			<em>京公网备案信息110105000809</em><span>|</span>
			<a target="_blank" href="http://www.miibeian.gov.cn/" rel="nofollow">京ICP证060405</a><span>|</span>
			<em>乙测资字11018014</em><span>|</span>
			<a target="_blank" href="//ir.58.com" rel="nofollow">Investor Relations</a><span>|</span>
			<em>违法信息举报：4008135858&nbsp;&nbsp;jubao@58.com</em>
		</div>
	    <div class="fotBtmIcon">
	        <a target="_blank" id="fotBm_1" href="//fanqizha.58.com"></a>
	        <a target="_blank" id="fotBm_2" href="http://www.12377.cn/"></a>
	        <a target="_blank" id="fotBm_3" href="http://www.12377.cn/node_548446.htm"></a>
	        <a target="_blank" id="fotBm_4" href="https://credit.szfw.org/CX20120918001650001720.html"></a>
	        <a target="_blank" id="fotBm_5" href="//img.58cdn.com.cn/ui6/index/qyxinyong.jpg?v=3"></a>
	        <a target="_blank" id="fotBm_6" href="//about.58.com/fqz/index.html"></a>
	    </div>
	</footer>


<script type="text/javascript" src='//j1.58cdn.com.cn/js/5_1/comm_js/boot_listpage_zuche_v20200602183328.js'></script>

<script type="text/javascript" src='//j1.58cdn.com.cn/crop/biz/base/js/avoidCheat_pc_list_v20171106151755.js'></script>

<script type="text/javascript" src='//j1.58cdn.com.cn/componentsLoader/dist/ComponentsLoader_v20201229143446.js'></script>
<script type="text/javascript">
    document.domain = '58.com';
    var win = new GetToolTipWindow('tooltipdiv', 'keyword', '3','1','', true);win.setSearchButton('searchbtn');
    var win1 = new GetToolTipWindow('tooltipdiv1', 'keyword1', '2','1','');win1.setSearchButton('searchbtn1');
    var win2 = new GetToolTipWindow('tooltipdiv2', 'mudidi', '7', '1', '');win2.setSearchButton('searchbtn1');
    boot.require('util.cookie', function(Frame, cookie){cookie.set('city', 'bj');});
</script>

<script type="text/javascript">boot.require('huangye.topbar', function(Frame, topbar){topbar.init();});</script>

<script type="text/javascript" >
    // JavaScript Document
    function toggle(classname){
        $(classname).hover(function() {$(this).addClass("hover");}, function() {$(this).removeClass("hover");});
    }

    $(document).ready(function(){
        toggle(".haschild");
        toggle(".sec-time span");
        toggle(".prifilter span");
        toggle(".header-search input");
        toggle(".filterRange");
        toggle(".header-search input");
        toggle(".p-help");
        $(".tooltip li").hover(function() {$(this).addClass("selected");}, function() {$(this).removeClass("selected");});
        $(".tbimg tr").hover(function() {$(this).css('background','#f5f5f5');}, function() {$(this).css('background','#fff');});
        $(".small-tbimg tr").hover(function() {$(this).css('background','#f5f5f5');}, function() {$(this).css('background','#fff');});
    });

    $("#shangmenfw a").each(function(){
        if("到店服务"== $(this).text().trim()){
            $(this).remove();
        }
    });

    BAIDU_CLB_fillSlot("567863");

    var alldt = $("#relateSelect>dt");
    function showNear(dt){
        alldt.each(function(){
            if(this == dt) {$(this).next("dd").show();$(this).addClass("select");}
            else {$(this).next("dd").hide();$(this).removeClass("select");}
        });
    }
    alldt.each(function(){
        $(this).click(function(){showNear(this);}).hover(function(){showNear(this);});
    });

                showNear(alldt[0]);
    
    // 
    //信息hover
    // boot.require('dom, anim.hover', function(Frame, dom, hover){
    //     var cc = dom.get('infolist'),
    //             trs = cc.getElementsByTagName('tr');
    //     Frame.each(trs, function(item){
    //         hover.bind(item, 'bg');
    //     });
    // });

    $(".selector").each(function(){
        $(this).hover(function(){$(this).addClass("hover")},function(){$(this).removeClass("hover")});
    });

</script>

<div style="display:none;">
    <script type="text/javascript">
        var _bdhmProtocol = (("https:" == document.location.protocol) ? " https://" : " http://");
                document.write(unescape("%3Cscript src='" + _bdhmProtocol + "hm.baidu.com/h.js%3F3bb04d7a4ca3846dcc66a99c3e861511' type='text/javascript'%3E%3C/script%3E"));
            </script>
</div>






<!-- 
$pageModel.show("google_js")
 -->

<script>
    var currentDate = new Date();
</script>


<script src="//j1.58cdn.com.cn/lbg/58.com/ui8/modules/js/huangye/common_foot-5ec7b9ca47.min.js"></script>


<script>
            var jingzhun_ads = [];
            var jingzhun_ads_param = "";
    </script>
<script type="text/javascript">
    $("#more_ObjectType").click(function(){
    	console.log("ok");
    	var showmoreBrand = document.getElementById("moreBrandct"),moreBrandbtn = document.getElementById("moreBrandbtn");
	    showmoreBrand.style.display = (showmoreBrand.style.display=="block")?"none":"block";
	    moreBrandbtn.className = (moreBrandbtn.className=="moreBrand")?"moreBrand hold":"moreBrand";
    });
    var pemorealldt = $("#moreBrandct dt");
	function showpemorecon(dt){
		pemorealldt.each(function(){
		    if(this == dt) {$(this).next("dd").show();$(this).addClass("select");}
            else {$(this).next("dd").hide();$(this).removeClass("select");}
        });
    }
    pemorealldt.each(function(){
		$(this).click(function(){showpemorecon(this);}).hover(function(){showpemorecon(this);});
    });
    showpemorecon(pemorealldt[0]);
</script>
<script type="text/javascript">
function cpcget(url){
    if(url != '') {
        $.ajax({
            url: url,
            type: "get",
            dataType: "jsonp",
            success: function (data) {

            }
        });
    }
}
</script>
    <script>
    var A2B_SSP = "A";
    var sspDate = new Date();
    var sspDateNumber = String(sspDate.getFullYear()) + String((sspDate.getMonth()+1)) + String(sspDate.getDate()) + String(sspDate.getHours());
    document.write("<scr"+"ipt src='//adshow.58.com/js/ssp_init.js?r=" + sspDateNumber + "'></sc"+"ript>");
    </script>
	    <script>
    $(function(){
        var adData = {1000106:[],1000019:[],1000114:[],1000001:[],1002464:[]} || "";
        if(adData){
			getAdurl(1000616);
			getAdurl(1000617);
			getAdurl(1000620);
			getAdurl(1000621);
			getAdurl(1000583);

			getAdurl(1000640);
			getAdurl(1000641);
			getAdurl(1000586);

			getAdurl(1000663);
			getAdurl(1000664);

			getAdurl(1000698);
			getAdurl(1000699);
			getAdurl(1000695);

        };
        function getAdurl(num){
            if(adData[num] && adData[num].length!=0){
                for(var i=0; i<adData[num].length; i++ ){
                    $.get(adData[num][i]);
                }
            }
        };
    })
    </script>
	
<script>
var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "//hm.baidu.com/hm.js?e15962162366a86a6229038443847be7";
  var s = document.getElementsByTagName("script")[0];
  s.parentNode.insertBefore(hm, s);
})();
</script>

<div style="display:none">
    <script type="text/javascript">
    	    	    ____json4fe.supplycount = {level1:'', level2:'', level3:'', sortid:'',results:"no_supply"};
    	    var _gaq = _gaq || [],site_name = "58",post_count = "5750",page_type = "list";
    <!-- shouji -->
                ____json4fe.link_abtest = 'PCFlowSupport_B';
	var _trackURL = "{'cate':'5,36','area':'1','version':'','pagetype':'list','page':'shouji','GA_pageview':'/bj/sale/shouji/listversion/','pagesize':'35','pagenum':'1','post_count':'5750','jiage':'','abVersion':'PCFlowSupport_B', 'filterlocal':'','filtercate':'','nodeType':'info','tabFrom':'info','tabName':'全部'}";
    var _setmy = "mmmm";
        _gaq.push(['pageTracker._setAccount', 'UA-877409-4']);
        _gaq.push(['pageTracker._setDomainName', '.58.com']);
        _gaq.push(['pageTracker._setAllowLinker', true]);
        _gaq.push(['pageTracker._addOrganic', 'sogou', 'query']);
        _gaq.push(['pageTracker._addOrganic', 'baidu', 'word']);
        _gaq.push(['pageTracker._addOrganic', 'soso', 'w']);
        _gaq.push(['pageTracker._addOrganic', 'youdao', 'q']);
    	boot.require('business.gaq, util.js', function (Frame, gaq, js) {   
        _gaq.push(['pageTracker._trackPageview', gaq("sale", "/bj/sale/shouji/listversion/")]);
        _gaq.push(['pageTracker._trackPageLoadTime']);
		 Frame.init(function () {
            js.append('//tracklog.58.com/referrer4.js');
        });
    });
    </script>
    </div>









<!-- 抢单神器下线-->
<!-- 抢单神器下线-->
<!-- 装修搜索弹窗  S -->
<!-- 招标弹框start -->
<link rel="stylesheet" href='//c.58cdn.com.cn/ds/zhaobiao/pc_search_win/list_win_zb_v20200603095131.css' />
<script src="//j1.58cdn.com.cn/ds/zhaobiao/v1/js/fastpost/jquery.md5.js"></script>
<script>
 var cityid=1;
 if($('.top1guangg').length){
   window.clickLog && clickLog('from=lbg_list_adv_show');
 }
</script>

    
    <!-- 招标弹框end -->		<!-- 装修搜索弹窗  E -->

<!-- <script type="text/javascript" src='//j1.58cdn.com.cn/crop/ecom/m/tcb/ideaShow/main_v20180917142647.js'></script> -->
</body>
</html>`
	//resp, err := http.Get("https://bj.58.com/shouji/?from=ganji")
	//HandleError(err,"http.get")
	//defer resp.Body.Close()
	//
	//bytes, err := ioutil.ReadAll(resp.Body)
	//HandleError(err,"ioUtils.Read")
	//println(string(bytes))

	compile := regexp.MustCompile(`<div class="new-long-tit new-long-tit2">
				(.*)
			</div>`)
	findAllSubmatch := compile.FindAllStringSubmatch(testVarrr, 10)
	for _, str := range findAllSubmatch {
		fmt.Println(str[1])
	}
	fmt.Println("==================")

	compile2 := regexp.MustCompile(`<b class='pri'>(.*)</b>`)
	submatch2 := compile2.FindAllStringSubmatch(testVarrr, 10)
	for _, str := range submatch2 {
		fmt.Println(str[1])
	}

}

func TestPaCh() {
	strtest := `djakslfjal666uiosusioaiuoudjakslfjal999uiosusioadhaudadjakslfjal606uiosusioa`
	compile := regexp.MustCompile(`djakslfjal(.*)uiosusioa`)
	submatch := compile.FindAllStringSubmatch(strtest, -1)
	fmt.Println(submatch)
}
