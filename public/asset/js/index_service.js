var index = angular.module('services.index', ['ngSanitize']);
if (typeof String.prototype.startsWith != 'function') {
  String.prototype.startsWith = function (prefix){
    return this.slice(0, prefix.length) === prefix;
  };
}
index.factory('MyData', function($websocket, $sce) {
    // Open a WebSocket connection
    
    var collection = [];
    var debugLogs   = [];
    var showDetail_ = false;
    var extColorColums_ = [];
    var clients = [];
    var spClient_;
    var lastMesg;

    var ws = $websocket("ws://"+location.hostname+":"+ location.port+"/ws");
    ws.onMessage(function(event) {

        // var res;
        // try {
        //     res = JSON.parse(event.data);
        // } catch (e) {
        //     res = {
        //         'username': 'anonymous',
        //         'message': event.data
        //     };
        // }
        var message = event.data;

        var debugPattern =  /.*=-= ==(\w*)==.*/;
        var networkPattern =  /.*==HTTP==.*/;
        
        if(debugPattern.test(message)) {
            paraseDebugLog(message);
        }
        else if(networkPattern.test(message)) {
            
        }
        else {
            paraseServerLog(message);
        }
        paraseClients(message);
    });

    var paraseClients = function(message) {
        var matches = message.split(" =-= ");
        if(matches.length>1 && clients.indexOf(matches[0])==-1) {
            clients.push(matches[0]);
        }

    }
    var paraseDebugLog = function(message){
        
        var found = message.match(/^==(\w*)==/);

        debugLogs.push({
            content: message,
            level: found[1],
            timeStamp: event.timeStamp
        });
    };

    var randomColor = function() {
        let colors = ['palette-turquoise', 'palette-emerald', 'palette-peter-river', 'palette-amethyst', 'palette-carrot','palette-wet-asphalt','palette-midnight-blue','palette-sun-flower'];
        return colors[Math.floor(Math.random()*colors.length)];
    }

   var safehtml = function(str) {
        return String(str).replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');
    }

    var paraseServerLog = function(message){

        var aHtml = message;

        if(spClient_.length>0 && !message.startsWith(spClient_)) {


            var clients = spClient_.toString().split(",");
            var bShow = false;
            clients.every(function(element, index){
                if(message.toString().startsWith(element)) {
                    bShow = true;
                    return false;
                }
                return true;
            })
            if(!bShow) return;
        }

        colums = ["ACT", "EVENT", "PSRC",  "PT", "LCN", "LTYPE",]; //"TYPE",
        colums.forEach(function(element){
            var pattern =  eval("/[<|]("+element+":[^|]+)/gi");
            var match = pattern.exec(aHtml);
            if(match) {
                aHtml = aHtml.replace(match[1], "<b class=\"" + randomColor()  +" demo_highlight\">"+safehtml(match[1])+"</b>" );
            }
        });


        if(!showDetail_) {
            var ignorepatter = /[<|](REALTIME|PROD|VER|PLAT|FROM|SRC|UUID|IDFA|UI|DEV|JAILB|OSV|CIP|DEP|NE|BLUETOOTH_NAME|BLUETOOTH_TYPE|LOCATIONID|BLKCNT|ENDTYPE|PRET|GPS|DOWNLOAD|BG|ONLYWIFI|BLKTM|PROJECT|APP|BR|T|CA|FMT|ISVIP|ISMPU|NE_TYPE|TeamID|DELAY|NOWPLAY|NOWPLAY_TYPE|PTYPE|DES):[^|]*/gi;
            aHtml = aHtml.replace(ignorepatter, "");
        }
        else {
            console.log("显示详情");
        }

        if(extColorColums_.length > 0) {
            console.log("进入扩展过滤");
            extColorColums_.forEach(function(element){
                var pattern =  eval("/[<|]("+element+":[^|]+)/gi");
                var match = pattern.exec(aHtml);
                if(match) {
                    aHtml = aHtml.replace(match[1], "<b class=\"" + randomColor()  +"  demo_highlight\">"+match[1]+"</b>" );
                }
            });
        }
        var ditto = aHtml == lastMesg
        lastMesg = aHtml

        collection.push({
            content: '<div class="demo-line">' +aHtml+"</div>",
            ditto:ditto,
            timeStamp: event.timeStamp
        });
    };

    ws.onError(function(event) {
        console.log('connection Error', event);
    });
    ws.onClose(function(event) {
        console.log('serverLogClient closed', event);
        ws.send('serverLogClient closed');
    });
    ws.onOpen(function() {
        console.log('connection open');
        ws.send('serverLogClient connect');
    });
    // setTimeout(function() {
    //   ws.close();
    // }, 500)
    setInterval(function(){
        if(collection.length>0 && collection[collection.length-1].content != "--") {
            collection.push({
                content: "--",
                timeStamp: "---"
            });
        }

        if(debugLogs.length>0 && debugLogs[debugLogs.length-1].content != "--") {
            debugLogs.push({
                content: "--",
                timeStamp: "---"
            });
        }
    
    }, 1000);

    var factory = {
        collection: collection,
        clients:clients,
    };
    factory.extendColumn = function(columns) {
        extColorColums_ = columns;
    };
    factory.doClientFilter = function(cli) {
        spClient_ = cli;
    };
    factory.showDetail = function(showDe) {
        showDetail_ = showDe;
    };
    factory.status = function() {
        return ws.readyState;
    };
    factory.send = function(message) {
        if (angular.isString(message)) {
            ws.send(message);
        } else if (angular.isObject(message)) {
            ws.send(JSON.stringify(message));
        }
    };
    factory.clear = function () {
        collection.length = 0;
    };

    return factory;
})
