(function () {

    //Feature 6261014
    function generateUID() {
        //https://stackoverflow.com/questions/105034/how-to-create-a-guid-uuid/2117523#2117523
        return ([1e7]+-1e3+-4e3+-8e3+-1e11).replace(/[018]/g, c =>
            (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
        );
    };

    function appendGreenIDFrame(id) {
        var frame = document.createElement("iframe");
        frame.setAttribute("src", "https://fpt.microsoft.com/tags?session_id=" + id);
        frame.style.width = "100px";
        frame.style.height = "100px";
        frame.style.cssText = "display: none; color: rgb(0,0,0); float:left; position:absolute; top:-200px; left:-200px; border:0px";
        frame.title = "greenID";
        frame.setAttribute("id", "greenID");
        document.body.append(frame);
    }

    var uuid = generateUID();
    appendGreenIDFrame(uuid);
})();
