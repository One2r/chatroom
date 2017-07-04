<div class="navbar">
    <div class="navbar-inner">
            {{ if .isLogin}}
            <ul class="nav pull-right">
                <li id="fat-menu" class="dropdown">
                    <a href="#" role="button" class="dropdown-toggle" data-toggle="dropdown">
                        <i class="icon-user"></i> {{ .isLogin }}
                        <i class="icon-caret-down"></i>
                    </a>
                    <ul class="dropdown-menu">
                        <li class="divider visible-phone"></li>
                        <li><a tabindex="-1" href="/admin/signout.html">Logout</a></li>
                    </ul>
                </li>
            </ul>
            {{end}}
            <a class="brand" href="/admin/dashboard.html"><span class="first">{{.Appname}}</span> <span class="second">v{{.Appver}}</span></a>
    </div>
</div>