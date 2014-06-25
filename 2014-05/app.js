if (!String.prototype.fill) {
    String.prototype.fill = function() {
        var args = arguments;
        return this.replace(/{(\d+)}/g, function(match, number){
            return typeof args[number] != 'undefined'
                ? args[number] : match;
        });
    };
}


(function($){
    var stack = {};

    function listen(key, fn) {
        if (!stack.hasOwnProperty(key)) stack[key] = [];
        stack[key].push(fn);
        return "registered new function for {0}".fill(key);
    }

    function tell(key, p) {
        p = p || null;
        if (!stack.hasOwnProperty(key)) return "nothing listening.";
        $.each(stack[key], function(i, fn){
            fn(p);
        });
        return "called {0} with {1}".fill(key, JSON.stringify(p));
    }

    // export to global
    window.listen = listen;
    window.tell = tell;

}(jQuery));



(function($){
    var current = 0,
        pages = $('.kui-page');

    function forward() {
        if (current >= pages.length -1) return;
        $(pages[current]).hide();
        $(pages[current+1]).show();
        current++;
    }

    function backward() {
        if (current <= 0) return;
        $(pages[current]).hide();
        $(pages[current-1]).show();
        current--;
    }

    function bind() {
        $('.kui-nav button').bind('click', function(e){
            e.preventDefault();
            var cmd = $(this).data('control');
            tell(cmd);
        });
    }

    function first() {
        $(pages[0]).show();
    }

    listen('forward', forward);
    listen('backward', backward);
    listen('dom.ready', bind);
    listen('page.one', first);
}(jQuery));


(function($){
    $(function(){
        tell('page.one');
        tell('poll.start');
        tell('dom.ready');
    });
}(jQuery));

