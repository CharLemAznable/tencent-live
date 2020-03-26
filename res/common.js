layui.use(['jquery'], function () {
    let $ = layui.$;

    $('.live-header .layui-nav-item a').each(function () {
        let href = $(this).attr('href');
        let path = location.pathname;
        if (path.startsWith(href)) $(this).parent().addClass('layui-this');
    });
});