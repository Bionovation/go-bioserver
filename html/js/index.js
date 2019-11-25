var viewer = null;

function createViewer(host,path, w, h) {
    viewer = OpenSeadragon({
        id:            "contentDiv",
        prefixUrl:     "/openseadragon/images/",
        navigatorSizeRatio: 0.25,
        wrapHorizontal:     false,
		showNavigator:  true, 
        tileSources:   {
            height: h,
            width:  w,
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

/* function Init(sampleId) {
    var path = config.datas[sampleId];
    var url = config.host + "/GetImageInfo?path="+path; 
    $.get(url,function(data,status){
        if(status == "success"){
            createViewer(config.host, path, data.width, data.height);
        }else{
            alert("failed.");
        }
        
  });
} */

function Init(path) {    
    var url = config.host + "/slideinfo?path="+path; 
    $.get(url,function(data,status){
        if(status == "success"){
            createViewer(config.host, path, data.width, data.height);
        }else{
            alert("failed.");
        }
        
  });
}

