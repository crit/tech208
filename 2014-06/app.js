(function($, R){

    var stage = {
        w: $(window).width(),
        h: $(window).height() - 75,
        col: [],
        row: []
    };

    var a = 75, b = a/2, c = a/25;

    for (var i = 0; i < 12; i++) {
        stage.col.push((stage.w/12) * (i + 1) );
        stage.row.push((stage.h/12) * (i + 1));
    }

    var paper = R(0, 75, stage.w, stage.h);

    function node(x, y, text) {
        paper.rect(x - b, y - b, a, a).attr({"fill": "white"});
        paper.text(x, y, text);
        return {x: x, y: y};
    }

    function connect(x1, y1, x2, y2) {
        return paper.path('M' + x1 + ' ' + y1 + ' L' + x2 + ' ' + y2).attr({"stroke-dasharray": "- "}).toBack();
    }

    function rconnect(node1, node2) {
        return connect(node1.x + b, node1.y + c, node2.x - b, node2.y - c);
    }

    function lconnect(node1, node2) {
        return connect(node1.x - b, node1.y + c, node2.x + b, node2.y - c);
    }

    var slide = [], server = null, user = [], conn = [], cdn = [];

    slide.push(function slide1() {
        server = node(stage.col[5], stage.row[4], "ORIGINAL\nSERVER");
    });

    slide.push(function slide2() {
        user = [
            node(stage.col[0], stage.row[0], "CLIENT\nCOMPUTER"),
            node(stage.col[0], stage.row[4], "CLIENT\nCOMPUTER"),
            node(stage.col[0], stage.row[8], "CLIENT\nCOMPUTER"),
            node(stage.col[10], stage.row[0], "CLIENT\nCOMPUTER"),
            node(stage.col[10], stage.row[4], "CLIENT\nCOMPUTER"),
            node(stage.col[10], stage.row[8], "CLIENT\nCOMPUTER"),
        ];
    });

    slide.push(function slide3() {
        conn = [
            rconnect(user[0], server).attr({stroke: "red"}),
            rconnect(user[1], server).attr({stroke: "red"}),
            rconnect(user[2], server).attr({stroke: "red"}),
            lconnect(user[3], server).attr({stroke: "red"}),
            lconnect(user[4], server).attr({stroke: "red"}),
            lconnect(user[5], server).attr({stroke: "red"}),
        ];
    });

    slide.push(function slide4() {
        cdn = [
            node(stage.col[3], stage.row[2], "CDN\nSERVER"),
            node(stage.col[3], stage.row[6], "CDN\nSERVER"),
            node(stage.col[7], stage.row[2], "CDN\nSERVER"),
            node(stage.col[7], stage.row[6], "CDN\nSERVER"),
        ];

    });

    slide.push(function slide5() {
        $.each(conn, function(i, item){
            item.remove();
        });
        conn = [];

        conn = [
            rconnect(cdn[0], server).attr({stroke: "red"}),
            rconnect(cdn[1], server).attr({stroke: "red"}),
            lconnect(cdn[2], server).attr({stroke: "red"}),
            lconnect(cdn[3], server).attr({stroke: "red"}),
            rconnect(user[0], cdn[0]).attr({stroke: "blue"}),
            rconnect(user[1], cdn[0]).attr({stroke: "blue"}),
            rconnect(user[2], cdn[1]).attr({stroke: "blue"}),
            lconnect(user[3], cdn[2]).attr({stroke: "blue"}),
            lconnect(user[4], cdn[3]).attr({stroke: "blue"}),
            lconnect(user[5], cdn[3]).attr({stroke: "blue"}),
        ];
    });

    function play() {
        var loc = 0;
        var id = setInterval(function(){
            slide[loc]();
            loc++;
            if (loc == slide.length) clearInterval(id);
        }, 1000);
    }

    var floc = 0;
    function next() {
        if (floc >= slide.length) return;
        slide[floc]();
        floc++;
    }

    function reset() {
        document.location = document.location;
    }

    // on dom load
    $(function(){
        $('#play').click(play);
        $('#next').click(next);
        $('#reset').click(reset);
    });

}(jQuery, Raphael));
