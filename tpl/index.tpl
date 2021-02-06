<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>表格</title>
<script src="https://cdn.staticfile.org/jquery/1.10.2/jquery.min.js"></script>
</head>
<body>
<div>
    <div id="left-frame">
        <table border="1">
           <tr>
             <th>表名称</th>
             <th>描述</th>
             <th>操作</th>
           </tr>

         {{range .}}
           <tr>
             <td class="tableName">{{.Name}}</td>
             <td>{{.Comment}}</td>
             <td>
                 <button type="submit" id="getTableBtn">查看表结构</button>
                 <button type="submit" id="getTableRecordBtn">查看数据</button>
             </td>
           </tr>
         {{end}}
         </table>
     </div>

     <div id="right-frame"></div>
 </div>

 <script>
    var domain = "http://127.0.0.1:8081";

    $(document).ready(function(){
        // 查看表结构
        $("#getTableBtn").click(function(){
            tableName = $(this).parent().parent().find(".tableName").text();
            url = domain+"/getTable?tableName=" + tableName;
            $.get(url,function(data,status){
                $("#right-frame").html(data);
            });
        });

        // 查看表数据
        $("#getTableRecordBtn").click(function(){
            tableName = $(this).parent().parent().find(".tableName").text();
            url = domain+"/getTableRecords?tableName=" + tableName
            $.get(url,function(data,status){
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
 <style>
    #right-frame {
         margin-top: 20px;
    }
 </style>

</body>
</html>