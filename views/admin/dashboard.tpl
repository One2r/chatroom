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
            <h1 class="page-title">Dashboard</h1>
        </div>
        
        <ul class="breadcrumb">
            <li><a href="dashboard.html">Home</a> <span class="divider">/</span></li>
            <li class="active">Dashboard</li>
        </ul>

        <div class="container-fluid">
            <div class="row-fluid">
                <div class="row-fluid">
                    <div class="block">
                        <a href="#page-stats" class="block-heading" data-toggle="collapse">Latest Stats</a>
                        <div id="page-stats" class="block-body collapse in">
                            <div class="stat-widget-container">
                                <div class="stat-widget">
                                    <div class="stat-button">
                                        <p class="title">{{.Statis.roomNum}}</p>
                                        <p class="detail">总房间数</p>
                                    </div>
                                </div>
                                <div class="stat-widget">
                                    <div class="stat-button">
                                        <p class="title">{{.Statis.online}}</p>
                                        <p class="detail">总在线人数</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="row-fluid">
                    <div class="block span6">
                        <a href="#tablewidget" class="block-heading" data-toggle="collapse">房间统计<span class="label label-warning">+10</span></a>
                        <div id="tablewidget" class="block-body collapse in">
                            <table class="table">
                            <thead>
                                <tr>
                                <th>房间ID</th>
                                <th>状态</th>
                                <th>当前在线人数</th>
                                <th>最高在线人数</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                <td>Mark</td>
                                <td>Tompson</td>
                                <td>the_mark7</td>
                                </tr>
                                <tr>
                                <td>Ashley</td>
                                <td>Jacobs</td>
                                <td>ash11927</td>
                                </tr>
                                <tr>
                                <td>Audrey</td>
                                <td>Ann</td>
                                <td>audann84</td>
                                </tr>
                                <tr>
                                <td>John</td>
                                <td>Robinson</td>
                                <td>jr5527</td>
                                </tr>
                                <tr>
                                <td>Aaron</td>
                                <td>Butler</td>
                                <td>aaron_butler</td>
                                </tr>
                                <tr>
                                <td>Chris</td>
                                <td>Albert</td>
                                <td>cab79</td>
                                </tr>
                            </tbody>
                            </table>
                        </div>
                    </div>
                </div>
                <footer>
                    <hr>
                    <p class="pull-right">Design by <a href="http://www.portnine.com" target="_blank">Portnine</a></p>
                </footer>
            </div>
        </div>
    </div>
    <script src="/static/theme/lib/bootstrap/js/bootstrap.js"></script>
  </body>
</html>