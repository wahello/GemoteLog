angular.module('LogChecker', ['ngWebSocket', 'services.network', 'luegg.directives', 'ngSanitize'])
.config(function($interpolateProvider) {
    $interpolateProvider.startSymbol('{[{');
    $interpolateProvider.endSymbol('}]}');
  })
.controller('networkController', function($scope, MyData) {
    $scope.MyData = MyData;
    $scope.glued = true;
    $scope.colorColums = "";
    $scope.filterValue = "==HTTP==";
    $scope.curLevel = 4;
    var isFileter = false;
    $scope.clientIp= "clientIP";
    $scope.clientsTxt = "(" + MyData.clients.length +")";

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

    $scope.myfilter = function (msg) {
        if(!isFileter) {
            return true;
        }
        console.log("begin filter" + msg);


        console.log("begin filter" + msg);

        var pattern = new RegExp($scope.filterValue);
        return msg.content.match(pattern);
    }

    $scope.doClear = function () {
        $scope.MyData.clear();
    }

    $scope.doFilter = function (msg) {
        isFileter = true;
    }

})
.filter('htmlTrust',function($sce){
    return function(item){
        return $sce.trustAsHtml(item);
    }
})
;