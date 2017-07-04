<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Bootstrap Admin</title>
     {{ template "admin/public/meta.tpl" }}
  </head>

  <!--[if lt IE 7 ]> <body class="ie ie6"> <![endif]-->
  <!--[if IE 7 ]> <body class="ie ie7 "> <![endif]-->
  <!--[if IE 8 ]> <body class="ie ie8 "> <![endif]-->
  <!--[if IE 9 ]> <body class="ie ie9 "> <![endif]-->
  <!--[if (gt IE 9)|!(IE)]><!--> 
  <body class=""> 
  <!--<![endif]-->
    {{ template "admin/public/header.tpl" }}
    <div class="row-fluid">
        <div class="dialog">
            <div class="block">
                <p class="block-heading">Sign In</p>
                <div class="block-body">
                    <form method="post" action="/admin/signin.html"> 
                        <label>Username</label>
                        <input type="text" name="username" class="span12">
                        <label>Password</label>
                        <input type="password" name="password" class="span12">
                        {{if .showMsg}}
                            <a href="javascript:;" data-toggle="collapse">账号或密码错误</a>
                        {{end}}
                        <button type="submit"  href="index.html" class="btn btn-primary pull-right">Sign In</button>
                        <div class="clearfix"></div>
                    </form>
                </div>
            </div>
            <p class="pull-right" style=""><a href="javascript:;" target="blank">Theme by Portnine</a></p>
        </div>
    </div>
    <script src="/static/theme/lib/bootstrap/js/bootstrap.js"></script>
  </body>
</html>