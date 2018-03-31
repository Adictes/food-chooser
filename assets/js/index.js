function showMap() {
    document.getElementById("row-container").style.height = "10vh";
    document.getElementById("hidden-wrapper").style.display = "block";
}

function initMap() {
    var map = new google.maps.Map(document.getElementById("map"), {
        center: { lat: 53.3548, lng: 83.7698 },
        zoom: 16,
    });

    navigator.geolocation.getCurrentPosition(
        function(position) {
            var pos = {
                lat: position.coords.latitude,
                lng: position.coords.longitude,
            };
            map.setCenter(pos);

            var marker = new google.maps.Marker({
                position: pos,
                map: map,
            });
        },
        function() {
            console.log("Проблемы с геолокацией");
        }
    );
}
