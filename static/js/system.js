layui.use(['form','laydate','layer'], function(){
    var form = layui.form;
    var layer = layui.layer;
    var laydate = layui.laydate;
    var laypage = layui.laypage;

    laydate.render({
        elem: '#start_time',
        type: 'datetime'
    });

    laydate.render({
        elem: '#end_time',
        type: 'datetime'
    });

    form.on('switch(Status)', function(data){
        if (data.elem.checked){
            data.elem.value = 1
        } else {
            data.elem.value = 0
        }
    });

});

function showAllContent(o,data) {
    layer.open({
        type: 1,
        area: ['600px', '360px'],
        shadeClose: true, //点击遮罩关闭
        content: '\<\div style="padding:20px;display:block;word-break: break-all;word-wrap: break-word;line-height:22px">'+data+'\<\/div>'
    });
}

// 单条删除
function deleteDataByOne(obj,id)
{
    layer.confirm('确认要执行操作吗？',{btn:['确定','取消']},function(index){
        $.ajax({
            url: decodeURI(deleteUrl),
            data: {"id":id},
            type: "get",
            dataType: "json",
            success: function (data) {
                var messge = "网络繁忙...";
                if(data.Message) {
                    messge = data.Message;
                }
                layer.msg(messge,{icon:1,time:1000},function () {
                    if(data.Code > 0){
                        window.location.href=data.Data;
                    }
                });
            }
        });
        return false;
    });
}

// 批量删除
function deleteDataByBatch()
{
    var idArr = [];
    $(".layui-form-checked").each(function () {
        var currDataId = $(this).attr("data-id");
        if ("undefined" != typeof(currDataId)){
            idArr.push(currDataId)
        }
    });

    if(!idArr.length){
        layer.msg("未选中记录",{icon:1,time:1000});
        return
    }

    layer.confirm('确认要执行操作吗？',{btn:['确定','取消']},function(index){
        $.ajax({
            url: decodeURI(deleteUrl),
            data: {"id":idArr.join()},
            type: "get",
            dataType: "json",
            success: function (data) {
                var messge = "网络繁忙...";
                if(data.Message) {
                    messge = data.Message;
                }

                layer.msg(messge,{icon:1,time:1000},function () {
                    if(data.Code > 0){
                        window.location.href=data.Data;
                    }
                });
            }
        });
        return false;
    });
}

// 单条修改状态
function modifyStatusByOne(obj,id,status)
{
    layer.confirm('确认要执行操作吗？',{btn:['确定','取消']},function(index){
        $.ajax({
            url: decodeURI(modifyStatusUrl),
            data: {"id":id,"status":status},
            type: "get",
            dataType: "json",
            success: function (data) {
                var messge = "网络繁忙...";
                if(data.Message) {
                    messge = data.Message;
                }

                layer.msg(messge,{icon:1,time:1000},function () {
                    if(data.Code > 0){
                        window.location.href=data.Data;
                    }
                });
            }
        });
        return false;
    });
}


// 批量修改状态
function modifyStatusByBatch(status)
{
    var idArr = [];
    $(".layui-form-checked").each(function () {
        var currDataId = $(this).attr("data-id");
        if ("undefined" != typeof(currDataId)){
            idArr.push(currDataId)
        }
    });

    if(!idArr.length){
        layer.msg("未选中记录",{icon:1,time:1000});
        return
    }

    layer.confirm('确认要执行操作吗？',{btn:['确定','取消']},function(index){
        $.ajax({
            url: decodeURI(modifyStatusUrl),
            data: {"id":idArr.join(),"status":status},
            type: "get",
            dataType: "json",
            success: function (data) {
                var messge = "网络繁忙...";
                if(data.Message) {
                    messge = data.Message;
                }

                layer.msg(messge,{icon:1,time:1000},function () {
                    if(data.Code > 0){
                        window.location.href=data.Data;
                    }
                });
            }
        });
        return false;
    });
}

function formSubmit(url, data) {
    console.log(url)
    $.ajax({
        url: url,
        data: data,
        type: "post",
        dataType: "json",
        success: function (data) {

            var messge = "网络繁忙...";

            if(data.Message) {
                messge = data.Message;
            }

            layer.alert(messge, {icon: 6,time: 5000},function () {
                // 获得frame索引
                window.parent.location.reload();
                var index = parent.layer.getFrameIndex(window.name);
                //关闭当前frame
                parent.layer.close(index);
            });
        },
        error: function (XMLHttpRequest, textStatus, errorThrown) {
            layer.alert(messge, {icon: 6},function () {
                // 获得frame索引
                var index = parent.layer.getFrameIndex(window.name);
                //关闭当前frame
                parent.layer.close(index);
            });
        },
        beforeSend: function () {
        },
        complete: function () {
        }
    });
}