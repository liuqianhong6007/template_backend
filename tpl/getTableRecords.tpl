<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>表格数据</title>
<script src="https://cdn.staticfile.org/jquery/1.10.2/jquery.min.js"></script>
</head>
<body>

 <table border="1">
 {{range .}}
    <tr>
         {{range .}}
             <td>{{.Value}}</td>
         {{end}}
         <td>
            <button type="submit">修改</button>
            <button type="submit">删除</button>
         </td>
    </tr>
 {{end}}
 <tr>
    <td>
        <button type="submit">新增</button>
    </td>
 </tr>
 </table>

 <script>
    $(document).ready(function(){
        $("tr").hover(function(){
            $(this).css("background-color","yellow");
        },function(){
            $(this).css("background-color","white");
        });
    });
 </script>

</body>
</html>