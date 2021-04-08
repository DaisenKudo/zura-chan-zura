const CLASSNAME = "visible";
const WAIT = 0;

setInterval(() => {
    $(".bg-1").addClass(CLASSNAME);
    $(".bg-2").addClass(CLASSNAME);
}, WAIT);