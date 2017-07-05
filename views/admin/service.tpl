<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{.Appname}} admin</title>
    {{ template "admin/public/meta.tpl" }}
  </head>

  <!--[if lt IE 7 ]> <body class="ie ie6"> <![endif]-->
  <!--[if IE 7 ]> <body class="ie ie7 "> <![endif]-->
  <!--[if IE 8 ]> <body class="ie ie8 "> <![endif]-->
  <!--[if IE 9 ]> <body class="ie ie9 "> <![endif]-->
  <!--[if (gt IE 9)|!(IE)]><!--> 
  <body class=""> 
  <!--<![endif]-->
    {{ template "admin/public/header.tpl" .}}
    {{ template "admin/public/sidebar.tpl"}}
    
    <div class="content">
        <div class="header">
            <h1 class="page-title">服务更新</h1>
        </div>
        
        <ul class="breadcrumb">
            <li><a href="dashboard.html">Home</a> <span class="divider">/</span></li>
            <li class="active">服务更新</li>
        </ul>

        <div class="container-fluid">
            <div class="row-fluid">
                <div class="row-fluid">
                    <div class="block">
                        <a href="#page-stats" class="block-heading" data-toggle="collapse">快捷操作</a>
                        <div id="page-stats" class="block-body collapse in">
                            <div class="stat-widget-container">
                                <div class="stat-widget">
                                    <div class="stat-button">
                                        <a class="btn btn-primary btn-large sensitive-update-btn">刷新敏感词<i class="icon-refresh"></i></a>
                                    </div>
                                </div>
                                <div class="stat-widget">
                                    <div class="stat-button">
                                        <a class="btn btn-primary btn-large replace-update-btn">刷新替换词<i class="icon-refresh"></i></a>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                {{ template "admin/public/footer.tpl"}}
            </div>
        </div>
    </div>
    {{ template "admin/public/js.tpl"}}
    <script src="/static/artDialog-7.0.0/dialog.js"></script>
    <script>
        $(".sensitive-update-btn").click(function(){
            $.ajax({
                url:"/admin/sensitive/update",
                type:"GET",
                success:function(result){ 
                    if(result.Error === undefined){
                        if(result.Data == true){
                            dialog({
                                title: '提示',
                                content: '刷新敏感词成功！',
                                cancelValue: '关闭',
                                cancel: function () {
                                    window.location.href=window.location.href;
                                }
                            }).width(320).show();
                            return;
                        }
                    }
                    dialog({
                        title: '提示',
                        content: '系统错误，请稍候再试！',
                        cancelValue: '关闭',
                        cancel: function () {}
                    }).width(320).show();
                }
            })
        });

        $(".replace-update-btn").click(function(){
            $.ajax({
                url:"/admin/replace/update",
                type:"GET",
                success:function(result){ 
                    if(result.Error === undefined){
                        if(result.Data == true){
                            dialog({
                                title: '提示',
                                content: '刷新替换词成功！',
                                cancelValue: '关闭',
                                cancel: function () {
                                    window.location.href=window.location.href;
                                }
                            }).width(320).show();
                            return;
                        }
                    }
                    dialog({
                        title: '提示',
                        content: '系统错误，请稍候再试！',
                        cancelValue: '关闭',
                        cancel: function () {}
                    }).width(320).show();
                }
            })
        });
    </script>
  </body>
</html>