var viewer = null;


function addItemToSlideList(img)
{
	console.info("new item added.")
}

function createList(imglist)
{
	// 循环请求
	for(var i = 0; i < imglist.length; ++i){
		var url = config.host + "/GetImageThumbNil?" + "path=" + imglist[i]
		  $.get(url,function(img,status){
		      if(status == "success"){
		          addItemToSlideList(img)
		      }else{
		          console.error("get image thumbnil failed. ")
		      }
		});
	}
}

function createViewer(host, path, w, h) {
    viewer = OpenSeadragon({
        id:            "openseadragon1",
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
                return host + "/GetImageTile?" + 
                "path=" + path + 
                "&level="+(level-8)+
                "&x="+x+
                "&y="+y;
            }
        }
    });
}

function Init() {
    var url = config.host + "/GetImageList"; 
    $.get(url,function(data,status){
        if(status == "success"){
            createList(data)
        }else{
            alert("服务器连接失败。");
        }
  });
}

