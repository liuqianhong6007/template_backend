<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>表格</title>
<script src="https://cdn.staticfile.org/jquery/1.10.2/jquery.min.js"></script>
</head>
<body>
    <header>模板后台</header>
    <nav>
        <span id="getTablesSpan" class="functionName">表查询</span>
    </nav>
    <aside></aside>
    <footer></footer>
    <script>
        var domain = "http://127.0.0.1:8081";

        $(document).ready(function(){
            // table 列表
            $("#getTablesSpan").click(function(){
                url = domain + "/getTables";
                $.get(url,function(data,status){
                    $("aside").html(data);
                })
            });

            // 表格样式
            $("tr").hover(function(){
                $(this).css("background-color","yellow");
            },function(){
                 $(this).css("background-color","white");
            });
         });
    </script>
    <style>
        header{
            background-color: LightSteelBlue ;
            color: black;
            line-height: 50px;
            text-align: center;
            marin: auto;
        }
        nav{
            background-color: AliceBlue;
            color: black;
            height: 20px;
            text-align: left;
        }
        aside{
            background-color: Linen ;
            color: black;
            height: 700px;
            width: 100%;
            float: left;
        }
        footer{
           background-color: LightSteelBlue;
           color: black;
           height: 50px;
           line-height: 50px;
           text-align: center;
        }
        #getTablesSpan{
            margin: 5px;
            color: red;
        }
        .functionName{
            background-color: LightSkyBlue;
        }
    </style>
</body>
</html>