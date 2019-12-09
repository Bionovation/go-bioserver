var viewer = null;

function createViewer(host,path,data) {
    viewer = OpenSeadragon({
        id:            "contentDiv",
        prefixUrl:     "/openseadragon/images/",
        navigatorSizeRatio: 0.2,
        wrapHorizontal:     false,
		showNavigator:  true, 
        minScrollDeltaTime: 25,
        maxZoomPixelRatio: 5.0,
        slideInfo: data,
        scrollWheelInSlider: function (scale) {
            $("#vertical-slider").slider("value", parseInt(scale));  
        },
        tileSources: {
            height: data.PhysicalHeight,
            width: data.PhysicalWidth,
            tileSize: 256,
            minLevel: 8,
            getTileUrl: function( level, x, y ){
                return host + "/slidetile?" + 
                "path=" + path + 
                "&level="+(level-8)+
                "&x="+x+
                "&y="+y;
            }
        }
    }); 
} 

function sliderInit(_slideInfo) {
    var SourceLens = _slideInfo.SourceLens;//viewer.viewport.
    $(".slider")
        .slider({
            min: 2,
            max: SourceLens*5,
            step: 0.1,
            value: 10,
            orientation: "vertical"
        })
        .slider("pips", {
            //rest: "label"
            rest: false
        })
        .slider("float", {
            suffix: "x"
        })
        .on("slidechange", function (e, ui) {
            //console.log(e);
            //console.log(ui);
            //console.log(e.originalEvent);
            if (typeof (e.originalEvent) != "undefined") {
                viewer.viewport.zoomtoscale(ui.value);
            }
            //viewer.viewport.zoomtoscale(ui.value);
        });
}



function Init(path) {    
    var url = config.host + "/slideinfo?path="+path; 
    $.get(url,function(data,status){
        if (status == "success") {
            sliderInit(data);
            createViewer(config.host, path, data);
        }else{
            alert("failed.");
        }
        
  });
}

