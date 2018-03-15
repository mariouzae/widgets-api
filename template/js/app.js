var widgets = angular.module("widgets", []);
      // http interceptor to add token over all requests
      widgets.factory('httpRequestInterceptor', function () {
        return {
          request: function (config) {
            config.headers['Authorization'] = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHQiOjE1MjEwMDA4MzAsIm5hbWUiOiIifQ.E_l-58kMM63ei1pd1yU-h3yrwjZ2LNaWotrSG8GhvWc';
            return config;
          }
        };
      });

      widgets.config(function ($httpProvider) {
        $httpProvider.interceptors.push('httpRequestInterceptor');
      });

      widgets.controller('widgetController', function($scope, $http) {
        var self = this;
        var isPost = true;
        $scope.title = "Create Widget";
        $http.get('https://localhost:4000/api/v1/widgets').then(function(resp){
            $scope.widgets = resp.data;
          }, function(err){
            console.error("error");
          })

          var opts = ["red", "purple", "black", "green", "magenta", "white", "depends on the viewing angle"]
          $scope.options = opts
          $scope.selectedOption = $scope.options[0]

          $scope.buttom = "create"
          $scope.GetItem = function(index) {
            $scope.id = $scope.widgets[index].id
            $scope.name = $scope.widgets[index].name
            $scope.price = $scope.widgets[index].price
            $scope.selectedOption = $scope.options[opts.indexOf($scope.widgets[index].color)]
            $scope.melts = $scope.widgets[index].melts
            $scope.inventory = $scope.widgets[index].inventory
            $scope.buttom = "update"
            $scope.title = "Update Widget"
            isPost = false;
          }

          $scope.submit = function() {
            console.log($scope.selectedOption);
            var obj = {
              "name" : $scope.name,
              "price" : parseFloat($scope.price),
              "color" : $scope.selectedOption,
              "inventory" : parseInt($scope.inventory),
              "melts" : $scope.melts
            }

            if(isPost) {
              $http.post("https://localhost:4000/api/v1/widgets", obj).then(function(resp){
                console.log(resp);
              }, function(err){
                console.error(err);
              })
            } else {
              $http.put("https://localhost:4000/api/v1/widgets/" + $scope.id, obj).then(function(resp){
                console.log(resp);
              }, function(err){
                console.error(err);
              })
            }
          }

      });
         
      widgets.controller('dashboardController', function($scope, $http) {
        $http.get('https://localhost:4000/api/v1/users').then(function(resp){
           $scope.users = resp.data;
           $scope.usersLength = Object.keys(resp.data).length;
         }, function(err){
           console.error("error");
         });

         $http.get('https://localhost:4000/api/v1/widgets').then(function(resp){
           $scope.widgets = resp.data;
           $scope.widgetsLength = Object.keys(resp.data).length;
         }, function(err){
           console.error("error");
         });
      });      
      
      widgets.controller('usersController', function($scope, $http) {
         $http.get('https://localhost:4000/api/v1/users').then(function(resp){
           $scope.users = resp.data;
         }, function(err){
           console.error("error");
         })

         $scope.search = function(text) {
          valToSend = $scope.searchText
          $http.get('https://localhost:4000/api/v1/users/' + valToSend).then(function(resp){
            var u = [];
            u.push(resp.data)
            $scope.users = u;
          }, function(err){
            console.error("error");
          })
         }
      });