layui.config({
	base : "../../js/"
}).use(['flow','form','layer','upload'],function(){
    var flow = layui.flow,
        form = layui.form,
        layer = parent.layer === undefined ? layui.layer : top.layer,
        upload = layui.upload,
        $ = layui.jquery;


    function init(){
        initCombobox()
    }
    init()
    function initCombobox(){
        var str = '<option value="">请选择一个类型</option>';
        // var str = ''
        $.ajax({
            url:"/labelList/all",
            method:"get",
            success:function(data){
                 if(data){
                     data = eval('(' + data + ')');
                     if (data.code === 200){
                         let result = data.result;
                         for (let k in result){
                             str+='<option value="'+result[k]["id"]+'">'+result[k]["name"]+'</option>'
                         }

                         $('select').append(str);
                         layui.form.render('select','rtcombox');
                         return
                     }
                 }
                 alert("获取下拉框时发生错误！")
            }
        })
        layui.form.on('select(rtcombox)', function (data) {
            console.log(data);
            let text = $("#tableValue").val()
            $("#input_rtId").val(data.value)
            let data1 = {value:data.value,text:text}
            initUpload(data1)

        })
    }


    function initUpload(data){

        //多图片上传
        upload.render({
            elem: '.uploadFile',
            url: '/fileManager/batch?RtId='+data.value+'&Content='+data.text,
            accept: 'file',
            multiple: true,
            before: function(obj){
                //预读本地文件示例，不支持ie8
                obj.preview(function(index, file, result){
                    // $('#Images').prepend('<li><img layer-src="'+ result +'" src="'+ result +'" alt="'+ file.name +'" class="layui-upload-img"><div class="operate"><div class="check"><input type="checkbox" name="belle" lay-filter="choose" lay-skin="primary" title="'+file.name+'"></div><i class="layui-icon img_del">&#xe640;</i></div></li>')
                    // //设置图片的高度
                    // $("#Images li img").height($("#Images li img").width());
                    // form.render("checkbox");
                });
            },
            done: function(res){
                //上传完毕
            }
        });
    }
    //流加载图片
    var imgNums = 15;  //单页显示图片数量
    flow.load({
        elem: '#Images', //流加载容器
        done: function(page, next){ //加载下一页
            $.get("../../json/images.json",function(res){
                //模拟插入
                var imgList = [],data = res.data;
                var maxPage = imgNums*page < data.length ? imgNums*page : data.length;
                setTimeout(function(){
                    for(var i=imgNums*(page-1); i<maxPage; i++){
                        imgList.push('<li><img layer-src="../../'+ data[i].src +'" src='+ data[i].thumb +'"../../.." alt="'+data[i].alt+'"><div class="operate"><div class="check"><input type="checkbox" name="belle" lay-filter="choose" lay-skin="primary" title="'+data[i].alt+'"></div><i class="layui-icon img_del">&#xe640;</i></div></li>');
                    }
                    next(imgList.join(''), page < (data.length/imgNums));
                    form.render();
                }, 500);
            });
        }
    });

    //设置图片的高度
    $(window).resize(function(){
        $("#Images li img").height($("#Images li img").width());
    })








    //弹出层
    $("body").on("click","#Images img",function(){
        parent.showImg();
    })

    //删除单张图片
    $("body").on("click",".img_del",function(){
        var _this = $(this);
        layer.confirm('确定删除图片"'+_this.siblings().find("input").attr("title")+'"吗？',{icon:3, title:'提示信息'},function(index){
            _this.parents("li").hide(1000);
            setTimeout(function(){_this.parents("li").remove();},950);
            layer.close(index);
        });
    })

    //全选
    form.on('checkbox(selectAll)', function(data){
        var child = $("#Images li input[type='checkbox']");
        child.each(function(index, item){
            item.checked = data.elem.checked;
        });
        form.render('checkbox');
    });

    //通过判断是否全部选中来确定全选按钮是否选中
    form.on("checkbox(choose)",function(data){
        var child = $(data.elem).parents('#Images').find('li input[type="checkbox"]');
        var childChecked = $(data.elem).parents('#Images').find('li input[type="checkbox"]:checked');
        if(childChecked.length == child.length){
            $(data.elem).parents('#Images').siblings("blockquote").find('input#selectAll').get(0).checked = true;
        }else{
            $(data.elem).parents('#Images').siblings("blockquote").find('input#selectAll').get(0).checked = false;
        }
        form.render('checkbox');
    })

    //批量删除
    $(".batchDel").click(function(){
        var $checkbox = $('#Images li input[type="checkbox"]');
        var $checked = $('#Images li input[type="checkbox"]:checked');
        if($checkbox.is(":checked")){
            layer.confirm('确定删除选中的图片？',{icon:3, title:'提示信息'},function(index){
                var index = layer.msg('删除中，请稍候',{icon: 16,time:false,shade:0.8});
                setTimeout(function(){
                    //删除数据
                    $checked.each(function(){
                        $(this).parents("li").hide(1000);
                        setTimeout(function(){$(this).parents("li").remove();},950);
                    })
                    $('#Images li input[type="checkbox"],#selectAll').prop("checked",false);
                    form.render();
                    layer.close(index);
                    layer.msg("删除成功");
                },2000);
            })
        }else{
            layer.msg("请选择需要删除的图片");
        }
    })

})