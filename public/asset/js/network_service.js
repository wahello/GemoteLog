var debug = angular.module('services.network', ['ngSanitize']);

debug.factory('MyData', function($websocket, $sce) {
    var ws = $websocket("ws://"+location.hostname+":19908/");
    var debugLogs   = [];
    var clients = [];
    var spClient_;

    ws.onMessage(function(event) {

        var message = event.data;
        // console.log("RB:" + message);
        var networkPattern =  /.*==HTTP==.*/;
        if(networkPattern.test(message)) {
            paraseDebugLog(message);
        }

        paraseClients(message);
    });

    var paraseClients = function(message) {
        var matches = message.split(" =-= ");
        if(matches.length>1 && clients.indexOf(matches[0])==-1) {
            clients.push(matches[0]);
        }

    }

    var randomColor = function() {
        let colors = ['palette-turquoise', 'palette-emerald', 'palette-peter-river', 'palette-amethyst', 'palette-carrot','palette-wet-asphalt','palette-midnight-blue','palette-sun-flower','palette-ALIZARIN'];
        return colors[Math.floor(Math.random()*colors.length)];
    }

    var paraseDebugLog = function(message){
        
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

        var pattern =  /[<|]URL:([^|]+)/gi
        var match = pattern.exec(message);
        if(match) {
            message = message.replace(match[1], "<a class=\"" + randomColor()  +" demo_highlight\" target=\"_blank\" href=\""+match[1]+"\">"+match[1]+"</a>" );
        }


        var newDate = new Date();
        var dateString = newDate.toLocaleTimeString();
        
        debugLogs.push({
            content: '<div class="demo-line">' +message+"</div>",
            timeStamp: dateString
        });
    };


    ws.onError(function(event) {
        console.log('connection Error', event);
    });
    ws.onClose(function(event) {
        console.log('connection closed', event);
        ws.send('debugClient closed');
    });
    ws.onOpen(function() {
        console.log('debugClient connect');
        ws.send('debugClient connect');
    });
    // setTimeout(function() {
    //   ws.close();
    // }, 500)
    setInterval(function(){
        if(debugLogs.length>0 && debugLogs[debugLogs.length-1].content != "--") {
            debugLogs.push({
                content: "--",
                timeStamp: "---"
            });
        }
    
    }, 1000);
    return {
        debugLogs: debugLogs,
        clients:clients,
        status: function() {
            return ws.readyState;
        },
        send: function(message) {
            if (angular.isString(message)) {
                ws.send(message);
            } else if (angular.isObject(message)) {
                ws.send(JSON.stringify(message));
            }
        },
        doClientFilter : function(cli) {
        spClient_ = cli;
        },
        clear: function () {
            debugLogs.length = 0;
        }
    };
})