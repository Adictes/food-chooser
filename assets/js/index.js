var frws = new WebSocket("wss://" + window.location.host + "/frws");

frws.onmessage = function (event) {
    var data = JSON.parse(event.data);
    console.log(data);

    for (const i in data) {
        var marker = new google.maps.Marker({
            position: data[i].geometry.location,
            map: map,
        });
        var node = document.createElement("li");
        node.classList.add("list-group-item");
        var text = document.createTextNode(data[i].name);
        node.appendChild(text);
        document.getElementById("info-list").appendChild(node);
    }
};

function sendRequest() {
    navigator.geolocation.getCurrentPosition(
        function (position) {
            var pos = {
                lat: position.coords.latitude,
                lng: position.coords.longitude,
            };
            map.setCenter(pos);

            var query = document.getElementById("request_field").value;
            var str =
                position.coords.latitude.toString() +
                "|" +
                position.coords.longitude.toString() +
                "|" +
                query;
            frws.send(str);
        },
        function (error) {
            var infoWindow = new google.maps.InfoWindow({
                map: map,
                position: map.getCenter(),
            });
            switch (error.code) {
                case error.PERMISSION_DENIED:
                    infoWindow.setContent("User denied the request for Geolocation.");
                    break;
                case error.POSITION_UNAVAILABLE:
                    infoWindow.setContent("Location information is unavailable.")
                    break;
                case error.TIMEOUT:
                    infoWindow.setContent("The request to get user location timed out.")
                    break;
                case error.UNKNOWN_ERROR:
                    infoWindow.setContent("An unknown error occurred.")
                    break;
            }
            frws.send("error");
        }
    );
    document.getElementById("row-container").style.height = "10vh";
    document.getElementById("hidden-wrapper").style.display = "block";
}

var map;
function initMap() {
    map = new google.maps.Map(document.getElementById("map"), {
        center: { lat: 53.3548, lng: 83.7698 },
        zoom: 16,
    });
}
