
function JSONToCSVConvertor(JSONData, ReportTitle, ShowLabel) {
    var arrData = typeof JSONData != 'object' ? JSON.parse(JSONData) : JSONData;
    var fileName = "";
    fileName += ReportTitle.replace(/ /g,"") + "";
    var userAgent = navigator.userAgent; //取得浏览器的userAgent字符串
    var isOpera = userAgent.indexOf("Opera") > -1; //判断是否Opera浏览器
    var isIE = userAgent.indexOf("compatible") > -1 && userAgent.indexOf("MSIE") > -1 && !isOpera; //判断是否IE浏览器
    var CSV = 'data:text/csv;charset=utf-8,\ufeff';
    if(isIE){
        CSV = '\ufeff';
    }
    //CSV += ReportTitle + '\r\n\n';
    if (ShowLabel) {
        var row = "";
        for (var index in arrData[0]) {
			switch (index){
				case 'HiDst':
					row += "当前进程全路径" + ',';
					break;
				case 'HiItn':
					row += "窗口标题" + ',';
					break;
				case 'HiCle':
					row += "当前进程子进程参数" + ',';
					break;
				case 'HiMD5PAR':
					row += "父进程MD5" + ',';
					break;
				case 'HiOrn':
					row += "当前进程名" + ',';
					break;
				case 'HiScl':
					row += "当前进程启动参数" + ',';
					break;
				case 'HiUzp':
					row += index + ',';
					break;
				case '_row':
					row += "查询标识" + ',';
					break;
				case 'clientip':
					row += "客户端IP" + ',';
					break;
				case 'esid':
					row += "客户端唯一识别号" + ',';
					break;
				case 'filelevel':
					row += "文件等级" + ',';
					break;
				case 'filepath':
					row += "文件路径" + ',';
					break;
				case 'hisrc':
					row += "父进程绝对路径 " + ',';
					break;
				case 'md5':
					row += index + ',';
					break;
				case 'time':
					row += "时间" + ',';
					break;
				case 'domain':
					row += "域名" + ',';
					break;
				case 'ip':
					row += index + ',';
					break;
				case 'ndtype':
					row += "ND类型" + ',';
					break;
				case 'rp':
					row += index + ',';
					break;
				case 'url':
					row += "URL地址" + ',';
					break;
				case '_ts':
					row += index + ',';
					break;
				case 'count':
					row += "访问次数" + ',';
					break;
				case 'host':
					row += "Host的URL"+ ',';
					break;
				case 'first_seen':
					row += "第一次访问时间" + ',';
					break;
				case 'last_seen':
					row += "最后一次访问时间" + ',';
					break;
				case 'date':
					row += "访问时间" + ',';
					break;
				default:
					row += index + ',';
			}
        }
        row = row.slice(0, -1);
        CSV += row + '\r\n';
    }
    for (var i = 0; i < arrData.length; i++) {
        var row = "";
        for (var index in arrData[i]) {
            row += '"' + unescape(arrData[i][index]) + '",';
        }
        row.slice(0, row.length - 1);
        CSV += row + '\r\n';
    }
    if (CSV == '') {
        alert("Invalid data");
        return;
    }
    //IE只能使用Blob文件下载，不支持URI，详见:
    //https://technet.microsoft.com/ZH-CN/Library/hh779016.aspx
    if(isIE){
        var blob = new Blob([decodeURIComponent(encodeURI(CSV))], {
            type: "text/csv;charset=utf-8;"
        });
        window.navigator.msSaveBlob(blob, fileName + ".csv")
    }
    else{
        var link = document.createElement("a");
        link.href = encodeURI(CSV);
        link.style = "visibility:hidden";
        link.download = fileName + ".csv";
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }
}