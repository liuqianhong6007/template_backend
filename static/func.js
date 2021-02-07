var domain = "http://127.0.0.1:8081";

// jquery 入口
$(document).ready(function(){
    // table 列表
    $("#getTablesSpan").click(function(){
        url = domain + "/getTables";
        $.get(url,function(data,status){
            $("aside").html(data);
        })
    });

    // 弹出选项卡
    $("aside").on("mouseenter",".table-name",function(){
        $("#table-option-pop").html(`
            <button type="submit" id="getTableBtn">查看表结构</button>
            <button type="submit" id="getTableRecordBtn">查看数据</button>
        `);
    });
    $("aside").on("mouseleave",".table-name",function(){
        setTimeout(function () {
            $("#table-option-pop").html("");
        },4000);
    });

    // 查看表结构
    $("aside").on("click","#getTableBtn",function () {
        tableName = $(this).parent().parent().find(".table-name").text().trim();
        url = domain+"/getTable?tableName=" + tableName;
        $.get(url,function(data,status){
            $("main").html(data);
        });
    });

    // 查看表数据
    $("aside").on("click","#getTableRecordBtn",function () {
        tableName = $(this).parent().parent().find(".table-name").text().trim();
        url = domain+"/getTableRecords?tableName=" + tableName
        $.get(url,function(data,status){
            $("main").html(data);
        });
    });

    // 更新表数据
    $("main").on("click","#updateTableRecordBtn",function () {
        tableName = $(this).parent().parent().parent().parent().attr("id").trim();
        url = domain+"/updateTableRecord";

        columns = $(this).parent().siblings();
        var data = "";
        for (i = 0;i < columns.length;i++){
            data = data + $(columns[i]).attr("class").trim() + "=" + $(columns[i]).text() + ";"
        }
        data = data.substr(0,data.length-1);

        $.post(url, {
            "tableName": tableName,
            "data": data,
        },function(data,status){
            $("main").html(data);
        });
    });

    // 删除表数据
    $("main").on("click","#deleteTableRecordBtn",function () {
        tableName = $(this).parent().parent().parent().parent().attr("id").trim();
        url = domain+"/deleteTableRecord";

        columns = $(this).parent().siblings();
        var data = "";
        for (i = 0;i < columns.length;i++){
            data = data + $(columns[i]).attr("class").trim() + "=" + $(columns[i]).text() + ";"
        }
        data = data.substr(0,data.length-1);

        $.post(url, {
            "tableName": tableName,
            "data": data,
        },function(data,status){
            $("main").html(data);
        });
    });

    // 新增表数据
    $("main").on("click","#addTableRecordBtn",function () {
        columns = $(this).parent().find("table tr th")
        var data = "<tr>";
        for (i = 0;i < columns.length - 1;i++){
            data = data + `<td class=` + $(columns[i]).text().trim() + ` contentEditable="true"> </td>>`
        }
        data += `<td>
                    <button type="submit" id="submitAddTableRecordBtn">提交</button>
                    <button type="submit" id="cancelAddTableRecordBtn">取消</button>
                 </td>`
        data += `</tr>`;
        $(this).parent().find("table").append(data)
    });

    // 提交新增表数据
    $("main").on("click","#submitAddTableRecordBtn",function () {
        tableName = $(this).parent().parent().parent().parent().attr("id").trim();
        url = domain+"/addTableRecord";

        columns = $(this).parent().siblings();
        var data = "";
        for (i = 0;i < columns.length;i++){
            data = data + $(columns[i]).attr("class").trim() + "=" + $(columns[i]).text() + ";"
        }
        data = data.substr(0,data.length-1);

        $.post(url, {
            "tableName": tableName,
            "data": data,
        },function(data,status){
            $("main").html(data);
        });
    })

    // 取消新增表数据
    $("main").on("click","#cancelAddTableRecordBtn",function () {
        $(this).parent().parent().remove();
    })

    // 表格样式
    $("tr").hover(function(){
        $(this).css("background-color","yellow");
    },function(){
        $(this).css("background-color","white");
    });
});

function getTableColumn(table) {

}
