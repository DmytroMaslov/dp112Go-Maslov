angular.module('golang')
	.constant('DEFAULTS', function() {
  		return {
			"task1": {
				"width": 8,
				"height": 4,
				"symbol": "#"
			}, 
			"task2": {"En1":{"S1":4, "S2":5},"En2":{"S1":10, "S2":20}}
		};
	}());
/*
//{"En1":{"S1":4, "S2":5},"En2":{"S1":10, "S2":20}}
angular.module('golang')
    .constant('DEFAULTS', function() {
        return {
            "task1": {
                "width": 8,
                "height": 4,
                "symbol": "#"
            },
            "task2": [{
                "width": 8,
                "height": 5
            }, {
                "width": 6,
                "height": 9
            }]
        };
    }());
    */