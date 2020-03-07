var index = angular.module('services.index', ['ngSanitize']);

angular.module('LogChecker', ['ngWebSocket', 'services.index', 'luegg.directives', 'ngSanitize'])
.config(function($interpolateProvider) {
    $interpolateProvider.startSymbol('{[{');
    $interpolateProvider.endSymbol('}]}');
  })
.controller('indexController', function($scope, MyData) {
    $scope.MyData = MyData;
    $scope.glued = true;
    $scope.colorColums = "";
    $scope.filterValue = "ACT:VIDEO_LOG";
    $scope.simpleBtnText = "Detail";
    $scope.clientIp= "clientIP";
    $scope.clientsTxt = "(" + MyData.clients.length +")";

    $scope.isFileter = false;
    $scope.isnFileter = true;
    $scope.isSimple = false;

    //$.toaster({ priority : 'success', title : 'Notice', message : 'Your message here'});

    $scope.doColorize = function() {
        var cols = $scope.colorColums.split(",");
        $scope.MyData.extendColumn(cols);
    }

    $scope.$watch('clientIp', function(newVal, oldVal){
        $scope.clientIp = newVal;

        if(!newVal ||newVal.length<1 || newVal.toString().startsWith("clientIP")) { 
            $scope.MyData.doClientFilter("");
        }
        else {
            $scope.MyData.doClientFilter(newVal);
        }
      });

    $scope.$watch('MyData.clients.length', function(newVal, oldVal){
        if(parseInt(newVal) > parseInt(oldVal))
            $.toaster({ priority : 'danger', title : 'Notice', message : 'new client ' + $scope.MyData.clients[$scope.MyData.clients.length - 1] + ' connect'});
        else {
            $.toaster({ priority : 'info', title : 'Notice', message : 'a client leave'});
        }

        if(newVal != oldVal && $scope.clientIp.toString().startsWith("clientIP")) {
            $scope.clientsTxt = ("(" + newVal + ")");
        }

      });

    $scope.showSimple = function() {

        $scope.isSimple = !$scope.isSimple;
        $scope.MyData.showDetail($scope.isSimple);
        if($scope.isSimple)
            $scope.simpleBtnText = "Simple";
        else 
            $scope.simpleBtnText = "Detail";
    }

    $scope.clientFilter = function(msg) {
        if($scope.clientIp.length>0) {

            var clients = $scope.clientIp.split(",");
            var bShow = ture;
            clients.every(function(element, index){
                if(msg.toString().startsWith(element)) {
                    bShow = false;
                    return false;
                }
                return true;
            })

            return bShow;
        }
        else {
            return true;
        }
    }

    $scope.matchInArray = function(string, expressions) {

        var len = expressions.length,
            i = 0;
        for (; i < len; i++) {
            if (string.match(expressions[i])) {
                return true;
            }
        }
        return false;
    }

    $scope.myfilter = function (msg) {
        // console.log("begin filter" + msg);

        if(!$scope.isFileter) {
            return true;
        }

        // console.log("begin filter" + msg);
        var cols = $scope.filterValue.split(",");

        return $scope.matchInArray(msg.content, cols);
        // var pattern = new RegExp($scope.filterValue);
        // return msg.content.match(pattern);
    }

    $scope.randomColor = function() {
        let colors = ['palette-turquoise', 'palette-emerald', 'palette-peter-river', 'palette-amethyst', 'palette-carrot'];
        return colors[Math.floor(Math.random()*colors.length)];
    }   

    $scope.doClear = function () {
        $scope.MyData.clear();
    }

    $scope.doFilter = function (msg) {
        $scope.isFileter = !$scope.isFileter;
        $scope.isnFileter = !$scope.isnFileter;
    }

})
.filter('htmlTrust', ['$sce', function($sce){
    return function(item){
        return $sce.trustAsHtml(item);
    }
}])
.filter('colorize', ['$sce', function($sce){ //colorize:this
    return function(item, scope){
        if(item.length < 1) return;

        return $sce.trustAsHtml(item);
    }
}])