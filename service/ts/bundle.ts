const CLASSNAME = "visible";
const WAIT = 0;

setInterval(() => {
    $(".bg1").addClass(CLASSNAME);
    $(".bg2").addClass(CLASSNAME);
}, WAIT);