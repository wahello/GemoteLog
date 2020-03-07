var debug = angular.module('services.debug', ['ngSanitize']);

debug.factory('MyData', function($websocket, $sce) {
    // Open a WebSocket connection
    var ws = $websocket("ws://"+location.hostname+":19908/");
    var debugLogs   = [];
    ws.onMessage(function(event) {

        var message = event.data;
        console.log("RB:" + message);
        var debugPattern =  /^==(\w*)==.*/;
        if(debugPattern.test(message)) {
            paraseDebugLog(message);
        }
    });

    var paraseDebugLog = function(message){
        
        let errorLevel = ["ERROR", "WARNNING", "INFO", "DEBUG", "VERBOSE"];

        var found = message.match(/^==(\w*)==/);
        console.log(found);

        var levelStr = "";
        var levelNum = 10;
        if(found && found.length>0) {
            levelStr = found[1];
            var idx = errorLevel.indexOf(levelStr);
            if(idx > -1) {
                levelNum = idx;
            }
        }

        var newDate = new Date();
        var dateString = newDate.toLocaleTimeString();
        
        debugLogs.push({
            content: message,
            level:levelStr,
            levelN:levelNum,
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
        clear: function () {
            debugLogs.length = 0;
        }
    };
})