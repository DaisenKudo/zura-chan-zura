const CLASSNAME = "-visible";
const TIMEOUT = 0;
const $target = $(".bg");
setInterval(() => {
    $target.addClass(CLASSNAME);
    setTimeout(() => {
        $target.removeClass(CLASSNAME);
    }, TIMEOUT);
}, TIMEOUT * 2);
