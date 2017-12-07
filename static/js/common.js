//json格式字符串判断
//function isJSON(str, pass_object) {
//	if(pass_object && isObject(str)) return true;
//
//	if(!isString(str)) return false;
//
//	str = str.replace(/\s/g, '').replace(/\n|\r/, '');
//
//	if(/^\{(.*?)\}$/.test(str))
//		return /"(.*?)":(.*?)/g.test(str);
//
//	if(/^\[(.*?)\]$/.test(str)) {
//		return str.replace(/^\[/, '')
//			.replace(/\]$/, '')
//			.replace(/},{/g, '}\n{')
//			.split(/\n/)
//			.map(function(s) {
//				return isJSON(s);
//			})
//			.reduce(function(prev, curr) {
//				return !!curr;
//			});
//	}
//
//	return false;
//}
function isJSON(str) {
    if (typeof str == 'string') {
        try {
            JSON.parse(str);
            return true;
        } catch(e) {
            console.log(e);
            return false;
        }
    }
    console.log('It is not a string!')    
}
//获取url中的参数
function getUrlParam(name) {
	var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
	var r = window.location.search.substr(1).match(reg); //匹配目标参数
	if(r != null) return unescape(r[2]);
	return null; //返回参数值
}
//设置当前页按钮样式
function setCurPageStyle(name) {
	if(getUrlParam(name) > 1) {
		//获取url参数
		var pVal = getUrlParam(name);
		//				console.log(pVal);
		//分页按钮
		var aBtns = $('.pagination').find('a.ui.button');
		for(var i = 0; i < aBtns.length; i++) {
			if(aBtns[i].innerText == pVal) {
				aBtns[i].className += " gcur";
			}
		}
	} else if((!getUrlParam(name)) || getUrlParam(name) == 1) {
		var aBtns = $('.pagination').find('a.ui.button');
		for(var i = 0; i < aBtns.length; i++) {
			if(aBtns[i].innerText == 1) {
				aBtns[i].className += " gcur";
			}
		}
	}
}
//添加左侧菜单样式表示当前是哪个页
function setCurPage(UrlDoms) {
	var domainUrl = [];
	for(var i = 0; i < UrlDoms.length; i++) {
		domainUrl.push($(UrlDoms[i]).attr("href").split("/cmdb/")[1]);
	}
	console.log(domainUrl);
	//获取当前页名称

	$(".gcitem").each(function(i) {
		if(window.location.href.toLowerCase().split('cmdb/')[1] == domainUrl[i]) {
			console.log(1);
			$(this).addClass("gcuri");
		}
	});
}
// Cookie functions
function createCookie(name, value, days) {
	if(days) {
		var date = new Date();
		date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
		var expires = "; expires=" + date.toGMTString();
	} else var expires = "";
	document.cookie = name + "=" + value + expires + "; path=/";
}

function readCookie(name) {
	var nameEQ = name + "=";
	var ca = document.cookie.split(';');
	for(var i = 0; i < ca.length; i++) {
		var c = ca[i];
		while(c.charAt(0) == ' ') c = c.substring(1, c.length);
		if(c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
	}
	return null;
}

function eraseCookie(name) {
	createCookie(name, "", -1);
}
//
//function setMessage(txt, type) {
//	var msgDom = $('<div id="s-msg" class="alert" role="alert"></div>');
//	msgDom.html('<span class="s-text">' + txt + '</span>');
//	//	msgDom.html('<span class="s_close">x</span><span class="s-text">' + txt + '</span>');
//	switch(type) {
//		case 0:
//			msgDom.addClass("alert-danger");
//			break;
//		case 1:
//			msgDom.addClass('alert-warning');
//			break;
//		case 2:
//			msgDom.addClass('alert-success');
//			break;
//		default:
//			msgDom.addClass('alert-info');
//	}
//	$('body').append(msgDom);
//	showMessage();
//	setTimeout(hideMessage, 3000);
//}
function setMessage(txt, type) {
	var msgDom = $('<div id="s-msg" class="alert" role="alert"></div>');
	msgDom.html('<span class="s-text">' + txt + '</span>');
	//	msgDom.html('<span class="s_close">x</span><span class="s_text">' + txt + '</span>');
	switch(type) {
		case 0:
			//成功,绿色
			msgDom.addClass('alert-success');
			break;
		case 1:
			//补充,蓝色改白色样式
			msgDom.addClass('alert-info');
			break;
		case 2:
			//警告，黄色
			msgDom.addClass('alert-warning');
			break;
		case 3:
			//错误，红色
			msgDom.addClass("alert-danger");
			break;
	}
	$('body').append(msgDom);
	showMessage();
	setTimeout(hideMessage, 3000);
}
//错误,红色
function setError(txt) {
	setMessage(txt, 3);
}
//警告,黄色
function setWarning(txt) {
	setMessage(txt, 2);
}
//成功,绿色
function setInfo(txt) {
	setMessage(txt, 0);
}
//补充，白色
function setMsg(txt) {
	setMessage(txt, 1);
}

function showMessage() {
	$('#s-msg').show();
}

function hideMessage() {
	$('#s-msg').hide();
}