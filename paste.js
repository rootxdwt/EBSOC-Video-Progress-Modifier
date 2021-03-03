//paste this code into the ebs console
    function getCookie(cookieName){
        var cookieValue=null;
        if(document.cookie){
    	  var array=document.cookie.split((escape(cookieName)+'='));
    	  if(array.length >= 2){
    	    var arraySub=array[1].split(';');
    	    cookieValue=unescape(arraySub[0]);
    	  }
        }
        return cookieValue;
    }
function getParam(sname) {
            var params = location.search.substr(location.search.indexOf("?") + 1);
            var sval = "";

            params = params.split("&");

            for (var i = 0; i < params.length; i++) {
                temp = params[i].split("=");
                if ([temp[0]] == sname) { sval = temp[1]; }
            }
            if( sval ) {
                return decodeURIComponent(sval);
            } else {
                return sval;
            }
        }

`${getCookie("memberSeq")}|${getCookie("access")}|${getParam("lctreLrnSqno")}|${"100"}`

//it will show you a long string. paste it into the program.
//if you want to use a different progress, modify the "100" to a number you want
