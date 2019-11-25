$('.show_userinfo').html('<a href=\"javascript:void(0);\" onclick=\"AjaxLog()\" target=\"_self\">登录</a> &nbsp;|&nbsp; <a href=\"/e/member/register/index.php?groupid=1&button=下一步\">新用户注册</a>');
var WebUrl="/";
$("#login-sub").live('click',function(){
	$('#errmsg').hide();
	var user=$("#logusername").val();   //获取用户名输入框的值
	var pass=$("#logpassword").val();   //获取密码输入框的值
	if(document.getElementById("keeplogin").checked==true){
	var lifetime=315360000;
	}
	if(user==""){
		$('#errmsg').show().html("用户名不能为空！");
		$("#username").focus();
		return false
	};
	if(user=="输入用户名"){
		$('#errmsg').show().html("用户名不能为空！");
		$("#username").focus();
		return false
	};
	if(pass==""){
		$('#errmsg').show().html("密码不能为空！");
		$("#password").focus();
		return false
	};
	$.ajax({
		type:"POST",url:WebUrl+"e/member/ajaxlogin/",dataType:"html",data:{
			'username':user,'password':pass,enews:'login','lifetime':lifetime                      //提交字段
		}
		,beforeSend:function(){
			$('#errmsg').show().html('<img src="/images/loading.gif" /> 正在登陆...')
		}
		,success:function(data){
			if(data=='failuserid'){
				$('#errmsg').show().html('用户名错误!')
			}
			else if(data=='failpassword'){
				$('#errmsg').show().html('密码错误!')
			}
			else if(data=='failcheck'){
				$('#errmsg').show().html('你的账号还没有激活，<a href="'+WebUrl+'e/member/register/regsend.php">重发激活邮件！</a>')
			}
			else{
				$("#errmsg").hide();
				$('.show_userinfo').show().html(data);
				$("#lean_overlay").fadeOut(200);
				$("#js_login").hide();
				$('#shortinfo').html('<a  href=/e/space/?userid= target=_blank></a>&nbsp;&nbsp;<a href=javascript:void(0) id=logout onclick=LogOut(1)>退出</a>')
			}
		}
	})
});
function LogOut(flag){
    var aj = "";
    if(flag){
        aj = "&aj=1";
    }
	$.post(WebUrl+"e/member/ajaxlogin/login.php?enews=exit"+aj,function(msg){
		if(msg=='exitsuccess'){
			$(".show_userinfo").html('<a href=\"javascript:void(0);\" onclick=\"AjaxLog()\" target=\"_self\">登录</a> &nbsp;|&nbsp; <a href=\"/e/member/register/index.php?groupid=1&button=下一步\">新用户注册</a>');
			$('#shortinfo').html("<a href=\"javascript:void(0);\" onclick=\"AjaxLog()\" target=\"_self\">登录</a> &nbsp;|&nbsp; <a href=\"/e/member/register/index.php?groupid=1&button=下一步\">新用户注册</a>");
			alert('退出成功!');
		}
	})
}
function AjaxReg(){
  //self.location=WebUrl+'e/member/register/index.php?groupid=1';
  AjaxLog();
  switchTag('_xka','_xka_list',2,2,'conxk1','conxk2');
}
function CheckReg(key,value){
	if(value!=''){
		var data = {},key=key;
		data['enews']='register';
		data[key]=value;
	    if(key=='repassword'){
			var pass=$("#password").val();   //获取密码输入框的值
			data['password']=pass;
		}
		$.ajax({
			type:"POST",url:WebUrl+"e/member/ajaxlogin/",dataType:"html",data:data
			,beforeSend:function(){
				$('.'+key).html('<img src="/images/loading.gif" />')
			}
			,success:function(msg){
				if(!msg){
					$('.'+key).html('<font color=green>√</font>');   //这里的代码已函数checkpass()里面的一段代码相对应！
					checkpass();
				}else{
					$('.'+key).html('<font color=red>'+msg+'</font>');
					checkpass('all');
					return false;
				}
			}
		});

	}else{
		$('.'+key).html('<font color=red>不允许为空！</font>');
	}

}
function checkpass(all){
	var user=$("#username").val();
	var username=$(".username").html();
	var pass=$(".password").html();
	var pass2=$(".repassword").html();
	var email=$(".email").html();
	if(username==pass && pass2==email && pass==pass2 && document.getElementById("checkin").checked==true && username!='' && user!=''){
		$('#Reg-sub').removeAttr('disabled');
		$('#checkreg').show().html('全部填写正确！');
	}else{
		document.getElementById("Reg-sub").disabled = true;
		$('#checkreg').hide();
		if(all){
			$('#checkreg').show().html('<font color=red>有必选项没填写正确！</font>');
		}
	}
}
(function($){
	$.fn.extend({
		leanModal:function(options){
			var defaults={
				top:200,overlay:0.5,closeButton:null
			};
			var overlay=$("<div id='lean_overlay'></div>");
			$("body").append(overlay);
			options=$.extend(defaults,options);
			return this.each(function(){
				var o=options;
				$(this).click(function(e){
					var modal_id=$(this).attr("href");
					$("#lean_overlay").click(function(){
						close_modal(modal_id)
					});
					$(o.closeButton).click(function(){
						close_modal(modal_id)
					});
					var modal_height=$(modal_id).outerHeight();
					var modal_width=$(modal_id).outerWidth();
					$("#lean_overlay").css({
						"display":"block",opacity:0
					});
					$("#lean_overlay").fadeTo(200,o.overlay);
					$(modal_id).css({
						"display":"block","position":"fixed","opacity":0,"z-index":11000,"left":50+"%","margin-left":-(modal_width/2)+"px","top":o.top+"px"
					});
					$(modal_id).fadeTo(200,1);
					e.preventDefault()
				})
			});
			function close_modal(modal_id){
				$("#lean_overlay").fadeOut(200);
				$(modal_id).css({
					"display":"none"
				})
			}
		}
	})
})(jQuery);
function switchTag(tag,content,k,n,stylea,styleb){for(i=1;i<=n;i++){if(i==k){document.getElementById(tag+i).className=stylea;document.getElementById(content+i).className="showbox"}else{document.getElementById(tag+i).className=styleb;document.getElementById(content+i).className="hidden"}}};
/*登录框代码*/
document.writeln("<div id=\"js_login\" class=\"js_login\" style=\"display:none\">");
document.writeln("<div class=\"js_title\">");
document.writeln("<ul>");
document.writeln("<li class=\"conxk1\" id=\"_xka1\" onclick=\"switchTag(\'_xka\',\'_xka_list\',1,2,\'conxk1\',\'conxk2\');\">");
document.writeln("用户登录</li>");

document.writeln("</ul>			<span class=\"close\">×</span>");
document.writeln("</div>");
document.writeln("<div class=\"js_content\">");
document.writeln("<div class=\"con_dak fixed clear\">");
document.writeln("<div id=\"_xka_list1\" class=\"showbox\">");
document.writeln("<div class=\"de_list\">");
document.writeln("<!----------login start------>");
document.writeln("<div  id=\"login\">");
document.writeln("<div class=\"inputbg\" id=\"logtext\">");
document.writeln("<input name=\"logusername\" id=\"logusername\" type=\"text\" onfocus=\"if(this.value==\'输入用户名或邮箱\')value=\'\';\" value=\"输入用户名或邮箱\" placeholder=\"输入用户名或邮箱\" class=\"inputsub\" />");
document.writeln("</div>");
document.writeln("<div class=\"inputbg\" id=\"pass\">");
document.writeln("<input name=\"logpassword\" placeholder=\"输入密码\" type=\"password\"  value=\"\" id=\"logpassword\" class=\"inputsub\" />");
document.writeln("</div>");
document.writeln("<div class=\"inputbg\" id=\"checkbox\" style=\"border:none;height:10px\">");
  document.writeln("<input type=\"checkbox\" id=\"keeplogin\" checked=\"checked\"/>保持登录");
document.writeln("</div>");
document.writeln("<div class=\"buttons\">");
document.writeln("<button class=\"sub_btn1\" id=\"login-sub\">登 录</button>");



document.writeln("</div>");
document.writeln("<div class=\"clr\"></div>");
document.writeln("<div class=\"reg\">");
document.writeln("忘记密码？<a href=\""+WebUrl+"e/member/GetPassword/\" >找回密码</a> <span id=errmsg></span>");
document.writeln("</div>");
document.writeln("</div>");
document.writeln("<!--------login end--------->");
document.writeln("</div>");
document.writeln("</div>");

document.writeln("</div>");
document.writeln("</div> ");
document.writeln("</div>");
document.writeln("<a href=\"#js_login\" id=\"AjaxLog\"></a><script>$('#AjaxLog').leanModal({ top: 200, overlay: 0.2, closeButton: '.close' })</script>");
function AjaxLog()
{
document.getElementById('AjaxLog').click();
switchTag('_xka','_xka_list',1,2,'conxk1','conxk2');
}
$('.close').click(function(){$("#lean_overlay").fadeOut(200);
$("#js_login").hide();
}
<!--location reload(); -->
);    
