<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>表格</title>
<script src="https://cdn.staticfile.org/jquery/1.10.2/jquery.min.js"></script>
</head>
<body>
    <table border="1" id="{{.TableName}}">
        <tr>
        {{range .Columns}}
            <th>{{.}}</th>
        {{end}}
        <th>操作</th>
        </tr>

        {{range .Records}}
        <tr>
             {{range .Values}}
             <td>{{.}}</td>
             {{end}}
             <td>
                <button type="submit" id="updateTableRecordBtn">修改</button>
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
                // 查看表数据
                $("#updateTableRecordBtn").click(function(){
                   tableName = $(this).parent().parent().parent().attr("id");
                   url = domain+"/updateTableRecord";
                   $.post(url,
                   {
                        "tableName": tableName,
                        "data": data,
                   },function(data,status){
                       $("#right-frame").html(data);
                   });
                });

                // 表格样式
                $("tr").hover(function(){
                    $(this).css("background-color","yellow");
                },function(){
                     $(this).css("background-color","white");
                });
         });
     </script>
</body>
</html>