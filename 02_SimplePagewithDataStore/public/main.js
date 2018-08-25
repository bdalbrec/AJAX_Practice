var el = document.querySelector("#enter-area");

var speedLimit = 250;
var lastTime = 0;

el.addEventListener("input", function(e){
    
    // wrap AJAX request it a timer to cut down on server calls
    var currentTime = new Date().getTime();
    if (currentTime - lastTime > speedLimit) {
        // create AJAX request
        var xhr = new XMLHttpRequest();
    
        // send AJAX request 
        xhr.open("post", "/api/check");
        console.log("Sending: ", e.target.value);
        xhr.send(e.target.value);

        lastTime = currentTime;
    

        //receive AJAX response
        xhr.addEventListener("readystatechange", function(){
            if (xhr.readyState === 4 && xhr.status === 200) {
                var taken = xhr.responseText;
                console.log("Received from server: ", xhr.responseText);

            }
        });
    }
});
