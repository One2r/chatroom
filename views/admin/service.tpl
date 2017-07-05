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
                        <a href="#page-stats" class="block-heading" data-toggle="collapse">最新统计</a>
                        <div id="page-stats" class="block-body collapse in">
                            <div class="stat-widget-container">
                                <div class="stat-widget" style="width: 33.3%;">
                                    <div class="stat-button">
                                        <p class="title">{{.Statis.roomNum}}</p>
                                        <p class="detail">总房间数</p>
                                    </div>
                                </div>
                                <div class="stat-widget" style="width: 33.3%;">
                                    <div class="stat-button">
                                        <p class="title">{{.Statis.online}}</p>
                                        <p class="detail">当前线人数</p>
                                    </div>
                                </div>
                                 <div class="stat-widget" style="width: 33.3%;">
                                    <div class="stat-button">
                                        <p class="title">{{.Statis.MaxOnline}}</p>
                                        <p class="detail">最高线人数</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="row-fluid">
                    <div class="block">
                        <a href="#tablewidget" class="block-heading" data-toggle="collapse">房间信息</a>
                        <div id="tablewidget" class="block-body collapse in">
                            <table class="table">
                            <thead>
                                <tr>
                                <th>房间ID</th>
                                <th>全员禁言</th>
                                <th>当前在线人数</th>
                                <th>最高在线人数</th>
                                </tr>
                            </thead>
                            <tbody>
                                {{range $index, $elem := .Statis.rooms}}
                                <tr>
                                <td>{{$index}}</td>
                                <td>{{if $elem.Silence}}是{{else}}否{{end}}</td>
                                <td>{{$elem.online}}</td>
                                <td>{{$elem.MaxOnline}}</td>
                                </tr>
                                {{end}}
                            </tbody>
                            </table>
                        </div>
                    </div>
                </div>
                {{ template "admin/public/footer.tpl"}}
            </div>
        </div>
    </div>
    {{ template "admin/public/js.tpl"}}
  </body>
</html>